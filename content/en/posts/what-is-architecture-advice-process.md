---
title: What is the Architecture Advice Process (AAP)?
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

There is a summary document compiled easily in LT, so please refer to the following if you want to know the summary.

cf. [Researching AAP](https://speakerdeck.com/bmf_san/aapnituitediao-betemita)

## What is the Architecture Advice Process?

In modern software development, while team autonomy and development speed are emphasized, consistency in architecture is also required. The Architecture Advice Process (AAP) responds to these conflicting demands. AAP is a Technique/Trial published in Thoughtworks Technology Radar 2025-04.

AAP is a decentralized approach where **anyone can make architectural decisions**. However, it is a condition that **advice must be sought from those affected** and **relevant experts** before making a decision. Importantly, this is **not an approval process**. While advice is sought, there is no need to obtain permission.

In traditional centralized architecture management, a small number of architects made all decisions. However, AAP decentralizes decision-making authority, allowing those affected and relevant experts to be involved through "advice." This optimizes decision-making authority while involving the appropriate number of people, building trust, and replacing delays waiting for approval with necessary advisory conversations.

The core of AAP is simple. It consists of **one rule** and **one constraint**. The rule is that "anyone can make architectural decisions." The constraint is to "consult two groups before making a decision." First, all people significantly affected by that decision. Second, experts in the decision domain.

## Comparison with Architecture Review Board

Traditional architecture governance is often conducted by a centralized organization known as the Architecture Review Board (ARB). The ARB is a governance body responsible for reviewing and approving the architectural aspects of software development projects within the organization.

The main purpose of the ARB is to ensure that enterprise applications are designed, developed, and maintained in alignment with the organization’s strategic goals, standards, best practices, and architectural guidelines. Typically, it consists of experienced architects, senior developers, and key stakeholders who deeply understand the organization’s business and technical requirements.

The ARB indeed provides advantages such as ensuring consistency, risk mitigation, alignment with business goals, quality assurance, and knowledge sharing. By applying architectural standards, patterns, and best practices, it ensures consistency in application development across the organization and facilitates collaboration, maintenance, and software development among teams.

However, centralized architectural review has significant issues. It creates queues waiting for approvals, which can **hinder the flow of value**. In fact, the State of DevOps Report states that external change approvals (such as Change Advisory Boards: CAB) are "counterproductive" and negatively correlate with software delivery performance, while there is no evidence that formal approval processes lead to lower change failure rates. Thoughtworks' Technology Radar presents AAP as a practical alternative to this situation. It evaluates that the decentralized approach of "anyone can decide, but advice must be obtained from stakeholders and experts" can **optimize flow without compromising quality**.

cf. [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)

## Rules of the Architecture Advice Process

The implementation of AAP is very simple. Its core is summarized in one rule and one constraint.

**Rule: Anyone can make architectural decisions**

**Constraint: Before making a decision, seek advice from those meaningfully affected and domain experts**

This concept is innovative in that it institutionalizes **conversations rather than meetings** and standardizes **advice gathering** rather than **waiting for approvals**. Decision-makers do not need to wait for approval. What is needed is dialogue with the appropriate people and documentation of that dialogue. There is no obligation to adopt the advice, but it should be listened to and recorded in the ADR.

### Distinction between "Advice" and "Opinion"

In AAP, it is important to distinguish between **advice** and **opinion**. Advice is a recommendation based on experience and evidence, always accompanied by reasoning. In contrast, opinion is a subjective judgment based on preference or weak rationale. Effective AAP requires well-founded advice.

Example: "X does not meet the internal SLO" → advice (with basis) / "I dislike X" → opinion (weak basis)

## Benefits of the Architecture Advice Process

AAP brings various benefits to organizations.

**Elimination of Bottlenecks**: Since decision-making authority lies with the parties (individuals or teams), development does not stop due to veto power. There are no queues waiting for approvals as in the past.

**Increased Speed**: Since **consensus is not a requirement**, progress can be made after collecting advice. Decision-making is expedited as there is no need to obtain agreement from everyone.

**Clarification of Boundaries**: The process of identifying who is affected helps clarify the **responsibility boundaries of the system**. Analyzing the scope of impact also validates the appropriateness of system design.

**Increased Ownership**: The principle that the decision-maker is the executor clarifies accountability. There is a heightened sense of responsibility for one's decisions, leading to more careful and effective decision-making.

**Social Trust**: The mutual agreement of seeking and giving advice becomes part of the organizational culture. This deepens trust among teams and promotes knowledge sharing.

## What the Architecture Advice Process Does Not Aim For

To understand AAP, it is important to clarify what it does not aim for.

**It is not an approval process**: Advice is sought, but there is no need to obtain **permission**. Decision-making authority lies with the decision-maker. Advisors do not have veto power.

**It is not consensus-building**: Even if there are opposing opinions, they can be **recorded**, and necessary guardrails (such as phased releases) can be established to **move forward**. It is not a process that seeks consensus.

**It does not eliminate the need for architects**: The role of architects shifts from **decision-makers** to **facilitators**. They provide value through supporting activities such as creating advisory spaces and establishing principles and visualizations.

These characteristics fundamentally differentiate AAP from traditional governance processes.

## Supporting Elements of the Architecture Advice Process

1. **ADR (Architecture Decision Records)**: Document the context, options, trade-offs of decisions, and **who provided what advice and how it was handled**.
2. **Architecture Advisory Forum**: A **space for advice, not an approval body**. Bring ongoing decisions to the forum to publicly gather advice.
3. **Team-sourced Architectural Principles**: **Design principles** developed by the team. The “North Star” of autonomy and alignment.
4. **Internal Tech Radar**: **Visualize** the internal technology portfolio to inform decision-making. (Proposed by Andrew Harmel-Law)

cf. [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)

## Metrics for the Architecture Advice Process

To measure the effectiveness of AAP, appropriate metrics are needed. Based on Xapo's practices, the following metrics are effective.

**Decision-making Lead Time**: Measured by the state transitions of ADR (Draft→Proposed→Accepted→**Adopted** [production reflection]). By tracking the time from decision to implementation completion, the efficiency of the process can be evaluated.

**Observability of Advice**: Continuously record advisors, summaries, and dates in the **Advice section** of ADR. This visualizes the transparency and participation of the advisory process.

**Engineering Outcomes**: Check trends alongside **team metrics** such as deployment frequency and lead time. This provides a comprehensive evaluation of the impact of AAP on team performance.

In Xapo's practice, an **Adopted** status was established in ADR to track the time from decision to implementation. This visualizes the entire process until the decision actually brings value.

## Case Study of the Architecture Advice Process

Xapo, an online bank for Bitcoin, combined **DDD/Team Topologies/AAP** and simultaneously launched the **Architecture Advice Forum** and **ADR**. From the first session, they created a **public advisory habit** using “retroactive ADR” based on real cases, involving various teams and InfoSec/regulatory/operations/product. They also migrated ADR operations to Jira, visualizing the **lead time from decision to production reflection**. Innovations to avoid a return to consensus-seeking and to include product strategy have also been reported.

cf. [Decentralizing the Practice of Architecture at Xapo Bank - martinfowler.com](https://martinfowler.com/articles/xapo-architecture-experience.html)

## Patterns of Failure in the Architecture Advice Process

Understanding key patterns that lead to failure in the implementation of AAP is important.

**Good Failures**: Small failures that occur when inexperienced individuals make decisions are actually valuable learning opportunities. Such failures increase transparency and allow for quick identification and correction. Decision-makers can safely reflect on and share learnings during implementation. These failures should be accepted and visualized and shared as learning in the advisory forum.

**Participation Bias**: The most dangerous failure mode is when only a core group participates, and those who should be involved do not. It may seem successful in the early stages, but in reality, diverse perspectives are not utilized. It is important to pay attention to who is contributing and actively seek opinions from less vocal members.

**Bypassing the Process**: Decisions may occur without being addressed in the advisory forum or recorded in ADR. This should be viewed as a learning opportunity, necessitating a willingness to improve the process with those who made the decisions. Various reasons may exist, such as external pressure or misperception of importance.

**Shadow Architecture**: The most dangerous scenario is when architects superficially support AAP while continuing traditional approvals behind the scenes. This nullifies all the benefits of AAP and undermines trust. Architects must commit to being facilitators who engage in the right conversations with the right people at the right time.

## Conclusion

The Architecture Advice Process is a practical answer to the challenge of balancing "speed and quality" in modern software development. It can eliminate the bottlenecks of traditional centralized architecture governance while maintaining architectural consistency and quality.

The success of AAP depends on transforming organizational culture. A shift from a culture of seeking approval to one of seeking advice, from consensus-building to a transparent decision-making process, and from architects as decision-makers to architects as facilitators is necessary.

However, this cannot be achieved overnight. The introduction of appropriate tools, continuous education, and fostering a culture of learning from failures are essential. Particularly, understanding failure patterns such as participation bias and process bypass, and maintaining a continuous improvement mindset is crucial.

AAP is highly compatible with modern development practices such as Agile development and DevOps. This approach, which emphasizes team autonomy while maintaining organizational alignment, is likely to become increasingly important in the future.

## References

### [Architecture advice process - thoughtworks.com](https://www.thoughtworks.com/en-us/radar/techniques/architecture-advice-process)
Thoughtworks Technology Radar 2025-04 presents AAP as a **Trial**. It states that centralized ARB hinders workflows and correlates with low performance, recommending decentralized decision-making of "anyone can decide, but advice must be sought from stakeholders and experts." By combining ADR and advisory forums, it claims that flow can be optimized while maintaining quality, with increasing success stories in regulated industries.

### [Scaling the Practice of Architecture, Conversationally - martinfowler.com](https://martinfowler.com/articles/scaling-architecture-conversationally.html)
A comprehensive AAP guide by Andrew Harmel-Law. It advocates for a shift from centralized approval to decentralized decision-making supported by conversation, detailing the core of AAP and the four supporting elements (ADRs, advisory forums, team-sourced principles, internal Tech Radar). It also presents specific failure patterns and application procedures.

### [Decentralizing the Practice of Architecture at Xapo Bank - martinfowler.com](https://martinfowler.com/articles/xapo-architecture-experience.html)
A report on practices at Xapo, a Bitcoin online bank. It combined DDD, Team Topologies, and AAP, launching the Architecture Advice Forum and ADR simultaneously. By inviting a wide range of stakeholders to weekly forums and enhancing the Advice section of ADR, it achieved alignment and learning. It also introduced operational metrics by managing ADR in Jira and visualizing decision lead times by adding the **Adopted** status.

### [Architecture Advice Process - archicionado.com](https://archicionado.com/p/architecture-advice-process/)
An implementation guide for AAP by Romain Vasseur. Inspired by Harmel-Law's ideas, it presents the purpose set of AAP (building a knowledge ecosystem, transparency and trust, team autonomy, lightweight but effective mechanisms). It shows implementation ideas using reader-prioritized ADR design and GitHub Issue/Discussion (ADR hub, advisory threads, labeling), proposing lightweight and traceable operations.

### [Introducing the Architecture Advice Process - linkedin.com](https://www.linkedin.com/pulse/introducing-architecture-advice-process-lindsey-tibbitts-jwzsc)
An introduction to AAP from a practitioner's perspective by Lindsey Tibbitts. It lists the effects of the decentralized decision-making model that seeks advice rather than approval, such as eliminating bottlenecks, increasing speed without consensus, strengthening system boundaries, increasing ownership, and fostering social trust. It emphasizes the difference between advice and opinion (evidence-based recommendations vs. unfounded preferences) and the importance of a mindset of "seeking knowledge, deciding close to the work, and trusting each other."

### [Facilitating Software Architecture - facilitatingsoftwarearchitecture.com](https://facilitatingsoftwarearchitecture.com/)
A book by Andrew Harmel-Law published by O'Reilly. Centered on "decentralization and empowerment," it serves as a guide to establish architecture as "everyone's work." Chapter 4 organizes concepts, introduction, and operations with supplementary materials like one-pagers on the Advice Process. It includes implementation guides divided into chapters on ADRs, advisory forums, principles, and Tech Radar, presenting a philosophy of starting with a minimal set and nurturing through documentation and space.

### [DORA: Accelerate — State of DevOps 2019（PDF） - services.google.com](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf)
An annual report by DevOps Research and Assessment (DORA). It demonstrates that external change approvals (such as Change Advisory Boards: CAB) are "counterproductive" and negatively correlate with software delivery performance, providing theoretical support for AAP.