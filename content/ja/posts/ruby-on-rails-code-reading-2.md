---
title: Ruby on Railsのコードリーディング第2回-Railsのリクエスト処理
slug: ruby-on-rails-code-reading-2
date: 2024-09-02T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
  - Ruby on Rails
translation_key: ruby-on-rails-code-reading-2
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
5. コントローラーにindexメソッドを実装
```ruby
  def index
    binding.pry
    render json: { message: 'Hello World!' }
  end
```
6. ルーティングを設定
```ruby
  get "example" => "example#index"
```

# コードリーディング
Railsがリクエストを処理する流れを追う。

`http://127.0.0.1:3000/example`にアクセスして、コンソールで`pry-backtrace`をすると、スタックトレースを見ることができる。

大量に出力されるため、すべてを追いきれないので端折ってみていく。

1. pumaがリクエストを受け付けて、Railsアプリケーションを呼び出す
  - [puma/puma - lib/puma/request.rb#L99](https://github.com/puma/puma/blob/796d8c6af139a22e29e57e156c05c66ed3082634/lib/puma/request.rb#L99)
- [railties/lib/rails/engine.rb#L536](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/engine.rb#L536)
  - Rack APIの定義
- [rack/rack - lib/rack/sendfile.rb#L113](https://github.com/rack/rack/blob/main/lib/rack/sendfile.rb#L113)
  - Rackアプリケーションが実行される
2. リクエストの情報を元に、ルーティングを解決する
  - [rails/rails - actionpack/lib/action_dispatch/journey/router.rb#L126](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_dispatch/journey/router.rb#L126)
  - [rails/rails - actionpack/lib/action_dispatch/routing/route_set.rb#L66](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_dispatch/routing/route_set.rb#L66)
    - ルーティングでマッチしたcontrollerの処理を呼び出す
3. コントローラーのindexを呼び出す
  - [rails/rails - actionpack/lib/action_controller/metal/basic_implicit_render.rb#L7](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_controller/metal/basic_implicit_render.rb#L7)

想像していたよりも複雑で、Rubyに慣れていないのもあり、雰囲気程度しか読み切れなかった。。。

# 参考
- [speakerdeck.com - Ruby on Rails Hacking Guide](https://speakerdeck.com/a_matsuda/ruby-on-rails-hacking-guide)
- [magazine.rubyist.net - RubyOnRails を使ってみる 【第 4 回】 ActionPack](https://magazine.rubyist.net/articles/0008/0008-RubyOnRails.html)



