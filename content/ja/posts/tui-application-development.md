---
title: x／termで作るTUIアプリケーション開発
slug: tui-application-development
date: 2025-07-16T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - TUI
  - CLI
  - ゲーム
translation_key: tui-application-development
---


# はじめに

最近、Goのx/termパッケージを使ってターミナルベースのタイピングゲームを作ってみた。この記事では、x/termパッケージの特徴や、TUIアプリケーション開発で気づいたことを共有していく。

x/termを使った実践的なTUIアプリケーションとして、[ggc](https://github.com/bmf-san/ggc)というgitのクライアントツールを開発しているので、よければStarを押してほしい。

# x/termパッケージとは

`x/term`は、Goの実験的なパッケージの一つで、ターミナル操作のための低レベルな機能を提供している。以前は`golang.org/x/crypto/ssh/terminal`だったが、現在は`golang.org/x/term`として独立したパッケージとなっている。

## 主な機能

1. ターミナルのサイズ取得
```go
width, height, err := term.GetSize(int(os.Stdout.Fd()))
```

2. ローレベルなキー入力の取得
```go
oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
if err != nil {
    log.Fatal(err)
}
defer term.Restore(int(os.Stdin.Fd()), oldState)
```

3. エコーの制御
4. ターミナルモードの制御

これらの機能を使えば、カーソル位置の制御やキー入力の即時検出など、インタラクティブなTUIアプリケーションが作れる。

# タイピングゲームの実装

## 1. 基本設計

タイピングゲームの核となる機能はこんな感じである：

- ランダムな英文の表示
- キー入力の即時検出
- 正確性の測定
- ミス回数のカウント

## 2. ターミナル状態の管理

x/termを使う上で最も気をつけたいのはターミナルの状態管理である：

```go
// ターミナルを生モードに設定
oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
if err != nil {
    log.Fatal(err)
}

// プログラム終了時に元の状態に戻す
defer term.Restore(int(os.Stdin.Fd()), oldState)

// シグナルハンドリング
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
go func() {
    <-sigCh
    term.Restore(int(os.Stdin.Fd()), oldState)
    os.Exit(0)
}()
```

## 3. 画面表示の制御

ANSIエスケープシーケンスを使って、画面表示を制御している：

```go
// カーソルを非表示
fmt.Print("\033[?25l")
defer fmt.Print("\033[?25h") // 終了時に表示を戻す

// 画面クリア
fmt.Print("\033[2J")

// カーソル位置を移動（x=10, y=5の位置へ）
fmt.Printf("\033[%d;%dH", 5, 10)
```

## 4. パフォーマンス最適化

表示の安定性を確保するためのポイントをいくつか紹介する：

- バッファリングの制御
- 画面更新の最適化
- goroutineを使った非同期処理

# 実装のポイント

## 1. エラー処理とリカバリ

```go
// パニック時のリカバリ処理
defer func() {
    if r := recover(); r != nil {
        term.Restore(int(os.Stdin.Fd()), oldState)
        fmt.Printf("Recovered from panic: %v\n", r)
    }
}()
```

## 2. クロスプラットフォーム対応

```go
var clear string
if runtime.GOOS == "windows" {
    clear = "cls"
} else {
    clear = "clear"
}
cmd := exec.Command(clear)
cmd.Stdout = os.Stdout
cmd.Run()
```

# まとめ

x/termを使ったTUIアプリケーション開発は、低レベルな制御が必要だが、その分自由度が高い。今回作ったタイピングゲームの実装例は基本的なパターンを示すものだが、これを応用すれば様々なTUIアプリケーションが作れる。

# 参考リンク

- [golang.org/x/term](https://pkg.go.dev/golang.org/x/term)
