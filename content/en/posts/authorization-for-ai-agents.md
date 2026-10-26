---
title: "What is Authorization for AI Agents? Delegation Chains, Human Approval, and Audit with MCP, AuthZEN, AIMS, and A2A/AP2"
slug: authorization-for-ai-agents
date: 2026-10-30
author: bmf-san
categories:
  - Architecture
tags:
  - Authorization
  - OAuth
  - IAM
  - Generative AI
  - AI
description: "What changes in authorization when AI agents act for users? This post covers delegation chains, least privilege, prompt injection, human approval, and audit attribution, then maps MCP authorization, AuthZEN AARP and COAZ, AIMS, OBO/AAuth, ID-JAG, and A2A/AP2 with their status as of October 2026."
translation_key: authorization-for-ai-agents
draft: false
---

# Overview

An AI agent takes a goal from a user and decides on its own which tools to call and in what order. That differs at the root from a traditional application that runs a fixed workflow, and it changes how we handle permissions.

This post answers three questions in order.

- What changes in authorization when AI agents enter the picture
- Which standards respond to those changes, and at which layer
- Which standard to use for what when you design a system

The standards covered are MCP authorization, the new AuthZEN drafts (the AARP profile and the COAZ framework with its MCP binding), AIMS, the OAuth extensions for delegation (OBO and AAuth), ID-JAG, and A2A and AP2. I checked versions and status against primary sources as of October 2026. Most of these are still drafts.

This post assumes the definition of authorization and the IAM subdomains from [the post on authorization and IAM subdomains](/posts/authorization-domain-iam-subdomains/) and the overview in [the map of OAuth-related standards](/posts/oauth-standards-map-authorization-platform/).

# Why AI agents change authorization

## How agents differ from traditional applications

| Aspect | Traditional application | AI agent |
|---|---|---|
| Steps | Runs a workflow that developers defined | Decides the means (which tools, in what order) at runtime |
| Callees | Known in advance | Discovers tools and agents at runtime and calls them |
| Behavior | Deterministic | Non-deterministic; may generate wrong arguments or the wrong target |
| Trust in input | Input and instructions stay separate | Documents and tool results it reads may act as instructions |
| Path of authority | User → app → API | User → agent A → agent B → tool, as a chain |
| On denial | Ends with an error | May try another route around the denial |

These differences create five new authorization problems.

## 1. Delegation chains make it impossible to verify whose authority applies

Say a user asks agent A, A hands work to agent B, and B calls a tool. If A just relays the token it received, the tool sees only the user as the caller, and A and B vanish from the record. The token also arrives with the full scope the user gave A.

This is a problem because the tool cannot decide whether to allow the request. It cannot check whether the user delegated this operation all the way to B, so it lets through operations the user never delegated. Each hop must represent who acts on whose authority and narrow the scope as it passes it on.

## 2. Agents cannot know their permissions up front, so they tend to get broad ones

An agent finds tools at runtime and learns which permissions it needs during the task. Nobody can narrow the permissions in advance, so teams tend to grant broad ones at the start.

This is a problem because the damage from an agent's mistakes or misuse spreads as far as the permissions it holds. Agents need a way to get only the scope they need, when they need it, step by step.

## 3. Authentication cannot stop requests driven by prompt injection

Suppose an email or web page that the agent reads contains a hidden instruction: "send the customer list to this external address." If the agent follows it, the outbound call arrives as **a request with a valid token from a correctly authenticated client**.

This is a problem because authentication only checks who sends the request. The sender is legitimate, so authentication cannot tell that the request is malicious. Preventing prompt injection completely is also hard. So assume that someone can manipulate the agent, and narrow its permissions so that the damage stays limited when that happens. I covered how to protect prompts themselves in "[Considering Engineering in the Age of AI](/posts/engineering-in-ai-reflections/)".

## 4. A confirmation prompt in the UI is no evidence of human approval

You want a human to approve high-risk operations such as payments or deletions. If the agent only asks "run this?" in its own UI, the approval happens outside the authorization system.

This is a problem because the authorization server and PDP cannot see the approval. No evidence remains, and if the agent skips the prompt or fakes the approval, nothing on the verifying side notices. The approval must take a form that an authorization server or PDP can verify.

## 5. Without separating users from agents, you cannot assign responsibility afterward

