---
title: RSpecのコードリーディング
slug: rspec-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - RSpec
  - Ruby
translation_key: rspec-code-reading
---


# 概要
RSpecのコードリーディングをする。

# 準備
1. RSpecのリポジトリをクローンする。
- https://github.com/rspec/rspec-core

# コードリーディング
1. RSpecの呼び出し
- [rspec/rspec-core/blob/main/exe/rspec#L4](https://github.com/rspec/rspec-core/blob/main/exe/rspec#L4)
  - エントリーポイント
- [lib/rspec/core/runner.rb#L43](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L43)
  - Runnerクラスのクラスメソッドであるinvokeを呼び出す
  - `disable_autorun`メソッドは、自動実行機能を無効にする
3. RSpecの実行
- [lib/rspec/core/runner.rb#L64](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L64)
  - Runnerのクラスメソッドであるrunを呼び出す
  - `trap_interrupt`メソッドは、Ctrl+Cなどの中断時の処理を行う
  - optionsにrunnerがあれば`call`、なければ`new.run`
- [lib/rspec/core/runner.rb#L85](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L85)
  - テスト実行が早期に終了した場合はレポーティングの処理を呼び出す
  - 早期終了しない場合は、`run_specs`メソッドを呼び出す
- [lib/rspec/core/runner.rb#L113](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L113)
  - テストを順次実行して結果をレポーティングし、終了コードを返す
