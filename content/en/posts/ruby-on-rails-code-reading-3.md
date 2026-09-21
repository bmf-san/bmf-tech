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
description: "The third installment of my code reading notes for Ruby on Rails. Using a project created with rails new, it traces how a request flows through routing until the action of the WelcomeController is invoked, following the request handling step by step."
translation_key: ruby-on-rails-code-reading-3
draft: false
---


# Overview
Documenting the code reading process for Ruby on Rails.

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
3. Run `bundle config set path '.bundle'` and then execute `bundle install`.
4. `rails generate controller Example`
5. Add `binding.pry` to [railties/lib/rails/welcome_controller.rb#L9](https://github.com/rails/rails/blob/8bac99ad7a403ef52a5c97e7afa73c7bbcc67110/railties/lib/rails/welcome_controller.rb#L9):
```ruby
  def index
    binding.pry
  end
```

# Code Reading
Let's explore how Rails invokes WelcomeController#index.

config/routes.rb does not define WelcomeController#index; Rails seems to define it by default, likely through the autoload mechanism.

[railties/lib/rails.rb#L33](https://github.com/rails/rails/blob/5385580ac82797167382ffcd79095a4bb973c666/railties/lib/rails.rb#L33)

Autoloading it here registers WelcomeController in the routing.

Here is the implementation of WelcomeController.
[railties/lib/rails/welcome_controller.rb#L5](https://github.com/rails/rails/blob/2b0ae167eee81d0d31b1d2f88c3f6c596c61ea8c/railties/lib/rails/welcome_controller.rb#L5)

Although I couldn't follow the code in detail, Rails appears to leverage the autoload mechanism to resolve routing without explicitly registering it.
