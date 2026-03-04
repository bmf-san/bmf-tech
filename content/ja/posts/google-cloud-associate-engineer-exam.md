---
title: "Google Cloud認定アソシエイトクラウドエンジニアを受験した"
slug: "google-cloud-associate-engineer-exam"
date: 2023-06-07
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Google Cloud Platform"
draft: false
---

# 概要
Google Cloud認定アソシエイトクラウドエンジニアを受験して合格したので、再受験するときや別の試験を受けるときのために勉強した過程を振り返っておく。

# 前提
ソフトウェアエンジニア7~8年目くらい。
GCPの経験は2年くらいあるかな・・どちらかというとAWSのほうが触っている感じ。

# モチベーション
業務でAWSとGCPの両方を触る機会があり、AWSの方はアソシエイトレベルの認定を取得していたのだが、GCPの方はまだだったので取得に向けて勉強することにした。

AWSのほうは2年くらい前に取得していた。
cf. [AWS認定ソリューションアーキテクトアソシエイトを受験した](https://bmf-tech.com/posts/AWS%e8%aa%8d%e5%ae%9a%e3%82%bd%e3%83%aa%e3%83%a5%e3%83%bc%e3%82%b7%e3%83%a7%e3%83%b3%e3%82%a2%e3%83%bc%e3%82%ad%e3%83%86%e3%82%af%e3%83%88%e3%82%a2%e3%82%bd%e3%82%b7%e3%82%a8%e3%82%a4%e3%83%88%e3%82%92%e5%8f%97%e9%a8%93%e3%81%97%e3%81%9f)

元々AWSの認定の後に流れでGCPの方も受験しようと考えていたのだが、なんやかんやあってこのタイミングになってしまった。

今年はKubernetesをいい加減ちゃんとキャッチアップしようと考えていて、K8Sの運用ならEKSよりGKEがいいかな（※個人の感想）と思うところがあり、GKE含めGCPの各種サービスについて知り、GCP上でアーキテクチャ設計ができるようになりたいという動機も受験を後押しした。

AWSの方は資格の有効期限が3年あるが、GCPの方は2年とやや短いが、それについては特にモチベーションに影響しなかった。

# 受験した感想
当日は横浜の試験センターで受験した。

AWSと同じくオンラインでも受験できるのだが、オンサイトのほうが環境について気にすること、気をつけることが少ないのでオンサイトにした。

平日だったこともあって全然人がいなくて思いっきり集中できた。

試験内容のあれこれについては言えないので、一言だけ感想をいうと、勉強したことは十分発揮できたように感じた。

正式な試験結果が予定よりも遅延したのだが、どうやら近い時期に受験した人も同じく遅延したらしい。
cf. https://note.com/aiue408/n/n8d5587f7362a

# 勉強期間
期間としては2~3ヶ月くらい。

育休のすきま時間で勉強していたが、以前から少し勉強を進めていたのものあって実際は2ヶ月弱くらいな気がする。

３ヶ月以内の期間で合格する計画を立てていたのだが、2週間くらい予定を前倒しできて良かった。

# 勉強したこと
勉強した内容の覚え書きは[GCPについての覚え書き](https://bmf-tech.com/posts/GCP%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6%e3%81%ae%e8%a6%9a%e3%81%88%e6%9b%b8%e3%81%8d)にまとめた。要点を頭にインプットするようの内容でラフに書きまとめておいたのだが、結構役立った。

公式の情報は充実しているが、模擬試験の問題集（参考書）がAWSよりは少ないので、試験の素振り回数が少なくなるがちなのが不安になる部分だった。

出題範囲、傾向を把握して、公式の情報をベースに勉強を重ねていけばなんとかなった。

## 書籍
- GCPの教科書
- GCPの教科書2

GCPの教科書は3も出版されているが、Cloud AIプロダクトについてはあまり深い問いがなさそうだったので読むのをサボっている。。

まず入門したいという感じであれば読んでも良いが、最短で試験勉強ということであれば読まなくても良いかなと思う。

## ドキュメント
- [ドキュメント](https://cloud.google.com/docs?hl=ja) 
  - ある程度目を通しておいた。
  - AWS認定の方も同様だが、試験に最短で合格するなら全部目を通さず、模擬試験を問いて足りないところを埋めていく形で参照するくらいで十分だと思う。自分は興味があったので一通り読んだ。
- [クラウドアーキテクチャのガイダンス](https://cloud.google.com/architecture?hl=ja)
  - リファレンス、ガイダンス、ベストプラクティスが提唱されている各種資料
  - 全部読むのは大変なので、気になるものだけ読んだ
  - 試験に関わらずGCPを使うのであれば[Google Cloud アーキテクチャ フレームワーク](https://cloud.google.com/architecture/framework?hl=ja)は一度読んでおいても損ないかなと思ったので、これだけでも良いかも。
- [gcloud CLI クイック リファレンス](https://cloud.google.com/sdk/docs/cheatsheet?hl=ja)
  - 全部は見ていない。模擬試験や試験範囲に関わるよく使われそうなコマンドだけ参照した
  - Udemyの模擬試験の傾向から判断して、iam、compute、container、config、app周りのコマンドは入念に見ておいた

## ブログ
- [GCPSketchnote](https://github.com/priyankavergadia/GCPSketchnote)
  - サービスについての要点やTipsがわかりやすく端的にまとめられている記事群
  - 一通り呼んだが、最短で試験勉強するのであれば読まなくても良いと思う
- [Google Cloud Japan Advent Calendar 2022- 今から始める Google Cloud](https://zenn.dev/google_cloud_jp/articles/12bd83cd5b3370#%E4%BB%8A%E3%81%8B%E3%82%89%E5%A7%8B%E3%82%81%E3%82%8B-google-cloud)
  - 最短で試験勉強するのであれば読まなくても良いと思うが、Google Cloud Japanの中の人が分かりやすい記事を書いているので勉強になる。
- [Associate Cloud Engineer試験対策マニュアル。出題傾向・勉強方法](https://blog.g-gen.co.jp/entry/associate-cloud-engineer)
  - 諸々の勉強をした後に自分の学習に抜け漏れがないか確認するために参考にした
  - 試験の要点が抑えられている
  - g-genさんはGCPについて分かりやすい記事を色々書いているのでこの記事以外も色々読ませて頂いた
  - 公式の情報が一次ソースなので、当然そちらを確認した上で参考にすること

# Youtube
- [【GCP】Google Cloud Platform認定 Associate Cloud Engineer解説動画【Google Cloud】](https://www.youtube.com/watch?app=desktop&v=7-IZv9o15t8)
  - 要点をバーっと知りたかったので2倍速で見た

# Cousera
- [Architecting with Google Compute Engine 日本語版専門講座
](https://www.coursera.org/specializations/gcp-architecture-jp?action=enroll&authMode=signup#courses)
  - これはやらなかった
  - おそらく後述のGoogle Skills Boostと内容重複する？ような感じだったので、受講する必要はなさそうだと判断

# Google Skills Boost
- [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training)
  - GCP ACE取得に向けた学習プランが無料で提供されている
  - 何から手を付けるか迷う場合は、[Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737)流れに従って学習していくのが良さそう
    - 最初からこれに従っておけばよかったかも
    - テーマごとに関連資料が用意されているので、学習しやすい

以下は[Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737)でも関連資料として記載があったコース。

 [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training)で紹介のあった下記コースも複数消化した。

- [Essential Google Cloud Infrastructure: Foundation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/50)
- [Essential Google Cloud Infrastructure: Core Services | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/49)
- [Elastic Google Cloud Infrastructure: Scaling and Automation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/178)
- [Getting Started with Google Kubernetes Engine | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/2)
- [Google Cloud Fundamentals: Core Infrastructure | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/60)

上記のコースで学びつつ、苦手な領域はQwicklabsでカバーするようにした。

# Google Cloud Certified Jumpstartプログラム
- [第 2 回 Google Cloud Certification Jumpstart プログラム](https://cloudonair.withgoogle.com/events/jp-cert-2)
  - 公式提供のプログラム
  - 試験日数日前に知ったのでチラ見しかしていないが、効率的に学習できそうなので良さそう

# 問題集
- [【GCP認定資格】Google Cloud Platform Associate Cloud Engineer模擬問題集](https://www.udemy.com/course/gcp-ace-mogi/)
  - 練習問題が4つあって問題集としてはおそらくこれが一番充実している
  - 合格率の目標に達するまで回答と復習を繰り返した。といっても2周しかしていないが..
  - 後述のGoogle提供の模擬試験より難しい気はするが、これをベースに勉強することで有意義に勉強することができたと思う
- [Associate Cloud Engineer 模擬試験](https://docs.google.com/forms/d/e/1FAIpQLSc7bkUHpDbFShBI5xE4u8OO2vl99DrP0htnswa-la9DQynToA/viewform?hl=ja&hl=ja)
  - Googleが無料で用意しているGoogle Formで回答する形式の模擬試験
  - 受験日の前々日くらいに最後の仕上げとして取り組んだ

# その他
[Qwiklabs](https://nvlabs.qwiklabs.com/journeys)でGCP関連のものをいくつか取り組んだ。

手を動かすほうが頭に入るし、実践的なので良い。AWSの試験のときにもお世話になった。

試験ではgcloudのコマンドが問われる設問もあるため、Qwicklabsでコマンドに触れておくのは有意義だと思う。

後は触ったことがなかったサービスの理解を深めるためにも役に立った。

個別のクレジット購入よりサブスクリプションのほうがコスパが良かったので課金した。かれこれ人生で4回くらい課金しているが、年間プランをサブスクする程ではなかったりするので都度課金している。。。（年中触らないので・・）

# 所感
一発で合格して良かった。

AWSの認定試験と違って、再受験ポリシーというものが設けられているため、不合格になるとペナルティような制限（最大受験回数や再受験までのクールタイムがある。また再受験は再度受験費用を負担することになる。）を受けるので、その度にプレッシャーが高まる。

不合格だったとしても、出題傾向や難易度を把握できるのでもしも再受験することになったら合格率は高まるかなとポジティブに考えるようにして、不合格を意識しないようにしていたが、試験の回答を提出するときだけはちょっと手が震えたw

AWSやGCPの認定試験は単なる暗記試験ではなく、実用的な知識が問われる試験であると思っているので、基礎知識の獲得と表明のためにコスト（時間・お金）を払う価値が十分にあると考えているので、今後も継続的に試験を受けたいと思う。
