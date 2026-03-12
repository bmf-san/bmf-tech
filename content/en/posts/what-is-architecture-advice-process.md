---
title: What is the Architecture Advice Process (AAP)?
description: An in-depth exploration of What is the Architecture Advice Process (AAP)?, covering design principles, trade-offs, and practical applications.
slug: what-is-architecture-advice-process
date: 2025-08-16T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Advice Process
  - Architecture Review Board
  - Architecture
translation_key: what-is-architecture-advice-process
---



There is a brief summary in the LT materials, so if you want to know the gist, please refer to the following.

cf. [Investigating AAP](https://speakerdeck.com/bmf_san/aapnituitediao-betemita)

## What is the Architecture Advice Process

In modern software development, while team autonomy and development speed are prioritized, architectural consistency is also demanded. The Architecture Advice Process (AAP) addresses these conflicting requirements. AAP is a Techniques/Trial listed in the Thoughtworks Technology Radar 2025-04.

AAP is a decentralized approach where **anyone can make architectural decisions**. However, the condition is to **seek advice** from **affected people** and **relevant experts** before making a decision. Importantly, this is **not an approval process**. Advice is sought, but permission is not required.

In traditional centralized architecture management, a few architects made all the decisions. However, AAP decentralizes decision-making authority, involving affected people and relevant experts through "advice." This optimizes decision-making authority while building trust and replacing approval delays with necessary and sufficient advisory conversations.

The core of AAP is simple. It consists of **one rule** and **one constraint**. The rule is that "anyone can make architectural decisions." The constraint is to "consult two groups before making a decision." First, all people significantly affected by the decision. Second, experts in the decision area.

## Comparison with Architecture Review Board

Traditional architecture governance is often conducted by a centralized organization called the Architecture Review Board (ARB). ARB is a governance body responsible for reviewing and approving the architectural aspects of software development projects within an organization.

The main purpose of ARB is to ensure that enterprise applications are designed, developed, and maintained in line with the organization's strategic goals, standards, best practices, and architectural guidelines. It usually consists of experienced architects, senior developers, and key stakeholders who deeply understand the organization's business and technical requirements.

ARB indeed provides benefits such as ensuring consistency, risk mitigation, alignment with business goals, quality assurance, and knowledge sharing. By applying architectural standards, patterns, and best practices, it ensures consistency in application development across the organization, facilitating collaboration, maintenance, and software development.

However, centralized architecture review has significant issues. It creates queues waiting for approval, which can easily **impede the flow of value**. In fact, the State of DevOps Report states that external change approvals (such as Change Advisory Board: CAB) are "counterproductive" and negatively correlated with software delivery performance, while there is no evidence that formal approval processes reduce change failure rates. Thoughtworks' Technology Radar presents AAP as a practical alternative to this situation. By adopting a decentralized approach where "anyone can decide, but advice is sought from stakeholders and experts," it is evaluated as being able to **optimize flow without compromising quality**.

cf. [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)

## Rules of the Architecture Advice Process

The implementation of AAP is very simple. Its core is condensed into one rule and one constraint.

**Rule: Anyone can make architectural decisions**

**Constraint: Before deciding, seek advice from people who will be meaningfully affected and knowledgeable experts in the area**

This concept is innovative in that it institutionalizes **conversation instead of meetings** and makes **advice gathering** the standard process instead of **waiting for approval**. Decision-makers do not need to wait for approval. What is needed is dialogue with the right people and documentation of that. Advice is not mandatory to adopt, but it should be listened to and recorded in the ADR on how it was handled.

### Distinction between "Advice" and "Opinion"

In AAP, it is important to distinguish between **advice** and **opinion**. Advice is a recommendation based on experience and evidence, always accompanied by reasons. On the other hand, an opinion is a subjective judgment based on preference or weak grounds. Effective AAP requires advice with a basis.

Example) "X does not meet internal SLO" → Advice (with basis) / "I dislike X" → Opinion (weak basis)

## Benefits of the Architecture Advice Process

AAP brings numerous benefits to organizations.

**Elimination of Bottlenecks**: Since decision-making authority lies with the parties involved (individuals or teams), development is not halted by veto power. There are no queues waiting for approval as in the past.

**Speed Improvement**: **Consensus is not a requirement**, so progress can be made after gathering advice. Decisions are expedited as there is no need to obtain everyone's agreement.

