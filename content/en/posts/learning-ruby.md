---
title: Learning Ruby
description: 'Master Ruby fundamentals using official documentation, classic books, object-oriented design patterns and practical implementation practice.'
slug: learning-ruby
date: 2024-05-16T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: learning-ruby
---



# Overview
This post discusses what someone who has experience with PHP and Go did to learn Ruby.

# Approach
## Official Documentation
First and foremost, start with the official documentation.

- [What is Ruby](https://www.ruby-lang.org/ja/about/)
  - Concisely describes the features of Ruby.
  - For specifications you want to investigate deeply, this is a good reference.
- [Ruby from Other Languages](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/)
  - Describes the unique parts of Ruby's language specifications. It summarizes the points to learn that are not present in other languages.
  - For specifications you want to investigate deeply, this is also a good reference.
- [From PHP to Ruby](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/to-ruby-from-php/)
  - There was no Go version. As a former PHPer, I referred to the PHP page. The content is quite straightforward.
- [Ruby in 20 Minutes](https://www.ruby-lang.org/ja/documentation/quickstart/)
  - Perfect for a light practice.
- [Object-Oriented Scripting Language Ruby Reference Manual (Ruby 3.3 Reference Manual)](https://docs.ruby-lang.org/ja/3.3/doc/index.html)
  - A reference manual by volunteers. Much appreciated.

## Books
Picked up some classic ones.

- [Fun with Ruby 6th Edition](https://amzn.to/3JZamMX)
- [Ruby for Professionals](https://amzn.to/4dG4ciA)
  - cf. [Ruby for Professionals](https://bmf-tech.com/posts/%e3%83%97%e3%83%ad%e3%82%92%e7%9b%ae%e6%8c%87%e3%81%99%e4%ba%ba%e3%81%ae%e3%81%9f%e3%82%81%e3%81%aeRuby%e5%85%a5%e9%96%80)
- [Perfect Ruby](https://amzn.to/3K0wLcR)
  - cf. [Perfect Ruby](https://bmf-tech.com/posts/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88Ruby)
- [Perfect Ruby on Rails](https://amzn.to/3yfO0nL)
  - Although it's a framework, if you're doing Ruby, you'll likely have opportunities to work with Rails.
  - cf. [Perfect Ruby on Rails](https://bmf-tech.com/posts/%e3%83%91%e3%83%bc%e3%83%95%e3%82%a7%e3%82%af%e3%83%88Ruby%20on%20Rails)

I also picked up some books related to object-oriented programming, but haven't had the time to read them.

## Ruby Tips
Read through articles related to tips.

- [Techniques for Writing Declarative Programs in Ruby](https://qiita.com/getty104/items/41d4309dac1da41f14fc)
- [Some Ruby Tips](https://gist.github.com/kyohei-shimada/9aa61358abdc10e38bfa)
- [Ruby Competitive Programming Tips (Basics, Traps, Speedup 108 2.7x2.7)](https://zenn.dev/universato/articles/20201210-z-ruby)
  - The content is substantial.
- [Ruby's Unique Practices](https://norix.tokyo/ruby-tips/16/#outline__4)
- [[Ruby, Rails] Tips for Refactoring (Beginner)](https://qiita.com/NaokiKotani/items/36283ca922d9f96c4a11)
- [Summary of Ruby Argument Types](https://qiita.com/pink_bangbi/items/f85456db344b468ef758#%E8%AB%B8%E6%B3%A8%E6%84%8F)
  - There are many types of arguments!
  - I want to remember them while writing and reading code, as I can't memorize them all at once.
- [qiita.com - Building a Web Framework in Ruby](https://qiita.com/ta1m1kam/items/0a2658776d3dffa1cc86)
- [The Rails Doctrine: 8 Principles by the Creator of Rails](https://postd.cc/rails-doctrine/)
- [What is Ruby on Rails and How to Deal with It?](https://speakerdeck.com/yasaichi/what-is-ruby-on-rails-and-how-to-deal-with-it)
- [Ruby Memo, Especially Confusing Parts, Omitting Parentheses in Method Calls, etc.](https://qiita.com/kamiya-kei/items/fd1dad1ca8810acea9a7)
- [Modern User Authentication Model Structure Explained by the Author of Perfect Rails](https://joker1007.hatenablog.com/entry/2020/08/17/141621)
- [How DHH Organizes His Rails Controllers](https://postd.cc/how-dhh-organizes-his-rails-controllers/)

## Blog
Learned about language specifications that caught my interest and summarized them in a blog.

- [About Ruby Symbols](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%82%b7%e3%83%b3%e3%83%9c%e3%83%ab%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Block Syntax](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%83%96%e3%83%ad%e3%83%83%e3%82%af%e6%a7%8b%e6%96%87%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Proc and Lambda](https://bmf-tech.com/posts/Ruby%e3%81%aeProc%e3%81%a8lamda%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Singleton Classes and Methods](https://bmf-tech.com/posts/Ruby%e3%81%ae%e7%89%b9%e7%95%b0%e3%82%af%e3%83%a9%e3%82%b9%e3%83%bb%e7%89%b9%e7%95%b0%e3%83%a1%e3%82%bd%e3%83%83%e3%83%89%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Modules](https://bmf-tech.com/posts/Ruby%e3%81%aeModule%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)

## Coding Quizzes
Solved coding quizzes for practice.

- [HackerRank - Prepare > Ruby](https://www.hackerrank.com/domains/ruby)
  - Worked on the Ruby tutorial.

I intended to solve some problems on LeetCode but skipped it due to time constraints.

## Data Structures and Algorithms
Practicing data structures and algorithms is a good way to learn how to use the language.

Converted what I wrote in [Go](https://github.com/bmf-san/road-to-algorithm-master/tree/master) to Ruby. ChatGPT and Copilot provided significant support.

- [bmf-san/ruby-algorithm-and-datastructure-practice](https://github.com/bmf-san/ruby-algorithm-and-datastructure-practice)

## Design Patterns
Since Ruby is an object-oriented language, I practiced some patterns.

Referred to [davidgf/design-patterns-in-ruby](https://github.com/davidgf/design-patterns-in-ruby) and transcribed a few.

# Impressions
I was able to grasp the unique language specifications of Ruby, so now it's just a matter of writing code.

After learning, I felt that the verbosity of code can vary greatly depending on proficiency. While this is also true for PHP, it wasn't something I felt much with Go, so I might struggle a bit when reading code.

On the other hand, I got the impression that the simplicity in appearance and writing style is appealing.

Having been away from object-oriented languages for a while, I felt unfamiliar with handling classes, so I want to practice more.

Looking at the environment surrounding Ruby, being a domestic language, there are many Japanese committers, and the level of RubyKaigi is high, giving a sense of the community's enthusiasm. I expect there will be much to learn, so I want to continue improving.

The cultural aspect of different community atmospheres by language is quite fascinating.
