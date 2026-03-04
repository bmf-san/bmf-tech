---
title: "Ruby on Railsのコードリーディング第1回-Railsの起動"
slug: "ruby-on-rails-code-reading-1"
date: 2024-09-02
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Ruby"
  - "Ruby on Rails"
draft: false
---

# 概要
Ruby on Railsのコードリーディング作業を記録する。

# 準備
1. `rails new RailsCodeReading`で新規プロジェクトを作成する。
2. Gemfileに以下を追加
```ruby
gem 'pg'
gem 'pry-rails'
gem 'pry-doc'
gem 'pry-byebug'
gem 'byebug'
```
3. `bundle config set path '.bundle'`を実行してから、`bundle install`を実行する。

# コードリーディング
`rails server`のコマンド実行後、Railsが起動するまでのサーバーの実行部分までコードを読む。

1. [rails/rails - railties/lib/rails/commands/server/server_command.rb#L132](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/commands/server/server_command.rb#L132)
  - serverコマンドの実体
  - [rails/rails - railties/lib/rails/commands/server/server_command.rb#L32](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/commands/server/server_command.rb#L32)
    - サーバー起動の処理
    - `super()`で、`Rackup::Server`の`start`が呼び出される
2. [rails/rails - railties/lib/rails/rackup/server.rb#L8](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/rackup/server.rb#L8)
  - `rackup/server`をrequireしている
3. [rack/rackup - lib/rackup/server.rb#L300](https://github.com/rack/rackup/blob/eaea24a3d64a1b117df943a9d06779e659bb61af/lib/rackup/server.rb#L300)
  - [rack/rackup - lib/rackup/server.rb#L341](https://github.com/rack/rackup/blob/eaea24a3d64a1b117df943a9d06779e659bb61af/lib/rackup/server.rb#L341)
    - `server.run`でサーバー起動処理が呼び出される
4. [puma/puma - lib/rack/handler/puma.rb#L67](https://github.com/puma/puma/blob/9ee922d28e1fffd02c1d5480a9e13376f92f46a3/lib/rack/handler/puma.rb#L67)
  - `server.run`で呼び出されるサーバーはアプリケーションが利用するサーバーに依る
  - pumaを使っている場合はここの`run`が呼び出される

# 参考
- [railsguides.jp - Rails の初期化プロセス](https://railsguides.jp/initialization.html)
