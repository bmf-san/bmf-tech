---
title: "How Much Authorization Platform Do You Need? Scope, Ownership, Stages, and Externalized Decisions"
slug: authorization-platform-scope-stages
date: 2026-10-25
author: bmf-san
categories:
  - Architecture
tags:
  - Authorization
  - IAM
  - Architecture
  - System Design
description: "How much authorization platform does your organization need? Why authorization spans the whole organization yet resists centralization, what to standardize centrally versus leave to domains, decision axes, pain signals, four stages, signs not to build, and whether to externalize decisions."
translation_key: authorization-platform-scope-stages
draft: false
---

# Overview

Discussions about an authorization platform tend to collapse into two options: "build one platform that manages every permission in the company" or "leave it to each application." In practice, several stages sit between those two, and your organization needs one of them.

This article lays out the material you need to decide how much authorization platform your organization needs. You will learn:

- Why an authorization platform spans the whole organization, and why you still should not gather it in one place
- How to split what you standardize across the organization and what each business domain owns
- The decision axes and present-day pain signals that tell you which stage you need
- Four stages, and the signs that you should not build a platform
- How to decide whether to move authorization decisions out of applications (externalize them)

[What is Authorization? Mapping AuthN, AuthZ, and IAM Subdomains](/posts/authorization-domain-iam-subdomains/) covers what authorization is, how it differs from authentication, and the IAM subdomains. This article builds on that and focuses on how much an organization should build.

The stages and decision axes in this article reflect my own framing. No standard set of criteria exists for this decision.

Here is the big picture first.

```mermaid
flowchart TB
  Q["How much authorization<br/>platform do you need?"] --> S["Scope: spans the organization<br/>but not centralized in one place"]
  S --> D["Ownership: mechanisms and contracts are shared<br/>meaning stays in each domain"]
  D --> A["Decision axes and pain signals"]
  A --> L0["Stage 0: each app implements its own"]
  A --> L1["Stage 1: shared conventions and libraries"]
  A --> L2["Stage 2: externalized decisions"]
  A --> L3["Stage 3: platform including governance"]
  L2 -.-> X["Whether to externalize<br/>trade-off with latency and availability"]
```

# An Authorization Platform Spans the Whole Organization

Authorization does not stay inside one team or one service. Six reasons explain why.

## Authorization happens on every request, for every resource

You can verify authentication once at the entry point and carry the result in a token. Authorization asks, for each resource, "may this subject perform this action on this resource?" That means you need enforcement points (PEPs, Policy Enforcement Points) everywhere resources live: gateways, every service, and the data layer.

The OWASP Authorization Cheat Sheet also recommends checking permissions on every request. The OWASP Top 10:2025 places Broken Access Control at A01.

[Authorization Platform Architecture and Design Decisions](/posts/authorization-platform-architecture-decisions/) covers where to place PEPs along the request path.

## The "meaning" of a permission lives in each business domain

Only the team that runs a business process knows what "may approve this invoice" means. An authorization platform can provide the decision mechanism. But what to allow requires alignment with every business domain. An authorization platform crosses not only technical boundaries but also team boundaries.

[The Abstraction Trap in Architecture Design](/posts/abstraction-trap-in-architecture/) also touched on how authorization is a bundle of domain rules.

## The information a decision needs comes from across the organization

The inputs to a decision do not sit in one place.

| Information | Main source |
|---|---|
| Subject attributes (department, title, employment type) | HR systems, directories |
| Permission assignments | Approvers and administrators (access governance) |
| Risk signals | Security team, detection systems |
| Contract state (plan, contract period) | Contract management |
| Resource attributes (owner, visibility, state) | Each business domain |

Authorization becomes the point where these sources meet. A single team cannot decide who holds which information and how fresh it must be when it reaches a decision.

## Governance is the organization's own process

Who may grant permissions, separation of duties, periodic access reviews, and revoking permissions when someone leaves all tie directly to the organizational structure and internal controls. NIST SP 800-53 Rev. 5 also lists Account Management (AC-2), Separation of Duties (AC-5), and Review of User Privileges (AC-6(7)) as controls.

Technology alone does not determine whether you can reliably remove a departed employee's permissions from every system. Cross-organizational processes determine it.

## Delegation chains across services

Permissions travel across services: from a user to an agent, to an MCP server, to a business API. If token audience design, scope structure, how services get new tokens, and how you record the user and the agent are not consistent, one loose implementation anywhere becomes a hole in the whole system. A typical example: a service that forwards the token it received straight to the next service.

