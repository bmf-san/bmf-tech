---
title: Rubyを学ぶ
description: Rubyを学ぶ
slug: learning-ruby
date: 2024-05-16T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: learning-ruby
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
  - cf. [プロを目指す人のためのRuby入門](https://bmf-tech.com/posts/%e3%83%97%e3%83%ad%e3%82%92%e7%9b%ae%e6%8c%87%e3%81%99%e4%ba%ba%e3%81%ae%e3%81%9f%e3%82%81%e3%81%aeRuby%e5%85%a5%e9%96%80)
- [パーフェクトRuby](https://amzn.to/3K0wLcR)
  - cf. [パーフェクトRuby](https://bmf-tech.com/posts/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88Ruby)
- [パーフェクトRuby on Rails](https://amzn.to/3yfO0nL)
  - Rubyではなくフレームワークだが、RubyやるならRailsも触る機会があるだろう
  - cf. [パーフェクトRuby on Rails](https://bmf-tech.com/posts/%e3%83%91%e3%83%bc%e3%83%95%e3%82%a7%e3%82%af%e3%83%88Ruby%20on%20Rails)

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

- [Rubyのシンボルについて](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%82%b7%e3%83%b3%e3%83%9c%e3%83%ab%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [Rubyのブロック構文について](https://bmf-tech.com/posts/Ruby%e3%81%ae%e3%83%96%e3%83%ad%e3%83%83%e3%82%af%e6%a7%8b%e6%96%87%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [RubyのProcとlamdaについて](https://bmf-tech.com/posts/Ruby%e3%81%aeProc%e3%81%a8lamda%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [Rubyの特異クラス・特異メソッドについて](https://bmf-tech.com/posts/Ruby%e3%81%ae%e7%89%b9%e7%95%b0%e3%82%af%e3%83%a9%e3%82%b9%e3%83%bb%e7%89%b9%e7%95%b0%e3%83%a1%e3%82%bd%e3%83%83%e3%83%89%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
- [RubyのModuleについて](https://bmf-tech.com/posts/Ruby%e3%81%aeModule%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)

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
