---
title: Considering Architecture Strategy
description: 'Understand architecture strategy as systematic policy for building and evolving systems, covering vision, principles, and scalability planning.'
slug: architecture-strategy-discussion
date: 2025-02-16T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture
  - Organizational Design
  - Design
  - System Design
  - Architecture Strategy
translation_key: architecture-strategy-discussion
---



# Considering Architecture Strategy
In software development, even if you are not in a clear position like CTO or architect, there are opportunities to feel the "necessity of architecture strategy" and think about it.

- "I want to consolidate technical policies on-site, but I don't know where to start."
- "I want an architecture that considers future scalability and organizational growth."
- "I want to know how to deal with the technical debt that has already accumulated."

This article organizes the **necessity and utilization methods of architecture strategy** to answer such questions.

## Background and Disclaimer
There are times when I wonder, "What strategy is this plan being executed under?" or "What is strategy in the first place?"

At such times, abstract thoughts swirl in my head, but since I couldn't properly verbalize them, I decided to organize my thoughts and study on my own to summarize the content.

Since I have no experience as a CTO or architect, and little practical experience in thinking about architecture strategies, it may just be a figment of my imagination.

So, if there is anything clearly wrong, I would be happy if you could let me know quietly.

## 1. Introduction
Before talking about architecture strategy, let's first organize the words "architecture" and "strategy." These are concepts often used in the field of software development, but they are surprisingly ambiguous and can be perceived differently by different people.

This article will outline the following flow.

- **What is architecture strategy?**
- **Why is it necessary, and what are the benefits?**
- **How should it be formulated (process and specific examples)?**
- **Necessary skills, points to note, reference materials**

## 2. What is Architecture Strategy?
### 2.1 What is Architecture?
In software development, "architecture" refers to the **overall structure and design policy of a system**. Specifically, it includes elements such as:

1. **System Structure**
- Components that make up the system (e.g., microservices, databases, API gateways, etc.)
- Interrelationships between components (data flow, communication methods, etc.)
2. **Design Principles and Patterns**
- Design methods such as layered architecture, clean architecture, microservices, DDD, etc.
- Best practices like SOLID principles, CQRS, Event Sourcing
3. **Consideration of Non-functional Requirements**
- Scalability, availability, security, performance, maintainability, and extensibility
4. **Technology Selection**
- Selection of languages, frameworks, databases, cloud providers, etc.
- API design (REST, GraphQL, gRPC, etc.)
5. **Development and Operation Processes**
- CI/CD pipeline design, monitoring/logging, team development flow (agile or waterfall)

These elements are examples and are not exhaustive.

**Architecture** is a higher-level concept than **design**, indicating the "framework of the system" or "design policy." Based on that guideline, the stage of concretizing individual modules, classes, data models, etc., is "design."

### 2.2 What is Strategy?
Strategy refers to the **major direction or plan to achieve goals**. While tactics indicate "individual specific means" or "measures," strategy defines the direction of the entire organization or project from a more long-term and comprehensive perspective.

