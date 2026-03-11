---
title: モノレポについて
slug: monorepo-overview
date: 2023-08-11T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - モノレポ
translation_key: monorepo-overview
---


# 概要
モノレポについてまとめる。

# モノレポとは
モノレポとは複数プロジェクトのコードを単一のレポジトリで管理したもの。対して複数のレポジトリで管理するものをポリレポ、またはマルチレポと呼ぶ。
マイクロサービスの管理方針の１つではあるものの、マイクロサービスを前提としたものではない。
モノリスとは同義ではない。

# モノレポの観点
モノレポの運用における観点を整理してみた。

## コードの自治
複数チームで運用する場合、チーム管轄外のコードも変更できてしまう。

GitHubであれば、CODEOWNERSによる管理で管轄範囲を整備できる。何かしらのツールに頼ったレギュレーションを敷く必要がありそう。

## 大きな泥団子
コード間の依存関係が複雑化することによる泥団子化。

これも何かしらのツールに頼った解決が必要そう。例えば、Nxではpublic APIを持つライブラリを作ることができたり、依存関係をグラフで可視化することができる。

## スケーラビリティ
コード量が増加していくあたり、ビルドやテスト、デプロイが遅くなる、Gitの管理上の問題が発生する。

前者については、個々に実行可能なCIパイプラインやデプロイフロー等の整備により解決できる。ツールによる解決が可能な範囲である。

後者については、少し悩み所かもしれない。cloneやpullが辛くなる状態を迎える時は何か対応を検討する必要がありそう。

Microsoftが開発しているGVFSというスケーラブルなGitを使う、Git LFSを活用する、諦めてリポジトリを分割するなど。

## 技術選択の自由度
特定の技術（プログラミング言語など）に縛られるという制約は基本的にはない。

コード管理が単一であるだけで、CI・CD等は複数のパイプライン管理を想定しているので特に懸念にはならなそう。

ビルドツールが対応していない言語や環境があると技術選択の幅が制限されるという可能性があるかもしれない。

## ブランチ戦略
featureブランチとの相性があまり良くないので、トランクベースの戦略を導入するのが望ましそう。フィーチャートグルも合わせて導入したいところ。

## その他運用上の懸念
GitHubであれば、IssueやPullRequestの運用方針を整備することに気を配る必要があるかもしれない。

# モノレポのメリット・デメリット
## メリット
- 全体像が把握しやすい
  - 1箇所にまとめて管理されていることにより、プロジェクトやサービスに関わるコードの全体を観測しやすい
- サイロ化の防止
  - 複数チームによるモノレポの運用の場合、チーム間の共有コストがポリレポよりも低い。（たぶん）
- 再利用しやすい
  - コードの再利用、統一がポリレポより容易
- 標準化が推進しやすい
  - レポジトリの運用方針をチーム間で揃えることができるので、ガバナンスを効かせやすい
  - これは個人的に大きいメリットだと感じている
    - ポリレポを基本としている開発組織において、レポジトリごとに運用方針が微妙に異なるようなチーム自治な組織形態を取っているとチーム間の人の移動の際に学習コストがオーバーヘッドになるという問題あると思うが、モノレポだとそのオーバーヘッドを減らすことができそう
    - ポリレポの場合、各リポジトリに同じような変更を加えたい（例えばセキュリティ的な問題で特定のCIサービスを使っているリポジトリにおいて同じような変更をしたいとき、など）時に手間が掛かるが、モノレポだと負担が減りそう

## デメリット
- 共通管理部分のメンテナンスコスト
  - CI・CDのパイプラインは上手く運用する必要がある
    - 複雑化する恐れ
- 依存関係の管理
  - 依存関係を可視化することで依存関係を把握することはしやすいと思うが、依存によりコードの変更しやすさが損なわれると開発生産性が落ちそう
