---
title: Rubyを学ぶ
description: "PHPやGoを触ってきた自分が、Rubyを学ぶ際に取り組んだことをまとめた記事です。まずは公式ドキュメントから入り、定番の書籍やオブジェクト指向の設計パターン、実際に手を動かす実践を通じて、Rubyの基礎を身につけていく道のりを紹介します。"
slug: learning-ruby
date: 2024-05-16T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: learning-ruby
draft: false
---


# 概要
PHPやGoを触ってきた人間はRubyを学ぶ際に取り組んだことを書く。

# 取り組み
## 公式ドキュメント
何はともあれはまずは公式ドキュメント。

- [Rubyとは](https://www.ruby-lang.org/ja/about/)
  - Rubyの特徴を端的に記している
  - 個別に深く調べておきたい仕様についてはここを参照すれば良さそう。
- [他言語からのRuby入門](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/)
  - Rubyの言語仕様で特徴的な部分を記している。他言語にはない部分で学ぶべきポイントがまとまっている。
  - 個別に深く調べておきたい仕様についてはこちらも参照すれば良さそう。
- [PHPからRubyへ](https://www.ruby-lang.org/ja/documentation/ruby-from-other-languages/to-ruby-from-php/)
  - Goはなかった。元PHPerだったのでPHPのページを参照してみた。内容としては結構あっさりしている。
- [20分ではじめるRuby](https://www.ruby-lang.org/ja/documentation/quickstart/)
  - 軽く素振りするのにちょうど良い。
- [オブジェクト指向スクリプト言語 Ruby リファレンスマニュアル (Ruby 3.3 リファレンスマニュアル)](https://docs.ruby-lang.org/ja/3.3/doc/index.html)
  - 有志によるリファレンスマニュアル。ありがたや。

## 本
定番系のものをピックアップした。

- [たのしいRuby 第6版](https://amzn.to/3JZamMX)
- [プロを目指す人のためのRuby入門](https://amzn.to/4dG4ciA)
  - cf. [プロを目指す人のためのRuby入門](/ja/posts/ruby-introduction-for-beginners/)
- [パーフェクトRuby](https://amzn.to/3K0wLcR)
  - cf. [パーフェクトRuby](/ja/posts/perfect-ruby/)
- [パーフェクトRuby on Rails](https://amzn.to/3yfO0nL)
  - Rubyではなくフレームワークだが、RubyやるならRailsも触る機会があるだろう
  - cf. [パーフェクトRuby on Rails](/ja/posts/perfect-ruby-on-rails/)

オブジェクト指向関連の本もピックアップしてはいたが、時間的都合で読めていない。

## Ruby Tips
Tips関連の記事を読み漁った。

- [Rubyで宣言的なプログラムを書くためのテクニックTips](https://qiita.com/getty104/items/41d4309dac1da41f14fc)
- [RubyのちょとしたTips](https://gist.github.com/kyohei-shimada/9aa61358abdc10e38bfa)
- [Ruby競プロTips(基本・罠・高速化108 2.7x2.7)](https://zenn.dev/universato/articles/20201210-z-ruby)
  - 内容が充実している
- [Ruby特有の作法](http://web.archive.org/web/20250814105230/https://norix.tokyo/ruby-tips/16/)
- [[Ruby, Rails] リファクタリングに役立つTips集 (初心者向け)](https://qiita.com/NaokiKotani/items/36283ca922d9f96c4a11)
- [Ruby の引数の種類をまとめてみた](https://qiita.com/pink_bangbi/items/f85456db344b468ef758#%E8%AB%B8%E6%B3%A8%E6%84%8F)
  - 引数の種類が多い・・！
  - 一度に覚えきれないのでコード書いたり読んだりする上で覚えていきたい
- [qiita.com - RubyでWebフレームワークを自作する](https://qiita.com/ta1m1kam/items/0a2658776d3dffa1cc86)
- [Railsの基本理念 : Railsの生みの親が掲げる8つの原則](https://postd.cc/rails-doctrine/)
- [Ruby on Railsの正体と向き合い方 / What is Ruby on Rails and how to deal with it?](https://speakerdeck.com/yasaichi/what-is-ruby-on-rails-and-how-to-deal-with-it)
- [Ruby備忘録、Rubyの特に分かりにくかったところ、メソッド呼び出しのカッコ省略等](https://qiita.com/kamiya-kei/items/fd1dad1ca8810acea9a7)
- [パーフェクトRails著者が解説するdeviseの現代的なユーザー認証のモデル構成について](https://joker1007.hatenablog.com/entry/2020/08/17/141621)
- [DHHはどのようにRailsのコントローラを書くのか](https://postd.cc/how-dhh-organizes-his-rails-controllers/)

## ブログ
気になった言語仕様について学んでブログにまとめておいた。

- [Rubyのシンボルについて](/ja/posts/ruby-symbols/)
- [Rubyのブロック構文について](/ja/posts/ruby-block-syntax/)
- [RubyのProcとlamdaについて](/ja/posts/ruby-proc-lambda/)
- [Rubyの特異クラス・特異メソッドについて](/ja/posts/ruby-singleton-class-methods/)
- [RubyのModuleについて](/ja/posts/ruby-modules/)

## コーディングクイズ
肩慣らしにコーディングクイズを解いておく。

- [HackerRank - Prepare > Ruby](https://www.hackerrank.com/domains/ruby)
  - Rubyのチュートリアルがあるので取り組んでおいた。

LeetCodeも数問解こうと思ったが時間的都合で一旦スキップした。

## データ構造とアルゴリズム
言語の使い方を覚えるにあたってちょうど良い題材なので素振りした。

昔[Goで書いたもの](https://github.com/bmf-san/road-to-algorithm-master/tree/master)をRubyに変換してみた。ChatGPTとCopilotがかなりサポートしてくれた。

- [bmf-san/ruby-algorithm-and-datastructure-practice](https://github.com/bmf-san/ruby-algorithm-and-datastructure-practice)

## デザインパターン
Rubyはオブジェクト指向言語であるので、いくつかパターンを素振りしておいた。

[davidgf/design-patterns-in-ruby](https://github.com/davidgf/design-patterns-in-ruby)を参考に、いくつかピックアップして写経した。

# 所感
おおよそRubyの特徴的な言語仕様を把握することができたので、あとはひたすらコードを書いていく。

一通り学んで思ったのは、練度によってコードの冗長さがかなりブレるだろうなという印象を持った。PHPでもそういう部分はあるが、Goではあまり感じない部分ではあったので、コード読む際にしばらく苦労するかもしれない。

一方で、見た目や書き味的なところはシンプルさがありそうという印象を持った。

オブジェクト指向言語からしばらく離れていたので、クラス周りの取り扱いに不慣れさを感じたので練習したい。

Rubyを取り巻く環境を見ると、国産言語だけあって日本人コミッターが多かったり、RubyKaigiのレベルが高かったり、Rubyのコミュニティの熱量が高さを感じた。学びが多いだろうなという期待感を感じたので今後とも精進していきたい。

言語によってコミュニティの雰囲気が違う文化的な側面は大変面白い。
