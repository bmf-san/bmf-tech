---
title: "gogocoin — セルフホスティング暗号資産の自動取引ボットの紹介"
description: 'Go製セルフホスティング型ビットコイン自動取引ボット『gogocoin』の詳細解説。Strategyインターフェース設計、EMAスキャルピングのリスク管理、ダブルチェックロックの残高キャッシュ、VPSへのsystemdデプロイ。'
slug: introducing-gogocoin
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - Golang
  - Bitcoin
  - Trading
  - Infrastructure
translation_key: introducing-gogocoin
---

# gogocoin — セルフホスティング暗号資産の自動取引ボットの紹介

## 作った理由

OSSの仮想通貨ボットや自動取引サービスは多く存在する。それでもこれを作ったのは、自分で思い通りに動くものを実装し、自分の資金で実際に利益を上げる体験をしてみたかったからだ。以前一度作ったことがあるが、今回はAIを活用しながら改めて作り直した。実際に運用してみると学びが多く、趣味として続けられるソフトウェアになっている。

ボットは[bitFlyer](https://bitflyer.com/)の公式REST/WebSocket API経由で動作する。サンプル実装として、XRP/JPYペアのEMAスキャルピング戦略を同梱している—小資本で始めやすい最小注文金額の低い通貨ペアだ。

## 使い始め方

**対応取引所はbitFlyer一択。** APIアクセスには自作の[`go-bitflyer-api-client`](https://github.com/bmf-san/go-bitflyer-api-client)ライブラリを使用しており、他取引所には対応していない。注文発注にはbitFlyerの現物専用エンドポイント（`/v1/me/sendchildorder`）を使うため、**信用取引（FX\_BTC\_JPY等）は非対応**だ。

前提条件はbitFlyer APIキーとDockerのみ。口座さえあればすぐに試せる。

```bash
git clone https://github.com/bmf-san/gogocoin.git && cd gogocoin

# APIキーを設定（.envにBITFLYER_API_KEY / BITFLYER_API_SECRETを記入）
cp .env.example .env

# 設定ファイル生成 → 起動
make init && make up

# → http://localhost:8080 でダッシュボードが開く
```

起動直後はXRP/JPY・200円/回のスキャルピングが動く。通貨ペアは`config.yaml`の`trading.symbols`、注文サイズは`strategy_params.scalping.order_notional`で変更できる。

## アーキテクチャ

コードベースは層構造を採用している。`cmd/gogocoin`がエントリーポイント、`internal/`にドメインロジック・ユースケース・Infrastructureアダプター、`pkg/strategy`が戦略コントラクト定義と内蔵スキャルピング実装を提供する公開パッケージだ。

![ダッシュボード](/assets/images/posts/introducing-gogocoin/01_dashboard.png)

## Strategyインターフェース

すべての取引戦略は`pkg/strategy/strategy.go`で定義された`Strategy`インターフェースを実装する必要がある。これによりエンジンは特定のアルゴリズムから分離される。

```go
type Strategy interface {
    // 最新の市場データポイントからシグナルを生成
    GenerateSignal(ctx context.Context, data *MarketData, history []MarketData) (*Signal, error)

    // 履歴データのバッチからシグナルを生成
    Analyze(data []MarketData) (*Signal, error)

    // ライフサイクル
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
    IsRunning() bool
    GetStatus() StrategyStatus
    Reset() error

    // メトリクス等
    GetMetrics() StrategyMetrics
    RecordTrade()

    // 設定
    Name() string
    Initialize(config map[string]interface{}) error
    UpdateConfig(config map[string]interface{}) error
    GetConfig() map[string]interface{}
}
```

`Initialize()`は`config.yaml`の`strategy_params.<name>`ブロックを`map[string]interface{}`として受け取る。`UpdateConfig()`はHTTP API経由でボットを再起動せずにライブでパラメータを更新できる。

## スキャルピング戦略 — リスク管理

内蔵戦略（`pkg/strategy/scalping`）はデュアルEMAクロスオーバーにより買・売シグナルを生成する。クロスオーバーシグナルが発生する前に、2つのリスク管理ガードが動作する。

```go
func (s *Strategy) isInCooldown() bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    if s.lastTradeTime.IsZero() {
        return false
    }
    return time.Since(s.lastTradeTime).Seconds() < float64(s.cooldownSec)
}

func (s *Strategy) isDailyLimitReached() bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    // JST日付で当日の取引数を判定
    today := time.Now().UTC().Add(9 * time.Hour).Format("2006-01-02")
    if s.lastTradeDate != today {
        return false
    }
    return s.dailyTradeCount >= s.maxDailyTrades
}
```

`isInCooldown()`は直前の取引から一定時間が経つ前に再ポジションを取るのを防ぐ。デフォルトのクールダウンは90秒だ。`isDailyLimitReached()`はJSTカレンダー日当たりの取引数を設定値（デフォルト3回）以内に抑制する。

利確・損切り価格はエントリー価格から設定パーセンテージを使って計算する。

```go
func (s *Strategy) GetTakeProfitPrice(entry float64) float64 {
    s.mu.RLock()
    pct := s.takeProfitPct
    s.mu.RUnlock()
    return entry * (1.0 + pct/100.0)
}

func (s *Strategy) GetStopLossPrice(entry float64) float64 {
    s.mu.RLock()
    pct := s.stopLossPct
    s.mu.RUnlock()
    return entry * (1.0 - pct/100.0)
}
```

すべての可変フィールドは`sync.RWMutex`で保護されており、HTTPハンドラからの`UpdateConfig()`とエンジンゴルーチンからの`GenerateSignal()`が競合しない。

## 残高キャッシュ — ダブルチェックロック

取引ループは残高情報を頻繁に問い合わせる。ティックごとにレート制限付きのREST API呼び出しは適切でない。`BalanceService`は60秒TTLのキャッシュを保持し、ダブルチェックロックパターンでthundering herd問題を防ぐ。

```go
func (s *BalanceService) GetBalance(ctx context.Context) ([]domain.Balance, error) {
    // 第1チェック: 書き込みロックなしで読み取る
    s.cache.mu.RLock()
    cacheTimestamp := s.cache.timestamp
    cacheData := s.cache.data
    s.cache.mu.RUnlock()

    if time.Since(cacheTimestamp) < CacheDuration && len(cacheData) > 0 {
        result := make([]domain.Balance, len(cacheData))
        copy(result, cacheData)
        return result, nil
    }

    // APIフェッチをシリアライズ
    s.fetchMu.Lock()
    defer s.fetchMu.Unlock()

    // 第2チェック: ロック取得後に再確認
    s.cache.mu.RLock()
    cacheTimestamp = s.cache.timestamp
    cacheData = s.cache.data
    s.cache.mu.RUnlock()
    if time.Since(cacheTimestamp) < CacheDuration && len(cacheData) > 0 {
        result := make([]domain.Balance, len(cacheData))
        copy(result, cacheData)
        return result, nil
    }

    // ... API呼び出しとキャッシュ更新
}
```

外側の`cache.mu`（`sync.RWMutex`）がフレッシュなキャッシュデータの同時読み取りを許可する。内側の`fetchMu`（`sync.Mutex`）がAPI呼び出しをシリアライズし、キャッシュ失効時に正確に1つのゴルーチンのみが取得する。

## Webダッシュボード

内蔵Web UIは`http://localhost:8080`で動作し、サイドバーナビゲーションで4ページに分かれている。

- **ダッシュボード** — 総損益・本日損益・勝率・本日取引回数の4枚のサマリーカードと、接続状態・実行中の戦略・稼働時間・監視中の通貨ペア価格をまとめたシステム状態パネル。
- **パフォーマンス** — 通貨別残高テーブル（総量・利用可能量）、シャープレシオ・プロフィットファクター・最大ドローダウンの3リスク指標、日別損益テーブル。
- **取引履歴** — 日時・通貨ペア・売買区分・価格・数量・手数料・損益の一覧テーブル。
- **システムログ** — ログレベル（DEBUG / INFO / WARN / ERROR）とカテゴリ（system / trading / API / strategy）でフィルタリングできるログビューア。

トップバーの「開始」「停止」ボタンでボットをその場で制御できる。設定（API資格情報、取引パラメータ）は`config.yaml`で行う—外部データベース不要（データはSQLiteに記録）。

## VPSデプロイ

gogocoinは単一の静的リンクバイナリ一本で配布する。[gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template)はConoHa VPSでの運用を想定したセットアップサンプルで、systemd設定やデプロイ手順の参考に使える。systemdユニットファイルは次の通りだ。

```ini
[Unit]
Description=gogocoin - bitFlyer 自動取引ボット
After=network.target

[Service]
Type=simple
User=gogocoin
WorkingDirectory=/opt/gogocoin
EnvironmentFile=/opt/gogocoin/.env
ExecStart=/opt/gogocoin/gogocoin
Restart=always
RestartSec=5
NoNewPrivileges=true
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

API資格情報は`/opt/gogocoin/.env`から読み込む（バイナリやconfigファイルには直書きしない）。`make deploy`を使った手動デプロイと、linux/amd64向けにビルドして`scp`でVPSへ転送するGitHub Actionsワークフローの両方に対応している。

## まとめ

gogocoinは透明性とコントロールを重視したミニマルなセルフホスティング型自動取引ボットだ。Strategyインターフェースが戦略を交換可能にし、リスク管理層が過大な損失を防ぐ。残高キャッシュはAPI呼び出しをレート制限内に充分抑える。

- **GitHub**: [bmf-san/gogocoin](https://github.com/bmf-san/gogocoin)
- **VPSテンプレート**: [bmf-san/gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template)