[Standardizing Workload Identity](/posts/workload-identity-wimse-transaction-tokens/) covers context propagation between services, and [What Changes in Authorization with AI Agents](/posts/authorization-for-ai-agents/) covers agent delegation.

## Hard to change later, and when it stops, everything stops

Authorization logic gets embedded in code everywhere. So when you change the permission model, you have to touch every service.

Conversely, once you externalize decisions, an outage of the decision service can stop every feature. The principle says to deny (fail-closed) rather than allow (fail-open) when you cannot reach a decision, so the decision service's availability drives the availability of the whole system. Either way, the blast radius covers the entire architecture.

# Do Not Centralize It: What to Standardize and What Domains Own

The reasons above make it tempting to say "manage all authorization centrally." But that collides with the second reason: the meaning of permissions lives in each domain. If a central platform team takes on the meaning of every business's permissions, the platform team becomes a bottleneck and its business knowledge goes stale.

So **standardize the mechanisms and contracts across the organization, and keep the meaning in each domain**.

| Standardized across the organization (owned by the platform) | Owned by each business domain |
|---|---|
| Token formats, conventions for audiences and scopes, how to carry delegation | Definitions of resources and actions, business rules |
| The decision API contract (such as AuthZEN) and how to write policies | The content of the domain's own policies |
| Audit log format, revocation mechanisms, delivery of continuous evaluation | Testing and changing policies |
| Guardrails (constraints no domain may break) | Decisions within the guardrails |

Guardrails are constraints every domain follows regardless of business meaning. Examples: "never allow access across tenants," "deny when you cannot reach a decision," and "always write privileged operations to the audit log."

```mermaid
flowchart TB
  subgraph Shared["Shared (owned by the platform)"]
    T["Token formats, scope conventions, delegation"]
    API["Decision API contract, policy style"]
    LOG["Audit log format, revocation, continuous evaluation"]
    G["Guardrails"]
  end
  subgraph DomainA["Business domain A"]
    PA["Resource and action definitions<br/>policy content and tests"]
  end
  subgraph DomainB["Business domain B"]
    PB["Resource and action definitions<br/>policy content and tests"]
  end
  Shared --> DomainA
  Shared --> DomainB
```

You can centralize authentication in an IdP almost entirely, but you cannot do the same with authorization. The organization can give one answer to "who is this?", but the answer to "what may they do?" differs by business. Because of this difference, an authorization platform spans the whole organization while the whole organization shares ownership of it.

[Viability Conditions for Building a Platform](/posts/platform-viability-conditions/) covers the general conditions for building a shared platform (being a supporting or generic subdomain, reuse by many teams, sustained ownership, and so on). An authorization platform can follow the same conditions, but if you put the meaning of permissions into the platform, the platform ends up holding core-domain knowledge. The line in the table above also keeps the platform on the supporting and generic side.

# Not "Do We Need It?" but "Which Stage?"

An authorization platform is not a yes-or-no choice. Ask instead **which stage of mechanism you need**.

## Decision axes

| Axis | Direction that calls for a platform | Direction that works without one |
|---|---|---|
| Scale and cross-cutting reach | Many applications and services need authorization, and you want the same rules to apply across several of them | Few applications, and each application's rules stay self-contained |
| Complexity of the permission model | Conditional permissions (time, amount, risk), sharing, hierarchy, and inheritance relationships, customer-defined roles, delegation | A few fixed roles suffice |
| Change frequency and who changes it | Policies change often, and people other than developers (business admins, customers) want to change them | Changes are rare, and only developers make them |
| Governance and regulatory requirements | You need audit trails, separation of duties, periodic access reviews, and explanations of "why we allowed this" | Light audit requirements |
| Diversity of subjects | Besides people, you have workloads, external partners, and AI agents, and delegation chains | Subjects are almost all people |
| Consistency requirements | You want the same person's permissions aligned across several products, and you want to reflect departures and permission changes in every system right away | One product, and propagation delays do not matter |
| Organizational capability and cost | A team can operate the platform and keep the decision service highly available (or buy one) | Nobody can operate it, and you cannot accept the latency of a decision round trip or the risk of lower availability |

The last axis, organizational capability and cost, differs in nature from the others. The other axes show need; this one shows feasibility. If the need is strong but capability falls short, consider options such as stepping down one stage or using a managed service.

[Authorization Architecture, Models, and Engines](/posts/authorization-architecture-models-engines/) covers how to choose a permission model (RBAC, ABAC, ReBAC, and others). [Exploring Permission Management Design](/posts/permission-management-design/) summarizes permission design within an application (functional scope and data scope, among others).

