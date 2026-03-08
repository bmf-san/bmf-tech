---
title: RubyとRailsのIO・CPU特性について
slug: ruby-rails-io-cpu-characteristics
date: 2025-06-14T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
  - Ruby on Rails
translation_key: ruby-rails-io-cpu-characteristics
---


# 概要
Rubyの並行モデルやGVLの役割、Pumaサーバのスレッド・プロセス設計、IO/CPUバウンドの捉え方、計測手法によるボトルネック把握、Rails/Pumaデフォルト設定変更背景などを整理し、適切なチューニング方針を提示する。

# Rubyの並行モデルとGlobal VM Lock（GVL）
## GVLの存在意義
Ruby（MRI/CRuby）にはGlobal VM Lock（GVL）があり、同一プロセス内のRubyコード実行を同時に一スレッドに制限する。GVLはRuby VMがC言語で実装されていることに起因し、内部のメモリ管理やオブジェクト管理、GC（ガベージコレクション）の整合性を保つために存在する。たとえば、オブジェクト割り当てや解放時のヒープ操作、マーク＆スイープ型GCでのオブジェクトトラバース、メソッドキャッシュの更新、内部テーブルの操作などはスレッドセーフではなく、GVLによって同時実行を防ぐことでクラッシュやデータ破損を回避している。

C拡張（ネイティブ拡張）も多くがGVL下で動作する前提で設計されており、GVLを外すには拡張側のスレッド安全性担保が必要となるため、VM全体の整合性維持コストが非常に高い。アプリケーションレベルのスレッド安全性は開発者がMutexなどで担保する必要がある一方、GVLはVM内部の一貫性確保のための大域的なロックとして機能している。

## GVLとパフォーマンスへの影響
GVL下ではCPUバウンドなRubyコードは同一プロセス内で複数スレッドが並列実行できず、一度に一スレッドのみが実行される。一方、DBアクセスや外部API呼び出しなどでI/O待ちが発生するとGVLが解放され、他スレッドが実行を継続できるため、I/Oバウンド混在ワークロードではスレッド並行が有効になる。しかしGVL争奪のオーバーヘッドやスレッド切り替え遅延、GC実行時の一時停止などが絡むと、見かけ上I/O待ちに見えても実はCPU飢餓による待ちが含まれるケースがある。

## 他実装との比較
TruffleRubyやJRubyなどはGVLを持たないが、VM内部やJVMによるメモリ管理・スレッド管理方式に依存している。MRIを単純にGVLなしへ改造するのは膨大かつ困難であり、Rails利用者はマルチプロセスと適度なスレッド並行を活用する運用モデルで大抵のWebワークロードを十分扱える。

# Pumaによる並行処理とデフォルト設定の動向
## Pumaのアーキテクチャ
PumaはRails標準サーバとして広く使われる。マスタープロセスが`fork`で複数ワーカープロセスを生成し、各プロセス内でスレッドプールを利用してリクエストを処理する。I/O待ちでGVLが解放されスレッド切り替えが活きる場面がある一方、CPUバウンド部分ではプロセス並列により並列性能を発揮する。

以下は新規Railsアプリで生成される`config/puma.rb`の該当部分の抜粋例。デフォルトではスレッド数が環境変数`RAILS_MAX_THREADS`で設定され、ワーカー数は`WEB_CONCURRENCY`で制御される。

```ruby
# config/puma.rb
threads_count = ENV.fetch("RAILS_MAX_THREADS") { 3 }.to_i
threads threads_count, threads_count

workers ENV.fetch("WEB_CONCURRENCY") { 2 }

preload_app!

on_worker_boot do
  ActiveRecord::Base.establish_connection if defined?(ActiveRecord)
end
```

forkモデル（プロセス並列）とスレッドモデル（スレッド並行）のメリット・デメリットを簡単にまとめると以下の通り。

|       モデル        |                                               メリット                                               |                                          デメリット                                           |
| ------------------- | ---------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| プロセス並列 (fork) | - GVL制限を回避し、CPUバウンド処理で真の並列性能を発揮<br>- メモリ空間が独立し、クラッシュ影響を隔離 | - メモリ使用量が増加しやすい<br>- プロセス起動コストがかかる                                  |
| スレッド並行        | - メモリオーバーヘッドが小さく軽量<br>- I/O待ち中に他スレッドが動作しやすい                          | - GVLの影響でCPUバウンド並列性能は制限される<br>- スレッド競合やGVL争奪による待ちが発生し得る |

上記を踏まえ、Pumaではプロセス数(`workers`)とスレッド数(`threads`)の組み合わせを、アプリのワークロード特性とインフラリソースに応じて調整することが重要である。プロセス数(`workers`)とスレッド数(`threads`)の組み合わせを、アプリのワークロード特性とインフラリソースに応じて調整することが重要である。

## デフォルトスレッド数変更の背景

