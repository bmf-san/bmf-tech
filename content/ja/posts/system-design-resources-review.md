---
title: システム設計関連の資料を読み漁った
description: システム設計関連の資料を読み漁ったについて、設計原則とトレードオフ、実践的な適用方法を詳しく解説します。
slug: system-design-resources-review
date: 2023-02-27T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - システム設計
  - リンク集
translation_key: system-design-resources-review
---


システム設計関連の資料をいくつか読み漁ったので、リンク集的にまとめておく。
動画系は見れていないものが多い。
あとあんまり関係なさそうなやや離れたトピックに関してのコンテンツも含んでいたりする。

# モチベーション
なぜ色々と調べてみたかというと、システム設計について体系的に学ぶことができないか知りたかった、学んでみたかったからである。
もっというと、システム設計についての能力を高める糸口が欲しかったといったところだ。

システム設計の能力というと、職人芸というか、経験に依るところが大きいと思っているのだが、知識として知っておいたほうが良いかも多々有るだろうと思ったので、幅広く色んな記事に目を通してみた。

当然読んでみたからといって能力が上がったわけではないし、そんな気は全くないが、知るべきことや考えるポイントみたいなところは結構学び得ることができたように思うので、今後にお役立ちな情報を得たような気はする。

# 資料一覧
## GitHub
- [donnemartin/system-design-primer](https://github.com/donnemartin/system-design-primer/blob/master/README-ja.md)
  - システム設計の基本概念が網羅されているような感じ。メンテされているようだが、404の記事や内容が古そうな記事のリンクなどもちらほらある
- [shashank88/system_design](https://github.com/shashank88/system_design)
  - システム設計インタビューの準備を始める上でのアプローチのため資料集
- [yangshun/tech-interview-handbook](https://github.com/yangshun/tech-interview-handbook)
  - システム設計ではなくコーディングインタビューに関しての内容であった
- [charlax/professional-programming](https://github.com/charlax/professional-programming) 
  - システム設計に特化ではなく、ソフトウェア開発全般のあれこれの知見について。システム設計の話もある。エンジニアリングに必要な領域が広くカバーされている。
- [karanpratapsingh/system-design](https://github.com/karanpratapsingh/system-design)
  - システムデザインの基礎的なことが端的にまとめられている
  - ベーシックな内容で、一通りのトピックが整理されていて良い
- [checkcheckzz/system-design-interview](https://github.com/checkcheckzz/system-design-interview)
  - システム設計関連の記事リンク集
- [binhnguyennus/awesome-scalability](https://github.com/binhnguyennus/awesome-scalability)
  - システム設計関連の記事リンク集。キーワードごとに数が豊富
- [System Design Cheatsheet](https://gist.github.com/vasanthk/485d1c25737e8e72759f)
  - チートシートというよりシステム設計のキーワードが記載されている感じ
- [ByteByteGoHq/system-design-101](https://github.com/ByteByteGoHq/system-design-101)
  - イメージ図付きで分かりやすい
- [mehdihadeli/awesome-software-architecture](https://github.com/mehdihadeli/awesome-software-architecture)
  - ソフトウェアアーキテクチャについて情報が集約されている

## 記事
- [Scalable Web Architecture and Distributed Systems](https://www.aosabook.org/en/distsys.html)
  - Webを分散システムとして構築するときのベストプラクティスの一例
- [Enterprise Integration Patterns - Messaging Patterns](https://www.enterpriseintegrationpatterns.com/patterns/messaging/toc.html)
  - メッセージングのパターンについてのいくつかの例をまとめて紹介している
- [tianpan.co](https://tianpan.co/notes/2016-02-13-crack-the-system-design-interview)
  - 解説は少なめで端的にアーキテクチャ事例やセオリーについて記載されている
- [hiredintech.com](https://www.hiredintech.com/system-design)
  - システム設計インタビューのアプローチについて順を追って解説。例題付き
- [Technical Interview Questions](https://www.interviewbit.com/technical-interview-questions/)
  - テクニカルインタビューについての学習ができるサイト。海外のテックの企業のjob huntingするならかなり役立ちそうなイメージ
  - いろんなクイズが用意されており、システム設計のカテゴリもあった。色んな質問や回答（ディスカッション？）があるようで、勉強になりそう
- [How to Nail the System Design Interview + Top System Design Interview Questions and Answers](https://blog.tryexponent.com/how-to-nail-the-system-design-interview/#)
  - システム設計インタビューについてのアプローチについて概要
- [System Design Interview Questions – Concepts You Should Know](https://www.freecodecamp.org/news/systems-design-for-interviews/)
  - システム設計インタビューに取り組む上で知っておきたいことのtips集
- [highscalability.com](https://highscalability.com/)
  - system-design-primerでもこのサイトのいくつかの記事が紹介されている。各企業の事例が豊富。
- [How to prepare for the System Design Interview in 2023](https://www.educative.io/blog/how-to-prepare-system-design-interview)
  - システム設計のインタビュー対策についてわかりやすくまとまった記事
- [Google System Design Interview Preparation Doc | Complete Guide](https://workat.tech/system-design/article/google-system-design-interview-prep-doc-m684to5zzkj4)
  - Googleのシステム設計インタビュー対策についてのサマリ
- [An Introduction to System Design Interviews
December 20, 2021](https://viglucci.io/introduction-to-system-design-interviews)
  - システム設計インタビューについて紹介記事。アプローチについて簡略的に記載されている。
- [System Design Interviews: A Step-By-Step Guide](https://designgurus.org/blog/step-by-step-guide) 
  - システム設計インタビューのアプローチについて記載
- [6 Interview Questions for System Designers (With Example Answers)](https://www.indeed.com/career-advice/interviewing/system-design-interview-questions)
  - システム設計インタビューについて簡易的な例題と回答例
- [Crack the System Design interview: tips from a Twitter software engineer](https://zhiachong.medium.com/how-to-system-design-dda63ed27e26)
  - Twitterのエンジニアが語るシステム設計インタビューのtips
- [System Design Interview Tutorial – The Beginner's Guide to System Design](https://www.freecodecamp.org/news/system-design-interview-practice-tutorial/)
  - Youtubeのシステム設計を例題としたインタビューのアプローチ方法
- [The complete guide to cracking the System Design interview](https://towardsdatascience.com/the-complete-guide-to-the-system-design-interview-ba118f48bdfc) 
  - システム設計インタビューのアプローチ方法と例題リンク集
- [System design interview prep (relax, start here)](http://web.archive.org/web/20250916043557/https://igotanoffer.com/blogs/tech/system-design-interview-prep)
  - システム設計インタビューの準備、基本について
- [System Design Interview Preparation Tips](https://www.interviewkickstart.com/blog/system-design-interview-preparation-tips)
  - システム設計インタビューのtips
- [Architecture Issues Scaling Web Applications](http://venkateshcm.com/2014/05/Architecture-Issues-Scaling-Web-Applications/)
  - アプリケーションのスケーリングとパフォーマンス・チューニングついて端的に要点を述べた記事
- [性能負荷テストにおけるリトルの法則](https://qiita.com/hiro0107@github/items/6154b412c7ff29c8785f)
  - 性能負荷テストにおけるリトルの法則について
  - 性能を考えるときにこの法則を思い出すと良さそう
- [How To Determine Web Application Thread Pool Size](http://venkateshcm.com/2014/05/How-To-Determine-Web-Applications-Thread-Poll-Size/)　
  - webアプリケーションのスレッドプールサイズをどのよおうに決めるのか
  - リトルの法則を念頭に考えるアプローチについて
- [principles.design](https://principles.design/examples/)
  - システム設計に限らない、設計の原則を集めている

## 本
- [ソフトウェアアーキテクチャの基礎](https://amzn.to/3y6m7hS)
  - ソフトウェアのアーキテクチャの評価観点やアーキテクチャパターンについて書かれた本
- [ソフトウェアアーキテクチャ Hardparts](https://amzn.to/44mx0Z6)
  - 分散システムにおけるアプローチは難しい問題について書かれた本
  - 読み直したい・・
  - 訳文レビューのボランティアで手伝った本でもあり、思い入れがある
- [データ指向アプリケーションデザイン](https://amzn.to/3UC7RGD)
  - システム設計のトピックから大分離れるような気がするが、分散システムの設計に通ずる部分が多々あると思うので記載
- [システム設計の面接試験](https://amzn.to/44oiVdI)が出版されたので読んだ
    - [System Design Interview – An insider's guide](https://amzn.to/3N8srJE)の和訳本
    - [donnemartin/system-design-primer](https://github.com/donnemartin/system-design-primer/blob/master/README-ja.md)を読んでいるとある程度知っている内容ではあるが、この類の本は日本ではあまり出版されていないような気がするので、手元において置きたい一冊
    - [ByteByteGo](https://bytebytego.com/)
      - 著者が運営しているサービス
      - 一部無料で基本は有料のサービスのようだが、勉強になりそうなので課金してみても良いかも
      - 著者が運営しているシステム設計についての話をするDiscordチャンネルがあとがきに記載されていたので参加しておいた 

## 動画
- [System Design - Gaurav Sen](https://www.youtube.com/playlist?list=PLMCXHnjXnTnvo6alSjVkgxV-VH6EPyvoX)
  - システム設計における主要なアプローチ、事例と解説している動画集
- [SYSTEM DESIGN INTERVIEW PREPARATION SERIES](https://www.codekarle.com/)
  - いくつかの企業のサービスを題材にシステムデザインを解説している動画集
- [CS75 (Summer 2012) Lecture 9 Scalability Harvard Web Development David Malan](https://www.youtube.com/watch?v=-W9F__D3oY4&t=6s)
  - ハーバード大でのスケーラビリティをテーマにした授業の動画
  - 多分結構有名

## スライド
- [「疎結合」を実現するメッセージングサービスの選択と利用](https://pages.awscloud.com/rs/112-TZM-766/images/DevAx_connect_season1_Day2_MessagingService_%E9%85%8D%E5%B8%83.pdf)
  - メッセージングサービスの比較・利用検討について整理されたスライド
  - AWS製品の紹介含む
- [30分でわかるデータ指向アプリケーションデザイン - Data Engineering Study #18](https://speakerdeck.com/xerial/30fen-dewakarudetazhi-xiang-apurikesiyondezain-data-engineering-study-number-18)
  - 理解するのが難しいあの本の要点をわかりやすく整理している