**Healthy Boundaries**: The process of identifying who is affected clarifies the **system's responsibility boundaries**. Analyzing the impact range also verifies the validity of system design.

**Increased Ownership**: The principle that the person who decides = the person who executes clarifies accountability. A sense of responsibility for one's decisions increases, leading to more careful and effective decisions.

**Social Trust**: The mutual agreement of seeking and giving advice becomes part of the organizational culture. This deepens trust between teams and promotes knowledge sharing.

## What the Architecture Advice Process Does Not Aim For

To understand AAP, it is important to clarify what it does not aim for.

**Not an Approval Process**: Advice is sought, but there is no need to obtain **permission**. Decision-making authority lies with the decision-maker. Advisors do not have veto power.

**Not Consensus Building**: Even if there are opposing opinions, they should be **recorded**, and necessary guardrails (such as phased releases) should be established to **move forward**. It is not a process that seeks consensus.

**Not an Architect Elimination**: The role of architects shifts from **decision-makers** to **facilitators**. They add value through support tasks such as creating advisory opportunities, organizing principles, and visualization.

These characteristics make AAP a fundamentally different approach from traditional governance processes.

## Supporting Elements of the Architecture Advice Process

1. **ADR (Architecture Decision Records)**
   Record the context, options, and trade-offs of decisions, as well as **who gave what advice and how it was handled**.
2. **Architecture Advisory Forum**
   **Not an approval body but a place for advice**. Bring ongoing decisions and gather advice openly.
3. **Team-sourced Architectural Principles**
   Design principles developed by the team. The "North Star" of autonomy and alignment.
