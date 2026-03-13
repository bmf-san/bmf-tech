---
title: Considering Engineering in the Age of AI
description: 'Establish prompt governance with Golden Dataset validation, LLM-as-Judge evaluation, and prompt injection security.'
slug: engineering-in-ai-reflections
date: 2026-02-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - LLM as Judge
  - Prompt as Code
  - Prompt Ops
  - Context Engineering
  - Prompt Engineering
  - Continuous AI
translation_key: engineering-in-ai-reflections
---

As the utilization of large language models (LLMs) expands from individuals to organizations, the need to manage prompts with the same rigor as traditional source code is increasing.

This article discusses three organizational practices: codifying prompts, quantitative evaluation using a Golden Dataset, and establishing a security governance framework, along with the transformation of the engineer's role.

We will organize thoughts on transforming AI from a "personal convenience tool" to a "trusted asset for organizations."

## Prompt as Code: Three Pillars of Organizational Practice

### 1. Codifying Prompts and Central Management

Prompts are "sets of instructions for AI" and are critical assets that determine system behavior. They need to be managed in Git, just like source code, rather than in Excel or document tools.

**Implementation Pattern:**

```yaml
# prompts/customer-support.yaml
version: 1.2.0
description: Response generation for customer support
template: |
  You are the customer support representative for {{company_name}}.
  Please respond to the following inquiry in a {{tone}} manner.

  Inquiry: {{user_question}}

  Please keep your response within {{max_length}} characters.

variables:
  company_name: "Company XX"
  tone: "polite and friendly"
  max_length: 200
```

By defining prompts in YAML or JSON format and allowing variables to be injected from external sources, the following benefits can be achieved:

- **Reusability**: Use of the same prompt in multiple places
- **Testability**: Automated testing with variable changes
- **Change History**: Change tracking via Git
- **Review**: Quality control through pull requests

### 2. Quantitative Evaluation Using Golden Dataset

A system is needed to determine whether "AI output has improved/deteriorated" based on data rather than subjective judgment.

**Definition of Golden Dataset:**
A test dataset that includes inquiries that actually occurred in the production environment, typical use cases, and edge cases (boundary conditions).

| Input Example | Expected Output | Evaluation Criteria |
|---------------|----------------|---------------------|
| "I want to return this" | Explanation of return policy + procedure | Accuracy, Tone |
| "When will it be delivered?" | How to check delivery status | Conciseness |
| "This is a complaint!" | Apology + escalation | Appropriate response |

Whenever a prompt is changed, it is essential to run tests with the Golden Dataset to confirm that there is no performance regression. By incorporating this into the CI/CD (Continuous Integration/Continuous Delivery) pipeline, we can prevent quality degradation.

**LLM-as-a-Judge (AI Evaluation):**
It is not realistic for humans to manually check hundreds of outputs. Therefore, using a more powerful LLM as an evaluator has become mainstream.

Example of an evaluation prompt:
```
Compare the following two responses and evaluate which is superior.

Evaluation Criteria:
- Accuracy: Is it based on facts?
- Tone: Is the language appropriate?
- Conciseness: Is it not verbose?

Response A: {{response_a}}
Response B: {{response_b}}

Choose the superior one and explain why.
```

### 3. Establishing Security and Governance Framework

An organizational framework is needed to address the unique security risks of AI systems.

**Key Risks:**

| Risk | Description | Countermeasures |
|------|-------------|-----------------|
| **Prompt Injection** | Overwriting AI instructions with malicious input | Input validation, context separation |
| **Personal Information Leakage** | Improper handling of PII in input/output | Automatic PII detection and masking |
| **Inappropriate Output** | Generation of discriminatory or harmful content | Introduction of guardrail models |

**Organizational Measures:**

1. **Establishment of AI Governance Committee**: A cross-functional approval process involving IT, legal, risk management, and business units.
2. **Clarification of Responsibilities**: Appointment of "AI System Owners" for each AI use case.
3. **Retention of Audit Logs**: Ensuring records of all AI communications and traceability.