## Pain signals you can see today

You consider the axes including future requirements, but present-day pain also tells you something. The following conditions signal that the time has come to move up a stage.

- Authorization code gets copied into each application, and behavior differs between applications
- Authorization defects (such as seeing someone else's data) keep recurring
- Adding a role or changing a permission requires development and a release every time
- You cannot answer "who can see this resource?" and someone investigates by hand at every audit
- Permission requests run on tickets and spreadsheets

A rough guide to which axis and stage each signal relates to:

| Signal | Related axis | Stage most likely to help |
|---|---|---|
| Authorization code copied, behavior differs | Scale and cross-cutting reach | Stage 1 |
| Authorization defects keep recurring | Scale and cross-cutting reach, complexity of the permission model | Stage 1, Stage 2 |
| Permission changes need development and a release | Change frequency and who changes it | Stage 2 |
| You cannot answer who can see a resource | Governance and regulatory requirements | Stage 2, Stage 3 |
| Requests run on tickets and spreadsheets | Governance and regulatory requirements | Stage 3 |

## The four stages

| Stage | What it includes | When it fits |
|---|---|---|
| 0. Each app implements its own | Each application decides in code | Small scale, few applications, fixed roles |
| 1. Shared conventions and libraries | Standardize token formats, scope structure, audit log format, and revocation mechanisms across the organization. Decisions stay in each application | Applications and teams have started to multiply |
| 2. Externalized decisions | Share a decision engine (PDP, Policy Decision Point) and manage policies as code. Each domain owns the content of its policies | Scale and cross-cutting reach, permission model complexity, and change frequency are strong |
| 3. Platform including governance | Add access requests and approvals, access reviews, continuous evaluation, and management of delegation chains | In addition, governance and regulation, diversity of subjects, and consistency requirements are strong |

```mermaid
flowchart LR
  L0["Stage 0<br/>Each app implements its own"] --> L1["Stage 1<br/>Shared conventions and libraries"]
  L1 --> L2["Stage 2<br/>Externalized decisions"]
  L2 --> L3["Stage 3<br/>Platform including governance"]
```

Each stage builds on the previous one. The PDP in Stage 2 relies on the token and audit log formats you standardized in Stage 1. If you skip Stage 1 and introduce a PDP, each PEP still represents subjects and writes logs in its own way, and you lose much of the consistency and auditability that motivated externalization.

You do not need to move every domain at once when you move up a stage. Start with the domains that feel the most pain, and firm up the shared contracts as you go. That keeps the platform close to real requirements.

[Case Study on Authorization Management Infrastructure](/posts/permission-management-case-study/) covers freee's example (from articles published in 2023) of a dedicated team that builds and runs a permission management platform.

## Align Stage 1 early

In my view, once you have several applications or teams, you should align Stage 1 early. How you handle tokens and audit logs forms a shared contract between teams. If you do not align on them as conventions, even without a platform team, unifying them later tends to become costly.

OAuth and related standards can express much of what Stage 1 standardizes (token formats, audiences and scopes, revocation). [A Map of OAuth and Related Standards for an Authorization Platform](/posts/oauth-standards-map-authorization-platform/) organizes which standard decides what.

# Signs You Should Not Build One

Some situations call for not moving up a stage.

- You have one product, with few and stable roles
- The latency of moving decisions out, or the risk that a decision service outage stops everything, outweighs the benefits
- Even if you build a platform, teams have no incentive to use it (they keep deciding inside their applications)

Building a platform too early tends to add availability risk and operational burden before it delivers any benefit. The third sign also means you fail the condition that "several teams or several products will reuse it" from [Viability Conditions for Building a Platform](/posts/platform-viability-conditions/).

Not building a platform does not mean doing nothing, though. As the previous section explains, aligning Stage 1 conventions early pays off once you have several applications or teams. The decision not to build applies mainly to Stage 2 and beyond.

# Should You Externalize Authorization Decisions?

Moving from Stage 1 to Stage 2 comes down to whether you take decisions out of your applications.

## The problem with decisions scattered across applications

Sprinkling checks like `if user.role == "admin"` through each application's code works as the simplest approach while you have few applications. As applications multiply, though, teams write the same rule in several places independently, and the copies drift apart.

David Brossard, co-chair of the AuthZEN Working Group, also said in an OpenID Foundation announcement that siloed, hardcoded, and brittle authorization is already a problem in today's apps.

## What externalized authorization means

Externalized authorization separates decision logic from applications and gathers it in a PDP. The application (PEP) asks "may Alice read doc 123?" and enforces the decision it gets back.

```mermaid
sequenceDiagram
  participant U as User
  participant PEP as Application (PEP)
  participant PDP as Decision engine (PDP)
  participant PIP as Attribute source (PIP)
  U->>PEP: I want to read doc 123
  PEP->>PDP: May Alice read doc 123?
  PDP->>PIP: Fetch attributes needed for the decision
  PIP-->>PDP: Department, resource owner, and so on
  PDP-->>PEP: Permit or deny
  PEP-->>U: Result with the decision enforced
```

Externalizing does not necessarily mean calling a decision service over the network. If you run the PDP as an in-process library or a sidecar, and the logic lives outside the application code as managed policy, you have still externalized it. [Authorization Architecture, Models, and Engines](/posts/authorization-architecture-models-engines/) covers the roles of PEP, PDP, PIP, and PAP, and [Implementation of Access Control System for Microservices Using OPA](/posts/microservices-access-control-opa/) shows an implementation with OPA.

## What externalization gives you

- **Central management, auditing, and reuse of policy**: rules live in one place as policy rather than scattered through code, and you can record each decision
- **Consistency across applications**: the same rules, evaluated by the same engine, apply to several applications
- **Localized change**: changing a rule means changing a policy, decoupled from application releases

These benefits act directly on three of the decision axes: scale and cross-cutting reach, change frequency and who changes it, and governance and regulatory requirements.

## What externalization costs you

Externalization has costs: latency from a network round trip per decision, and the PDP's availability directly affecting the availability of the whole system. Mitigations:

| Cost | Mitigation |
|---|---|
| Latency from decision round trips | Cache decision results. Run the PDP nearby as a sidecar or in-process library |
| A PDP outage stops everything | Distribute PDPs close to each PEP. Absorb short outages with caching. Design timeouts assuming fail-closed |
| List views multiply the number of decisions | Use partial evaluation to turn policy into conditions such as SQL (OPA partial evaluation). If the PDP supports it, use the Resource Search API to fetch permitted resources in bulk (AuthZEN) |

[Authorization Platform Architecture and Design Decisions](/posts/authorization-platform-architecture-decisions/) covers PEP and PDP placement and non-functional design such as caching and consistency.

## A standard between PEP and PDP: AuthZEN

One concern with externalization: the PEP implementation may get locked into a specific PDP product. The OpenID Foundation's AuthZEN Working Group standardizes the interface between PEP and PDP. OpenID Foundation members approved the version of Authorization API 1.0 dated January 11, 2026 as an OpenID Final Specification, and the foundation announced the approval on January 12, 2026.

Authorization API 1.0 defines these main APIs:

- Access Evaluation API: takes a Subject, Action, Resource, and an optional Context, and returns a decision
- Access Evaluations API: requests several decisions together
- Search APIs (Subject Search, Resource Search, Action Search): find permitted subjects, resources, or actions

The Access Evaluation API is the core of the spec and the only endpoint the metadata requires. The Access Evaluations API and the Search APIs are optional, and a PDP that publishes metadata signals support by listing their endpoints.

The policy language, architecture, and state management of a PDP are out of the specification's scope. So as long as a PDP implements the specification, the PEP can query it the same way, whatever runs inside. [What is the AuthZEN Authorization API?](/posts/authzen-authorization-api/) covers the specification in detail.

Engines and products you can use as the PDP for externalized authorization include OPA, Cedar (with Amazon Verified Permissions as a managed service built on Cedar), Cerbos, and OpenFGA. Not every one of them accepts the AuthZEN API out of the box. Whether and how each product supports it (natively, as an experimental feature, or through a separate plugin) varies.

## Externalized, but each domain still owns its policy content

Externalization differs from centralization. You share the decision mechanism (the evaluator, the decision API contract, the policy style, and the distribution mechanism), while each domain writes, tests, and changes its own policy content. The only policies the platform team owns are guardrails common to every domain.

This split applies the earlier rule (share mechanisms and contracts across the organization, keep meaning in each domain) to externalization. If the platform team starts writing every domain's policies because decisions now live outside the applications, the problems of central management come right back.

## Deciding whether to externalize

Externalization fits when several of these conditions hold together:

- You want the same rules to apply to several applications or services
- Policies change often, and you want to decouple them from application releases
- You want to explain and audit "why we allowed this" for each decision
- You have many conditional permissions or relationship-based permissions (sharing, hierarchy, inheritance)

Keeping decisions inside the application works better when:

- The rules stay within one application and rarely change
- You cannot afford a decision round trip on the synchronous path, and in-process evaluation is not an option either
- You have no team to run a highly available PDP, and you cannot use a managed service

If you are unsure, consider starting with an in-process library or a sidecar rather than a central service over the network. You gain the benefits of externalization (separated logic and managed policy) while keeping the latency and availability costs small.

# Summary

- Authorization happens wherever resources live, its meaning lives in each business domain, and its inputs come from across the organization. So an authorization platform spans the whole organization
- Even so, do not gather it in one place. Standardize mechanisms and contracts (tokens, the decision API contract, audit logs, revocation, guardrails) across the organization, and keep meaning (resources and actions, policy content) in each domain
- An authorization platform is not a yes-or-no choice. Ask which stage you need: Stage 0 (each application), Stage 1 (shared conventions and libraries), Stage 2 (externalized decisions), or Stage 3 (a platform including governance)
- The seven decision axes are scale and cross-cutting reach, permission model complexity, change frequency, governance, diversity of subjects, consistency, and organizational capability. Present-day pain also tells you where you stand
- Align Stage 1 early once you have several applications or teams. Skip Stage 2 and beyond when you have one stable product, when latency and availability risks outweigh the benefits, or when nobody would use the platform
- Externalizing decisions trades latency and availability costs for central management, consistency, and localized change. AuthZEN Authorization API 1.0 standardizes the interface between PEP and PDP, and each domain still owns its policy content after externalization

# Related Posts

- [What is Authorization? Mapping AuthN, AuthZ, and IAM Subdomains](/posts/authorization-domain-iam-subdomains/)
- [What Are Authorization Architectures and Models? PEP/PDP, RBAC/ABAC/ReBAC, and Choosing a Policy Engine](/posts/authorization-architecture-models-engines/)
- [What is the AuthZEN Authorization API? A Standard Between PEP and PDP](/posts/authzen-authorization-api/)
- [Which OAuth Standards Does an Authorization Platform Need? A Map of Delegation, Tokens, and Continuous Evaluation](/posts/oauth-standards-map-authorization-platform/)
- [What is Workload Identity? WIMSE, Transaction Tokens, and SPIFFE Client Authentication for Service Identity and Context](/posts/workload-identity-wimse-transaction-tokens/)
- [What is Authorization for AI Agents? Delegation Chains, Human Approval, and Audit with MCP, AuthZEN, AIMS, and A2A/AP2](/posts/authorization-for-ai-agents/)
- [How Should You Design Authorization for MCP? Common Incidents and Decisions That Are Hard to Reverse](/posts/mcp-authorization-design/)
- [What Does an Authorization Platform Architecture Look Like? Containers, Key Sequences, PEP/PDP Placement, and Non-Functional Decisions](/posts/authorization-platform-architecture-decisions/)
- [Exploring Permission Management Design](/posts/permission-management-design/)
- [Case Study on Authorization Management Infrastructure](/posts/permission-management-case-study/)
- [Viability Conditions for Building a Platform](/posts/platform-viability-conditions/)
- [The Abstraction Trap in Architecture Design](/posts/abstraction-trap-in-architecture/)
- [Implementation of Access Control System for Microservices Using OPA](/posts/microservices-access-control-opa/)

# References

- [OpenID Foundation - Authorization API 1.0](https://openid.net/specs/authorization-api-1_0.html)
- [OpenID Foundation - Authorization API 1.0 Final Specification Approved](https://openid.net/authorization-api-1-0-final-specification-approved/)
- [OpenID Foundation - AuthZEN Working Group](https://openid.net/wg/authzen/)
- [OpenID Foundation - OpenID Foundation advances authorization for the agent era with new AuthZEN Working Group Drafts](https://openid.net/openid-foundation-advances-authorization-for-the-agent-era-with-new-authzen-working-group-drafts/)
- [AuthZEN Interop](https://authzen-interop.net/)
- [OWASP Cheat Sheet Series - Authorization Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html)
- [OWASP Top 10:2025 - A01 Broken Access Control](https://owasp.org/Top10/2025/A01_2025-Broken_Access_Control/)
- [NIST SP 800-53 Rev. 5 - Security and Privacy Controls for Information Systems and Organizations](https://csrc.nist.gov/pubs/sp/800/53/r5/upd1/final)
- [NIST SP 800-162 - Guide to Attribute Based Access Control (ABAC) Definition and Considerations](https://csrc.nist.gov/pubs/sp/800/162/upd2/final)
- [Open Policy Agent - Data Filtering with OPA (partial evaluation)](https://www.openpolicyagent.org/docs/filtering)
