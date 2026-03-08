---
title: PostgreSQL インストール時の 'configure' エラーに関するトラブルシューティング
slug: postgresql-install-configure-error-troubleshooting
date: 2025-03-07T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - PostgreSQL
  - Ubuntu
  - CentOS
translation_key: postgresql-install-configure-error-troubleshooting
---


# PostgreSQL インストール時の 'configure' エラーに関するトラブルシューティング
## 1. はじめに

PostgreSQL のソースコードをコンパイルする際に、`make` を実行した際に `You need to run the 'configure' program first.` というエラーが発生した。

トラブルシューティングをメモしておく。

## 2. エラーの詳細

エラーメッセージ:

```sh
root@7a46fccdd51a:/postgresql-14.7# make -j$(nproc)
You need to run the 'configure' program first. See the file
'INSTALL' for installation instructions.
Makefile:19: recipe for target 'all' failed
make: *** [all] Error 1
```

このエラーは、PostgreSQL のソースコードディレクトリに `configure` スクリプトが実行されていないために発生する。

## 3. 解決策

### Step 1: ********************`configure`******************** を実行する

`configure` スクリプトは PostgreSQL のビルド環境をセットアップするために必要である。まず、PostgreSQL のソースディレクトリに移動し、以下のコマンドを実行する。

```sh
./configure
```

もし `configure` 実行時に追加のオプションが必要な場合は、以下のように実行する。

```sh
./configure --prefix=/usr/local/pgsql --enable-debug --enable-cassert
```

また、利用可能なオプションを確認する場合は、次のコマンドを実行する。

```sh
./configure --help
```

### Step 2: ********************`make`******************** を再実行

`configure` が正常に完了したら、以下のコマンドでビルドを実行する。

```sh
make -j$(nproc)
```

これにより、利用可能な CPU コア数を活用して高速にコンパイルが行われる。

### Step 3: ********************`make install`******************** を実行（必要な場合）

`make` が成功したら、次に PostgreSQL をインストールする。

```sh
make install
```

これにより、指定したプレフィックスディレクトリ（デフォルトは `/usr/local/pgsql`）に PostgreSQL がインストールされる。

## 4. ********************`configure`******************** 実行時の追加の問題と対処法

### 問題 1: `configure: command not found`

#### 原因

- `configure` スクリプトが存在しない
- 必要なビルドツールが不足している

#### 解決策

まず、`configure` スクリプトが存在するか確認する。

```sh
ls -l configure
```

もし `configure` が見つからない場合、`autoconf` を実行して再生成する。

```sh
autoreconf -fi
./configure
```

それでもエラーが発生する場合は、必要なパッケージがインストールされているか確認する。

#### Ubuntu/Debian の場合

```sh
apt update && apt install -y build-essential libreadline-dev zlib1g-dev flex bison
```

#### CentOS/RHEL の場合

```sh
yum groupinstall -y "Development Tools"
yum install -y readline-devel zlib-devel flex bison
```

### 問題 2: ********************`configure`******************** 実行時にライブラリのエラーが発生する

エラーメッセージ例:

```sh
configure: error: readline library not found
```

#### 解決策

不足しているライブラリをインストールする。

```sh
apt install -y libreadline-dev
```

または

```sh
yum install -y readline-devel
```

## 5. まとめ

PostgreSQL のコンパイル時に `make` が失敗する原因には、`configure` が未実行であったり、必要なライブラリが足りていないことに起因するものがある。本記事で紹介した手順を実行することで、以下の問題を解決できるかもしれない。

1. `configure` を実行し、環境を適切に設定する
2. `make` を実行して PostgreSQL をビルドする
3. 必要に応じて `make install` を実行してインストールする
4. `configure` の実行時に発生する問題（ライブラリ不足など）を適切に解決する


