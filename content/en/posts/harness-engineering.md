---
title: "What is Harness Engineering? Design Patterns for Long-Running AI Agents"
slug: harness-engineering
date: 2026-11-04
author: bmf-san
categories:
  - Architecture
tags:
  - Generative AI
  - AI
  - Context Engineering
  - LLM as Judge
description: "Running AI agents for long, reliable sessions takes more than a good model: it takes a well-designed harness around it. Based on primary sources from Anthropic, OpenAI, and martinfowler.com, this post covers guides and sensors, context, state, task decomposition, human involvement, and evals."
translation_key: harness-engineering
draft: false
---

# Overview

When you hand an AI agent work that takes hours or days, model capability alone stops being enough. Agents run out of context mid-task, declare victory before they finish the work, and repeat the same mistakes. You have to prevent these failures outside the model.

The machinery outside the model is the harness, and harness engineering is the practice of designing that machinery. This post organizes the building blocks and design patterns of a harness, based on material that Anthropic, OpenAI, and martinfowler.com published between December 2024 and April 2026. It ends with a mapping to Claude Code features.

The scope is the harness that agent users build for their own work. Model training and the internals of agent products are out of scope.

# What is a harness?

Birgitta Böckeler frames it as "Agent = Model + Harness", where the harness is everything except the model: system prompts, tools, files and notes, checks, and the points where humans step in.

OpenAI reported a five-month experiment that built a product with Codex and no hand-written code. Three engineers shipped about one million lines and about 1,500 pull requests. OpenAI says the main job of the engineers shifted from writing code to designing environments, intent, and feedback loops.

In the context of evals, Anthropic notes that evaluating an agent means evaluating the harness and the model working together. The same model produces different results under a different harness.

# Why you need a harness

Anthropic reported these failures when an agent continued long work across many context windows.

| Failure | What happens |
|---|---|
| Trying to do everything at once | The context runs out mid-task, and the next session inherits half-finished, undocumented work |
| Declaring victory early | The agent sees progress and concludes the whole job is done |
| Marking work done without checking | Unit tests or `curl` pass, but the feature does not work end-to-end |
| Spending time getting oriented | The agent guesses what happened in the last session and spends time repairing a broken state |

Research also reports context rot: as the context grows, the model recalls its contents less accurately. Context is a finite resource, and more of it does not mean better results.

Anthropic's multi-agent research system adds that small errors in a stateful agent compound. Restarting from scratch is expensive, so the design needs to resume from where the error happened.

# Guides and sensors

Böckeler classifies harness controls along two axes.

| | Computational (deterministic) | Inferential (LLM judgment) |
|---|---|---|
| Guides (steer before the action) | Types, templates, scaffolds | Convention docs, AGENTS.md, skills |
| Sensors (check after the action) | Tests, linters, type checkers, structural tests | AI review, LLM as a judge |

Guides alone leave you unable to tell whether a rule worked. Sensors alone leave you fixing the same mistake every time. Combine both so the agent can fix problems before a human sees them.

Computational sensors are fast and stable. LLM judgment is slower and less stable, but it can check meaning. Run the fast checks before commit and the expensive ones after integration to keep quality left.

Write sensor output so that an LLM can read it and act on it. OpenAI puts remediation instructions in custom linting error messages so they land in the agent's context. Anthropic's article on tool design also recommends error responses that state what to change, not opaque codes.

# Designing context

Anthropic's guiding principle for context is to find the smallest set of high-signal tokens that maximizes the likelihood of the desired outcome.

- Write system prompts at the right altitude: neither brittle if-else logic nor vague generalities.
- Pick a few canonical examples instead of a laundry list of edge cases.
- Do not load everything up front. Keep file paths and links, and let the agent load data with tools when it needs it.

OpenAI gave four reasons why a single large AGENTS.md failed.

- It crowds out the context.
- When everything is important, nothing is.
- It goes stale fast.
- Mechanical verification gets hard.

They turned AGENTS.md into a table of contents of about 100 lines and made a structured `docs/` directory the system of record. The idea is to give the agent a map, not a 1,000-page manual.

Agent Skills follow the same idea. At startup the agent reads only names and descriptions, loads the body when a skill becomes relevant, and opens bundled files only when needed. Because of this progressive disclosure, you can install many skills without crowding the context.

Claude Code's documentation recommends writing in CLAUDE.md only what Claude cannot learn by reading the code: the commands that verify work, places where the code misleads, and decisions that live outside the repository.

# State and notes

Across context windows, the agent has no memory of earlier sessions. Anthropic compares this to engineers working in shifts where each new engineer arrives with no memory of the previous shift.

Anthropic's harness for long-running agents carries state in three places.

1. A JSON list of features. Every item starts as failing, and the agent may change only the pass/fail field.
2. A plain-text progress file that the agent appends to at the end of each session.
3. Git commits at session boundaries, so the agent can revert bad changes.

They chose JSON for the list because the model was less likely to overwrite JSON inappropriately than Markdown. Each session starts the same way: check the working directory, read the progress file, read `git log`, read the feature list, and run a basic test before starting new work.

