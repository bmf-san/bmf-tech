---
title: "Monorepos Explained: Benefits, Drawbacks, and When to Use Them"
slug: monorepo-overview
date: 2023-08-11T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Monorepo
translation_key: monorepo-overview
---



# Overview
Summarizing about monorepos.

# What is a Monorepo?
A monorepo is a single repository that manages the code of multiple projects. In contrast, managing with multiple repositories is called a polyrepo or multi-repo. Although it is one of the management strategies for microservices, it is not premised on microservices. It is not synonymous with a monolith.

# Perspectives on Monorepo
Let's organize the perspectives on operating a monorepo.

## Code Autonomy
When operated by multiple teams, code outside the team's jurisdiction can also be changed.

In GitHub, management with CODEOWNERS can organize the jurisdiction. It seems necessary to establish regulations relying on some tools.

## Big Ball of Mud
The complexity of dependencies between codes can lead to a big ball of mud.

This also seems to require a solution relying on some tools. For example, Nx allows you to create libraries with public APIs and visualize dependencies in a graph.

## Scalability
As the amount of code increases, builds, tests, and deployments become slower, and Git management issues arise.

The former can be resolved by organizing CI pipelines and deployment flows that can be executed individually. It is within the range that can be solved by tools.

The latter might be a bit of a concern. When cloning or pulling becomes burdensome, some countermeasures need to be considered.

Using scalable Git developed by Microsoft called GVFS, utilizing Git LFS, or giving up and splitting the repository are some options.

## Freedom of Technology Choice
There is basically no constraint to be bound by specific technologies (such as programming languages).

Since code management is single, but CI/CD, etc., are assumed to be managed by multiple pipelines, it doesn't seem to be a particular concern.

There might be a possibility that the range of technology choices is limited if there are languages or environments not supported by the build tools.

## Branch Strategy
Since it doesn't match well with feature branches, it seems desirable to introduce a trunk-based strategy. It would be good to introduce feature toggles as well.

## Other Operational Concerns
In GitHub, it might be necessary to pay attention to organizing the operation policy of Issues and Pull Requests.

# Advantages and Disadvantages of Monorepo
## Advantages
- Easy to grasp the overall picture
  - By being managed in one place, it's easy to observe the entire code related to projects or services.
- Prevention of Silos
  - In the case of monorepo operation by multiple teams, the sharing cost between teams is lower than polyrepo. (Probably)
- Easy to Reuse
  - Code reuse and unification are easier than polyrepo.
- Easy to Promote Standardization
  - Since the repository operation policy can be aligned between teams, governance can be enforced easily.
  - I personally feel this is a big advantage.
    - In a development organization based on polyrepo, if each repository has slightly different operation policies, it can become an overhead in learning costs when people move between teams, but monorepo can reduce that overhead.
    - In the case of polyrepo, when you want to make similar changes to each repository (for example, when you want to make similar changes to repositories using a specific CI service due to security issues), it can be cumbersome, but monorepo can reduce the burden.

## Disadvantages
- Maintenance Cost of Common Management Parts
  - CI/CD pipelines need to be operated well
    - There is a risk of becoming complex
- Dependency Management
  - While visualizing dependencies makes it easier to grasp them, if dependencies make code changes difficult, development productivity might drop.
- Git Scalability
- Dependency on Build Tools and Catch-up Cost
  - Introducing some build tools seems to be a prerequisite to starting a monorepo.
  - It might not be a disadvantage per se, but whether it bothers you to have one more tool for repository management.
  - The learning cost might not be negligible depending on the tool.
    - For example, adopting Bazel might require a Bazel expert.
    - If developers can't handle the tool, operations might collapse.
    - In the Kubernetes project, Bazel was removed...
      - [Remove Bazel](https://github.com/kubernetes/kubernetes/pull/99561)
- Ecosystem Support Status
  - For example, the possibility that IDEs might not fully support build tools.

# Tools for Monorepo
[monorepo.tools - Many solutions, for different goals](https://monorepo.tools/#tools-review) provides a detailed organization.

Bazel, Nx, and Pants seem to be strong candidates.

# Impressions
- To leverage the strengths of a monorepo, it's not just about simply consolidating separately managed repositories into one.
- There seems to be a perspective on whether code unification is necessary in a monorepo.
  - Since it's about managing a single repository, and not about being bound by architectural policies, whether to unify or not should be judged as appropriate.
  - Depending on the architectural policy, such as microservices or modular monoliths, what and how much to unify or not is up to the policy, and unification itself is not mandatory, but it seems that some level of unification is a prerequisite for pipelines like build and deploy.
- The difficulty of operation seems to change significantly depending on whether it's operated by one team or multiple teams. Especially in the latter case, if regulations are not properly organized, it might easily become a big ball of mud.
  - In the case of multiple teams, how to handle and lead cross-cutting concerns (like CI pipeline management) might become an organizational issue.
- I felt that the first consideration is what unit or scope to make a monorepo.
  - It seems good to make a monorepo with a configuration that can be developed in a full cycle, but whether it's one service (a specific system), one product, or domain unit, etc., needs careful consideration.
- I was a bit curious whether transitioning from monorepo to polyrepo or from polyrepo to monorepo involves similar struggles.
  - Probably, dismantling an integrated CI pipeline seems more troublesome, so I think transitioning from monorepo to polyrepo might be more challenging...

# References
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
