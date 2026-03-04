---
title: "Go製Git操作ツール「ggc」の紹介"
slug: "go-git-ggc"
date: 2025-06-15
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "Git"
  - "CLI"
  - "TUI"
draft: false
---

# Go製Git操作ツール「ggc」の紹介

## ggcとは何か

[ggc](https://github.com/bmf-san/ggc)は、Go言語で実装されたGit操作支援ツールである。「覚えやすく、使いやすく、作業効率の向上を図る」ことを目的としており、日常的なGit操作をより快適にすることを意図している。

既存のGitクライアントツールには、機能が豊富すぎて学習コストが高いもの、あるいはシンプルすぎて実用に耐えないものが存在する。ggcはこのギャップを埋めるべく、**日常的に使用する機能に特化し、シンプルかつ記憶しやすいコマンド体系**を提供する。

### 特徴

1. **二重インターフェース**：CLIによる高速操作、直感的操作
2. **複合コマンド**：複数のGit操作を一つのコマンドで実行可能
3. **インクリメンタルサーチ**：コマンドを暗記せずに選択可能

### 利便性の比較

| 通常のGit操作                                         | ggcによる操作                     |
| ------------------------------------------------ | ---------------------------- |
| `git add .` → `git commit -m "..."` → `git push` | `ggc add-commit-push`        |
| `git branch` → `git checkout ブランチ名`              | `ggc branch checkout`（対話的選択） |
| `git stash` → `git pull` → `git stash pop`       | `ggc stash-pull-pop`         |

上記のように、一般的な操作を1コマンドで簡潔に実行可能である。

## 主な機能

* **二重インターフェース**：引数ありでコマンドを直接実行、引数なしでインタラクティブモードの起動
* **対話的操作**：ブランチ・ファイル選択やコミットメッセージ入力などに対応
* **豊富なコマンド群**：Gitの基本操作を網羅
* **複合コマンド**：`add-commit-push`、`stash-pull-pop`等を提供
* **軽量設計**：Go標準ライブラリと `golang.org/x/term` のみを使用
* **動作環境**：macOS（Apple Silicon/Intel）にて確認済み

## 使用例

```bash
# 最新の状態に更新
ggc pull current

# 新しいブランチで作業開始（対話的に選択）
ggc branch checkout
```

```bash
# 一括で変更をプッシュ
ggc add-commit-push
```

```bash
# 安全にマージ
ggc stash-pull-pop
```

## インストール手順

### `go install`による導入

最も簡単な導入方法は以下の通りである：

```sh
go install github.com/bmf-san/ggc@latest
```

必要に応じてPATHを設定する：

```sh
export PATH=$PATH:$(go env GOBIN)
```

### ソースコードからのビルド

```sh
git clone https://github.com/bmf-san/ggc
cd ggc
make build
```

ビルド後、生成されたバイナリをPATHの通ったディレクトリに配置する。

## 使い方

### CLI とインタラクティブモードの切り替え

引数の有無により自動的にCLIまたはインタラクティブモードが起動される。

```sh
# CLI（コマンドを直接指定）
ggc branch current

# インタラクティブモード
ggc
```

1つのバイナリで両モードをサポートしており、用途に応じた柔軟な操作が可能である。

### インタラクティブモードでのコマンド選択

`ggc`を引数なしで実行すると、インクリメンタルサーチによるコマンド選択画面が表示される。

```sh
ggc
```

表示例：

```
コマンドを選択してください（インクリメンタルサーチ：文字入力で絞り込み、ctrl+n:下移動, ctrl+p:上移動, enter:実行, ctrl+c:終了）
Search: branch

> branch current
  branch checkout
  branch checkout-remote
  branch delete
  branch delete-merged
```

操作手順：

* 文字を入力することで候補が絞り込まれる
* `Ctrl+n`/`Ctrl+p`で上下移動
* `Enter`で実行
* 引数が必要な場合はプロンプト表示
* 実行後、結果確認後に選択画面へ復帰

コマンドを記憶する必要はなく、入力に応じて候補が表示されるため、直感的な操作が可能である。

### 代表的なコマンド

| ggcコマンド                      | 実際のgitコマンド                                                          | 説明                     |
| ---------------------------- | ------------------------------------------------------------------- | ---------------------- |
| `ggc add <file>`             | `git add <file>`                                                    | ファイルをステージング            |
| `ggc add .`                  | `git add .`                                                         | 全ファイルをステージング           |
| `ggc add -p`                 | `git add -p`                                                        | 対話的ステージング              |
| `ggc branch current`         | `git rev-parse --abbrev-ref HEAD`                                   | 現在のブランチ名取得             |
| `ggc branch checkout`        | `git branch ... → git checkout <選択>`                                | 対話的ブランチ切り替え            |
| `ggc branch checkout-remote` | `git branch -r ... → git checkout -b <n> --track <remote>/<branch>` | リモートブランチから新規作成・切り替え    |
| `ggc branch delete`          | `git branch ... → git branch -d <選択>`                               | 対話的にローカルブランチ削除         |
| `ggc push current`           | `git push origin <branch>`                                          | 現在のブランチをプッシュ           |
| `ggc pull current`           | `git pull origin <branch>`                                          | 現在のブランチをプル             |
| `ggc log simple`             | `git log --oneline`                                                 | シンプルなログ表示              |
| `ggc commit <message>`       | `git commit -m <message>`                                           | コミット作成                 |
| `ggc fetch --prune`          | `git fetch --prune`                                                 | 古いリモート追跡ブランチを削除しつつフェッチ |
| `ggc clean files`            | `git clean -f`                                                      | ファイルのクリーンアップ           |
| `ggc remote add <n> <url>`   | `git remote add <n> <url>`                                          | リモート追加                 |
| `ggc stash`                  | `git stash`                                                         | 作業内容を一時退避              |
| `ggc rebase interactive`     | `git rebase -i`                                                     | 対話的リベース                |

### 複合コマンドの例

| ggcコマンド                       | 実行されるGit操作                             | 説明                      |
| ----------------------------- | -------------------------------------- | ----------------------- |
| `ggc add-commit-push`         | `git add . → git commit → git push`    | ステージ → コミット → プッシュを一括実行 |
| `ggc commit-push-interactive` | 対話的ステージ → コミット → プッシュ                  |                         |
| `ggc pull-rebase-push`        | `git pull → git rebase → git push`     | プル → リベース → プッシュを一括実行   |
| `ggc stash-pull-pop`          | `git stash → git pull → git stash pop` | 一時退避 → プル → 復元の一括操作     |

## 補完スクリプト

BashおよびZsh向けの補完スクリプトが同梱されている。

### 設定方法

```sh
# bashの場合
source /path/to/ggc/tools/completions/ggc.bash

# zshの場合（同一スクリプトを使用可能）
source /path/to/ggc/tools/completions/ggc.bash
```

これを`.bashrc`や`.zshrc`に追記することで、ターミナル起動時に補完が有効化される。

## まとめ

* コマンドを暗記せずに直感的操作が可能
* 定型作業を1コマンドで実行できる
* ブランチやファイルの選択も対話的に対応
* 複合コマンドによる作業効率の向上が期待できる

### 関連リンク

* **GitHubリポジトリ**：[https://github.com/bmf-san/ggc](https://github.com/bmf-san/ggc)
* **Issue・機能要望・バグ報告など**：[https://github.com/bmf-san/ggc/issues](https://github.com/bmf-san/ggc/issues)
