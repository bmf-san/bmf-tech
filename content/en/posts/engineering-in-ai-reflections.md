---
title: Rethinking Engineering in the Age of AI
slug: engineering-in-ai-reflections
date: 2026-02-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - LLM as judge
  - Prompt as code
  - Prompt Ops
  - Context engineering
  - Prompt engineering
  - Continuous AI
description: Exploring the transformation of engineering practices in the era of large language models (LLMs) and organizational AI adoption.
translation_key: engineering-in-ai-reflections
---

As the use of large language models (LLMs) expands from individuals to organizations, the need to manage prompts with the same rigor as traditional source code becomes increasingly important.

This article discusses three organizational practices—prompt codification, quantitative evaluation using a Golden Dataset, and establishing security governance frameworks—as well as the evolving role of engineers.

It organizes thoughts on transforming AI from a "personal convenience tool" to a "reliable organizational asset."

## Prompt as Code: The Three Pillars of Organizational Practices

### 1. Codifying and Centrally Managing Prompts

Prompts are "instruction sets for AI" and represent critical assets that determine system behavior. These should be managed in Git, just like source code, rather than in Excel or document tools.

**Implementation Pattern:**

```yaml
# prompts/customer-support.yaml
version: 1.2.0
description: Response generation for customer support
template: |
  You are a customer support representative for {{company_name}}.
  Please respond to the following inquiry in a {{tone}} manner.

  Inquiry: {{user_question}}

  Please keep the response within {{max_length}} characters.

variables:
  company_name: "Company XYZ"
  tone: "polite and friendly"
  max_length: 200
```

Defining prompts in formats like YAML or JSON and enabling external variable injection provides the following benefits:

- **Reusability**: Use the same prompt in multiple places
- **Testability**: Conduct automated tests by changing variables
- **Change History**: Track changes via Git
- **Review**: Ensure quality through pull requests

### 2. Quantitative Evaluation Using a Golden Dataset

A system is needed to determine whether "AI output has improved or degraded" based on data rather than subjective judgment.

**Definition of a Golden Dataset:**
A test dataset that includes real-world inquiries from production, typical use cases, and edge cases (boundary conditions).

| Input Example       | Expected Output                  | Evaluation Criteria       |
|---------------------|----------------------------------|---------------------------|
| "I want to return this." | Explanation of return policy + steps | Accuracy, tone            |
| "When will my delivery arrive?" | Instructions to check delivery status | Conciseness               |
| "This is a complaint!" | Apology + escalation          | Appropriate response      |

Whenever prompts are updated, tests must be run against the Golden Dataset to ensure no performance regressions occur. Integrating this into CI/CD (Continuous Integration/Continuous Delivery) pipelines prevents quality degradation.

**LLM-as-a-Judge (AI-based Evaluation):**
Manually checking hundreds of outputs is impractical. Therefore, using a more advanced LLM as an evaluator has become a mainstream approach.

Example evaluation prompt:

```
Compare the following two responses and evaluate which one is better.

Evaluation Criteria:
- Accuracy: Is it factually correct?
- Tone: Is the language appropriate?
- Conciseness: Is it free of redundancy?

Response A: {{response_a}}
Response B: {{response_b}}

Select the better response and explain your reasoning.
```

### 3. Establishing Security and Governance Frameworks

Organizational frameworks are required to address the unique security risks of AI systems.

**Key Risks:**

| Risk                  | Description                                  | Countermeasures                    |
|-----------------------|----------------------------------------------|------------------------------------|
| **Prompt Injection**  | Overwriting AI instructions with malicious input | Input validation, context isolation |
| **Data Leakage**      | Mishandling of PII in inputs/outputs         | Automatic PII detection and masking |
| **Inappropriate Output** | Generation of discriminatory or harmful content | Implementation of guardrail models |

**Organizational Measures:**

1. **Establish an AI Governance Committee**: A cross-functional approval process involving IT, legal, risk management, and business units
2. **Define Accountability**: Appoint an "AI System Owner" for each AI use case
3. **Maintain Audit Logs**: Record and ensure traceability of all AI communications

**Formalizing Knowledge: AGENTS.md and SKILL.md**

A system for documenting organizational standard processes and enabling AI agents to learn from them.

- **AGENTS.md**: The "constitution" of the project. Defines agent roles, available tools, and decision-making criteria
- **SKILL.md**: Specific workflows, such as "PR review procedures" or "incident response flows"

This prevents reliance on individual expertise and streamlines onboarding for new team members.

## The Evolving Role of Engineers: Three Paradigm Shifts

### From Writing Code → Designing Environments

Traditional engineering focused on "writing precise instructions (code)." Engineering in the AI era focuses on "designing environments where AI can function effectively."

Prompts, data, tools, memory, and evaluation criteria—all these are components of the "environment":

- **Prompts**: Clarify instructions for AI
- **Data (RAG)**: Organize reference information for AI
- **Tools (MCP)**: Define external functions accessible to AI
- **Memory**: Design the information AI should retain (short-term/working/long-term memory)
- **Evaluation Criteria**: Define metrics for assessing AI output quality (Golden Dataset)

The engineer's role is to combine these elements effectively to build an "environment" where AI can consistently create value. Coding skills remain important but are only a subset of the overall responsibilities.

### From Individual Skill → Organizational Process Design

While individual ability to write excellent prompts is important, the essence lies in building systems that enable the entire team to maintain a consistent level of quality.

In traditional software development, processes like code reviews, testing, and CI/CD have been used to prevent reliance on individual expertise and ensure quality. A similar approach is needed for AI development:

- Git management and change tracking for prompts
- Automated regression testing with a Golden Dataset
- Formalizing organizational knowledge with AGENTS.md and SKILL.md
- Standardizing security checks (PII detection, injection prevention)

### From Striving for Perfection → Driving Continuous Improvement

Traditional code is deterministic, always returning the same output for the same input. However, AI is probabilistic, and perfect initial settings do not exist.

What matters is the ability to rapidly iterate through monitoring, evaluation, and improvement cycles:

1. **Monitoring**: Real-time monitoring of AI output quality (drift detection)
2. **Evaluation**: Regular performance measurement using the Golden Dataset
3. **Improvement**: Adjusting prompts and retesting when issues are detected

Organizations that can execute this "Build-Measure-Learn" loop on a weekly or daily basis will gain a competitive edge.

Moreover, AI systems cannot be managed by engineers alone. Collaboration with legal, compliance, domain experts, and product managers is essential. Beyond technical accuracy, engineers are also responsible for organizational consensus-building and selecting tools accessible to non-engineers (e.g., no-code editing environments like PromptLayer).

## A Platform Engineering Perspective

Practices such as prompt code management and evaluation using a Golden Dataset align closely with the principles of platform engineering.

Instead of reinventing the wheel, verified prompt templates can be centrally provided as a "Golden Path" for engineers to use via self-service when needed. This reduces time spent on individual trial-and-error and improves the overall developer experience (DX) across the organization.

## Conclusion: The Paradigm Shift in Engineering

Engineering in the AI era is transforming from "writing code" to "designing environments where AI can function effectively."

**Three Key Transformations:**

1. **Prompts are the new code**: Rigorous management and testing are essential
2. **Evaluation must be data-driven, not subjective**: Golden Dataset and LLM-as-a-Judge
3. **From individual skills to organizational systems**: Formalizing knowledge with AGENTS.md/SKILL.md

**The Evolution Path for Engineers:**
- **AI Users** → **AI Trainers** → **AI Co-Designers**

By establishing the right organizational frameworks and equipping engineers with the necessary skills, organizations can transform LLMs from experimental tools into assets that continuously generate business value.