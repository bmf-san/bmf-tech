---
title: Ruby on Rails Code Reading Part 1 - Starting Rails
slug: ruby-on-rails-code-reading-1
date: 2024-09-02T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
  - Ruby on Rails
translation_key: ruby-on-rails-code-reading-1
---

# Overview
Record the code reading work for Ruby on Rails.

# Preparation
1. Create a new project with `rails new RailsCodeReading`.
2. Add the following to the Gemfile:
```ruby
gem 'pg'
gem 'pry-rails'
gem 'pry-doc'
gem 'pry-byebug'
gem 'byebug'
```
3. Run `bundle config set path '.bundle'`, then execute `bundle install`.

# Code Reading
Read the code up to the server execution part until Rails starts after executing the `rails server` command.

1. [rails/rails - railties/lib/rails/commands/server/server_command.rb#L132](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/commands/server/server_command.rb#L132)
  - Implementation of the server command
  - [rails/rails - railties/lib/rails/commands/server/server_command.rb#L32](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/commands/server/server_command.rb#L32)
    - Server startup processing
    - `super()` calls `Rackup::Server`'s `start`
2. [rails/rails - railties/lib/rails/rackup/server.rb#L8](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/rackup/server.rb#L8)
  - Requires `rackup/server`
3. [rack/rackup - lib/rackup/server.rb#L300](https://github.com/rack/rackup/blob/eaea24a3d64a1b117df943a9d06779e659bb61af/lib/rackup/server.rb#L300)
  - [rack/rackup - lib/rackup/server.rb#L341](https://github.com/rack/rackup/blob/eaea24a3d64a1b117df943a9d06779e659bb61af/lib/rackup/server.rb#L341)
    - Server startup processing is called with `server.run`
4. [puma/puma - lib/rack/handler/puma.rb#L67](https://github.com/puma/puma/blob/9ee922d28e1fffd02c1d5480a9e13376f92f46a3/lib/rack/handler/puma.rb#L67)
  - The server called by `server.run` depends on the server used by the application.
  - If using Puma, the `run` here will be called.

# References
- [railsguides.jp - Rails Initialization Process](https://railsguides.jp/initialization.html)