If logs do not separate the user from the agent, every operation looks like the user's.

This is a problem because, after an incident, you cannot tell an agent's mistake from the user's intent. Without the cause, you cannot decide whose permissions to narrow or who is responsible. Record enough to reconstruct who did what through which agent.

## Design principles

The standards in this post mostly agree on the following principles.

1. **Keep credentials away from the LLM.** The agent's runtime holds credentials, and the LLM never sees them.
2. **Do not use static API keys.** Use short-lived credentials bound to a key.
3. **Least privilege and down-scoping.** Request fine-grained permissions, and pass even narrower tokens internally.
4. **Record the delegation chain.** Represent the user as `sub` and the acting agent as `act` or `client_id`.
5. **Turn human approval into an authorization event.** Use an approval that the authorization server or PDP can verify, not a UI prompt.
6. **Keep evaluating after issuance.** Enforce revocations and downgrades without undue delay, and narrow permissions when risk changes.

I wrote about AI becoming the main actor in development in "[Engineering in the AI Era: Evolution from Prompts to Autonomous Systems](/posts/engineering-in-ai-evolution-to-autonomous-systems/)". This post deals with how to constrain that autonomy through permissions.

# The big picture: which standard covers which layer

```mermaid
flowchart LR
    U["User"]
    IDP["Enterprise IdP"]
    A["AI agent<br/>(MCP client)"]
    AS["Authorization server"]
    G["Gateway<br/>(PEP, if present)"]
    PDP["PDP<br/>(AuthZEN)"]
    T["MCP server<br/>(tools)"]
    RA["Remote agent"]
    PAY["Merchant / payments"]

    U -- "request / delegation" --> A
    IDP -- "ID-JAG" --> A
    AS -- "access token" --> A
    A -- "MCP" --> G --> T
    G -- "decision via COAZ-MCP" --> PDP
    PDP -. "requestable denial (AARP)" .-> G
    A -- "A2A" --> RA
    A -- "AP2 (mandates)" --> PAY
```

| Standard | Question it answers | Status (October 2026) |
|---|---|---|
| MCP authorization | How an agent gets and validates a token to access tools | MCP 2026-07-28 (current version) |
| AuthZEN COAZ / COAZ-MCP | How to map a request to an authorization question (SARC) | Draft 1 of the OpenID AuthZEN WG |
| AuthZEN AARP | How to obtain approval and re-evaluate when access is not allowed yet | Draft 1 (2026-10-13) |
| AIMS | How to combine existing standards to identify, authenticate, authorize, and monitor agents | draft-ietf-wimse-aims-00 in the IETF WIMSE WG |
| OBO, AAuth | How to express delegation from a user to an agent in OAuth | Both are expired individual drafts |
| ID-JAG | How an enterprise IdP centrally manages an agent's access to SaaS | draft-ietf-oauth-identity-assertion-authz-grant-04 in the OAuth WG |
| A2A | How agents discover each other and delegate tasks | A2A 1.0 (Linux Foundation) |
| AP2 | How to bind an agent's payment to the user's approval | v0.2.0 |

The sections below go through them in this order.

# MCP authorization

MCP (Model Context Protocol) connects AI applications to tools and data. In the 2026-07-28 version, authorization is a profile of OAuth 2.1. The MCP server acts as the resource server, and the MCP client acts as the client.

- The client restricts the token's audience to a specific MCP server with `resource` (RFC 8707). The server **MUST NOT pass a received token through to upstream APIs**.
- When permissions fall short, the server indicates the required scopes with `insufficient_scope`, and the client requests them incrementally (step-up).

MCP authorization covers only obtaining and validating tokens. Deciding "may this tool run on this data on behalf of this user", separating the user from the agent, and human approval all sit outside the authorization spec. The tools spec says servers MUST put access controls in place, but it does not define how to decide. [The post on MCP authorization design](/posts/mcp-authorization-design/) covers the design points.

# New AuthZEN drafts: filling the gaps before and after a decision

The AuthZEN Authorization API 1.0 is a standard API through which a PEP asks a PDP "may this subject perform this action on this resource?" It became an OpenID Final Specification in January 2026. [The AuthZEN Authorization API post](/posts/authzen-authorization-api/) covers the API itself.

In June 2026, the OpenID Foundation announced a new set of AuthZEN drafts aimed at agents. COAZ fills the gap before a decision, and AARP fills the gap after it.