4. **In-house Tech Radar**
   **Visualize** the internal technology portfolio and use it as a basis for decision-making.
   (Andrew Harmel-Law's proposal set)

cf. [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)

## Outcome Indicators of the Architecture Advice Process

To measure the effectiveness of AAP, appropriate indicators are necessary. Based on Xapo's practice, the following indicators are effective.

**Decision Lead Time**: Measure the state transition of ADR (Draft→Proposed→Accepted→**Adopted** [production reflection]). By tracking the time from decision-making to implementation completion, the efficiency of the process can be evaluated.

**Observability of Advice**: Continuously record advisors, summaries, and dates in the **Advice section** of ADR. Visualize the transparency and participation of the advisory process.

**Engineering Outcomes**: Check trends in **team indicators** such as deployment frequency and lead time. Evaluate the impact of AAP introduction on team performance comprehensively.

In Xapo's practice, an **Adopted** status was added to ADR to track the process from decision-making to implementation. This visualizes the entire process until the decision actually brings value.

## Case Study of the Architecture Advice Process

Xapo, an online bank for Bitcoin, combined **DDD/Team Topologies/AAP** and simultaneously launched the **Architecture Advice Forum** and **ADR**. From the first session, they created a **public advisory habit** by using real projects as subjects and involving teams, InfoSec/regulation/operations/product. They also transferred ADR operations to Jira, visualizing **lead time from decision to production reflection**. They reported learnings such as **avoiding a return to consensusism** and **product strategy participation**.

cf. [Decentralizing the Practice of Architecture at Xapo Bank - martinfowler.com](https://martinfowler.com/articles/xapo-architecture-experience.html)

## Patterns of Failure in the Architecture Advice Process

Understanding the main patterns that lead to failure in AAP implementation is important.

**Good Failures**: Small failures that occur when inexperienced people make decisions are actually valuable learning opportunities. These failures increase transparency, allowing for quick identification and correction. Decision-makers recognize issues during implementation, enabling safe reflection and sharing of learning. These failures should be accepted and visualized/shared as learning in the advisory forum.

**Participation Bias**: The most dangerous failure mode is when only a core group participates, and those who should be involved do not. It may seem successful in the early stages, but it actually fails to leverage diverse perspectives. It is important to pay attention to who is contributing and actively seek input from less vocal members.

**Process Bypass**: Decisions may occur without being brought up in the advisory forum or recorded in ADR. This should be seen as a learning opportunity, with a stance of improving the process together with those who made the decisions. Various reasons such as external pressure or misperception of importance may exist.

**Shadow Architecture**: The most dangerous is when architects superficially support AAP while continuing traditional approvals behind the scenes. This nullifies all the benefits of AAP and breaks trust. Architects need to commit to their role as facilitators, having the right conversations with the right people at the right time.

## Conclusion

The Architecture Advice Process is a practical answer to the challenge of "balancing speed and quality" in modern software development. It can maintain architectural consistency and quality while eliminating the bottlenecks of traditional centralized architecture governance.

The success of AAP depends on transforming organizational culture. A shift is needed from a culture of seeking approval to one of seeking advice, from consensus-building to a transparent decision-making process, and from architects as decision-makers to architects as facilitators.

However, this cannot be achieved overnight. The introduction of appropriate tools, continuous education, and fostering a culture of learning from failures are essential. It is particularly important to understand and continuously improve failure patterns such as participation bias and process bypass.

AAP is highly compatible with modern development practices such as Agile development and DevOps. This approach, which emphasizes team autonomy while maintaining organizational alignment, will become increasingly important in the future.

## References

### [Architecture advice process - thoughtworks.com](https://www.thoughtworks.com/en-us/radar/techniques/architecture-advice-process)
Thoughtworks Technology Radar 2025-04 lists AAP as a **Trial**. It states that centralized ARB impedes workflow and correlates with low performance, recommending decentralized decision-making where "anyone can decide but advice is sought from influencers and experts." By combining ADR and advisory forums, it states that flow can be optimized while maintaining quality, with increasing success examples even in regulated industries.

### [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)
A comprehensive AAP guide by Andrew Harmel-Law. Declares "I stopped making architectural decisions" and advocates a shift from centralized approval to conversation-supported decentralization. Details the core of AAP and the learning and alignment mechanisms supported by four elements (ADRs, advisory forums, team-sourced principles, in-house Tech Radar). Also presents failure patterns and application procedures concretely.

### [Decentralizing the Practice of Architecture at Xapo Bank - martinfowler.com](https://martinfowler.com/articles/xapo-architecture-experience.html)
A practice report at Xapo, an online bank for Bitcoin. Combines DDD and Team Topologies with AAP, simultaneously launching the Architecture Advice Forum and ADR. Weekly forums inviting a wide range of stakeholders from development to management and enriching the Advice section of ADR to gain alignment and learning. Introduces operational measurement innovations such as managing ADR with Jira and adding an "Adopted" state to visualize decision lead time.

### [Architecture Advice Process - archicionado.com](https://archicionado.com/p/architecture-advice-process/)
An AAP implementation guide by Romain Vasseur. Inspired by Harmel-Law's ideas, presents a set of AAP objectives (building a knowledge ecosystem, transparency and trust, team autonomy, lightweight but effective mechanisms). Proposes implementation ideas using GitHub Issue/Discussion (ADR hub, advice threads, labeling, etc.) for lightweight and traceable operations.

### [Introducing the Architecture Advice Process - linkedin.com](https://www.linkedin.com/pulse/introducing-architecture-advice-process-lindsey-tibbitts-jwzsc)
An introduction to AAP from a practitioner's perspective by Lindsey Tibbitts. Lists the effects of a decentralized decision-making model that seeks advice instead of approval, such as eliminating bottlenecks, speeding up without consensus, strengthening system boundaries, increasing ownership, and fostering social trust. Emphasizes the difference between advice and opinion (evidence-based recommendation vs. baseless preference) and the importance of a stance of "seeking knowledge, deciding close to the work, and trusting each other."

### [Facilitating Software Architecture - facilitatingsoftwarearchitecture.com](https://facilitatingsoftwarearchitecture.com/)
A book published by O'Reilly by Andrew Harmel-Law. Centers on "decentralization and empowerment," providing a guide to making architecture "everyone's job." Systematizes the concept, introduction, and operation of the Advice Process with supplementary materials such as a one-pager in Chapter 4. Includes implementation guides with chapters on ADRs, advisory forums, principles, Tech Radar, and presents an implementation philosophy of starting with a minimal set and nurturing it through records and forums.

### [DORA: Accelerate — State of DevOps 2019 (PDF) - services.google.com](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf)
An annual report by DevOps Research and Assessment (DORA). Demonstrates that external change approvals (such as Change Advisory Board: CAB) are "counterproductive" and negatively correlated with software delivery performance. States there is no evidence that formal approval processes reduce change failure rates, providing theoretical support for AAP.