OpenAI treats plans as first-class artifacts. Complex work gets an execution plan with progress and a decision log, checked into the repository. Knowledge that lives in Slack or in people's heads does not exist for the agent.

Anthropic's multi-agent research system saves its plan to external memory. Subagents write their output to external storage and return only lightweight references to the lead agent. Passing everything through the lead loses information, like a game of telephone.

# Decomposing work

## One thing per session

The most effective change in Anthropic's long-running harness was asking the agent to work on a single feature per session. When it finishes, it leaves the environment ready for the next session and stops.

## Workflows and agents

Anthropic distinguishes workflows, where code paths orchestrate the LLM, from agents, where the LLM directs its own process. It recommends starting with the simplest solution and adding complexity only when it demonstrably improves outcomes.

| Pattern | What it does | Good fit |
|---|---|---|
| Prompt chaining | Each step processes the previous output, with optional gates in between | Fixed subtasks, such as writing an outline, checking it, then writing the document |
| Routing | Classifies input and sends it to a specialized path | Distinct categories that need different handling |
| Parallelization | Splits work into parallel parts, or runs the same task several times and votes | Speed, or several perspectives |
| Orchestrator-workers | A central LLM breaks down tasks, delegates them, and synthesizes results | Subtasks you cannot predict up front |
| Evaluator-optimizer | One LLM generates and another evaluates, in a loop | Clear evaluation criteria |

## When to use subagents

In Anthropic's research system, a lead agent that plans and delegates searches to parallel subagents outperformed a single agent by 90.2% on their internal eval. It also used about 15 times as many tokens as a chat. It suits high-value research that parallelizes well and does not fit in one context window.

When delegating, state the goal, output format, sources and tools, and task boundaries. Short instructions led to duplicated or missed work.

Claude Code's documentation says subagents give you "parallelism or a fresh pair of eyes", not a smarter answer. They help with independent parallel work, keeping long output out of the main conversation, reviewers that did not write the code, and work with restricted tools.

A subagent starts without the conversation history, so restate the rules it must follow in the delegation prompt. In Claude Code, subagents cannot use `AskUserQuestion`, the tool that asks the user questions with choices. Keep questions to humans in the main conversation.

# Enforcing rules deterministically

Instructions in prompts are not always followed. Claude Code's documentation suggests moving a CLAUDE.md line into a hook if Claude keeps missing it despite added emphasis. Conversely, remove hooks that only encode preferences and let the model decide.

Hooks are handlers that Claude Code runs at fixed points in its lifecycle.

- A `PreToolUse` hook that exits with code 2 blocks the tool call, and its stderr goes back to the model as the reason.
- A `Stop` hook that blocks the stop turns its reason into the next instruction, so you can keep the agent working until the tests pass.
- `SessionStart` also runs on resume, which makes it a good place to inject the latest state.

OpenAI promoted rules that documentation failed to enforce into code, using custom linters and structural tests. Instead of micromanaging details, they enforce invariants such as dependency direction, naming, and file size limits.

# Human involvement

A harness does not remove humans. Böckeler argues that a good harness directs human input to where it matters most.

## The steering loop

The human job is to improve the harness whenever an issue repeats: add a guide so it happens less, or a sensor that catches it earlier. AI can help with that work too.

Anthropic's article on skills also recommends asking the agent to record what worked and what failed back into the skill. Capturing what the work actually needed beats guessing the context up front.

## What stays with humans

Claude Code's documentation asks five questions before you run an agent unattended.

1. Is "done" defined (a goal condition, a spec file, failing tests to make pass)?
2. Can it verify its own work (tests, a Stop hook, a review step)?
3. What may it touch (permissions, sandbox, the one branch it may push)?
4. What stops it (max turns, budget limits)?
5. Where do results go, and who approves them (pull requests, the person who approves merges and deploys)?

What stays with humans is starting the run, deciding what done means this time, and approving the merge at the end.

## Humans are not the only reviewers

OpenAI moved almost all review to agent-to-agent loops. Humans set priorities, turn user feedback into acceptance criteria, and check outcomes. Agents escalate only when a decision needs human judgment. OpenAI warns that these results depend heavily on that repository's structure and tooling, and that readers should not expect them to generalize.

Böckeler points out that no sensor reliably catches misdiagnosis, overengineering, or misunderstood instructions. Unless a human makes the intent clear up front, correctness falls outside every sensor's coverage.

# Evals

Anthropic organizes agent evals as follows.

- Graders come in three kinds. Code is fast and impartial but brittle to valid variations. Models are flexible but need calibration against human judgment. Humans are the most reliable but slow and expensive. Prefer deterministic graders where you can.
- Grade what the agent produced, not the path it took. Checking the order of tool calls is brittle.
- Start early and small. Twenty to fifty tasks drawn from real failures are enough, because early changes have large effects.
- Measure variance. `pass@k` is the chance of at least one success in `k` trials, and `pass^k` is the chance that all `k` trials succeed. At 75% per trial, all three of three succeed only about 42% of the time.
- Read the transcripts. Do not take scores at face value until someone has.
- Automated evals are one layer. Combine them with production monitoring, user feedback, and human review.

