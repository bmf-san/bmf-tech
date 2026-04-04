---
title: "gogocoin — セルフホスティング暗号資産の自動取引ボットの紹介"
description: 'Go製セルフホスティング型ビットコイン自動取引ボット「gogocoin」の詳細解説。プラガブル戦略アーキテクチャ、レイヤー構成と依存ルール、取引フロー・データモデル・残高キャッシュの詳解。'
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

OSSの仮想通貨ボットや自動取引サービスは多く存在する。それでもこれを作ったのは、自分で思い通りに動くものを実装し、自分の資金で実際に利益を上げる体験をしてみたかったからである。以前一度作ったことがあるが、今回はAIを活用しながら改めて作り直した。実際に運用してみると学びが多く、趣味として続けられるソフトウェアになっている。

[gogocoin](https://github.com/bmf-san/gogocoin)

## 使い始め方

**対応取引所はbitFlyer一択。** APIアクセスには自作の[`go-bitflyer-api-client`](https://github.com/bmf-san/go-bitflyer-api-client)ライブラリを使用しており、他取引所には対応していない。注文発注にはbitFlyerの現物専用エンドポイント（`/v1/me/sendchildorder`）を使うため、**信用取引（FX\_BTC\_JPY等）は非対応**となっている。

gogocoinには2つの使い方がある。

**A. ライブラリとして使う（推奨）**

`example/` ディレクトリに完全な動作サンプルが入っている。自分のリポジトリに発展させる際の出発点として使える。

```bash
git clone https://github.com/bmf-san/gogocoin.git && cd gogocoin/example

# 設定ファイルを作成（APIキーを環境変数で指定）
cp configs/config.example.yaml configs/config.yaml
export BITFLYER_API_KEY=your_key
export BITFLYER_API_SECRET=your_secret

make run
# または: go run ./cmd/

# → http://localhost:8080 でダッシュボードが開く
```

`example/configs/config.example.yaml`をそのまま使うと、XRP/JPY・1000円/回のスキャルピングが動く。通貨ペアは`trading.symbols`、注文サイズは`strategy_params.scalping.order_notional`で変更できる。

`go get github.com/bmf-san/gogocoin@latest`で自分のリポジトリに組み込んで使うこともできる。

**B. Dockerで素早く試す**

`example/`に`Dockerfile`と`docker-compose.yml`が入っている。Aと同じEMA+RSIスキャルピング戦略が登録済みのバイナリをビルドして起動する。

```bash
git clone https://github.com/bmf-san/gogocoin.git && cd gogocoin/example
cp configs/config.example.yaml configs/config.yaml
# configs/config.yaml にAPIキーを記入
make up

# → http://localhost:8080 でダッシュボードが開く
```

Dockerfileのビルドコンテキストはリポジトリルートなので、`example/`ディレクトリから実行すること。

## アーキテクチャ

コードベースは4層のレイヤードアーキテクチャを採用している。`internal/`にドメインロジック・ユースケース・外部アダプター（bitFlyerクライアント・SQLiteリポジトリ・HTTPハンドラ等）、`pkg/strategy`がStrategyインターフェース定義とスキャルピングリファレンス実装を提供する公開パッケージになっている。Composition Root（全サービスの指定・組み立て）は呼び出し側のリポジトリに存在する（`example/cmd/main.go`がそのサンプル）。

### C4 Context — システム全体像

```mermaid
C4Context
    Person(operator, "オペレーター", "取引ボットの管理者")
    System(gogocoin, "gogocoin", "自動スキャルピング取引ボット")
    System_Ext(bitflyer, "bitFlyer", "仮想通貨取引所 REST / WebSocket API")
    System_Ext(sqlite, "SQLite", "ローカルデータベース")

    Rel(operator, gogocoin, "REST API で取引制御・状態確認")
    Rel(gogocoin, bitflyer, "注文発注・残高取得・マーケットデータ受信")
    Rel(gogocoin, sqlite, "取引記録・ポジション・パフォーマンス保存")
```

### C4 Container — 主要コンテナ

```mermaid
C4Container
    Person(operator, "オペレーター")

    System_Boundary(gogocoin, "gogocoin") {
        Container(example, "example/cmd", "Go", "Composition Root・起動/終了（呼び元が指定）")
        Container(http, "adapter/http", "Go net/http", "REST API サーバー")
        Container(worker, "adapter/worker", "Go goroutine", "バックグラウンドワーカー群")
        Container(usecase, "usecase/", "Go", "業務ロジック（trading / strategy / risk / analytics）")
        Container(domain, "domain/", "Go", "ドメインモデル・インターフェース定義")
        Container(infra_bf, "infra/exchange/bitflyer", "Go", "bitFlyer API クライアント")
        Container(infra_db, "infra/persistence", "Go + SQLite", "SQLite 永続化")
    }

    System_Ext(bitflyer_api, "bitFlyer API", "REST / WebSocket")
    SystemDb_Ext(sqlite, "SQLite")

    Rel(operator, http, "HTTP/JSON")
    Rel(example, http, "起動")
    Rel(example, worker, "起動")
    Rel(http, usecase, "uses")
    Rel(worker, usecase, "uses")
    Rel(usecase, domain, "uses")
    Rel(infra_bf, domain, "implements IFs")
    Rel(infra_db, domain, "implements IFs")
    Rel(infra_bf, bitflyer_api, "HTTPS / WSS")
    Rel(infra_db, sqlite, "SQL")
```

### C4 Component — usecase/trading

```mermaid
C4Component
    Container_Boundary(trading, "usecase/trading") {
        Component(trader, "BitflyerTrader", "Go", "注文発注・キャンセル・残高取得")
        Component(monitor, "OrderMonitor", "Go goroutine", "注文状態のポーリング監視")
        Component(pnl, "PnLCalculator", "Go", "約定後の損益計算・永続化")
        Component(balance, "BalanceService", "Go", "残高取得・キャッシュ")
        Component(order, "OrderService", "Go", "注文バリデーション・発注")
        Component(validator, "OrderValidator", "Go", "注文サイズ検証・残高チェック")
    }

    ComponentDb(tradeRepo, "TradeRepository", "domain.TradeRepository")
    ComponentDb(positionRepo, "PositionRepository", "domain.PositionRepository")
    ComponentDb(balanceRepo, "BalanceRepository", "domain.BalanceRepository")

    Rel(trader, monitor, "starts / watches")
    Rel(trader, order, "delegates PlaceOrder")
    Rel(trader, validator, "ValidateOrder / CheckBalance")
    Rel(monitor, order, "GetOrders（OrderGetter IF）")
    Rel(monitor, pnl, "saveTradeToDB → CalculateAndSave")
    Rel(monitor, balance, "UpdateBalanceToDB after fill")
    Rel(pnl, tradeRepo, "SaveTrade")
    Rel(pnl, positionRepo, "GetOpenPositions / UpdatePosition / SavePosition")
    Rel(trader, balance, "GetBalance")
    Rel(balance, balanceRepo, "SaveBalance / GetLatestBalances")
```

### 依存グラフ

```mermaid
graph LR
    cmd([cmd])

    cmd --> adp_http[adapter/http]
    cmd --> adp_worker[adapter/worker]
    cmd --> infra_bf[infra/exchange/bitflyer]
    cmd --> infra_db[infra/persistence]
    cmd --> domain([domain])

    adp_http --> uc_trading[usecase/trading]
    adp_http --> uc_analytics[usecase/analytics]
    adp_http --> domain
    adp_worker --> uc_trading
    adp_worker --> uc_strategy[usecase/strategy]
    adp_worker --> uc_risk[usecase/risk]
    adp_worker --> domain

    uc_trading --> domain
    uc_strategy --> domain
    uc_risk --> domain
    uc_analytics[usecase/analytics] --> domain

    adp_worker --> uc_analytics

    infra_bf --> domain
    infra_db --> domain
    logger --> domain
    cmd --> config[config]
    cmd --> logger
```

依存ルールはCIで強制している。

| ルール | 詳細 |
|---|---|
| `domain/` は内部importゼロ | stdlibのみ。インフラも usecase も知らない |
| `usecase/` は `infra/` をimportしない | `domain/` インターフェースにのみ依存する |
| `adapter/` は `infra/` の具体型を持たない | `domain/` インターフェースのみ使用 |
| `infra/` は `domain/` を実装する | `usecase/` や `adapter/` は知らない |
| Composition Root は呼び出し側リポジトリに存在する | `internal/` はwiring不要 |

公開API（セマンティックバージョニング対象）は `pkg/` 以下に分離している。`pkg/engine` がエンジン起動のエントリーポイント、`pkg/strategy` がStrategyインターフェースとレジストリを提供する。

## ユースケース

```mermaid
graph LR
    OP(["👤 オペレーター"])
    BF(["🏦 bitFlyer"])
    SYS(["⚙️ システム"])

    subgraph sys["gogocoin システム境界"]
        UC1(取引を開始する)
        UC2(取引を停止する)
        UC3(取引状態を確認する)
        UC4(ポジションを確認する)
        UC5(パフォーマンスを確認する)
        UC6(マーケットデータを確認する)
        UC7(残高を確認する)
        UC8(取引履歴を確認する)
        UC9(注文一覧を確認する)
        UC10(ログを確認する)
        UC11(設定を確認する)
        UC12(戦略をリセットする)
        UC13(スキャルピング戦略でシグナルを検知する)
        UC14(リスクをチェックする)
        UC15(注文を発注する)
        UC16(注文状態を監視する)
        UC17(損益を計算・記録する)
        UC18(古いデータをメンテナンスする)
        UC19(戦略パラメータを監視・更新する)
        UC20(注文タイムアウト・キャンセルを処理する)
    end

    OP --> UC1
    OP --> UC2
    OP --> UC3
    OP --> UC4
    OP --> UC5
    OP --> UC6
    OP --> UC7
    OP --> UC8
    OP --> UC9
    OP --> UC10
    OP --> UC11
    OP --> UC12
    UC13 --> UC14
    UC14 --> UC15
    UC15 --> UC16
    UC16 --> UC17
    BF -.->|"約定通知（ポーリング）"| UC16
    SYS --> UC13
    SYS --> UC18
    SYS --> UC19
    SYS --> UC20
```

オペレーターはHTTP API（WebダッシュボードのUIを含む）を通じて取引の制御と状態確認を行う。シグナル生成・注文・損益計算・データクリーンアップはシステムが自律的に実行する。

## 取引フロー

WebSocketでティックを受信してから注文が約定・損益記録されるまでの、システムの主要なシーケンスを示す。

### 6.1 スキャルピング取引フロー

```mermaid
sequenceDiagram
    participant SW as StrategyWorker
    participant ST as Strategy
    participant SigW as SignalWorker
    participant TC as TradingController
    participant RM as risk.Manager
    participant BP as balanceProvider
    participant TR as BitflyerTrader
    participant BF as bitFlyer API
    participant OM as OrderMonitor
    participant PNL as PnLCalculator
    participant BS as BalanceService
    participant DB as persistence
    participant CB as callback

    note over RM,BP: risk.Manager は balanceProvider ローカルIFに依存する。<br/>BP は BitflyerTrader が実装する。BitflyerTrader.GetBalance() は<br/>内部で BalanceService（TTL キャッシュ）に委譲する。
    note over SW,SigW: StrategyWorker はシグナルをチャネルに書き込む。<br/>SignalWorker がチャネルから受信してリスクチェック・発注を行う。

    SW->>ST: Analyze(history []MarketData)
    ST-->>SW: Signal(BUY)
    SW-)SigW: signalCh <- signal（チャネル送信）
    SigW->>TC: IsTradingEnabled()
    TC-->>SigW: true
    SigW->>RM: CheckRiskManagement(ctx, signal)
    RM->>BP: GetBalance(ctx)
    BP->>BS: GetBalance(ctx)
    note over BS: TTL キャッシュ確認（10秒）。<br/>キャッシュヒット時は BF を呼び出さない。
    alt キャッシュミス
        BS->>BF: GET /v1/me/getbalance
        BF-->>BS: balance
    end
    BS-->>BP: balance
    BP-->>RM: balance
    alt リスク違反（残高不足・ポジション過多）
        RM-->>SigW: non-nil error（残高不足・制限超過等）
        SigW->>SigW: skip（次のティックまで待機）
    else リスクOK
        RM-->>SigW: nil
        note over SigW: createOrderFromSignal() で domain.OrderRequest を生成
        SigW->>TR: PlaceOrder(ctx, order)
        TR->>BF: POST /v1/me/sendchildorder
        BF-->>TR: order_id
        note over TR,OM: MonitorExecution はgoroutineで起動。<br/>PlaceOrder は即座にreturnする（非同期）。
        TR-)OM: go MonitorExecution(ctx, result)

        loop ポーリング（最大90秒・15秒間隔）
            OM->>BF: GET /v1/me/getchildorders
            BF-->>OM: status=ACTIVE
        end

        BF-->>OM: status=COMPLETED
        note over OM,PNL: OrderMonitor.saveTradeToDB() が PnL を直接呼び出す。<br/>onOrderCompleted コールバックより前。
        OM->>PNL: CalculateAndSave(result)
        note over PNL,DB: SELL の場合 GetOpenPositions はトランザクション外（事前読み取り）。<br/>SQLite はデフォルトで serializable 相当の isolation を持つため<br/>phantom read リスクは実質なく、tx 開始前に読み取ることで<br/>tx 内の処理を最小化しデッドロックリスクを下げる。
        PNL->>DB: GetOpenPositions() [SELLのみ・tx外]
        PNL->>DB: BeginTx()
        PNL->>DB: SavePosition() [BUY] / UpdatePosition() [SELL]
        PNL->>DB: SaveTrade()
        PNL->>DB: Commit()
        PNL-->>OM: (pnl float64)
        OM->>BS: InvalidateBalanceCache()
        OM->>BS: UpdateBalanceToDB(ctx)
        BS->>BF: GET /v1/me/getbalance
        BS->>DB: SaveBalance(balance)
        OM->>CB: onOrderCompleted(result)
    end
```

`StrategyWorker` と `SignalWorker` はGoチャネル経由で非同期に連結している。`PlaceOrder()` は発注後すぐにreturnし、約定監視は `OrderMonitor` がgoroutineで行う。`PnLCalculator` はポジションと取引を同一トランザクション内で保存し、完了後に `OrderMonitor` が残高スナップショットを追記する。

### 6.2 REST API 取引制御フロー

```mermaid
sequenceDiagram
    participant C as HTTP Client
    participant H as adapter/http
    participant TC as TradingController
    participant DB as AppStateRepository

    C->>H: POST /api/trading/start
    H->>TC: SetTradingEnabled(ctx, true)
    TC->>DB: SaveAppState("trading_enabled", "true")
    DB-->>TC: nil
    TC-->>H: nil
    H-->>C: 200 OK

    C->>H: POST /api/trading/stop
    H->>TC: SetTradingEnabled(ctx, false)
    TC->>DB: SaveAppState("trading_enabled", "false")
    DB-->>TC: nil
    TC-->>H: nil
    H-->>C: 200 OK
```

### 6.3 マーケットデータ収集フロー

```mermaid
sequenceDiagram
    participant BS as bootstrap
    participant WM as WorkerManager
    participant WS as bitflyer WebSocket
    participant MW as MarketDataWorker
    participant DB as MarketDataRepository

    BS->>WS: Connect()
    BS->>WM: StartAll(ctx)
    WM-)MW: Run(ctx)

    loop ティックデータ受信
        WS-->>MW: Tick(price, volume, ...)
        MW->>DB: SaveMarketData(tick)
    end

    note over BS,WS: 切断時は bootstrap が再接続（WorkerManager のワーカーライフサイクルとは別）
```

### 6.4 注文タイムアウト / CANCELED・EXPIRED フロー

```mermaid
sequenceDiagram
    participant TR as BitflyerTrader
    participant BF as bitFlyer API
    participant OM as OrderMonitor
    participant PNL as PnLCalculator
    participant DB as persistence
    participant LOG as Logger

    TR->>BF: POST /v1/me/sendchildorder
    BF-->>TR: order_id
    note over TR,OM: MonitorExecution はgoroutineで起動（戻り値なし）。<br/>結果はonOrderCompletedコールバックで通知。
    TR-)OM: go MonitorExecution(ctx, result)

    alt タイムアウト（90秒経過）
        loop ポーリング継続中
            OM->>BF: GET /v1/me/getchildorders
            BF-->>OM: status=ACTIVE
        end
        OM->>BF: GET /v1/me/getchildorders（saveFinalOrderState）
        BF-->>OM: 最終ステータス確認
        OM->>LOG: Warn("Order monitoring timeout", order_id)
        note over OM: goroutine終了。PlaceOrderへの戻り値なし。
    else ターミナル状態（CANCELED / EXPIRED / REJECTED）
        OM->>BF: GET /v1/me/getchildorders
        BF-->>OM: status=CANCELED
        OM->>LOG: Warn("order terminal", status, order_id)
        note over OM,PNL: saveTradeToDB はCANCELED でも呼ばれてトレードを記録する。<br/>残高更新・onOrderCompleted コールバックは呼ばない。
        OM->>PNL: CalculateAndSave(result) [キャンセル記録]
        PNL->>DB: BeginTx()
        PNL->>DB: SaveTrade() [status=CANCELED]
        PNL->>DB: Commit()
    end
```

### 6.5 レート制限時のリトライフロー

```mermaid
sequenceDiagram
    participant UC as usecase
    participant BF as infra/exchange/bitflyer
    participant API as bitFlyer API

    UC->>BF: PlaceOrder(req)
    note over BF: Client.WithRetry() がリトライを管理する。<br/>usecase 層はリトライの存在を知らない。
    BF->>API: POST /v1/me/sendchildorder
    API-->>BF: 429 Too Many Requests
    loop MaxRetries 回まで（指数バックオフ）
        BF->>BF: exponential backoff 待機
        BF->>API: POST /v1/me/sendchildorder（retry）
    end
    alt リトライ成功
        API-->>BF: 200 OK
        BF-->>UC: order_id
    else リトライ上限超過
        BF-->>UC: domain.ErrRateLimitExceeded
        note over UC: errors.As(err, &apiErr) で *domain.Error に変換し<br/>apiErr.Type == domain.ErrTypeRateLimit で判定して上位に伝播させる
    end
```

### 6.6 MaintenanceWorker フロー

```mermaid
sequenceDiagram
    participant BS as bootstrap
    participant WM as WorkerManager
    participant MW as MaintenanceWorker
    participant DB as MaintenanceRepository
    participant LOG as Logger

    BS->>WM: StartAll(ctx)
    WM-)MW: Run(ctx)

    loop 定期実行（毎日深夜）
        MW->>DB: GetDatabaseSize()
        DB-->>MW: size bytes
        MW->>DB: CleanupOldData(retentionDays)
        DB-->>MW: deleted rows
        MW->>DB: GetTableStats()
        DB-->>MW: stats
        MW->>LOG: Info("maintenance done", stats)
    end

    note over MW: ctx.Done() 受信で即座に終了
```

## Strategyインターフェース

すべての取引戦略は`pkg/strategy/strategy.go`で定義された`Strategy`インターフェースを実装する必要がある。これによりエンジンは特定のアルゴリズムから分離される。

```go
// AutoScaleConfig は Strategy.GetAutoScaleConfig が返す注文サイズ自動スケール設定。
// エンジンがこれを使うことで、戦略固有の設定キーを直接読まずに注文金額を計算できる。
type AutoScaleConfig struct {
    Enabled     bool
    BalancePct  float64 // 利用可能なJPY残高に対する割合（0-100）
    MaxNotional float64 // JPYの上限。0 = 無制限
    FeeRate     float64
}

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
    InitializeDailyTradeCount(count int)

    // 設定
    Name() string
    Description() string
    Version() string
    Initialize(config map[string]interface{}) error
    UpdateConfig(config map[string]interface{}) error
    GetConfig() map[string]interface{}

    // 注文サイジング — エンジンが戦略固有の設定キーを直接参照しないよう各戦略が実装する
    GetStopLossPrice(entry float64) float64   // 0 = 損切なし
    GetTakeProfitPrice(entry float64) float64 // 0 = 利確なし
    GetBaseNotional(symbol string) float64
    GetAutoScaleConfig() AutoScaleConfig
}
```

`Initialize()`は`config.yaml`の`strategy_params.<name>`ブロックを`map[string]interface{}`として受け取る。`UpdateConfig()`はHTTP API経由でボットを再起動せずにライブでパラメータを更新できる。

戦略はグローバルレジストリへの自己登録で動作する。`database/sql`パッケージのドライバ登録と同じ仕組みで、各戦略パッケージの`register.go`が`init()`内で`strategy.Register("name", constructor)`を呼び出し、`main.go`はブランクインポート（`_ "github.com/bmf-san/gogocoin/pkg/strategy/scalping"`）で戦略を取り込む。新しい戦略の追加は`main.go`への`import`行の追加だけで済み、エンジンコードの変更は不要となっている。

## エンジンのリスク管理

エンジン（`StrategyWorker`）は戦略から独立して、毎ティックで損切り・利確を強制適用する。ティックのたびに`GetStopLossPrice` / `GetTakeProfitPrice`を呼び出し、価格が閾値を超えた瞬間にポジションをクローズする。シグナル発生を待つ必要はない。

`config.yaml`の`max_open_positions_per_symbol: 1`ガードはポジションの積み上がりを防ぐ。このガードがないと、下降トレンド中に複数のBUYシグナルが同一シンボルに積み上がり、損切り発動時にすべてが一度にクローズされて損失が倍増する。`1`に設定することで、同シンボルにオープンポジションが存在する場合はBUYを拒否する。

エンジンは`GetAutoScaleConfig()`を通じた残高比例注文サイジングもフレームワーク機能として持つ。有効化するには、自分の戦略実装で`GetAutoScaleConfig()`をオーバーライドして`Enabled: true`と割合・上限を返すよう実装する。

## 残高キャッシュ — ダブルチェックロック

取引ループは残高情報を頻繁に問い合わせる。ティックごとにbitFlyer REST APIを呼び出すと、分あたり50リクエストのレート制限をすぐに使い果たしてしまう。`BalanceService`は60秒TTLのキャッシュを保持し、ダブルチェックロックパターンでthundering herd問題を防ぐ。

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

## データモデル

取引データはSQLiteに保存する。テーブルとドメインモデルの対応は次の通りとなっている。

| テーブル | 内容 | 備考 |
|---|---|---|
| `trades` | 約定レコード | `order_id UNIQUE`で冪等性保証。不変（UPDATEなし） |
| `positions` | FIFOポジション | OPEN / PARTIAL / CLOSED の3状態。BUYで生成・SELLで更新 |
| `balances` | 残高スナップショット | 追記のみ（上書きなし）。通貨ごとに最新行を取得 |
| `market_data` | WebSocketティックデータ | UNIQUE(symbol, timestamp) |
| `performance_metrics` | 日次パフォーマンス指標 | 取引完了ごとにスナップショットを追記 |
| `logs` | 構造化ログ | `fields` カラムはJSON |
| `app_state` | 実行時フラグのKVストア | `trading_enabled`等 |

```mermaid
erDiagram
    TRADES {
        INTEGER id PK
        TEXT    symbol
        TEXT    side
        TEXT    type
        REAL    size
        REAL    price
        REAL    fee
        TEXT    status
        TEXT    order_id "UNIQUE"
        DATETIME executed_at
        DATETIME created_at
        DATETIME updated_at
        TEXT    strategy_name
        REAL    pnl
    }

    POSITIONS {
        INTEGER id PK
        TEXT    symbol
        TEXT    side
        REAL    size
        REAL    used_size
        REAL    remaining_size
        REAL    entry_price
        REAL    current_price
        REAL    unrealized_pl
        REAL    pnl
        TEXT    status
        TEXT    order_id
        DATETIME created_at
        DATETIME updated_at
    }

    BALANCES {
        INTEGER  id PK
        TEXT     currency
        REAL     available
        REAL     amount
        DATETIME timestamp
    }

    MARKET_DATA {
        INTEGER  id PK
        TEXT     symbol
        DATETIME timestamp
        REAL     open
        REAL     high
        REAL     low
        REAL     close
        REAL     volume
        DATETIME created_at
    }

    PERFORMANCE_METRICS {
        INTEGER  id PK
        DATETIME date
        REAL     total_return
        REAL     daily_return
        REAL     win_rate
        REAL     max_drawdown
        REAL     sharpe_ratio
        REAL     profit_factor
        INTEGER  total_trades
        INTEGER  winning_trades
        INTEGER  losing_trades
        REAL     average_win
        REAL     average_loss
        REAL     largest_win
        REAL     largest_loss
        INTEGER  consecutive_wins
        INTEGER  consecutive_loss
        REAL     total_pnl
    }

    LOGS {
        INTEGER  id PK
        TEXT     level
        TEXT     category
        TEXT     message
        TEXT     fields
        DATETIME timestamp
    }

    APP_STATE {
        TEXT     key PK
        TEXT     value
        DATETIME updated_at
    }

    POSITIONS ||--o{ TRADES : "symbol（FIFO・論理結合）"
```

外部キー制約は持たない。`positions` と `trades` の間が唯一のクロステーブル論理参照だが、`PnLCalculator` が両者を**同一トランザクション内**で書き込む（BeginTx → SavePosition/UpdatePosition → SaveTrade → Commit）。トランザクションのアトミック性が整合性を保証するため、DBレベルのFK制約は不要となっている。

## Webダッシュボード
![ダッシュボード](/assets/images/posts/introducing-gogocoin/01_dashboard.png)
内蔵Web UIは`http://localhost:8080`で動作し、サイドバーナビゲーションで4ページに分かれている。

- **ダッシュボード** — 総損益・本日損益・勝率・取引回数の4枚のサマリーカードと、接続状態・戦略・稼働時間・監視中の通貨ペア価格をまとめたシステム状態パネル。
- **パフォーマンス** — 残高、リスク指標（シャープレシオ・プロフィットファクター・最大ドローダウン）、日別損益。
- **取引履歴** — 日時・通貨ペア・売買・価格・数量・手数料・損益の一覧テーブル。
- **システムログ** — ログレベル（DEBUG / INFO / WARN / ERROR）とカテゴリ（システム・取引・API・戦略・UI・データ）でフィルタリングできるログビューア。

トップバーの「開始」「停止」ボタンでボットをその場で制御できる。設定（API資格情報、取引パラメータ）は`config.yaml`で行う—外部データベース不要（データはSQLiteに記録）。

## 運用例

gogocoinは単一の静的リンクバイナリ一本で配布する。[gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template)はConoHa VPSでの運用を想定したセットアップサンプルで、systemd設定やデプロイ手順の参考に使える。

`make setup`でVPSの初期セットアップ（systemdサービスのインストール等）を行い、デプロイは付属のGitHub Actionsワークフロー（`workflow_dispatch`）で自動化されており、linux/amd64向けにビルドして`rsync`でVPSへ転送する。

## まとめ

gogocoinは自分の取引戦略を自由に実装して実際の資金で動かせる、ミニマルなセルフホスティング型自動取引ボットである。自作のコードが直接損益に影響するというのはなかなかエキサイティングで、戦略のチューニングや機能の作り込みに際限なくハマれる面白さがある。興味があればぜひ触ってみてほしい。

- **GitHub**: [bmf-san/gogocoin](https://github.com/bmf-san/gogocoin)
- **VPSテンプレート**: [bmf-san/gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template)
