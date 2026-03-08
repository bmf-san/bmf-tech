---
title: About Monorepos
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
Summarizing monorepos.

# What is a Monorepo?
A monorepo is a single repository that manages the code for multiple projects. In contrast, managing multiple repositories is referred to as a polyrepo or multi-repo. While it is one of the management policies for microservices, it is not necessarily based on microservices. It is not synonymous with a monolith.

# Aspects of Monorepos
I organized the aspects of operating a monorepo.

## Code Autonomy
When operated by multiple teams, code outside a team's jurisdiction can be modified.

In GitHub, management can be organized through CODEOWNERS to define jurisdiction. It seems necessary to establish some regulations relying on certain tools.

## Big Mudball
The complexity of dependencies between codes can lead to a mudball effect.

This also seems to require some tool-based solutions. For example, Nx allows the creation of libraries with public APIs and visualizes dependencies in a graph.

## Scalability
As the amount of code increases, builds, tests, and deployments can slow down, and issues may arise in Git management.

For the former, this can be resolved by establishing individually executable CI pipelines and deployment flows. This is within the range of tool-based solutions.

For the latter, it may be a bit of a dilemma. When clone or pull becomes painful, it seems necessary to consider some response.

Using GVFS, a scalable Git developed by Microsoft, leveraging Git LFS, or reluctantly splitting the repository are options.

## Freedom of Technology Choice
There are generally no constraints tied to specific technologies (such as programming languages).

Since code management is singular, CI/CD, etc., are assumed to manage multiple pipelines, it does not seem to be a particular concern.

However, if there are languages or environments that the build tools do not support, the range of technology choices may be limited.

## Branch Strategy
The compatibility with feature branches is not very good, so it seems desirable to adopt a trunk-based strategy. It would also be good to introduce feature toggles.

## Other Operational Concerns
In GitHub, it may be necessary to pay attention to the operational policies for Issues and Pull Requests.

# Advantages and Disadvantages of Monorepos
## Advantages
- Easier to grasp the overall picture
  - By managing everything in one place, it is easier to observe the entirety of the code related to projects and services.
- Prevention of Silos
  - When multiple teams operate a monorepo, the sharing cost between teams is lower than with polyrepos. (Probably)
- Easier to Reuse
  - Code reuse and unification are easier than with polyrepos.
- Easier to Promote Standardization
  - Since the operational policies of the repository can be aligned between teams, governance can be more effective.
  - I personally feel this is a significant advantage.
    - In development organizations based on polyrepos, when teams take on a form of autonomy where operational policies slightly differ for each repository, the learning cost during personnel movement between teams can become an overhead issue. However, with a monorepo, it seems possible to reduce that overhead.
    - In polyrepos, when wanting to make similar changes across repositories (for example, when wanting to make similar changes in repositories using a specific CI service for security reasons), it can be cumbersome, but with a monorepo, the burden seems to be reduced.

## Disadvantages
- Maintenance Cost of Common Management Parts
  - CI/CD pipelines need to be operated effectively.
    - There is a risk of complexity.
- Management of Dependencies
  - While visualizing dependencies makes it easier to understand them, if dependencies hinder the ease of code modification, development productivity may decline.
