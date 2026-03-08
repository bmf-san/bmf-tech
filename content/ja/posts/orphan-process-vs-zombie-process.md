---
title: 孤児プロセスとゾンビプロセスの違い
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
- 処理が終了した子プロセス
- プロセステーブルに残った状態で親プロセスのwaitを待つ
- システムリソースは使用しないが、PIDは保持される
- ゾンビプロセスが大量に増えると、使用可能なPIDが減り、他のプロセスを起動できなくなる
- ゾンビプロセスの確認方法
  - `ps aux`でstatがZ、末尾がdefunctのものがゾンビプロセス
  - `ps -ef | grep defunct`でゾンビプロセスだけ出力
- ゾンビプロセスのkill
  - 親プロセスをkillする

# 孤児プロセス
- 親プロセスがwaitせずに終了してしまったプロセス
- initプロセス（PIDが1）の子プロセスとなり、initプロセスが親（≒里親）となる
- 孤児プロセスの確認方法
  - `ps -elf | head -1; ps -elf | awk '{if ($5 == 1 && $3 != "root") {print $0}}' | head`
- 孤児プロセスのkill
  - killコマンドでOK 

# 参考
- [Qiita - 【unix】ゾンビプロセス・孤児プロセスって何。](https://qiita.com/ninoko1995/items/582106e8507163b2c50b)
- [Hibariya - プロセスをforkするときのこと](https://note.hibariya.org/articles/20120326/a0.html)
- [tutorialspoint.com - zombie-and-orphan-processes-in-linux](https://www.tutorialspoint.com/zombie-and-orphan-processes-in-linux)
- [日経XTECH - ゾンビ・プロセス](https://xtech.nikkei.com/it/article/Keyword/20070727/278487/)
- [geekride.com - Orphan Process](http://www.geekride.com/orphan-zombie-process/)
- [makiuchi-d.gihub.io - Goで子プロセスを確実にKillする方法](http://makiuchi-d.github.io/2020/05/10/go-kill-child-process.ja.html)
  - Goで孤児プロセスを生むコードの再現がある

