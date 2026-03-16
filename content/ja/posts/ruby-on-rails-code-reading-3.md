---
title: Ruby on Railsのコードリーディング第3回-WelcomeControllerの呼び出し
description: Ruby on Railsのコードリーディング第3回-WelcomeControllerの呼び出し
slug: ruby-on-rails-code-reading-3
date: 2024-09-04T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby on Rails
  - Ruby
translation_key: ruby-on-rails-code-reading-3
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
4. `rails generate controller Example`
5. [railties/lib/rails/welcome_controller.rb#L9](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/welcome_controller.rb#L9)に`binding.pry`を記述する。
```ruby
  def index
    binding.pry
  end
```

# コードリーディング
WelcomeController#indexがどのように呼ばれるのかコードリーディングしてみる。

WelcomeController#indexはconfig/routes.rbには定義されておらず、デフォルトで定義されているように見えるが、autoloadの仕組みによるものだと思われる。

[railties/lib/rails.rb#L33](https://github.com/rails/rails/blob/5385580ac82797167382ffcd79095a4bb973c666/railties/lib/rails.rb#L33)

ここでautoloadされることによって、WelcomeControllerがルーティングにセットされる。

WelcomeControllerの実装はこれ。
[railties/lib/rails/welcome_controller.rb#L5](https://github.com/rails/rails/blob/2b0ae167eee81d0d31b1d2f88c3f6c596c61ea8c/railties/lib/rails/welcome_controller.rb#L5)

詳細にコードを追えなかったが、autoloadの仕組みに乗っかることでルーティングの登録を明示的にしなくてもルーティングが解決されるように見える。