[GitHub Issue #50450](https://github.com/rails/rails/issues/50450)で議論されたように、Railsの新規アプリ生成時のPumaスレッド数デフォルトは従来の5から3に変更された。IssueではDHHが自身の運用経験を基に「ワーカーあたりスレッド数1が低レイテンシに寄与する」と提案し、多数の開発者が自アプリのベンチマーク結果やAmdahlの法則を用いた考察を共有した。主な検討ポイントはレイテンシとスループットのトレードオフ、I/O/CPU特性別の最適スレッド数、Heroku Dynoやコンテナ環境などリソース制約下での安全マージン確保などであった。結果として多くのアプリで3スレッド程度がバランスの良い妥当値と合意され、Rails 7.2でデフォルトが5から3に引き下げられた。既存アプリは明示的に`RAILS_MAX_THREADS`や`WEB_CONCURRENCY`を設定している場合影響を受けず、新規プロジェクトではまず3スレッドで開始し、モニタリングやベンチマーク結果に応じて適宜調整することが推奨される。

# IOバウンド vs CPUバウンドの誤認と測定手法
## 見かけ上I/O待ちに見える落とし穴
RailsログやAPM計測で「Query took: XX ms」と記録される時間には、実際のDB応答時間以外にスレッドスケジューリング待ち、GVL待ち、GC実行時間などが含まれる可能性がある。これを「DB待ちが支配的」と誤認すると、スレッド数を過度に増やしてGVL争奪を悪化させ、逆にパフォーマンスを低下させる恐れがある。

## GC時間計測
Ruby 3.x以降では`GC.total_time`がナノ秒単位の累積カウンタとして提供され、特定ブロック前後の差分でGCに要した時間を把握できる。Rails 7.2以降ではActiveSupport::Notifications経由でリクエストログにGC時間が含まれるようになり、GC負荷の影響を可視化できる。

## GVL待ち時間の可視化
Ruby 3.2以降のGVL Instrumentation APIと専用gem（例: gvltoolsなど）を使い、I/O部分とGVL待ち時間を分離計測する手法がある。これにより、バックグラウンドでCPU負荷が高い状況下でのGVL待ち増大を具体的に把握し、誤認を減らせる。

## OSスケジューラ待ちの勘所
OSレベルのスケジューラ待ち時間もI/O計測に含まれる場合があるが、個別I/Oごとの正確な計測は困難。Linuxの`/proc/<pid>/schedstat`などを活用し、コンテナやホスト全体のrunqueue待ち状況を監視することで、プロセス数やスレッド数の過不足を判断する指針となる。

## プロファイリングの重要性
上記各種計測によりアプリケーションのI/O/CPU比率やGVL待ちの実態を把握し、Amdahlの法則的視点でスレッド数やプロセス数を決める。デフォルトに従うだけでなく、自身のワークロード特性（外部API呼び出し頻度、DBアクセスパターン、レンダリング負荷など）をプロファイリングして最適化することが重要である。

# バックグラウンドジョブと並行度設定
Sidekiqなどのジョブ処理ではI/O集約的な処理（外部API呼び出し、ファイル操作、メール送信など）が多いため、高めの並行度設定（例えばconcurrency: 10〜25程度）を採用するケースがある。しかし以下の点に注意が必要である。

- **Sidekiqのconcurrency設定例**
  - `sidekiq.yml`で設定可能:
    ```yaml
    :concurrency: 15
    ```
  - 環境変数で上書く場合:

    ```bash
    export SIDEKIQ_CONCURRENCY=15
    bundle exec sidekiq
    ```
  - 並行度を上げるとI/O待ち中に他スレッドが動作しやすくなり、理論的にはスループット向上が期待できるが、GVL争奪やGC負荷増大による副作用もある。
- **GVL影響の測定ケーススタディ（擬似例）**
  - 目的: ジョブ内で複数の並行タスクが発生した場合のGVL待ち時間やスレッドスタリングを把握する。
  - 手順例:
    1. テストジョブを用意し、I/O部（sleepや外部呼び出しを模した処理）とCPU部（計算処理）を組み合わせる。
       ```ruby
       class BenchmarkJob
         include Sidekiq::Job
         def perform
           start = Process.clock_gettime(Process::CLOCK_MONOTONIC)
           # I/O模擬: sleepや小規模HTTPリクエスト
           sleep 0.02
           # CPU模擬: 計算負荷
           (1..200_000).each { |i| i*i }
           duration = Process.clock_gettime(Process::CLOCK_MONOTONIC) - start
           logger.info("Job duration: #{(duration*1000).round(1)}ms")
         end
       end
       ```
    2. GVL計測ツール（例: gvltools）を導入し、ジョブ実行中のGVL待ち時間を計測する。
       ```ruby
       require 'gvltools'
       class BenchmarkJob
         include Sidekiq::Job
         def perform
           GVLTools::LocalTimer.enable
           start_io = GVLTools::LocalTimer.monotonic_time
           sleep 0.02
           io_wait = GVLTools::LocalTimer.monotonic_time - start_io

           start_cpu = GVLTools::LocalTimer.monotonic_time
           (1..200_000).each { |i| i*i }
           cpu_time = GVLTools::LocalTimer.monotonic_time - start_cpu

           gvl_wait = GVLTools::LocalTimer.gvl_wait_time
           logger.info("I/O time: #{io_wait.round(3)}s, CPU time: #{cpu_time.round(3)}s, GVL wait: #{gvl_wait.round(3)}s")
         ensure
           GVLTools::LocalTimer.disable
         end
       end
       ```
    3. 並行度（concurrency）を変えながら複数ジョブを同時に投入し、ログのI/O時間、CPU時間、GVL待ち時間を比較する。
       - 例えばconcurrency: 5, 10, 20の設定で、それぞれ10〜50並列ジョブを実行し、GVL待ちがどの程度増加するかを観察する。
       - GVL待ち時間が急増するポイントを特定し、実運用での安全な並行度上限を把握する。

- **監視指標の設定**
  - SidekiqダッシュボードやPrometheusなどでジョブ処理時間、スループット、キュー長を監視。
  - RubyプロセスのGC時間やメモリ使用量、CPU使用率、runqueue待ちなどをメトリクス収集し、並行度変更時の影響を可視化する。

- **ベンチマークとチューニング手順**
  1. 既存ジョブのプロファイリング: 実運用近似のワークロードで、ジョブ処理時間のうちI/O/CPU比率を把握する。
  2. Amdahlの法則的視点で並行度候補を算出: I/O比率が高ければスレッド並行優先、CPU比率が高ければプロセス分割やワーカー台数増加を検討。
  3. 実測ベンチ: 異なるconcurrency設定で負荷テストを行い、処理時間、GVL待ち、GC、CPU使用率を比較。
  4. 運用環境反映: テスト結果を踏まえた最適並行度をステージングや本番で段階的に適用し、安定性とパフォーマンスを確認。

これにより、SidekiqなどのバックグラウンドジョブでもGVL影響を把握し、最適な並行度設定を導き出すことが可能となる。

# Ruby実行性能への取り組み
## JIT（YJITなど）の恩恵
YJIT導入によるレイテンシ改善事例は多数あり、I/O待ちが多い前提でも多くのアプリで15-30%程度の改善が見られることから、Rubyコード実行コストも無視できない。

## GVL削除の検討
GVL削除の議論はあるものの、MRI RubyでGVLを完全に消すのはC拡張やVM内部変更を含め膨大かつリスクの高い作業となる。TruffleRuby/JRubyやPythonのGIL削除事例から学びつつ、多くのWebワークロードではGVL下でのマルチプロセス・適度なスレッド並行で十分対応可能である。

# 運用・チューニング指針
- Rails新規アプリではPumaスレッドをデフォルト3で開始し、モニタリング結果に基づき変更する。プロセス数（WEB\_CONCURRENCY）はCPUコア数やインフラ環境（コンテナ/Heroku Dyno等）の特性を踏まえて設定する。
- 本番相当負荷下で、リクエストログに含まれるDB時間、GC時間、GVL待ち（可能な場合）、外部API呼び出し時間などを集計し、I/O/CPU比率とスレッド・プロセス構成の影響を評価する。
- 複数スレッド構成（例: 1〜5程度）、複数プロセス構成でレイテンシ・スループットを計測し、Amdahlの法則的視点で最適ポイントを探る。I/O重視ならスレッド並行、CPU重視ならプロセス並列の比率を調整する。
- Rubyバージョンに応じたGC設定（RUBY\_GC\_HEAP\_\*, pausetimeなど）を検討し、GCによる一時停止を最小化する。ログからGC時間指標を可視化し、必要に応じてGCパラメータを調整する。
- ホスト/コンテナのCPU利用率、runqueue状況、メモリ使用状況、I/O待ち指標などを監視し、プロセス・スレッド数設定がホストリソースと整合しているか確認する。
- アプリの成長やトラフィック特性の変化に応じてプロファイリングと設定見直しを定期的に行う。新たなRuby/RailsバージョンやJITの進化、エコシステムの更新に注意を払う。

# まとめ
Ruby/Railsのパフォーマンス最適化にはGVLやスレッド、プロセス、I/O/CPUバウンド特性、GC、OSスケジューラ待ちなど多面的な理解が求められる。計測に基づく実態把握と適切なチューニングを継続的に行うことで、レイテンシやスループット要件に柔軟に対応できるシステムを構築できる。

# 参考記事
- [Railsスケーリング（1）Puma、コンカレンシー、GVLのパフォーマンスへの影響を理解する（TechRacho翻訳）](https://techracho.bpsinc.jp/hachi8833/2025_06_09/151182)
- [The Mythical IO-Bound Rails App（byroot氏記事）](https://byroot.github.io/ruby/performance/2025/01/23/the-mythical-io-bound-rails-app.html)
- [Instrumenting Thread Stalling in Ruby Applications（byroot氏記事）](https://byroot.github.io/ruby/performance/2025/01/23/io-instrumentation.html)
- [So You Want To Remove The GVL?（TechRacho翻訳）](https://techracho.bpsinc.jp/hachi8833/2025_03_03/148712)
- [GitHub Issue: Set a new default for the Puma thread count (rails/rails#50450)](https://github.com/rails/rails/issues/50450)