- Scalability of Git
- Dependency on Build Tools and Catch-Up Costs
  - It seems that introducing some build tool will be a prerequisite for starting a monorepo.
  - This may not be a disadvantage per se, but whether or not it is concerning that one more tool is added for repository management.
  - Depending on the tool, the learning cost may not be negligible.
    - For example, adopting Bazel may require a Bazel expert.
    - If developers cannot easily interact with the tool, operations may collapse.
    - There have even been cases where Bazel was removed from Kubernetes projects...
      - [Remove Bazel](https://github.com/kubernetes/kubernetes/pull/99561)
- Ecosystem Compatibility
  - For example, there may be cases where IDEs do not fully support the build tools.

# Tools for Monorepos
Detailed organization can be found at [monorepo.tools - Many solutions, for different goals](https://monorepo.tools/#tools-review).

It seems that Bazel, Nx, and Pants are strong candidates.

# Thoughts
- To leverage the strengths of monorepos, it is not simply about consolidating separately managed repositories into one.
- In monorepos, there seems to be a perspective on whether code standardization is necessary.
  - While managing code in a single repository does not mean being bound by architectural policies, I think the decision on whether to standardize should be made as appropriate.
  - Depending on architectural policies like microservices or modular monoliths, what and how much to standardize is up to the policy, and standardization itself is not mandatory. However, there seems to be a premise for some degree of standardization in pipelines like builds and deployments.
- The difficulty of operations seems to change significantly depending on whether it is operated by one team or multiple teams. Especially in the latter case, if regulations are not properly established, it is likely to become a mudball state.
  - In cases involving multiple teams, how to handle cross-cutting concerns (like CI pipeline management) and how to lead them may become organizational challenges.
- I felt that the initial consideration should be about what unit or scope to use for the monorepo.
  - I think it would be good to structure the monorepo in a way that allows full-cycle development, but whether that is for one service (a specific system), one product, or at the domain level, careful consideration is needed regarding the structural unit.
- I was curious whether transitioning from a monorepo to a polyrepo or vice versa involves similar struggles.
  - I think dismantling an integrated CI pipeline may be more troublesome, so transitioning from a monorepo to a polyrepo might be more challenging...

# References
- [monorepo.tools - monorepo.tools](https://monorepo.tools/)
- [circleci.com - Advantages vs Disadvantages of Monorepo Development](https://circleci.com/ja/blog/monorepo-dev-practices/)
- [www.graat.co.jp - Misconceptions about Monorepos: Monorepo != Monolith Translated](https://www.graat.co.jp/blogs/ck1099bcoeud60830rf0ej0ix)
- [zenn.dev - What is a Monorepo? Summarizing its relationship with related architectures](https://zenn.dev/burizae/articles/c811cae767965a)
- [hireroo.io - Strategies and Operations for Developing Microservices with Monorepos](https://hireroo.io/journal/tech/mono-repo-for-microservices)
- [gist.github.com - About Monorepos](https://gist.github.com/pipopotamasu/efe7097454d9668f80cd8b43068afafc)
- [blog.ojisan.io - Should we use a Monorepo or split the repository?](https://blog.ojisan.io/monorepo-vs-polyrepo/)
- [note.com - Development and Operations of Microservices through Monorepos](https://note.com/tinkermodejapan/n/nb14009fe837f)
- [caddi.tech - Introduction to Monorepos in AI Organizations](https://caddi.tech/archives/4187)
- [cam-inc.co.jp - What We Learned from Operations! Projects Suitable for Monorepos](https://cam-inc.co.jp/p/techblog/570556215432577985)
- [tech.asoview.co.jp - The Story of Migrating 120 Repositories to One Monorepo in 3 Months](https://tech.asoview.co.jp/entry/2022/12/23/095914)
- [qiita.com - Monorepos and Microservices](https://qiita.com/ytanaka3/items/6d8d960179bc046e38c0)
- [cloudsmith.co.jp - [Monorepo] React+Node.js+Typescript Monorepo Construction Memo](https://cloudsmith.co.jp/blog/frontend/2023/06/2396016.html)
- [engineering.mercari.com - Experience of Monorepo Development at Mercari Shops](https://engineering.mercari.com/blog/entry/20210817-8f561697cc/)
- [times.hrbrain.co.jp - Consolidated Nearly 30 Repositories into One Repository](https://times.hrbrain.co.jp/entry/monorepo)
- [docs.aws.amazon.com - Monorepo Build Configuration](https://docs.aws.amazon.com/ja_jp/amplify/latest/userguide/monorepo-configuration.html)
- [postd.cc - Advantages of Monolithic Version Control](https://postd.cc/monorepo/)
- [kk-web.link - Monorepo? No Thanks](https://kk-web.link/blog/20210507)
- [www.atlassian.com - Challenges and Tips for Handling Monorepos in Git](https://www.atlassian.com/ja/blog/monorepos-in-git)
- [cybozu.github.io - Does the Frontend Monorepo Structure Dream of Scaling?](https://cybozu.github.io/frontend-expert/posts/considerations-for-monorepo)
- [kiyobl.com - What is a Monorepo? | Monorepo Development Using Yarn Workspace](https://kiyobl.com/monorepo-basic/)
- [speakerdeck.com - Development and Operations of Microservices Architecture through Monorepos](https://speakerdeck.com/bananaumai/monorehoniyorumaikurosahisuakitekutiyanokai-fa-yun-yong)