## COAZ and COAZ-MCP: mapping MCP requests to authorization questions

The AuthZEN API defines the shape of the question but not how to build it from an MCP message. The COAZ Framework defines that mapping in a protocol-neutral way, and COAZ-MCP defines a default mapping for each MCP method. If the mapping fails, the PDP denies, or the PDP is unreachable, the operation does not proceed (fail-closed).

The default for `tools/call` stops at "may this caller call this tool." A tool that needs to decide which customer the caller may read declares a mapping (`x-authzen-mapping`) in its `inputSchema` and maps arguments to the resource.

Three points matter for agent authorization.

1. **It separates the user from the agent.** The subject comes from the token's `sub` (the user), and the agent goes into `context`. Policies can judge each one on its own.
2. **A gateway can act as the PEP.** Mappings arrive in the `tools/list` response, so a gateway that knows nothing about the business can still enforce the decisions that the tool's team defined.
3. **The PDP never sees arguments that the mapping leaves out.** If a payment tool's mapping omits the amount, no amount limit applies. The MCP server writes the mappings, so your threat model must cover how far you trust the server.

Draft 1 also predates MCP 2026-07-28 and lags behind it; for example, it has no default mapping for methods that the newer version added.

## AARP: turning "not allowed yet" into approval and re-evaluation

An agent that hits a boundary may try to route around it with another tool or permission. A legitimate path past the denial gives the agent less reason to look for a detour. The AuthZEN Access Request and Approval Profile (AARP) defines that path.

```mermaid
sequenceDiagram
    autonumber
    participant PEP as PEP (agent / gateway)
    participant PDP as PDP
    participant ARS as Access Request Service
    actor M as Approver

    PEP->>PDP: Request evaluation
    PDP-->>PEP: Denial with request info (context.access_request)
    PEP->>ARS: Submit access request
    ARS-->>PEP: Task handle (pending)
    ARS->>M: Ask for approval
    M-->>ARS: Approve
    PEP->>ARS: Check status
    ARS-->>PEP: approved, with approval result
    PEP->>PDP: Re-evaluate with context.approval
    PDP-->>PEP: Permit, or denial with next action
```

The core of the spec lies in its invariants:

- A requestable denial remains a denial.
- Neither submitting a request nor receiving approval grants access by itself.
- At enforcement time, the PDP decides again with the current policy and state. The approval acts only as an input to that decision.

The PDP stays the final authority. The agent can store the task handle and keep working while it waits for approval. Still, Draft 1 has a single author, and no spec yet defines how to go from a COAZ-MCP denial to an AARP request.

# AIMS: a framework that combines existing standards for agents

draft-ietf-wimse-aims is an IETF draft that collects best practices for authenticating and authorizing AI agents. The WIMSE WG published -00 as a WG document in September 2026. AIMS **does not define a new protocol**; it combines WIMSE, SPIFFE, OAuth, and Shared Signals.

| Layer | Key points |
|---|---|
| Identifiers and credentials | Assign one WIMSE identifier (a SPIFFE ID works) to each agent. Keep credentials short-lived, and avoid static API keys. **The LLM MUST NOT have access to credentials** |
| Authentication | mTLS, WIMSE Workload Proof Token (WPT), HTTP Message Signatures |
| Authorization | Put the agent in `client_id` and the user in `sub`. Narrow internal calls with Transaction Tokens, and use Identity Chaining or ID-JAG across domains. Bind human approval to a grant from the authorization server, for example via CIBA |
| Monitoring and remediation | Receive revocation and risk-change signals through SSF (CAEP, RISC), and enforce them without undue delay. Record the agent, the delegating principal, the resource, the action, and the decision in audit logs |

[The workload identity post](/posts/workload-identity-wimse-transaction-tokens/) covers each authentication method.

AIMS still has gaps. It does not discuss how to represent multi-hop delegation with `act`. It leaves out of scope how to derive the needed permissions from a natural-language goal. The draft itself admits that CIBA does not fit well for getting human approval in the middle of execution. That last gap overlaps with the problem AARP addresses from the decision side.

# OAuth extensions for delegation: OBO and AAuth

## OBO: including the agent in consent