Anthropic's research system also started with about 20 queries and used an LLM judge with a rubric of factual accuracy, citation accuracy, completeness, source quality, and tool efficiency. Manual testing was still essential. Humans noticed that early agents kept choosing SEO content farms over authoritative sources.

Böckeler leaves an open question: if sensors never fire, is that a sign of high quality or of weak detection? No established way yet exists to measure the quality of a harness itself.

# Mapping to Claude Code

The table maps the patterns above to Claude Code features as of November 2026.

| Pattern | Claude Code feature |
|---|---|
| Context loaded every time | `CLAUDE.md`, limited to what the code does not tell Claude |
| Procedures and templates loaded on demand | Skills (`.claude/skills/<name>/SKILL.md`). Use `disable-model-invocation: true` for steps with side effects so they run only when a human invokes them |
| Work in a separate context | Subagents (`.claude/agents/<name>.md`), with `tools` to restrict access |
| Deterministic enforcement | Hooks (`PreToolUse`, `Stop`, `SessionStart`) and permission deny rules |
| State and notes | Files and Git in the repository. Subagents can keep cross-session notes with `memory` |
| Unattended runs | `claude -p`, `--allowedTools`, `--permission-mode dontAsk`, `--max-turns`, `--max-budget-usd` |
| Resuming | `--resume` and `--continue` |

# Design guidelines

Here are the guidelines that emerge from these sources.

1. Start with the simplest setup, and add complexity only after evals show it helps.
2. Advance one thing per run, and stop in a state the next run can pick up right away.
3. Keep state in structured data that resists overwriting, history in notes, and boundaries in Git.
4. Put only a map in the context, and let the agent load details when needed.
5. Enforce rules you cannot compromise on with hooks and linters, not prompts. Include remediation in check output.
6. Use subagents for parallelism and fresh eyes, and keep questions to humans in the main conversation.
7. Keep humans on defining done and approving irreversible actions. Improve the harness when the same failure repeats.
8. Start evals with a few real cases, and read transcripts before trusting scores.

# Conclusion

- The harness is everything except the model, and the same model behaves differently under a different harness.
- Long-running work tends to fail through overreach, premature victory, and completion without checks. Advance one thing per run, write state outside the model, and use Git for boundaries.
- Design control as a combination of guides before the action and sensors after it, mixing computational and inferential checks.
- Do not remove humans. Focus their input on defining done, approving, and improving the harness.
- Start evals with a few real cases and grade outcomes. How to measure the quality of a harness itself is still an open question.

# Related posts

- [Context Engineering](/posts/context-engineering/)
- [Engineering in the AI Era: Evolution from Prompts to Autonomous Systems](/posts/engineering-in-ai-evolution-to-autonomous-systems/)
- [Considering Engineering in the Age of AI](/posts/engineering-in-ai-reflections/)
- [LLM Prompt Management and Evaluation Platform Using GitHub Agentic Workflow](/posts/github-agentic-workflow-llm-prompt-management/)
- [What is Authorization for AI Agents? Delegation Chains, Human Approval, and Audit with MCP, AuthZEN, AIMS, and A2A/AP2](/posts/authorization-for-ai-agents/)

# References

- [Building effective agents](https://www.anthropic.com/engineering/building-effective-agents) (Anthropic, 2024-12-19)
- [Effective context engineering for AI agents](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents) (Anthropic, 2025-09-29)
- [Writing effective tools for agents — with agents](https://www.anthropic.com/engineering/writing-tools-for-agents) (Anthropic, 2025-09-11)
- [Effective harnesses for long-running agents](https://www.anthropic.com/engineering/effective-harnesses-for-long-running-agents) (Anthropic, 2025-11-26)
- [How we built our multi-agent research system](https://www.anthropic.com/engineering/multi-agent-research-system) (Anthropic, 2025-06-13)
- [Equipping agents for the real world with Agent Skills](https://www.anthropic.com/engineering/equipping-agents-for-the-real-world-with-agent-skills) (Anthropic, 2025-10-16)
- [Building agents with the Claude Agent SDK](https://www.anthropic.com/engineering/building-agents-with-the-claude-agent-sdk) (Anthropic, 2025-09-29)
- [Demystifying evals for AI agents](https://www.anthropic.com/engineering/demystifying-evals-for-ai-agents) (Anthropic, 2026-01-09)
- [Harness engineering: leveraging Codex in an agent-first world](https://openai.com/index/harness-engineering/) (OpenAI, 2026-02-11)
- [Harness engineering for coding agent users](https://martinfowler.com/articles/exploring-gen-ai/harness-engineering.html) (Birgitta Böckeler, 2026-04-02)
- [Claude Code: Best practices](https://code.claude.com/docs/en/best-practices)
- [Claude Code: Create custom subagents](https://code.claude.com/docs/en/sub-agents)
- [Claude Code: Extend Claude with skills](https://code.claude.com/docs/en/skills)
- [Claude Code: Hooks reference](https://code.claude.com/docs/en/hooks)
- [Claude Code: Run Claude Code programmatically](https://code.claude.com/docs/en/headless)
- [Agent Skills](https://agentskills.io)

