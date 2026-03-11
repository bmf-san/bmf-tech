---
title: Ruby on Rails Code Reading Part 2 - Rails Request Processing
slug: ruby-on-rails-code-reading-2
date: 2024-09-02T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
  - Ruby on Rails
description: Recording the code reading work of Ruby on Rails.
translation_key: ruby-on-rails-code-reading-2
---



# Overview
Recording the code reading work of Ruby on Rails.

# Preparation
1. Create a new project with `rails new RailsCodeReading`.
2. Add the following to the Gemfile
```ruby
gem 'pg'
gem 'pry-rails'
gem 'pry-doc'
gem 'pry-byebug'
gem 'byebug'
```
3. Run `bundle config set path '.bundle'` and then execute `bundle install`.
4. `rails generate controller Example`
5. Implement the index method in the controller
```ruby
  def index
    binding.pry
    render json: { message: 'Hello World!' }
  end
```
6. Set up routing
```ruby
  get "example" => "example#index"
```

# Code Reading
Follow the flow of how Rails processes requests.

Access `http://127.0.0.1:3000/example` and use `pry-backtrace` in the console to view the stack trace.

Due to the large output, it's not possible to follow everything, so let's look at it in a condensed form.

1. Puma accepts the request and calls the Rails application
  - [puma/puma - lib/puma/request.rb#L99](https://github.com/puma/puma/blob/796d8c6af139a22e29e57e156c05c66ed3082634/lib/puma/request.rb#L99)
- [railties/lib/rails/engine.rb#L536](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/engine.rb#L536)
  - Definition of Rack API
- [rack/rack - lib/rack/sendfile.rb#L113](https://github.com/rack/rack/blob/main/lib/rack/sendfile.rb#L113)
  - Rack application is executed
2. Resolve routing based on request information
  - [rails/rails - actionpack/lib/action_dispatch/journey/router.rb#L126](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_dispatch/journey/router.rb#L126)
  - [rails/rails - actionpack/lib/action_dispatch/routing/route_set.rb#L66](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_dispatch/routing/route_set.rb#L66)
    - Call the process of the controller matched by routing
3. Call the index of the controller
  - [rails/rails - actionpack/lib/action_controller/metal/basic_implicit_render.rb#L7](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/actionpack/lib/action_controller/metal/basic_implicit_render.rb#L7)

It was more complex than expected, and since I'm not familiar with Ruby, I could only grasp the atmosphere to some extent...

# References
- [speakerdeck.com - Ruby on Rails Hacking Guide](https://speakerdeck.com/a_matsuda/ruby-on-rails-hacking-guide)
- [magazine.rubyist.net - RubyOnRails を使ってみる 【第 4 回】 ActionPack](https://magazine.rubyist.net/articles/0008/0008-RubyOnRails.html)
