---
title: Key Points to Keep in Mind When Writing ADRs
slug: adr-writing-key-points
date: 2026-01-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Decision Record
  - Design
  - Architecture
description: Learn essential tips for writing effective Architecture Decision Records (ADRs) to document critical software architecture decisions.
translation_key: adr-writing-key-points
---

ADR (Architecture Decision Record) is a document used to record important decisions related to software architecture.

Even if the content of an ADR is predetermined, writing it can lead to inconsistencies depending on the person, or it may become unclear what should be written, resulting in the document losing its purpose.

Based on my experience, I’ve summarized four key points that are crucial.

## 1: Focus on a Single Decision (Atomicity)

The first and most important rule is to **write only one decision per ADR**.

Do not combine multiple topics into a single document, such as "Selecting a database and authentication platform." If, in the future, only the database needs to be changed, it becomes overly complex to manage whether to mark the entire document as "Superseded" or partially revise it.

Additionally, mixing multiple decisions can lead to situations during review where opinions like "I agree with the database choice but disagree with the authentication" arise, halting consensus-building.

If you feel tempted to write "About A and B" in the title, that’s a sign to split it into two ADRs.

## 2: Focus on "What to Decide" Over "What to Write" (Decision Point)

An ADR is not a system manual. It’s a record of **decision points**.

The second law of software architecture, often discussed in foundational texts, states that **"Why" is more important than "How."** Implementation details (How) can be understood by looking at the code. What should be recorded in an ADR are the reasons behind choices and the options that were discarded, which cannot be inferred from the code.

The "Consequences" section is particularly important. It’s necessary to document not only the benefits but also the **trade-offs and risks** associated with the decision.

For example, a decision like "Prioritizing development speed at the expense of some consistency guarantees" becomes valuable information for the future.

## 3: Don’t Write Immediately—Organize and Discuss First

An ADR serves as a "record of decisions" that solidifies team discussions and agreements, ensuring they aren’t changed later. Instead of jumping straight into drafting, start by discussing with the team using a whiteboard or lightweight documents to build consensus.

When writing an ADR, begin by listing the key points in bullet form and start reviews based on an outline. This process helps improve team understanding and agreement, enhancing the quality of the ADR itself.

## 4: Practice Technical Writing

Since ADRs are technical documents, applying **technical writing principles** is effective.

Reference: [developers.google.com - Technical Writing](https://developers.google.com/tech-writing)

This principle isn’t limited to ADRs; it’s crucial for all technical documentation to ensure consistent quality regardless of the author.

## Notes on the Flexibility of ADR Formats

There are various ADR formats, but the format proposed by Nygard is highly flexible. This flexibility, especially in the "Consequences" section, can lead to uncertainty about what to write. When customizing formats for your organization or team, ensure the structure doesn’t overlook **key decision-making axes**.

For example, since trade-offs are central to decision-making, creating a dedicated "Trade-offs" section can prevent the evaluation of pros and cons from becoming superficial and ensures the rationale behind decisions is clearly documented.

## Conclusion

To ensure ADRs are not merely "records" but assets that support future decision-making, I’ve introduced four key points:

1. **Focus on a single decision** — Avoid complexity in management and reviews.
2. **Record "Why"** — Document reasons and trade-offs that aren’t visible in the code.
3. **Discuss before writing** — Go through the process of consensus-building to solidify team decisions.
4. **Use technical writing** — Communicate facts concisely and in a structured format.

Accurately documenting decisions helps build future assets and enhances team consistency and understanding.