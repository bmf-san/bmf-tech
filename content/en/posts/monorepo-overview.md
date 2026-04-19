---
title: "Monorepos Explained: Benefits, Drawbacks, and When to Use Them"
description: 'Understand what monorepos are, the tooling involved (Nx, Turborepo, Bazel), and the practical trade-offs between monorepos and polyrepos for growing engineering teams.'
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
This article summarizes monorepos.

# What is a Monorepo?
A monorepo is a single repository that manages the code for multiple projects. In contrast, managing multiple repositories is referred to as a polyrepo or multi-repo. While it is one of the management policies for microservices, it is not necessarily predicated on microservices. It is not synonymous with a monolith.

# Aspects of Monorepos
I have organized the aspects of operating a monorepo.

## Code Autonomy
When operated by multiple teams, code outside a team's jurisdiction can be modified.

On GitHub, management through CODEOWNERS can help establish jurisdictional boundaries. It seems necessary to establish some regulations relying on a tool.

## Large Mudball
The complexity of dependencies between codes can lead to a "mudball" effect.

This also seems to require some tool-based solutions. For example, Nx allows the creation of libraries with public APIs and can visualize dependencies in a graph.

## Scalability
As the amount of code increases, builds, tests, and deployments may slow down, and issues related to Git management may arise.

For the former, this can be resolved by establishing individually executable CI pipelines and deployment flows. This is within the range of tool-based solutions.

The latter may be a bit more challenging. When clone or pull operations become burdensome, it may be necessary to consider some responses.

Using GVFS, a scalable Git developed by Microsoft, leveraging Git LFS, or reluctantly splitting the repository are options.

## Freedom of Technology Choice
There are generally no constraints binding you to specific technologies (such as programming languages).

Since code management is centralized, CI/CD, etc., are assumed to involve managing multiple pipelines, so it does not seem to be a particular concern.

However, if there are languages or environments that the build tools do not support, it may limit the range of technology choices.

## Branch Strategy
The compatibility with feature branches is not very good, so it seems desirable to adopt a trunk-based strategy. It would also be good to implement feature toggles.

## Other Operational Concerns
On GitHub, attention may need to be paid to the operational policies for Issues and Pull Requests.

# Advantages and Disadvantages of Monorepos
## Advantages
- Easier to grasp the overall picture
  - By managing everything in one place, it becomes easier to observe the entirety of the code related to projects and services.
- Prevention of siloing
  - In the case of a monorepo operated by multiple teams, the shared costs between teams are lower than with polyrepos (probably).
- Easier to reuse
  - Code reuse and unification are easier than with polyrepos.
- Easier to promote standardization
  - Since the operational policies of the repository can be aligned among teams, governance can be more effectively enforced.
  - I personally feel this is a significant advantage.
    - In development organizations based on polyrepos, if the operational policies differ slightly for each repository, it can create a learning cost overhead when moving people between teams. However, with a monorepo, this overhead can be reduced.
    - In the case of polyrepos, when you want to make similar changes across different repositories (for example, when you want to make similar changes in repositories using a specific CI service for security reasons), it can be cumbersome, but with a monorepo, the burden seems to be reduced.

## Disadvantages
- Maintenance costs for common management parts
  - CI/CD pipelines need to be operated effectively.
    - There is a risk of complexity.
- Management of dependencies
  - While visualizing dependencies makes it easier to understand them, if dependencies hinder the ease of code modification, development productivity may decline.