- Gitのスケーラビリティ
- ビルドツールへの依存・キャッチアップコスト
  - モノレポを始めるために何かしらのビルドツールの導入が前提になりそう
  - デメリットというほどでもないかもしれないが、レポジトリの管理のためにツールが1つ増えるというのが気になるか否か
  - ツールによっては学習コストがあまり軽視できない
    - 例えばBazelを採用するとBazel職人が必要になる可能性
    - 開発者が誰でもツールを触れるようにしないと運用が破綻するかも
    - kubernetesのプロジェクトではBazelが削除されたなんてことも・・
      - [Remove Bazel](https://github.com/kubernetes/kubernetes/pull/99561)
- エコシステムの対応状況
  - 例えばIDEがビルドツールに対応しきれていない可能性など

# モノレポのためのツール
[monorepo.tools - Many solutions, for different goals](https://monorepo.tools/#tools-review)に詳しく整理されている。

Bazel、Nx、Pants辺りが有力候補なのかなといった印象。

# 所感
- モノレポの強みを発揮するには、バラバラに管理されているレポジトリを単に１つにまとめれば良い、というわけではない
- モノレポにおいて、コードの共通化が必要かどうか？という観点がありそう
  - コードを単一リポジトリを管理するのであって、アーキテクチャの方針が縛られるわけではないとは思うので、共通化するかどうかは適宜判断かなと思う
  - マイクロサービスやモジュラモノリスなどアーキテクチャの方針によって何をどこまで共通化するのか、しないのかは方針次第であって、共通化することは自体が必須ではないと思うが、ビルドやデプロイ等のパイプラインについてはある程度共通化が前提となるところはありそうだと思った
- ワンチームによる運用なのか、複数チームによる運用なのかによって運用の難しさが大きく変わりそう。特に後者の場合はレギュレーションをきっちりと整備しないとそれこそ泥団子状態になりやすいと思う
  - 複数チームでやる場合は、横断的な関心事（CIのパイプライン管理など）をどう扱うか、どのようにリードするかが組織課題になりそう
- どういう単位、スコープでモノレポにするかというのが最初の考え所だと感じた
  - フルサイクルに開発できる構成でモノレポにするのが良いかなと思うが、それが1サービス（特定の1つのシステム）なのか、1つのプロダクトなのか、ドメイン単位なのかなど構成単位をどうするかはよく考える必要がありそう
- モノレポからポリレポ、ポリレポからモノレポへの移行はどちらも同じような苦労が伴うか？というのはちょっと気になった
  - 多分、統合されたCIパイプラインを解体することのほうが面倒そうな気がするので、モノレポからポリレポに移行するほうが大変なのではと思う・・・

# 参考
- [monorepo.tools - monorepo.tools](https://monorepo.tools/)
- [circleci.com - Monorepo開発のメリット vs デメリット](https://circleci.com/ja/blog/monorepo-dev-practices/)
- [www.graat.co.jp - モノレポについての誤解 - Misconceptions about Monorepos: Monorepo != Monolith を翻訳しました](https://www.graat.co.jp/blogs/ck1099bcoeud60830rf0ej0ix)
- [zenn.dev - Monorepoって何なのか？と関連アーキテクチャとの関係をまとめてみた](https://zenn.dev/burizae/articles/c811cae767965a)
- [hireroo.io - モノレポでマイクロサービスを開発するための戦略と運用](https://hireroo.io/journal/tech/mono-repo-for-microservices)
- [gist.github.com - モノレポについて](https://gist.github.com/pipopotamasu/efe7097454d9668f80cd8b43068afafc)
- [blog.ojisan.io - モノレポにすべきか、レポジトリを分割すべきか](https://blog.ojisan.io/monorepo-vs-polyrepo/)
- [note.com - モノレポによるマイクロサービスの開発運用](https://note.com/tinkermodejapan/n/nb14009fe837f)
- [caddi.tech - AI 組織のモノレポ紹介 Technology](https://caddi.tech/archives/4187)
- [cam-inc.co.jp - 運用してわかった！モノレポに向いているプロジェクト](https://cam-inc.co.jp/p/techblog/570556215432577985)
- [tech.asoview.co.jp - 3ヶ月で120のリポジトリを1つのMonorepo(モノレポ/モノリポ)に移行した話](https://tech.asoview.co.jp/entry/2022/12/23/095914)
- [qiita.com - モノレポとマイクロサービス](https://qiita.com/ytanaka3/items/6d8d960179bc046e38c0)
- [cloudsmith.co.jp - [Monorepo] React+Node.js+Typescript モノレポ構築備忘録](https://cloudsmith.co.jp/blog/frontend/2023/06/2396016.html)
- [engineering.mercari.com - メルカリShopsでのmonorepo開発体験記](https://engineering.mercari.com/blog/entry/20210817-8f561697cc/)
- [times.hrbrain.co.jp - 30近いリポジトリを一つのリポジトリにまとめました](https://times.hrbrain.co.jp/entry/monorepo)
- [docs.aws.amazon.com - モノレポのビルドの設定](https://docs.aws.amazon.com/ja_jp/amplify/latest/userguide/monorepo-configuration.html)
- [postd.cc - モノリシックなバージョン管理の利点](https://postd.cc/monorepo/)
- [kk-web.link - モノレポ？いらんでしょ](https://kk-web.link/blog/20210507)
- [www.atlassian.com - Gitでmonorepoを扱う際の課題とヒント](https://www.atlassian.com/ja/blog/monorepos-in-git)
- [cybozu.github.io - フロントエンドのモノレポ構成はスケーリングの夢を見るか](https://cybozu.github.io/frontend-expert/posts/considerations-for-monorepo)
- [kiyobl.com - モノレポとは ？ | yarn workspace を使ったモノレポ開発](https://kiyobl.com/monorepo-basic/)
- [speakerdeck.com - モノレポによるマイクロサービスアーキテクチャの開発運用](https://speakerdeck.com/bananaumai/monorehoniyorumaikurosahisuakitekutiyanokai-fa-yun-yong)