According to [Good Strategy, Bad Strategy](https://amzn.to/4hZTc0Q), the quality of a strategy is described as follows:

- **Key Points of a Good Strategy**
  - There are specific action guidelines, and what to do is clear
  - Appropriate selection and focus are possible, and "what to do" and "what not to do" are clear
  - Consists of three elements: diagnosis (current situation analysis) → basic policy → action
- **Characteristics of a Bad Strategy**
  - Hollow and lacks specificity
  - Avoids facing major problems head-on and evades issues
  - Simply sets goals without clear methods to achieve them

### 2.3 Definition of Architecture Strategy
Combining these concepts of "architecture" and "strategy" results in **architecture strategy**. It can be defined as **"a policy or plan to systematically indicate how to build and evolve systems within an organization or project."**

Specifically, it includes elements such as:

1. **Clarification of Vision and Goals**
- Aligning with business strategy while indicating technical direction
- Examples: adopting cloud-native, event-driven architecture, transitioning to microservices, etc.
2. **Architecture Principles and Guidelines**
- Criteria for technology selection, principles of system design, standardization guidelines for operation and maintenance
3. **Technology Stack Selection**
- Establishing criteria for adopting development languages, DB, frameworks, etc.
4. **System Scalability and Extensibility**
- Scalable configuration, load balancing, service division strategy
5. **Security and Compliance**
- Authentication/authorization, data encryption, compliance (e.g., GDPR, SOC2)
6. **Operation and Monitoring Strategy**
- Standardization of logging/monitoring and incident response processes

## 3. Purpose and Benefits of Architecture Strategy
Formulating an architecture strategy provides the following benefits:

1. **Unification of Business and Technical Direction**
- Technical decision-making is not ad-hoc and can be chosen strategically and long-term
2. **Suppression of Technical Debt**
- Without planning, technical debt tends to accumulate. Pre-strategy can reduce debt
3. **Improvement of Development Speed and Productivity**
- With clear guidelines, unnecessary discussions within the team are reduced, allowing smooth progress
4. **Improvement of System Quality and Stability**
- By consciously addressing non-functional requirements (scalability, availability, security, etc.), risks of failures and security issues can be reduced
5. **Team Consensus**
- With common goals and policies, decision-making is faster, and the team can align more easily

A **good** architecture strategy eliminates all waste in the project and supports efficient resource allocation towards achieving business goals.

## 4. Areas Requiring Architecture Strategy
Architecture strategy can be formulated and applied at various scopes, from team units to the entire organization.

- **Organization-wide Level**
  - Align with company-wide IT strategy and define common infrastructure and security policies
  - Formulate a long-term roadmap aiming for overall optimization
- **Department Level**
  - Optimize according to department-specific business requirements
  - Decide policies considering system integration and data sharing within the department
- **Team Level**
  - Detailed architecture design for a specific service or application
  - Formulate strategies directly related to practice, such as technology selection, implementation methods, and testing policies

As the scope expands, the cost of coordination and the difficulty of consensus building increase.

## 5. Process of Formulating Architecture Strategy
The formulation of an architecture strategy is a process to clarify "what business objectives to achieve, with what technical approach, and how to operate and govern."

Whether the architecture strategy should be formulated top-down, bottom-up, or hybrid depends on the case and varies according to the organization's culture and business situation.

Below are example steps in the process of formulating an architecture strategy.

### 1. Confirmation of Business Strategy and Vision
1. **Understanding Business Goals and Strategy**
- Understand the business goals and strategies set by management and business divisions, and consider how architecture contributes to the business

2. **Identification of Stakeholders**
- Identify key stakeholders involved in the architecture strategy (management, development departments, business divisions, security personnel, etc.) and organize their requirements, constraints, and expectations

### 2. Understanding the Current Architecture (As-Is)
1. **Investigation of Current Systems and Architecture**
- Investigate and document the system configuration, data flow, technologies used, and operation processes in the target area

2. **Extraction of Issues and Risks**
- Identify business and technical bottlenecks, cost structures, personalization, operational load, technical debt, etc.

### 3. Formulation of Architecture Principles and Direction
1. **Setting Architecture Principles**
- Define which architecture characteristics (non-functional requirements) to prioritize
- Clarify trade-offs (e.g., balance between flexibility and performance)

2. **Policy for General Direction and Prioritization of the Roadmap**
- Determine which areas to address first, where the return on investment is high, etc., and draw a high-level direction
- Consider the priority of business requirements and technical feasibility

### 4. Definition of the Target State (To-Be)
1. **Architecture Design**
- Define the target state from perspectives such as business architecture (business flow, organizational structure), application architecture (structure of services and systems), and technology architecture (infrastructure and network configuration)
  - It is beneficial to use tools like the C4 model

2. **Consideration of Scenarios and Use Cases**
- Materialize major scenarios and use cases expected in business and verify if they align with the architecture
- Consider operational scenarios incorporating security and availability requirements

3. **Architecture Evaluation and Selection (Option Comparison)**
- Often compare multiple architecture proposals and evaluate their pros and cons
- Consider technical stacks (selection of public cloud vendors, types of databases, development languages, etc.) and operational models (on-premises/cloud/hybrid, etc.)

### 5. Gap Analysis and Roadmap Creation
1. **Gap Analysis between As-Is and To-Be**
- Identify differences between the current state and target architecture in terms of technology, organization, cost, skills, etc.
- Consider when and how to resolve each gap and prioritize

2. **Formulation of Migration Strategy**
- Design migration patterns, such as whether a large-scale overhaul is necessary, whether to transition gradually to microservices, or how to handle legacy systems
- Consider methods to reduce initial risks, such as the "Strangler Pattern" or "Lift & Shift"

3. **Creation of Roadmap**
- Set milestones for when and how to proceed with architecture changes
- Roughly plan deliverables, success indicators (KPI/KGI), required resources, budget, and structure for each phase

### 6. Preparation of Execution Plan
1. **Clarification of Execution Plan and Project Plan**
- Break down each measure into project units based on the roadmap and detail the schedule, resources, and structure
- Use project management tools (Jira, Confluence, etc.) to manage progress and risks

2. **Awareness and Consensus Building within the Organization**
- Present the strategy content to stakeholders, obtain approval, and build a cooperative system

### 7. Continuous Evaluation and Feedback
1. **Monitoring during Implementation and Operation Phases**
- Not only proceed according to the roadmap but also regularly check indicators (KPI and non-functional requirement monitoring results, etc.) and evaluate the situation
- Introduce mechanisms such as DevOps and CI/CD to ensure smooth implementation, testing, and release of technical changes

2. **Regular Review of Architecture**
- Regularly review the architecture in line with changes in the business environment and technology trends
- Incorporate improvement proposals at governance meetings (such as architecture review meetings) and update the roadmap

3. **Organizational Learning and Deployment**
- Share success and failure cases and accumulate know-how on architecture formulation and operation within the organization
- Update practices and standards and apply them to the next project

## 6. Example: Architecture Strategy in HR and Labor SaaS
Here, we present a partial example of an architecture strategy assuming "HR and Labor SaaS" as a simple sample.

### 1. Technical Vision
**"Achieve strict permission management for large-scale tenants while balancing rapid feature expansion and operational efficiency."**

- **Background**
  - Increasing use by tenants with complex organizational hierarchies, such as large enterprises.
  - A permission model that spans multiple functions, such as attendance management, payroll, and year-end adjustments, is necessary.
  - A **"common permission management platform"** that is easy to expand and maintain and can handle audits is essential.

### 2. Current Analysis
1. **Existing System**
- "Modular Monolith" structure where each function implements its own permission logic.
- Permission definitions are scattered, making role settings and access control for large organizations complicated.

2. **Problems and Challenges**
- **Development Load**: Need to add permission flags and roles every time a new feature is added.
- **Operational Risk**: Role names and permission requirements are inconsistent, making management and auditing difficult.
- **Scalability**: Unable to handle large-scale tenants with many users accessing simultaneously.

### 3. Architecture Principles
1. **Cloud-Native Centered on GCP**
- Utilize GKE (Google Kubernetes Engine) and Cloud Run for foundational resources.
- Use Cloud Logging/Cloud Monitoring as the basis for application logging/monitoring, while also using Datadog as needed.

2. **Single "Permission Service" Microservice**
- Separate authentication (Auth) and authorization, and make the authorization platform a common API.
- Combine RBAC (Role-Based Access Control) + ABAC (Attribute-Based) to accommodate complex hierarchies of large enterprises.

3. **Thorough Infrastructure as Code and CI/CD**
- Manage infrastructure with Terraform/Cloud Deployment Manager.
- Improve operational efficiency with Cloud Build or GitHub Actions + Cloud Run/GKE auto-deployment.

4. **Standardization of Operational and Audit Logs**
- Aggregate all permission changes and access requests in Cloud Logging.
- Accumulate audit data in BigQuery to facilitate reporting and analysis.

### 4. Example of Technology Stack Selection
1. **Authentication**
- Use Google Identity Platform or Auth0 to achieve SSO with OIDC/OAuth2.
- Configuration that easily supports internal/external ID integration.

2. **Permission Management (Authorization)**
- **Custom Permission Service** (based on Go or Python) running on GKE or Cloud Run.
- Manage role/policy definitions with Cloud SQL (PostgreSQL).
- Use Memorystore (Redis) for caching for fast reference.

3. **Each SaaS Module**
- Containerize function services such as attendance management, payroll, and year-end adjustments, and refer to the permission service.
- API Gateway: Integrate traffic control and authorization flow with Cloud Endpoints or Istio (Service Mesh).

4. **Monitoring and Log Collection**
- **Cloud Logging/Cloud Monitoring** is standard for all containers.
- Error alerts and SLO/SLA management may also be linked to Opsgenie/PagerDuty.

### 5. Roadmap (Example: 3 Phases)
1. **Phase 1 (up to 3 months)**
- **Design and PoC of Permission Platform**
  - Formulate data model for role/policy definition
  - Separate Auth (authentication) and Authorization (authorization) services
- **Build CI/CD Pipeline**
  - Demonstrate auto-deployment with Cloud Build or GitHub Actions + Terraform integration

2. **Phase 2 (up to 6 months)**
- **Introduction of Permission Service to Major Modules**
  - Migrate existing roles starting with attendance management and user management
  - Integrate some remaining on-premises systems (integrate with VPN connection → VPC peering, etc.)
- **Prepare Audit Logs**
  - Accumulate change history and access history in BigQuery and prototype reporting functions

3. **Phase 3 (up to 12 months)**
- **Application to All Modules**
  - Gradually transition microservices such as payroll and year-end adjustments to the permission service
- **Load Testing and Optimization for Large-Scale Tenants**
  - Verify Redis cache and scaling policies and deploy them into operation
- **Strengthen Governance for Long-Term Operation**
  - Fully operate policy management UI and automatic generation of audit reports

## 7. Skills and Knowledge Required for Consideration
The following skills and knowledge are important when considering and executing an architecture strategy.

1. **Basics of Software Design**
- Basic design methods such as clean architecture, microservices, monoliths
- Best practices like SOLID principles and DDD

2. **Perspective to Overlook the Entire System**
- Understanding of architecture characteristics such as scalability, security, performance, availability, and operational management, and the ability to analyze trade-offs

3. **Understanding of Business Requirements**
- Ability to explain why the architecture contributes to the business

4. **Grasp of Technology Trends**
- Ability to follow the latest technologies and industry trends to broaden architecture options

5. **Leadership and Communication**
- Skills in communication, consensus building, and decision-making with teams and stakeholders

6. **Understanding of Organizational Strategy**
- Ability to understand how architecture strategy connects to business strategy and formulate strategies aligned with organizational policies

## 8. Conclusion
This article introduced what **architecture strategy** is, its necessity, the formulation process, and specific examples. The key points can be summarized into the following three:

1. **Long-term Perspective Connecting Business and Technology**
- Not just technology selection, but showing policies and roadmaps linked to business goals
2. **Visualizing Technical Debt and Trade-offs and Building Consensus**
- Identifying non-functional requirements and risks, involving multiple stakeholders to find the optimal solution
3. **Continuous Verification and Improvement**
- Since the business environment and technology trends are constantly changing, regularly review the architecture strategy

By clarifying the architecture strategy, the team's direction aligns, and in the long term, development efficiency, quality, and business promotion power are expected to improve significantly. It is important to judge whether a strategy is necessary according to the situation of your product or organization and adopt an appropriate approach.

## 9. References
### Books
- [**The Path of the Staff Engineer: Architecting to Become an Excellent Technical Professional**](https://amzn.to/4aXEjd3)
  - Explains behaviors and techniques to lead teams and organizations as a software engineer
- **[Fundamentals of Software Architecture: A Systematic Approach Based on Engineering](https://amzn.to/3CVjMcv)**
  - Learn systematically about non-functional requirements and design decisions in architecture
- **[Design It!: An Introduction to Architecting for Programmers](https://amzn.to/42UWjTm)**
  - Acquire practical practices from an architect's perspective
- **[Decision-Making Techniques for Software Architects: Leadership, Technology, and Product Management Utilization](https://amzn.to/4b2kRMj)**
  - Specifically explains how architects make decisions and lead teams
- **[97 Things Every Software Architect Should Know](https://amzn.to/40RZisY)**
  - A book summarizing the experiences and sayings of veteran architects in short chapters
- **[The Architect's Handbook: Building Software Architecture That Creates Value](https://amzn.to/3CTb3aU)**
  - Understand the perspective and necessary knowledge and skills of an architect systematically
- **[Requirements-Optimized Architecture Strategy](https://amzn.to/4gCH3hj)**
  - Introduces the idea of strategically linking requirements definition and architecture design
- **[Good Strategy, Bad Strategy](https://amzn.to/4hZTc0Q)**
  - Recommended for those who want to deeply understand the concept of "strategy." Learn the differences between good and bad strategies through examples

### Organization and Team Related
- **[Five Dialogues to Change Organizations](https://amzn.to/41iZu6h)**
  - Dialogue techniques and facilitation methods to promote organizational change
- **[Invitation to Engineering Organization Theory: Facing Uncertainty and Refactoring Organizations](https://amzn.to/4hD5Sei)**
  - Introduces organizational building and team growth strategies to face uncertainty
- **[Team Topologies: Adaptive Organizational Design to Deliver Valuable Software Quickly](https://amzn.to/3Qoj2j2)**
  - Ideas for optimizing team structure and collaboration in software development organizations
- **[Dynamic Re-teaming 2nd Edition: Effective Team Formation with Five Patterns](https://amzn.to/4i1fGyz)**
  - Introduces practical patterns for advancing team reorganization

### Blogs and Sites
- [bmf-tech.com - System Design Related Materials](https://bmf-tech.com/posts/%e3%82%b7%e3%82%b9%e3%83%86%e3%83%a0%e8%a8%ad%e8%a8%88%e9%96%a2%e9%80%a3%e3%81%ae%e8%b3%87%e6%96%99%e3%82%92%e8%aa%ad%e3%81%bf%e6%bc%81%e3%81%a3%e3%81%9f)
  - Summarizes reference materials related to system design and architecture
- [jlhood.com - How to Set the Technical Direction for Your Team](https://jlhood.com/how-to-set-team-technical-direction/)
  - Practical examples of setting technical direction at the team level
- [sarahtaraporewalla.com - Defining a Tech Strategy](https://sarahtaraporewalla.com/agile/design/architecture/Defining-a-Tech-Strategy)
  - Simply introduces the key points of architecture strategy
- [leaddev.com - Technical Strategy Power Chords](https://leaddev.com/technical-direction/technical-strategy-power-chords)
  - An article focusing on how to create the "core" of a technical strategy
- [staffeng.com](https://staffeng.com/guides/learning-materials/)
  - Knowledge system and learning guide for staff engineers
- [zenn.dev - How Architecture, Design, and Development Differ](https://zenn.dev/fujishiro/scraps/1144522710389a)
  - Clearly explains the difference between architecture and design
