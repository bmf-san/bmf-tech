---
title: 孤児プロセスとゾンビプロセスの違い
description: 'ゾンビプロセス（defunct）と孤児プロセスの違いを解説。PIDの枯渇リスク・psコマンドによる確認手順・killの方法、親プロセスとwaitシステムコールの仕組みをDockerコンテナの文脈も交えて紹介します。'
slug: orphan-process-vs-zombie-process
date: 2020-08-24T00:00:00Z
author: bmf-san
categories:
  - コンピューターアーキテクチャ
tags:
  - UNIX
  - プロセス
translation_key: orphan-process-vs-zombie-process
---


# 概要
Dockerを触っていたらorphan（孤児の意）というプロセスの存在を知ったのでゾンビプロセスとの違いを調べてみた。

# ゾンビプロセスとは

ゾンビプロセスとは、**処理が終了した子プロセスがプロセステーブルに残り、親プロセスによる`wait()`システムコールを待っている状態**のプロセスである。

Linuxでは、子プロセスが終了しても、親プロセスが`wait()`（または`waitpid()`）を呼び出して終了ステータスを受け取るまで、プロセスの情報はプロセステーブルに保持される。この状態を「ゾンビ状態（defunct）」と呼ぶ。

**特徴**：
- CPUやメモリなどのシステムリソースは使用しない
- PID（プロセスID）だけが保持される
- ゾンビプロセスが大量に増えると、使用可能なPIDが枯渇し、新しいプロセスを起動できなくなる

**ゾンビプロセスの確認方法**：
```bash
# ps auxでstatがZのもの、またはコマンド末尾がdefunctのもの
ps aux | grep defunct

# もしくは
ps -ef | grep defunct
```

**ゾンビプロセスのkill**：

ゾンビプロセス自体は終了済みのため、`kill`コマンドは効かない。親プロセスが`wait()`を呼び出すか、**親プロセス自体をkillする**ことで解消される。親プロセスがkillされると、孤児プロセスがinitに引き取られてwaitが行われ、プロセステーブルから削除される。

```bash
# 親プロセスのPIDを確認（PPIDの欄）
ps -ef | grep defunct

# 親プロセスをkill
kill -9 <親プロセスのPID>
```

# 孤児プロセス

孤児プロセスとは、**親プロセスが`wait()`を呼ばずに終了してしまったことで、親を失ったプロセス**のことである。

Linuxでは、親プロセスが先に終了した場合、孤児プロセスはinitプロセス（PID: 1）またはサブリーパー（Dockerなら`tini`など）に引き取られる。initプロセスが里親となり、定期的に`wait()`を呼んでゾンビ状態を解消する。

**特徴**：
- 動作は継続する（ゾンビとは異なりリソースを使用する）
- PIDが1のinitに再ペアレントされる
- Dockerコンテナ内では、PID 1が`init`でない場合（`sh`や`bash`が起動している場合など）、適切にwaitされずゾンビが蓄積しやすい

**孤児プロセスの確認方法**：
```bash
# PPIDが1のプロセスのうちrootでないものを表示
ps -elf | head -1; ps -elf | awk '{if ($5 == 1 && $3 != "root") {print $0}}' | head
```

**孤児プロセスのkill**：
```bash
kill <PID>
```

# ゾンビプロセスと孤児プロセスの対比

| 項目 | ゾンビプロセス | 孤児プロセス |
|-----|--------------|------------|
| 状態 | 終了済み（defunctとして残存） | 実行中 |
| リソース消費 | なし（PIDのみ保持） | あり |
| 親プロセス | 存在するがwait()待ち | 存在しない（initが引き取る） |
| 問題のリスク | PID枯渇 | リソース消費続行 |
| 解消方法 | 親がwait()するか、親をkill | killコマンドで終了 |

# Dockerにおける注意点

Dockerコンテナ内のPID 1プロセスが通常のアプリ（`node`、`python`等）の場合、signal処理や`wait()`をデフォルトでは行わないため、ゾンビプロセスが蓄積しやすい。`docker run`に`--init`フラグを付けるか、`tini`などのinitプロセスマネージャーを使用することが推奨される。

```yaml
# docker-compose.ymlでの設定
services:
  app:
    image: myapp
    init: true  # tiniをPID 1として使用
```

# 参考
- [Qiita - 【unix】ゾンビプロセス・孤児プロセスって何。](https://qiita.com/ninoko1995/items/582106e8507163b2c50b)
- [Hibariya - プロセスをforkするときのこと](https://note.hibariya.org/articles/20120326/a0.html)
- [tutorialspoint.com - zombie-and-orphan-processes-in-linux](https://www.tutorialspoint.com/zombie-and-orphan-processes-in-linux)
- [日経XTECH - ゾンビ・プロセス](https://xtech.nikkei.com/it/article/Keyword/20070727/278487/)
- [geekride.com - Orphan Process](http://web.archive.org/web/20211025203533/http://www.geekride.com/orphan-zombie-process/)
- [makiuchi-d.gihub.io - Goで子プロセスを確実にKillする方法](http://makiuchi-d.github.io/2020/05/10/go-kill-child-process.ja.html)
  - Goで孤児プロセスを生むコードの再現がある

