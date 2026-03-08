---
title: Learning Ruby
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
As someone who has worked with PHP and Go, I will write about my experiences learning Ruby.

# Efforts
## Official Documentation
First and foremost, the official documentation.

- [What is Ruby](https://www.ruby-lang.org/ja/about/)
  - Concisely describes the features of Ruby.
  - For specifications that you want to research deeply, this seems to be a good reference.
- [Introduction to Ruby from Other Languages](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/)
  - Describes the distinctive parts of Ruby's language specifications. It summarizes points to learn that are not present in other languages.
  - For specifications that you want to research deeply, this can also be referenced.
- [From PHP to Ruby](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/to-ruby-from-php/)
  - There was nothing for Go. As a former PHP developer, I referred to the PHP page. The content is quite straightforward.
- [Getting Started with Ruby in 20 Minutes](https://www.ruby-lang.org/ja/documentation/quickstart/)
  - Just right for a light practice.
- [Object-Oriented Scripting Language Ruby Reference Manual (Ruby 3.3 Reference Manual)](https://docs.ruby-lang.org/ja/3.3/doc/index.html)
  - A reference manual by volunteers. Much appreciated.

## Books
I picked up some standard books.

- [Programming Ruby 1.9](https://amzn.to/3JZamMX)
- [Ruby Programming for Beginners](https://amzn.to/4dG4ciA)
  - cf. [Ruby Programming for Beginners](https://bmf-tech.com/posts/%e3%83%97%e3%83%ad%e3%82%92%e7%9b%ae%e6%8c%87%e3%81%99%e4%ba%ba%e3%81%ae%e3%81%9f%e3%82%81%e3%81%aeRuby%e5%85%a5%e9%96%80)
- [Perfect Ruby](https://amzn.to/3K0wLcR)
  - cf. [Perfect Ruby](https://bmf-tech.com/posts/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88Ruby)
- [Perfect Ruby on Rails](https://amzn.to/3yfO0nL)
  - Although it's a framework, if you're doing Ruby, you'll likely have opportunities to touch Rails as well.
  - cf. [Perfect Ruby on Rails](https://bmf-tech.com/posts/%e3%83%91%e3%83%bc%e3%83%95%e3%82%a7%e3%82%af%e3%83%88Ruby%20on%20Rails)

I also picked up some books related to object-oriented programming, but I haven't been able to read them due to time constraints.

## Ruby Tips
I read a lot of articles related to tips.

- [Techniques for Writing Declarative Programs in Ruby](https://qiita.com/getty104/items/41d4309dac1da41f14fc)
- [Some Ruby Tips](https://gist.github.com/kyohei-shimada/9aa61358abdc10e38bfa)
- [Ruby Competitive Programming Tips (Basics, Traps, Speedup 108 2.7x2.7)](https://zenn.dev/universato/articles/20201210-z-ruby)
  - The content is substantial.
- [Ruby-Specific Conventions](https://norix.tokyo/ruby-tips/16/#outline__4)
- [[Ruby, Rails] Tips for Refactoring (Beginner-Friendly)](https://qiita.com/NaokiKotani/items/36283ca922d9f96c4a11)
- [Summary of Argument Types in Ruby](https://qiita.com/pink_bangbi/items/f85456db344b468ef758#%E8%AB%B8%E6%B3%A8%E6%84%8F)
  - There are many types of arguments...!
  - Since I can't memorize them all at once, I want to remember them while writing and reading code.
- [qiita.com - Creating a Web Framework in Ruby](https://qiita.com/ta1m1kam/items/0a2658776d3dffa1cc86)
- [Basic Philosophy of Rails: Eight Principles from the Creator of Rails](https://postd.cc/rails-doctrine/)
- [Understanding Ruby on Rails and How to Deal with It](https://speakerdeck.com/yasaichi/what-is-ruby-on-rails-and-how-to-deal-with-it)
- [Ruby Memoirs: Particularly Confusing Aspects of Ruby, Omitting Parentheses in Method Calls, etc.](https://qiita.com/kamiya-kei/items/fd1dad1ca8810acea9a7)
- [Explanation of Modern User Authentication Model Structure by the Author of Perfect Rails](https://joker1007.hatenablog.com/entry/2020/08/17/141621)
- [How DHH Organizes His Rails Controllers](https://postd.cc/how-dhh-organizes-his-rails-controllers/)

## Blog
I learned about language specifications that caught my attention and summarized them in a blog.

- [About Ruby Symbols](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%82%b7%e3%83%b3%e3%83%9c%e3%83%ab%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Block Syntax](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%83%96%e3%83%ad%e3%83%83%e3%82%af%e6%a7%8b%e6%96%87%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Proc and Lambda](https://bmf-tech.com/posts/Ruby%e3%81%aeProc%e3%81%a8lamda%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby's Singleton Class and Singleton Method](https://bmf-tech.com/posts/Ruby%e3%81%ae%e7%89%b9%e7%95%b0%e3%82%af%e3%83%a9%e3%82%b9%e3%83%bb%e7%89%b9%e7%95%b0%e3%83%a1%e3%82%bd%e3%83%83%e3%83%89%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [About Ruby Modules](https://bmf-tech.com/posts/Ruby%e3%81%aeModule%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)

## Coding Quizzes
I solved some coding quizzes to warm up.

- [HackerRank - Prepare > Ruby](https://www.hackerrank.com/domains/ruby)
  - There are Ruby tutorials, so I worked on them.

I thought about solving a few problems on LeetCode, but I skipped it for now due to time constraints.

## Data Structures and Algorithms
This is a good subject to remember how to use the language, so I practiced.

I tried converting something I wrote in [Go](https://github.com/bmf-san/road-to-algorithm-master/tree/master) to Ruby. ChatGPT and Copilot supported me quite a bit.

- [bmf-san/ruby-algorithm-and-datastructure-practice](https://github.com/bmf-san/ruby-algorithm-and-datastructure-practice)

## Design Patterns
Since Ruby is an object-oriented language, I practiced a few patterns.

I referred to [davidgf/design-patterns-in-ruby](https://github.com/davidgf/design-patterns-in-ruby) and picked a few to transcribe.

# Impressions
I was able to grasp the distinctive language specifications of Ruby, so now I just need to keep writing code.

After learning everything, I felt that the redundancy of the code would vary significantly depending on the level of proficiency. There are parts like this in PHP, but I didn't feel it much in Go, so I might struggle a bit when reading code for a while.

On the other hand, I felt that there seems to be simplicity in terms of appearance and writing style.

Having been away from object-oriented languages for a while, I felt a lack of familiarity with handling classes, so I want to practice.

Looking at the environment surrounding Ruby, there are many Japanese committers since it's a domestic language, the level of RubyKaigi is high, and I felt the enthusiasm of the Ruby community. I have high expectations for learning, so I want to continue to improve.

The cultural aspect of how the atmosphere of the community differs by language is very interesting.