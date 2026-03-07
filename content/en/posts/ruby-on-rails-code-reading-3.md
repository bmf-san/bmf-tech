---
title: Ruby on Rails Code Reading Part 3 - Invoking WelcomeController
slug: ruby-on-rails-code-reading-3
date: 2024-09-04T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby on Rails
  - Ruby
translation_key: ruby-on-rails-code-reading-3
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
4. Run `rails generate controller Example`.
5. Write `binding.pry` in [railties/lib/rails/welcome_controller.rb#L9](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/welcome_controller.rb#L9).
```ruby
  def index
    binding.pry
  end
```

# Code Reading
Let's read the code to see how WelcomeController#index is invoked.

WelcomeController#index does not seem to be defined in config/routes.rb, but appears to be defined by default, likely due to the autoload mechanism.

[railties/lib/rails.rb#L33](https://github.com/rails/rails/blob/5385580ac82797167382ffcd79095a4bb973c666/railties/lib/rails.rb#L33)

By being autoloaded here, WelcomeController is set in the routing.

Here is the implementation of WelcomeController:
[railties/lib/rails/welcome_controller.rb#L5](https://github.com/rails/rails/blob/2b0ae167eee81d0d31b1d2f88c3f6c596c61ea8c/railties/lib/rails/welcome_controller.rb#L5)

I couldn't trace the code in detail, but it seems that by leveraging the autoload mechanism, routing can be resolved without explicitly registering it.