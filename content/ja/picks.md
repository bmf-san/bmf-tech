---
title: "ピックアップ記事"
description: "設計判断やトレードオフを扱う記事を中心に、長く参照できるものを筆者が選んだ記事の一覧。テーマ別にまとめている。"
slug: picks
date: 2026-03-21T00:00:00Z
author: bmf-san
template: picks.html
draft: false
translation_key: picks
listing_slugs:
  - jose-overview
  - oauth-2-1-security-bcp
  - oidc-discovery-metadata
  - access-token-validation-revocation
  - oidc-logout-security-event-token
  - dynamic-client-registration
  - sender-constrained-tokens-mtls-dpop
  - service-to-service-authentication
  - spa-token-storage-bff
  - platform-engineering-and-systems-thinking
  - systems-thinking-basics
  - accidental-complexity-and-missing-decision-records
  - distributed-transaction-patterns-2pc-saga-outbox
  - abstraction-trap-in-architecture
  - platform-viability-conditions
  - product-development-vs-platform-development
  - engineering-management-as-a-deck
  - migrating-bmf-tech-from-gobel-to-gohan
  - engineering-in-ai-reflections
  - github-agentic-workflow-llm-prompt-management
  - career-portfolio-creation-guide
  - good-bad-architecture-strategies
  - when-why-write-architecture-strategy
  - architecture-strategy-tactics-design
  - requirements-vs-constraints
  - adr-writing-key-points
  - team-topologies-introduction
  - mvv-importance-software-teams
  - platform-engineering-explained
  - terminal-specifications-for-tui-development
  - what-is-architecture-advice-process
  - importance-of-constraints-tradeoffs
  - cache-strategy-patterns
  - web-application-processing-model
  - acid-vs-base
  - cap-theorem-pacelc-theorem
  - reflection-framework-bmf
  - cohesion-coupling
  - access-control-patterns-using-opa
  - ruby-rails-io-cpu-characteristics
  - pagination-impact-solutions-opa
  - design-shelf-life-considerations
  - measuring-scalability
  - writing-architecture-documentation
  - system-specification-structure
  - bounded-context-explanation
  - event-storming-introduction
  - microservices-access-control-opa
  - story-points-using-cynefin-framework
  - architecture-strategy-discussion
  - load-testing-start-guide
  - set-theory
  - custom-http-router-to-new-servemux
  - mysql-locking
  - what-is-index
  - dealing-with-uncertainty-software-projects
  - go-performance-improvement
  - mysql-transaction-anomalies
  - transaction-overview
  - system-design-resources-review
  - bmf-tech-supporting-technologies
  - building-http-router-with-net-http
# listing_slugs の各記事をどのグループに載せるか。1 記事は 1 グループだけ。一致は make check-picks で検証する。
picks_sections:
  - title: "システムの技術"
    groups:
      - title: "アーキテクチャの意思決定"
        description: "何をどう決め、どう残すか"
        slugs:
          - accidental-complexity-and-missing-decision-records
          - good-bad-architecture-strategies
          - when-why-write-architecture-strategy
          - architecture-strategy-tactics-design
          - architecture-strategy-discussion
          - requirements-vs-constraints
          - importance-of-constraints-tradeoffs
          - adr-writing-key-points
          - what-is-architecture-advice-process
          - writing-architecture-documentation
          - system-specification-structure
          - design-shelf-life-considerations
      - title: "設計の原則とモデリング"
        description: "システムの構造をどう分けるか"
        slugs:
          - cohesion-coupling
          - abstraction-trap-in-architecture
          - bounded-context-explanation
          - event-storming-introduction
          - system-design-resources-review
      - title: "データと分散システム"
        description: "データの一貫性・保存・分散"
        slugs:
          - transaction-overview
          - acid-vs-base
          - mysql-transaction-anomalies
          - mysql-locking
          - what-is-index
          - cap-theorem-pacelc-theorem
          - distributed-transaction-patterns-2pc-saga-outbox
          - cache-strategy-patterns
      - title: "認証・認可"
        description: "誰が誰で、何を許すか"
        slugs:
          - jose-overview
          - oauth-2-1-security-bcp
          - oidc-discovery-metadata
          - access-token-validation-revocation
          - oidc-logout-security-event-token
          - dynamic-client-registration
          - sender-constrained-tokens-mtls-dpop
          - service-to-service-authentication
          - spa-token-storage-bff
          - access-control-patterns-using-opa
          - microservices-access-control-opa
          - pagination-impact-solutions-opa
      - title: "性能とスケーラビリティ"
        description: "どれだけ速く、どれだけ多く処理できるか"
        slugs:
          - web-application-processing-model
          - ruby-rails-io-cpu-characteristics
          - measuring-scalability
          - load-testing-start-guide
          - go-performance-improvement
      - title: "実装"
        description: "特定の技術で作った・移行した記録"
        slugs:
          - building-http-router-with-net-http
          - custom-http-router-to-new-servemux
          - terminal-specifications-for-tui-development
          - github-agentic-workflow-llm-prompt-management
          - migrating-bmf-tech-from-gobel-to-gohan
          - bmf-tech-supporting-technologies
  - title: "作る人と組織"
    groups:
      - title: "プラットフォームと組織"
        description: "チームと組織をどう構成するか"
        slugs:
          - platform-engineering-explained
          - platform-viability-conditions
          - product-development-vs-platform-development
          - platform-engineering-and-systems-thinking
          - team-topologies-introduction
          - mvv-importance-software-teams
          - engineering-management-as-a-deck
      - title: "仕事の進め方と考え方"
        description: "個人やチームがどう考え、どう進めるか"
        slugs:
          - systems-thinking-basics
          - set-theory
          - story-points-using-cynefin-framework
          - dealing-with-uncertainty-software-projects
          - reflection-framework-bmf
          - career-portfolio-creation-guide
          - engineering-in-ai-reflections
---

設計判断やトレードオフを扱う記事を中心に、長く参照できるものを筆者が選んでいる。テーマ別にまとめている。