draft-oauth-ai-agents-on-behalf-of-user (OBO) extends the Authorization Code grant so that a user can approve delegation to a specific agent on the consent screen. The authorization request names the agent (`requested_actor`), and the token request adds the agent's own authentication (`actor_token`). The resulting access token records the agent in an `act` claim (RFC 8693). The resource server can identify the user by `sub` and the acting agent by `act.sub`.

Still, the draft does not bind the token to the agent's key, and it does not handle multi-hop delegation. The latest version, -02, expired in February 2026.

## AAuth: delegation over channels without redirects

draft-rosenberg-oauth-aauth (AAuth) targets agents that serve customers over phone or SMS. Instead of a service account with access to every user's data, it identifies the user from personal details gathered in conversation and issues a low-privilege token for that user only. Higher privileges require human approval.

Still, how strongly it identifies the user depends on the personal details it collects, and the draft does not address attackers who know a victim's details. The latest version, -01, expired in April 2026. A separate draft with the same name, draft-hardt-oauth-aauth-protocol, takes a different approach.

Neither OBO nor AAuth counts as a standard. They serve as references for a design idea: do not give agents broad permissions; give them only the user's own scope, after approval.

# Enterprise-managed access: ID-JAG

Enterprises want to manage centrally, in their IdP, which SaaS products and which permissions the agents that employees use may reach. When users consent for each SaaS separately, the IdP cannot see those connections.

With ID-JAG (Identity Assertion JWT Authorization Grant), the client exchanges an ID token or similar assertion from the enterprise IdP for an ID-JAG through Token Exchange (RFC 8693). The IdP evaluates administrator-defined policy before issuing it. The client presents the ID-JAG to the SaaS authorization server as a JWT authorization grant (RFC 7523) and receives an access token. Enterprise-Managed Authorization, an official MCP extension, brings this mechanism to MCP.

Keep two caveats in mind. The `client_id` in an ID-JAG represents the agent's application, and the draft does not yet define how to represent individual agents or multi-hop delegation. Also, the IdP can stop only the issuance of new ID-JAGs; neither the ID-JAG draft nor the MCP extension defines a way to cut off tokens already issued. Keep their lifetimes short or propagate revocation through SSF/CAEP. ID-JAG sits at -04 (May 2026) as an OAuth WG draft.

# Agent-to-agent and payments: A2A and AP2

## A2A: collaboration between agents

A2A (Agent2Agent Protocol) lets agents from different vendors delegate tasks to each other. Version 1.0 shipped under the Linux Foundation in March 2026. It complements MCP rather than competing with it; the expected setup has a remote agent using MCP internally.

Agents reuse existing authentication methods such as OAuth 2.0 and mTLS, and declare them in their Agent Card. Each agent decides its own authorization model, but checking authorization on every operation is mandatory (MUST). When an agent needs more authorization mid-task, it pauses the task in `TASK_STATE_AUTH_REQUIRED`, but the credential exchange happens outside A2A. In short, A2A defines how agents connect but not whose authority flows along the delegation chain. OAuth mechanisms such as Token Exchange and Transaction Tokens fill that gap.

## AP2: binding payments to user approval

AP2 (Agent Payments Protocol) makes payments by agents verifiable against what the user approved. The latest version is v0.2.0 (April 2026).

AP2 states that **preventing prompt injection attacks is infeasible**, and treats every LLM and agent as a potential attacker. The user approves what to buy and the payment on a Trusted Surface, and the agent carries the signed evidence (mandates). When the user will not be present, the user hands over a mandate with constraints such as merchants and an amount range in advance. Verifiers such as the merchant and the payment processor each check the constraints independently.

In OAuth, an authorization server issues tokens. In AP2, the agent carries evidence of the user's approval and shows each verifier only the parts it needs.

# Answers to "not enough permission mid-execution"

One of the newer issues in agent authorization concerns what happens when permissions run short in the middle of a task. Each standard answers at a different layer.

| Mechanism | What it does | Who decides |
|---|---|---|
| MCP `insufficient_scope` and step-up | Indicates missing scopes and makes the client obtain a new token | Authorization server and the user's consent |
| CIBA | Gets the user's consent through a separate channel such as a mobile phone | The user |
| A2A `TASK_STATE_AUTH_REQUIRED` | Pauses the task and reports that more authorization is needed | Outside the protocol |
| AuthZEN AARP | Submits a request, obtains approval, and has the PDP decide again | Approvers (a person or an automated governance system), and finally the PDP |

