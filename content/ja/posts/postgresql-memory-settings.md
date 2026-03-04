---
title: "PostgreSQLのメモリ設定"
slug: "postgresql-memory-settings"
date: 2025-06-14
author: bmf-san
categories:
  - "データベース"
tags:
  - "PostgreSQL"
draft: false
---

# 概要
データベースの性能向上や安定運用には適切なメモリ設定が必要である。ディスクアクセスはメモリアクセスに比べ極めて遅く、可能な限りメモリから読み書きすることで応答性能を向上させたい一方、過度なメモリ割り当てはOOM（Out Of Memory、メモリ不足によるプロセス強制終了）リスクを高め、システム全体の停止につながる可能性がある。したがって、安定性を担保しつつ性能を確保するために、PostgreSQLのメモリ管理設定を慎重に行う必要がある。

本稿では、PostgreSQL公式ドキュメントや実運用の知見、Gihyo記事などを踏まえ、共有メモリ域とローカルメモリ領域の基本構成から主要パラメータの設定指針、運用上の検証手順までをまとめる。

# PostgreSQLのプロセス構成とメモリ領域の区分
PostgreSQLはマルチプロセスモデルを採用しており、サーバ起動時に生成されるマスタプロセス、WAL書き込みなどを担うバックグラウンドプロセス群、クライアント接続ごとに生成されるバックエンドプロセス（セッションプロセス）から構成される。各プロセスは独自にメモリを使用し、特に接続数に比例して増加するバックエンドプロセスのメモリ割り当てが全体のメモリ消費に大きく影響する。

メモリ管理は大きく以下の2つに分かれる。

1. **共有メモリ域（shared memory）**
   サーバ起動時に確保量が決定し複数プロセス間で共有される領域。主な設定項目には`shared_buffers`、`wal_buffers`、Free Space MapやVisibility Mapなどがある。
2. **ローカルメモリ領域（process-local memory）**
   各バックエンドプロセスごとに確保される作業用メモリ。ソートやハッシュ結合、メンテナンス操作などで利用される領域で、`work_mem`、`maintenance_work_mem`、`temp_buffers`などが該当し、動的にSET可能なものもある。

# shared_buffersの設定指針
`shared_buffers`はPostgreSQLがデータベースキャッシュとして使用する共有メモリの量を設定するパラメータである。デフォルト128MBは小さいため、専用サーバであればシステムメモリの約25%を初期値とし、OSキャッシュとのバランスを見ながら段階的に増加を試みる。

- デフォルト・単位: デフォルト128MB。単位未指定時はBLCKSZ（通常8kB）単位とみなされるが、実運用では`shared_buffers = '2GB'`など明示的に単位付きで指定することが推奨される。
- 推奨設定範囲: システムメモリの25%程度を初期設定とし、上限は40%を目安。
- 再起動要件: 設定変更には再起動が必要。
- OSカーネル設定: `shmmax`やTransparent Huge Pages無効化、NUMA最適化が必要な場合がある。
- WAL/チェックポイント関連: `shared_buffers`増加に伴い`max_wal_size`や`checkpoint_completion_target`の調整が必要。これによりチェックポイント時のwrite burstやI/Oスパイクを緩和できる。
- ワークロード特性: 読み取り主体のワークロードでは効果が高いが、書き込み中心の場合はI/O負荷やチェックポイントの影響に注意。

# work_memの設定指針
`work_mem`は並び替え、ハッシュ結合などの一時操作で使用可能な上限メモリを設定する。各クエリ実行プロセス、かつ各操作ごとに上限が適用されるため、実際の消費量は`work_mem × 一時操作数 × パラレルワーカー数 × 同時セッション数`といった要素に依存する。最悪の場合には大きなメモリ消費に至る可能性がある。ただしこれは理論上の最大想定であり、実際のクエリ内容やタイミングにより変動するため、あくまで目安として捉えるべきである。

- デフォルト・単位: デフォルト4MB。`work_mem = '16MB'`などと明示指定。
- 適用単位: 各クエリ・各操作ごと。パラレルクエリ使用時は各ワーカーにも適用され、`max_parallel_workers_per_gather`と連動して消費量が増加する。
- ハッシュメモリ乗数: `hash_mem_multiplier`によりハッシュ操作のメモリ利用上限が制御される。
- 同時接続数とリスク: 多数の同時接続や複雑なクエリが多い場合、OOMのリスクがある。

# その他の関連パラメータ
- `effective_cache_size`: プランナが使用可能と想定するキャッシュ量（実メモリ消費には直接影響しない）が、インデックス利用判断に大きく影響する。
- `maintenance_work_mem`: VACUUMやインデックス作成時の作業メモリ。
- `temp_buffers`: 一時テーブル用メモリ領域。
- `max_connections`とPgBouncer: 同時接続制御でメモリ消費を抑制。
- `max_parallel_workers_per_gather`: パラレルクエリのワーカー数制御。
- `replication_slot_max_wal_size`: レプリケーションスロットが保持できる最大WALサイズ。
- `autovacuum_work_mem`: 自動バキュームプロセスが使用するメモリ。デフォルトは`maintenance_work_mem`の値。
- `logical_decoding_work_mem`: 論理デコード時のメモリ使用上限（PostgreSQL 13以降）。
- `wal_buffers`: WAL書き込み用のバッファ。自動設定されるが、高負荷時には調整が効果的。
- `temp_file_limit`: セッションが利用可能な一時ファイルのサイズ上限（MB単位）。
- `bgwriter_lru_maxpages`, `bgwriter_lru_multiplier`: バックグラウンドライタのバッファ書き出し量と頻度制御。
- `shared_memory_type`: `mmap`/`sysv`/`windows` のどの方式で共有メモリを確保するか。
- `huge_pages`: OSのHuge Pages（Transparent Huge Pages）使用有無。

# まとめ
PostgreSQLのメモリ管理は`shared_buffers`と`work_mem`を中心に、プロセスごとの消費や同時接続数、パラレルクエリ特性を加味した総合設計が不可欠である。設定変更は段階的に行い、事前検証・リスク評価・継続監視の3点を徹底することで、安定性と性能を両立した運用を実現できる。

# 参考
- [PostgreSQL公式ドキュメント: Runtime Configuration - Resource Consumption](https://www.postgresql.org/docs/current/runtime-config-resource.html)
- [Gihyo連載：詳解PostgreSQL第2回 - PostgreSQLの内部構造](https://gihyo.jp/dev/feature/01/dex_postgresql/0002)
- [Katsusandブログ：PostgreSQLのメモリ管理を考える](https://katsusand.dev/posts/postgresql-memory/)


