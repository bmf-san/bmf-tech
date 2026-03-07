---
title: LLM Prompt Management and Evaluation Infrastructure Using GitHub Agentic Workflows
slug: github-agentic-workflow-llm-prompt-management
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
translation_key: github-agentic-workflow-llm-prompt-management
---

As the development of integrating LLMs into products increases, the incorporation of prompts into the software engineering process is becoming more common. Driven by the desire to "continuously manage the quality of prompts" and "explore approaches to optimize the management and evaluation of prompts," I considered a system that treats prompts with a process equivalent to code.

In this article, I will introduce the architecture and implementation of a prompt management and evaluation infrastructure using **GitHub Agentic Workflows (gh-aw)**.

## What is GitHub Agentic Workflows?

[GitHub Agentic Workflows](https://github.github.com/gh-aw/) is a system where **natural language instructions** written in `.github/workflows/*.md` are interpreted and executed by GitHub Copilot, Claude by Anthropic, OpenAI Codex, and others.

This is easier to understand when contrasted with traditional GitHub Actions (which define procedures in YAML).

**Traditional GitHub Actions:**

```yaml
- name: Run tests
  run: pytest tests/
```

**GitHub Agentic Workflows:**

```markdown
---
on:
  pull_request:
permissions:
  contents: read
engine: copilot
---

## Step 1: Run Tests
1. Run all Python tests in the tests/ directory
2. Summarize the results in a report format
3. Comment on the PR if there are any failed tests
```

You declare triggers, permissions, and allowed operations (Safe Outputs) in YAML front matter, and write natural language instructions in Markdown below. The Copilot agent reads and executes these instructions.

The `.md` file is compiled into `.lock.yml` (the actual working GitHub Actions YAML file) using the `gh aw compile` command.

## Architecture Overview

The directory structure is as follows:

```
.
├── .github/
│   └── workflows/
│       ├── evaluate-prompts.md        # Instructions for the Copilot agent (written by humans)
│       └── evaluate-prompts.lock.yml  # Compiled workflow (automatically generated)
├── prompts/                            # Prompts to be evaluated
│   └── code-review.md
│
└── tests/
    └── golden-dataset.yaml            # Evaluation criteria dataset
```

## Implementation of Each Component

### 1. Prompt Files (`prompts/`)

Prompts are managed in Markdown files.

```markdown
# Code Review Prompt

You are an experienced software engineer.
Please review the submitted code based on the following criteria.

### Review Criteria
1. **Accuracy**: Are there any bugs in the logic?
2. **Readability**: Are the naming and structure clear?
3. **Security**: Are there any concerns about vulnerabilities?
...
```

### 2. Golden Dataset (`tests/golden-dataset.yaml`)

The core of the evaluation is the **Golden Dataset**. It defines the evaluation criteria, such as "For this input, what quality of output is expected?"

```yaml
test_cases:
  - id: "code-review-001"
    category: "code_review"
    description: "Can it correctly review code that contains bugs?"
    input:
      user_message: "Please review the following code."
      context: |
        def divide(a, b):
            return a / b
    expected_output:
      criteria:
        - name: "Problem Identification"
          description: "Can it point out the zero division issue?"
          weight: 0.4
        - name: "Improvement Suggestions"
          description: "Can it suggest appropriate fixes?"
          weight: 0.4
        - name: "Clarity of Explanation"
          description: "Can it explain the issues clearly?"
          weight: 0.2
```

Each test case defines **evaluation criteria (criteria) and weights (weight)**. This clarifies the axes for the LLM to evaluate the output (LLM-as-Judge).

### 3. Instruction Workflow (`evaluate-prompts.md`)

The orchestration of the evaluation is defined in the natural language instruction. The structure consists of five steps.

```markdown
---
on:
  pull_request:
    paths:
      - 'prompts/**'

safe-outputs:
  add-comment:
    target: triggering
  add-labels:
    allowed: [needs-improvement, prompt-evaluation]

engine: copilot
---

### Step 1: Identify Changed Prompt Files
1. Identify files changed in the Pull Request:
   gh pr view ${{ github.event.pull_request.number }} --json files \
     --jq '.files[].path' | grep '^prompts/.*\.md$'

### Step 2: Load Golden Dataset
1. Load tests/golden-dataset.yaml
2. Check the structure of each test case

### Step 3: Evaluate Prompts
For each test case:
- Use the content of the prompt as the system prompt
- Generate output and score each criterion from 1-5

### Step 4: Generate Evaluation Report
Generate a report in Markdown format

### Step 5: Output Results
- Post as a comment on the PR (add-comment)
- If the score is below 3.0: add needs-improvement label
```

The `safe-outputs` configuration restricts the operations the agent can perform to only "posting comments on PRs" and "adding specific labels." Direct changes to code or pushes to the main branch are not permitted.

## Evaluation Flow

The flow from creating a PR to receiving feedback is illustrated below.

```mermaid
sequenceDiagram
    participant Dev as Developer
    participant PR as Pull Request
    participant GHA as GitHub Actions
    participant Agent as Copilot Agent

    Dev->>PR: Prompt changes
    PR->>GHA: Trigger workflow
    GHA->>Agent: Start agent
    Agent->>Agent: Read evaluate-prompts.md
    Agent->>Agent: Read golden-dataset.yaml
    Agent->>Agent: Execute each test case (LLM-as-Judge)
    Agent->>Agent: Generate report
    Agent->>PR: Post PR comment (add-comment)
    Agent->>PR: Add label (add-labels)
    PR->>Dev: Notification
```

The evaluation report is posted as a Markdown comment on the PR. If the score is below 3.0/5.0, the `needs-improvement` label is automatically applied.

## Security Design

The permission settings are designed to be minimal.

```yaml
permissions:
  contents: read  # Read-only access to the repository
```

Posting comments on PRs and adding labels are done via `safe-outputs`, so `pull-requests: write` is not needed. Even if the agent malfunctions, it cannot rewrite code or push branches.

## Application: Expanding to Other AI Assets

This design is centered around the idea of treating prompts with the same quality management process as code. The same mechanism can be applied to various AI assets.

| Target                               | Items to Place in `prompts/`         | Example Evaluation Criteria      |
| ------------------------------------ | ------------------------------------- | ------------------------------- |
| General AI Prompts                   | system prompt, few-shot examples      | Output quality, adherence to constraints |
| Claude skills/tools                  | tool's `description`, `input_schema` | Is the tool correctly selected? |
| MCP Server Descriptions              | tool/resource's description           | Can the LLM interpret and invoke it correctly? |
| RAG Query Templates                  | query generation prompts              | Search accuracy, recall rate    |
| Chatbot System Prompts               | system prompt                         | Persona, adherence to constraints |

The file format and evaluation logic for the evaluation targets can be flexibly adjusted by simply changing the natural language instructions in `evaluate-prompts.md`.

## Conclusion

By using GitHub Agentic Workflows, a workflow that can be described as Continuous AI—where "prompts can be reviewed in PRs and merged if they pass automated tests"—can be achieved with only natural language instructions, without writing a single line of YAML.

As prompt engineering becomes an organizational activity, the importance of such quality management infrastructure will increase. The architecture introduced here is just one example, but leveraging the flexibility of GitHub Agentic Workflows, there is potential for application in quality management of various AI assets.