Pick by the question: step-up when the token's scope falls short, CIBA when you need the user's consent through a separate channel, and AARP when a manager or administrator must approve. No spec connects these yet, so for now you must decide how to combine them yourself.

# A design map

Finally, here is a single table of the concerns in designing authorization for AI agents and the standards to use.

| Concern | Standards | Status (October 2026) |
|---|---|---|
| Agent identifiers and credentials | WIMSE (WIT and WIC as credentials, WPT as proof of possession), SPIFFE | WIMSE specs are IETF drafts |
| Access to tools | MCP authorization (OAuth 2.1, RFC 9728, RFC 8707) | MCP 2026-07-28 is current; OAuth 2.1 is a draft |
| Argument-level decisions | AuthZEN API 1.0, COAZ-MCP | API is Final; COAZ-MCP is Draft 1 |
| Delegation from user to agent | Authorization Code, `act` (OBO for reference) | OBO is an expired individual draft |
| Least privilege inside and across hops | Token Exchange (RFC 8693), Transaction Tokens | Transaction Tokens at -11 with WG consensus |
| Across domains | Identity Chaining, ID-JAG | Identity Chaining in the RFC Editor Queue; ID-JAG at -04 |
| Central management by the enterprise | ID-JAG, MCP Enterprise-Managed Authorization | The extension sits in the stable directory |
| Not allowed yet | AuthZEN AARP | Draft 1 |
| Human approval | CIBA, AARP | CIBA is an OpenID Final spec |
| Revocation and monitoring after issuance | SSF, CAEP, RISC, structured audit logs | SSF, CAEP, and RISC 1.0 are Final |
| Agent-to-agent collaboration | A2A | 1.0 |
| Payments | AP2 | v0.2.0 |
| Overall checklist | AIMS | draft-ietf-wimse-aims-00 |

## Where to start

You do not need everything at once. Decide first on what you cannot easily change later.

| When | What to do |
|---|---|
| Decide first | Represent users and agents with separate identifiers. Keep credentials away from the LLM. Never relay tokens, and restrict their audience (RFC 8707). Record both the user and the agent in audit logs |
| Add early | Keep scopes minimal and add more through step-up. Require human approval that the authorization server or PDP can verify for high-risk operations |
| Add when needed | Argument-level decisions at a gateway (COAZ-MCP), approval and re-evaluation (AARP), central management with ID-JAG, internal propagation with Transaction Tokens |

What you can defer is argument-level decisions at a gateway. Access control in the tool itself is a MUST in the MCP tools spec, so you need it from the start.

# Summary

- With AI agents, delegation chains, least privilege that you cannot fix up front, misuse of legitimate authority through prompt injection, human approval, and audit attribution become new problems.
- MCP authorization covers obtaining and validating tokens for tools. How to decide, separating users from agents, and human approval sit outside the authorization spec.
- AuthZEN COAZ-MCP maps MCP requests to authorization questions, down to the argument level when a tool declares a mapping. AARP connects a requestable denial to approval and re-evaluation. Both keep the PDP as the final authority.
- AIMS is a draft that collects best practices for agent authentication and authorization built from existing standards, and you can also use it as a design checklist.
- OBO and AAuth serve as references for expressing delegation (both remain expired individual drafts), ID-JAG handles central management by enterprises, and A2A and AP2 handle agent collaboration and payment approval.
- Most of these remain drafts. Adopt the design principles first, such as keeping credentials away from the LLM, recording the delegation chain, and turning human approval into an authorization event. Pin versions when you build to the wire formats.

# Related Posts