**Formalizing Knowledge: AGENTS.md and SKILL.md**

This is a mechanism to document organizational standard processes and teach AI agents.

- **AGENTS.md**: The "constitution" of the project. Defines the roles of agents, available tools, and criteria for judgment.
- **SKILL.md**: Specific workflows. Describes procedures like "PR review process" and "incident response flow."

This prevents individual dependency and streamlines onboarding for new members.

## Transformation of the Engineer's Role: Three Paradigm Shifts

### From Code Writer to Environment Designer

Traditional engineering was about "writing precise instructions (code)." Engineering in the AI era is about "designing an environment where AI can function effectively."

Prompts, data, tools, memory, and evaluation criteria—all of these are components of the "environment":

- **Prompts**: Clarifying instructions to AI
- **Data (RAG)**: Organizing information sources for AI to reference
- **Tools (MCP)**: Defining external functions available to AI
- **Memory**: Designing the information AI should retain (short-term/work/long-term memory)
- **Evaluation Criteria**: Measurement standards for AI output quality (Golden Dataset)

The engineer's role is to appropriately combine these elements to build an "environment" where AI can stably create value. Coding skills remain important, but they are just a subset of the whole.

### From Individual Technical Skills to Organizational Process Design Skills

While the ability of individuals to write excellent prompts is important, what is more essential is the ability to build a system where the entire team can maintain a certain level of quality.

In traditional software development, processes such as code reviews, testing, and CI/CD have been used to prevent individual dependency and ensure quality. A similar approach is necessary in AI development:

- Ensuring Git management of prompts and traceability of change history
- Automating regression testing with the Golden Dataset
- Formalizing organizational knowledge through AGENTS.md and SKILL.md
- Standardizing security checks (PII detection, injection countermeasures)

### From Seeking Perfection to Continuous Improvement

Traditional code is deterministic. The same input always returns the same output. However, AI is probabilistic, and there is no perfect initial setting.

What is important is the ability to rapidly cycle through monitoring, evaluation, and improvement:

1. **Monitoring**: Real-time monitoring of AI output quality (drift detection)
2. **Evaluation**: Regular performance measurement using the Golden Dataset
3. **Improvement**: Adjusting prompts and retesting when issues are detected

Organizations that can execute this "Build-Measure-Learn" loop weekly or daily will gain a competitive advantage.

Moreover, AI systems cannot be completed solely by engineers. Collaboration with legal, compliance, domain experts, and product managers is essential. Ensuring not only technical accuracy but also organizational consensus and selecting tools accessible to non-engineers (such as no-code editing environments like PromptLayer) also becomes an important responsibility for engineers.

## Perspective of Platform Engineering

Practices such as prompt code management and evaluation using Golden Dataset have a high affinity with the concept of platform engineering.

Instead of each person reinventing the wheel, a mechanism should be established to centrally provide verified prompt templates as a "Golden Path" (recommended path), allowing engineers to self-service when needed. This reduces the time spent on individual trial and error and enhances the overall developer experience (DX) of the organization.

## Conclusion: Shift in Engineering Paradigms

Engineering in the AI era is transforming from "code writers" to "designers of environments where AI functions effectively."

**Three Key Transformations:**

1. **Prompts are the new code**: Strict management and testing are essential.
2. **Evaluation is based on data, not intuition**: Golden Dataset and LLM-as-a-Judge.
3. **From individual skills to organizational systems**: Formalization through AGENTS.md/SKILL.md.

**Evolution Path of Engineers:**
- **Users of AI** → **Nurturers of AI** → **Designers collaborating with AI**

By establishing the appropriate organizational framework and enabling individual engineers to acquire necessary skills, it becomes possible to transform LLMs from mere experimental tools into assets that continuously generate business value.