---
title: "Go製Git操作ツール『ggc』の紹介（2026年版）"
description: 'ggc v8の全機能完全解説。CLI/インタラクティブ分離アーキテクチャ、Fuzzyサーチエンジンの実装、Workflow Mode内部構造、カスタマイズ可能なエイリアス、クロスプラットフォーム対応のキーバインドプロファイル。'
slug: introducing-go-git-tool-ggc
date: 2025-06-15T00:00:00Z
lastmod: 2026-03-20
author: bmf-san
categories:
  - ツール
tags:
  - Golang
  - Git
  - CLI
  - TUI
translation_key: introducing-go-git-tool-ggc
---

# Go製Git操作ツール『ggc』の紹介（2026年版）

## ggcとは

[ggc](https://github.com/bmf-san/ggc)はGo製のGitワークフローツールだ。日常的なGitサブコマンドを一貫したインターフェースで提供し、コマンド名を暗記せずとも操作できるインタラクティブFuzzy検索TUIを搭載する。v8では Workflow Mode、プレースホルダー対応のカスタムエイリアス、レイヤー山積キーバインドプロファイルが新たに導入された。本記事ではその全てを解説する。

## CLI/インタラクティブモードのアーキテクチャ

ggcには2つの実行パスがある。直接コマンドを実行する**CLIパス**と、フルスクリーンTUIを起動する**インタラクティブパス**だ。`cmd/execute.go`の`Execute()`メソッドがどちらのパスを連るかを決定する唯一のエントリーポイントだ。

```go
func (c *Cmd) Execute(args []string) error {
    if len(args) == 0 {
        c.Interactive()
        return nil
    }

    cmdName, cmdArgs := args[0], args[1:]

    // エイリアスかチェック
    if c.configManager != nil && c.configManager.GetConfig().IsAlias(cmdName) {
        return c.executeAlias(cmdName, cmdArgs)
    }

    // 通常コマンド
    return c.Route(args)
}
```

引数がゼロの場合は`Interactive()`でTUIを起動する。エイリアス名を検出した場合は`executeAlias()`で展開し実行する。それ以外は`Route()`に貸す。この`Route()`関数はCLIパス、エイリアス展開、Workflow実行エンジンの3つすべてが共有する実装だ。

## インタラクティブモードの設計

インタラクティブTUIは1つのターミナル画面を共有する2つのサブモードに分かれている。`Ctrl+t`でトグルする。

### Search Mode — Fuzzy Search

Search Modeがデフォルトだ。全コマンドの一覧を表示し、入力に応じてリアルタイムにフィルタリングする。スコアリングは`internal/interactive/fuzzy.go`の`matchPattern()`関数が担当する。

```go
func matchPattern(textRunes, patternRunes []rune) (bool, matchMetadata) {
    meta := matchMetadata{firstIndex: -1, lastIndex: -1}
    textIdx := 0
    patternIdx := 0

    for textIdx < len(textRunes) && patternIdx < len(patternRunes) {
        if textRunes[textIdx] == patternRunes[patternIdx] {
            if meta.firstIndex == -1 {
                meta.firstIndex = textIdx
            }
            if meta.lastIndex != -1 {
                meta.gapScore += textIdx - meta.lastIndex - 1
            }
            meta.lastIndex = textIdx
            patternIdx++
        }
        textIdx++
    }

    if patternIdx != len(patternRunes) {
        return false, meta
    }
    return true, meta
}
```

アルゴリズムは古典的な**部分列マッチング**だ。パターンの全文字がテキスト内に順序通り現れればマッチと判定する（連続不要）。`matchMetadata`構造体は3つのシグナル—`firstIndex`（マッチ開始位置）、`gapScore`（文字間の距離合計）、`lastIndex`—を蓄積し、`matchScore`としてソートに利用する。致密で早索引のマッチが、散居した遅いマッチより高ቂ5ンク付けされる。

### Workflow Mode

`Ctrl+t`を押すと`internal/interactive/ui_mode.go`の`ToggleWorkflowView()`が呼び出される。

```go
func (ui *UI) ToggleWorkflowView() {
    if ui.state.IsWorkflowMode() {
        ui.enterSearchMode()
        return
    }
    ui.enterWorkflowMode()
}
```

Workflow Modeでは、検索リストの任意アイテムで`Tab`を押すことでggcコマンドのシーケンス（例：`add` → `commit` → `push`）を組み立てられる。`x`を押すと`WorkflowExecutor.Execute()`が順次実行する。

```go
func (we *WorkflowExecutor) Execute(workflow *Workflow) error {
    steps := workflow.GetSteps()
    if len(steps) == 0 {
        return fmt.Errorf("workflow is empty")
    }

    for i, step := range steps {
        resolvedArgs, canceled := resolveStepPlaceholders(we.ui, step)
        if canceled {
            return ErrWorkflowCanceled
        }

        parts := append([]string{step.Command}, resolvedArgs...)
        if err := we.router.Route(parts); err != nil {
            return fmt.Errorf("step %d/%d failed: %w", i+1, len(steps), err)
        }
    }
    return nil
}
```

各ステップの前に`resolveStepPlaceholders()`が引数内の`<name>`トークンをスキャンし、存在する場合はインタラクティブに入力を求める。事前にコミットメッセージをハードコードする必要はなく、実行時にプロンプトで單やかに入力できる。

## v8のその他の新機能

### コマンドエイリアス

エイリアスは`~/.ggcconfig.yaml`で定義する。シンプル形式とシーケンス形式の両方をサポートする。

```yaml
aliases:
  br: branch          # シンプル — ggc br list
  ci: commit

  quick:              # シーケンス
    - status
    - add .
    - commit

  deploy:             # プレースホルダー付きシーケンス
    - "branch checkout {0}"
    - "pull current"
    - "push {0}"
```

### キーバインドプロファイル

インタラクティブUIは4つのビルトインキーバインドプロファイル（`default`、`emacs`、`vi`、`readline`）をサポートする。レイヤー順に解決される—デフォルト → プロファイル → プラットフォーム → ユーザー設定。

```yaml
interactive:
  profile: emacs
  keybindings:
    move_up: "ctrl+p"
    move_down: "ctrl+n"
    toggle_workflow_view: "ctrl+t"
    add_to_workflow: "tab"
```

### クロスプラットフォーム対応

GoReleaserを介してLinux、macOS、Windows（amd64 / arm64 / 386）すべてのpresetバイナリを配布している。インタラクティブTUIの全機能が3プラットフォームで動作する。

## インストール

```bash
# Homebrew (macOS/Linux)
brew install bmf-san/tap/ggc

# Go install
go install github.com/bmf-san/ggc/v8@latest
```

## まとめ

ggc v8は、明確な設計思想に基づく柔軟なGitワークフローツールだ。CLIパスとインタラクティブパスは同じ`Route()`層を共有しており、挙動が一賫している。Fuzzyサーチスコアラーは依存性ゼロの純粋関数であり、Workflow実行エンジンはコマンドディスパッチロジックを重複定義せずに`Route()`を再利用する。

- **GitHub**: [bmf-san/ggc](https://github.com/bmf-san/ggc)