- [What is Authorization? Mapping AuthN, AuthZ, and IAM Subdomains](/posts/authorization-domain-iam-subdomains/)
- [How Much Authorization Platform Do You Need? Scope, Ownership, and Decision Axes](/posts/authorization-platform-scope-stages/)
- [What Are Authorization Architectures and Models? PEP/PDP, RBAC/ABAC/ReBAC, and Choosing a Policy Engine](/posts/authorization-architecture-models-engines/)
- [What is the AuthZEN Authorization API? A Standard Between PEP and PDP](/posts/authzen-authorization-api/)
- [Which OAuth Standards Does an Authorization Platform Need? A Map of Delegation, Tokens, and Continuous Evaluation](/posts/oauth-standards-map-authorization-platform/)
- [What is Workload Identity? WIMSE, Transaction Tokens, and SPIFFE Client Authentication for Service Identity and Context](/posts/workload-identity-wimse-transaction-tokens/)
- [How Should You Design Authorization for MCP? Common Incidents and Decisions That Are Hard to Reverse](/posts/mcp-authorization-design/)
- [What Does an Authorization Platform Architecture Look Like? Containers, Key Sequences, PEP/PDP Placement, and Non-Functional Decisions](/posts/authorization-platform-architecture-decisions/)
- [What is the Big Picture of Service-to-Service Authentication? mTLS, Token Exchange, SPIFFE, and Zero Trust](/posts/service-to-service-authentication/)
- [What are Sender-Constrained Tokens? An Introduction to mTLS (RFC 8705) and DPoP (RFC 9449)](/posts/sender-constrained-tokens-mtls-dpop/)
- [Engineering in the AI Era: Evolution from Prompts to Autonomous Systems](/posts/engineering-in-ai-evolution-to-autonomous-systems/)
- [Considering Engineering in the Age of AI](/posts/engineering-in-ai-reflections/)
# References

- [MCP Specification 2026-07-28: Authorization](https://modelcontextprotocol.io/specification/2026-07-28/basic/authorization)
- [MCP Specification 2026-07-28: Tools](https://modelcontextprotocol.io/specification/2026-07-28/server/tools)
- [MCP Authorization Extensions (ext-auth)](https://github.com/modelcontextprotocol/ext-auth)
- [OpenID AuthZEN Working Group: Specifications](https://openid.net/wg/authzen/specifications/)
- [AuthZEN Authorization API 1.0](https://openid.net/specs/authorization-api-1_0.html)
- [AuthZEN Access Request and Approval Profile 1.0 - Draft 1](https://openid.github.io/authzen/authzen-access-request-approval-profile-1_0.html)
- [COAZ Framework 1.0 - Draft 1](https://openid.github.io/authzen/authzen-coaz-framework-1_0.html)
- [COAZ-MCP Binding 1.0 - Draft 1](https://openid.github.io/authzen/authzen-coaz-mcp-binding-1_0.html)
- [OpenID Foundation advances authorization for the agent era with new AuthZEN Working Group Drafts](https://openid.net/openid-foundation-advances-authorization-for-the-agent-era-with-new-authzen-working-group-drafts/)
- [draft-ietf-wimse-aims: AI Identity Management System](https://datatracker.ietf.org/doc/draft-ietf-wimse-aims/)
- [draft-ietf-wimse-wpt: WIMSE Workload Proof Token](https://datatracker.ietf.org/doc/draft-ietf-wimse-wpt/)
- [The Mission Shaping Problem](https://notes.karlmcguinness.com/notes/the-mission-shaping-problem/)
- [draft-oauth-ai-agents-on-behalf-of-user](https://datatracker.ietf.org/doc/draft-oauth-ai-agents-on-behalf-of-user/)
- [draft-rosenberg-oauth-aauth](https://datatracker.ietf.org/doc/draft-rosenberg-oauth-aauth/)
- [draft-hardt-oauth-aauth-protocol](https://datatracker.ietf.org/doc/draft-hardt-oauth-aauth-protocol/)
- [draft-ietf-oauth-identity-assertion-authz-grant (ID-JAG)](https://datatracker.ietf.org/doc/draft-ietf-oauth-identity-assertion-authz-grant/)
- [draft-ietf-oauth-identity-chaining](https://datatracker.ietf.org/doc/draft-ietf-oauth-identity-chaining/)
- [draft-ietf-oauth-transaction-tokens](https://datatracker.ietf.org/doc/draft-ietf-oauth-transaction-tokens/)
- [RFC 8693 OAuth 2.0 Token Exchange](https://www.rfc-editor.org/rfc/rfc8693)
- [OpenID Connect Client-Initiated Backchannel Authentication Flow - Core 1.0](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html)
- [OpenID Shared Signals Working Group: Specifications](https://openid.net/wg/sharedsignals/specifications/)
- [A2A Protocol Specification](https://a2a-protocol.org/latest/specification/)
- [Agent Payments Protocol (AP2)](https://github.com/google-agentic-commerce/AP2)