- Scalability of Git
- Dependence on build tools and catch-up costs
  - It seems that the introduction of some build tool will be a prerequisite for starting a monorepo.
  - This may not be a significant disadvantage, but whether the addition of a tool for repository management is a concern or not is something to consider.
  - Some tools may have a learning cost that cannot be overlooked.
    - For example, adopting Bazel may require Bazel experts.
    - If developers cannot easily use the tool, operations may collapse.
    - There have even been cases where Bazel was removed from Kubernetes projects...
      - [Remove Bazel](https://github.com/kubernetes/kubernetes/pull/99561)
- Ecosystem compatibility
  - For example, there may be cases where IDEs do not fully support the build tools.

# Tools for Monorepos
Detailed information is organized at [monorepo.tools - Many solutions, for different goals](https://monorepo.tools/#tools-review).

I have the impression that Bazel, Nx, and Pants are strong candidates.

# Personal Thoughts
- To leverage the strengths of a monorepo, it is not simply a matter of consolidating separately managed repositories into one.
- There seems to be a perspective on whether code unification is necessary in a monorepo.
  - Since it manages code in a single repository, it does not necessarily constrain architectural policies, so I think the decision to unify or not should be made as appropriate.
  - Depending on architectural policies such as microservices or modular monoliths, the decision on what to unify and to what extent is up to the policy, and unification itself is not mandatory. However, it seems that some degree of unification is a prerequisite for pipelines such as builds and deployments.
- The difficulty of operation seems to vary greatly depending on whether it is managed by one team or multiple teams. Especially in the latter case, if regulations are not properly established, it is likely to become a mudball.
  - When multiple teams are involved, how to handle cross-cutting concerns (such as CI pipeline management) and how to lead them will likely become an organizational challenge.
- I felt that determining the unit and scope for the monorepo is an initial consideration.
  - I think it would be good to structure the monorepo in a way that allows for full-cycle development, but whether that is for one service (a specific system), one product, or at the domain level, careful consideration is needed regarding the structural unit.
- I was curious whether transitioning from a monorepo to a polyrepo and vice versa involves similar struggles.
  - I think dismantling an integrated CI pipeline may be more troublesome, so transitioning from a monorepo to a polyrepo might be more challenging...

# References
- [monorepo.tools - monorepo.tools](https://monorepo.tools/)
- [circleci.com - Advantages vs Disadvantages of Monorepo Development](https://circleci.com/ja/blog/monorepo-dev-practices/)
- [www.graat.co.jp - Misconceptions about Monorepos: Translated Monorepo != Monolith](https://www.graat.co.jp/blogs/ck1099bcoeud60830rf0ej0ix)
- [zenn.dev - What is a Monorepo? Summarizing its relationship with related architectures](https://zenn.dev/burizae/articles/c811cae767965a)
- [hireroo.io - Strategies and Operations for Developing Microservices with Monorepos](https://hireroo.io/journal/tech/mono-repo-for-microservices)
- [gist.github.com - About Monorepos](https://gist.github.com/pipopotamasu/efe7097454d9668f80cd8b43068afafc)
- [blog.ojisan.io - Should You Use a Monorepo or Split Repositories?](https://blog.ojisan.io/monorepo-vs-polyrepo/)
- [note.com - Development and Operations of Microservices Using Monorepos](https://note.com/tinkermodejapan/n/nb14009fe837f)
- [caddi.tech - Introduction to Monorepos in AI Organizations](https://caddi.tech/archives/4187)
- [cam-inc.co.jp - What We Learned from Operations! Projects Suitable for Monorepos](https://cam-inc.co.jp/p/techblog/570556215432577985)
- [tech.asoview.co.jp - The Story of Migrating 120 Repositories to One Monorepo in 3 Months](https://tech.asoview.co.jp/entry/2022/12/23/095914)
- [qiita.com - Monorepos and Microservices](https://qiita.com/ytanaka3/items/6d8d960179bc046e38c0)
- [cloudsmith.co.jp - [Monorepo] React+Node.js+Typescript Monorepo Construction Memo](https://cloudsmith.co.jp/blog/frontend/2023/06/2396016.html)
- [engineering.mercari.com - Experience of Monorepo Development at Mercari Shops](https://engineering.mercari.com/blog/entry/20210817-8f561697cc/)
- [times.hrbrain.co.jp - Consolidating Nearly 30 Repositories into One Repository](https://times.hrbrain.co.jp/entry/monorepo)
- [docs.aws.amazon.com - Configuration for Monorepo Builds](https://docs.aws.amazon.com/ja_jp/amplify/latest/userguide/monorepo-configuration.html)
- [postd.cc - Advantages of Monolithic Version Control](https://postd.cc/monorepo/)
- ~~kk-web.link - Monorepo? Not Needed~~
- [www.atlassian.com - Challenges and Tips for Handling Monorepos in Git](https://www.atlassian.com/ja/blog/monorepos-in-git)
- [cybozu.github.io - Does the Frontend Monorepo Structure Dream of Scaling?](https://cybozu.github.io/frontend-expert/posts/considerations-for-monorepo)
- ~~kiyobl.com - What is a Monorepo? | Monorepo Development Using Yarn Workspace~~
- [speakerdeck.com - Development and Operations of Microservices Architecture Using Monorepos](https://speakerdeck.com/bananaumai/monorehoniyorumaikurosahisuakitekutiyanokai-fa-yun-yong)
