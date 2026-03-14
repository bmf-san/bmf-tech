---
title: Key Points to Consider When Writing an ADR
description: 'Master ADR writing by focusing on atomic decisions, documenting trade-offs, team discussions, and technical writing principles for architecture.'
slug: adr-writing-key-points
date: 2026-01-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Decision Record
  - Design
  - Architecture
translation_key: adr-writing-key-points
---

ADR (Architecture Decision Record) is a document used to record important decisions regarding software architecture.

Even if the content to be written in an ADR is decided, when actually writing it, the content may vary from person to person, or one might not know what to write at all, leading to its formalization.

Based on my experience, I have summarized four important points.

## 1: Focus on One Decision (Atomicity)

The first and most important rule is to **write only one decision per ADR**. 

You should not combine multiple topics into a single document, such as "Choosing a Database and Authentication Infrastructure." If in the future you need to change only the database, managing whether to mark the entire document as "Superseded" or partially modify it becomes very complex.

Additionally, if multiple decisions are mixed, during reviews, if someone says, "I agree with the DB but disagree with the authentication," consensus can come to a halt.

If you feel the urge to write "About A and B" in the title, that is a sign to split it into two ADRs.

## 2: Focus on What to Decide (Decision Point)

An ADR is not a manual for the system; it is a record of **decision points**. 

There is a saying in the second law of software architecture that states, **"The 'Why' is more important than the 'How.'"** The implementation details (How) can be seen in the code. What should be left in the ADR are the "reasons for the choices" and the "discarded options" that cannot be inferred from the code.

The section on "Consequences" is particularly important. It is necessary to record not only the benefits but also the **drawbacks (trade-offs) and risks** that arise from that decision.

Decisions like, "To prioritize development speed, we will sacrifice some consistency" will become valuable information in the future.

## 3: Don’t Write Immediately; Go Through Outline and Discussion Processes

An ADR is a "proof of decision" that solidifies team discussions and agreements, and should not be written immediately. It is important to first discuss with the team using a whiteboard or lightweight documents to form a consensus.

Even when writing the ADR, before formalizing it, it is good to first jot down the points to be covered in bullet points and start the review from an outline base.

Going through the process of divergence and convergence in discussions enhances the overall understanding and agreement within the team, improving the quality of the ADR itself.

## 4: Practice Technical Writing

Since an ADR is a technical document, applying **technical writing principles** is effective.

Reference: [developers.google.com - Technical Writing](https://developers.google.com/tech-writing)

This is not limited to ADRs, but it is important to ensure that the quality of the writing remains consistent, regardless of who writes the technical document.

## Caution Regarding the Flexibility of ADR Formats

There are various formats for ADRs, but the format proposed by Nygard is highly flexible, and particularly the "Consequences" section can be confusing regarding what to write. When customizing the format to fit the organization or team, it is necessary to be aware of the **structure that does not overlook important decision-making axes**.

For example, since trade-offs are at the core of decision-making, establishing a separate section for "Trade-offs" ensures that the examination of merits and demerits does not become formalized, and the basis for judgment remains clear.

## Conclusion

I introduced four points to ensure that ADRs do not end up as mere "records" but become assets that support future decision-making.

1. **Focus on One Decision** — Avoid complexity in management and review.
2. **Record the 'Why'** — Leave the reasons for choices and trade-offs that do not appear in the code.
3. **Write After Discussion** — Solidify as a team decision through the consensus-building process.
4. **Use Technical Writing** — Convey facts in a concise and structured format.

Accurately recording decision-making will become an asset for the future and help deepen the team's consistency and understanding.