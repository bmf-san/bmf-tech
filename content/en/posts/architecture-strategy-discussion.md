---
title: Considering Architecture Strategy
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
In software development, even if you are not in a clear position like a CTO or architect, there are opportunities to feel and think about the "necessity of architecture strategy."

- "I want to summarize the technical policy on-site, but I don't know where to start."
- "I want to create an architecture that takes into account future scalability and organizational growth."
- "I want to know how to deal with the technical debt that has already accumulated."

In this article, I will organize the **necessity and utilization of architecture strategy** to answer these questions.

## Background and Disclaimer
Sometimes I feel, "Huh? What kind of strategy is this plan being executed under?" or "What exactly is a strategy?"

At such times, abstract thoughts swirl in my mind, but I couldn't articulate them properly, so I thought I should organize my thoughts and study to summarize the content.

Since I have no experience as a CTO or architect, my practical experience in thinking about architecture strategy is limited, so it may just be a figment of my imagination.

That said, I would appreciate it if you could quietly contact me if there is anything clearly incorrect.

## 1. Introduction
Before discussing architecture strategy, let's first clarify the terms "architecture" and "strategy." These concepts are often used in software development, but they can be surprisingly vague or understood differently by different people.

This article will outline the following:

- **What is architecture strategy?**
- **Why is it necessary, and what benefits does it provide?**
- **How should it be formulated (process and specific examples)?**
- **Necessary skills, points of caution, and reference materials.**

## 2. What is Architecture Strategy?
### 2.1 What is Architecture?
In software development, "architecture" refers to the **overall structure and design policy of the system.** Specifically, it includes elements such as:

1. **System Structure**
   - Components that make up the system (e.g., microservices, databases, API gateways, etc.)
   - Interrelationships between components (data flow, communication methods, etc.)
2. **Design Principles and Patterns**
   - Design methodologies such as layered architecture, clean architecture, microservices, DDD, etc.
   - Best practices like SOLID principles, CQRS, and Event Sourcing.
3. **Consideration of Non-Functional Requirements**
   - Scalability, availability, security, performance, maintainability, and extensibility, etc.
4. **Technology Selection**
   - Selection of programming languages, frameworks, databases, cloud providers, etc.
   - API design (REST, GraphQL, gRPC, etc.)
5. **Development and Operations Processes**
   - Design of CI/CD pipelines, monitoring/logging, team development flow (agile or waterfall, etc.)

These elements are just examples and are not exhaustive.

**Architecture** is a higher-level concept than **design** and indicates the "broad framework of the system" or "design policy." Based on that guidance, the stage of concretizing individual modules, classes, and data models is called **design.**

### 2.2 What is Strategy?
Strategy refers to the **broad direction or plan for achieving goals.** While tactics indicate "specific means" or "measures," strategy sets the direction for the organization or project as a whole from a more long-term and comprehensive perspective.

According to [Good Strategy, Bad Strategy](https://amzn.to/4hZTc0Q), the quality of a strategy can be described as follows:

- **Key Points of a Good Strategy**
  - There are specific action guidelines, and it is clear what should be done.
  - Appropriate choices and focus are made, clearly defining "what to do" and "what not to do."
  - It consists of three elements: diagnosis (current situation analysis) → basic policy → action.
- **Characteristics of a Bad Strategy**
  - It is vague and lacks specificity.
  - It avoids facing significant problems directly and sidesteps issues.
  - It merely sets goals without a clear method for achieving them.

### 2.3 Definition of Architecture Strategy
Combining the concepts of "architecture" and "strategy" gives us **architecture strategy.** In other words, it can be defined as **"a policy or plan that systematically shows how to build and evolve systems within an organization or project."**

Specifically, it includes elements such as:

1. **Clarification of Vision and Goals**
   - Indicating technical direction while aligning with business strategy.
   - Example: Cloud-native transformation, adoption of event-driven architecture, migration to microservices, etc.
2. **Architecture Principles and Guidelines**
   - Criteria for technology selection, principles of system design, standardization guidelines for operations and maintenance, etc.
3. **Selection of Technology Stack**
   - Establishing criteria for adopting programming languages, databases, frameworks, etc.
4. **Scalability and Extensibility of the System**
   - Scalable configurations, load balancing, service segmentation strategies, etc.
5. **Security and Compliance**
   - Authentication/authorization, data encryption, compliance measures (e.g., GDPR, SOC2).
6. **Operations and Monitoring Strategy**
   - Standardization of logging/monitoring and incident response processes.

## 3. Objectives and Benefits of Architecture Strategy
By formulating an architecture strategy, the following benefits can be obtained:

1. **Unification of Business and Technical Direction**
   - Technical decision-making is not ad-hoc but can be made strategically and long-term.
2. **Suppression of Technical Debt**
   - Without a plan, technical debt tends to accumulate. A prior strategy can help mitigate this debt.
3. **Improvement of Development Speed and Productivity**
   - With clear guidelines, unnecessary discussions within the team can be reduced, allowing for smoother progress.
4. **Enhancement of System Quality and Stability**
   - By consciously addressing non-functional requirements (scalability, availability, security, etc.), it is possible to reduce failures and security risks.
5. **Team Consensus**
   - With common goals and policies, decision-making speeds up, and the team is more likely to align in the same direction.

A **good** architecture strategy eliminates all waste in a project and supports efficient resource allocation toward achieving business goals.

## 4. Areas Where Architecture Strategy is Required
Architecture strategy can be formulated and applied at various scopes, from team units to the entire organization.

- **Organization-Wide Level**
  - Aligning with the overall IT strategy and defining common infrastructure and security policies.
  - Formulating a long-term roadmap aimed at overall optimization.
- **Department Level**
  - Optimizing according to department-specific business requirements.
  - Deciding policies considering system integration and data sharing within the department.
- **Team Level**
  - Detailed architecture design for a specific service or application.
  - Formulating strategies directly related to practical matters such as technology selection, implementation methods, and testing policies.

It is important to note that as the scope widens, the costs of coordination and the difficulty of consensus increase.

## 5. Process for Formulating Architecture Strategy
The process of formulating an architecture strategy is to clarify "what business objectives to achieve, how to realize them through technical approaches, and how to operate and govern them."

Whether the architecture strategy should be formulated top-down, bottom-up, or through a hybrid approach depends on the case and varies according to the organization's culture and business situation.

Below is an example of steps in the process of formulating an architecture strategy.

### 1. Confirm Business Strategy and Vision
1. **Understanding Business Goals and Strategies**
   - Understand the business goals and strategies set by management and business units, and consider how architecture can contribute to the business.

2. **Identifying Stakeholders**
   - Identify key stakeholders involved in the architecture strategy (management, development departments, business units, security personnel, etc.) and organize their requirements, constraints, and expectations.

### 2. Understanding Current Architecture (As-Is)
1. **Investigating Current Systems and Architecture**
   - Investigate the system composition, data flow, utilized technologies, and operational processes in the target area and document them.

2. **Extracting Issues and Risks**
   - Identify bottlenecks in business and technology, cost structures, personnel dependencies, operational burdens, and technical debt.

### 3. Formulating Architecture Principles and Direction
1. **Setting Architecture Principles**
   - Define which architecture characteristics (non-functional requirements) to prioritize.
   - Clarify trade-offs (e.g., balancing flexibility and performance).

2. **Policy for Rough Direction and Prioritization of the Roadmap**
   - Determine which areas to address first and where the return on investment is high, outlining high-level direction.
   - Consider the priority of business requirements and technical feasibility together.

### 4. Defining the Desired State (To-Be)
1. **Designing Architecture**
   - Define the desired state from perspectives such as business architecture (business flow, organizational structure), application architecture (structure of services and systems), and technology architecture (infrastructure and network configuration).
     - Tools like the C4 model can be helpful.

2. **Examining Scenarios and Use Cases**
   - Specify major scenarios and use cases anticipated in business and verify their alignment with the architecture.
   - Consider operational scenarios that include security and availability requirements.

3. **Architecture Evaluation and Selection (Option Comparison)**
   - It is common to compare multiple architecture proposals and evaluate their respective merits and demerits.
   - Consider the technology stack (selection of public cloud vendors, types of databases, programming languages, etc.) and operational models (on-premises/cloud/hybrid, etc.).

### 5. Gap Analysis and Roadmap Creation
1. **Gap Analysis Between As-Is and To-Be**
   - Identify the differences between the current and target architecture in terms of technology, organization, cost, skills, etc.
   - Consider which gaps to close, when, and how, and prioritize them.

2. **Formulating Migration Strategy**
   - Design migration patterns, such as whether a large-scale overhaul is necessary, whether to migrate gradually to microservices, and how to handle legacy systems.
   - Consider methods that reduce risks during the introduction phase, such as the "Strangler Pattern" or "Rehosting (Lift & Shift)."

3. **Creating the Roadmap**
   - Plan when and how to proceed with architecture changes, setting milestones for each phase.
   - Roughly plan deliverables and success indicators (KPI/KGI), necessary resources, budget, and structure for each phase.

### 6. Preparing the Execution Plan
1. **Clarifying Execution and Project Plans**
   - Based on the roadmap, break down each measure into project units and detail schedules, resources, and structures.
   - Use project management tools (Jira, Confluence, etc.) to manage progress and risks.

2. **Informing and Building Consensus within the Organization**
   - Present the strategy content to stakeholders to gain approval and build a cooperative framework.

### 7. Continuous Evaluation and Feedback
1. **Monitoring During Implementation and Operation Phases**
   - Not only proceed according to the roadmap but also regularly check indicators (KPI and monitoring results of non-functional requirements) and evaluate the situation.
   - Introduce mechanisms like DevOps and CI/CD to ensure smooth implementation, testing, and release of technical changes.

2. **Regular Review of Architecture**
   - Regularly review the architecture in line with changes in the business environment and technology trends.
   - Incorporate improvement proposals in governance meetings (architecture review meetings, etc.) and update the roadmap.

3. **Organizational Learning and Deployment**
   - Share success and failure cases to accumulate know-how on architecture formulation and operation within the organization.
   - Update practices and standards to apply them to the next project.

## 6. Specific Example: Architecture Strategy in HR and Labor SaaS
Here, I will provide a simple example of an architecture strategy for a hypothetical "HR and Labor SaaS."

### 1. Technical Vision
**"Achieve strict authority management for large-scale tenants while balancing rapid feature expansion and operational efficiency."**

- **Background**
  - Increased usage in tenants with complex organizational hierarchies, such as large corporations.
  - A permissions model that spans multiple functions like attendance management, payroll calculation, and year-end adjustments is necessary.
  - A **"common authority management foundation"** that is easy to extend and maintain, and can also handle audits is essential.

### 2. Current Situation Analysis
1. **Existing Systems**
   - A "modular monolith" structure where each function implements its own authority logic.
   - Permission definitions are decentralized, making role settings and access control for large organizations cumbersome.

2. **Issues and Challenges**
   - **Development Burden**: Each time a new feature is added, it is necessary to increase permission flags or roles.
   - **Operational Risks**: Role names and permission requirements are inconsistent, making management and audits difficult.
   - **Scalability**: Unable to handle large-scale tenants with many simultaneous users.

### 3. Architecture Principles
1. **Cloud-Native Centered on GCP**
   - Utilize GKE (Google Kubernetes Engine) and Cloud Run for foundational resources.
   - For application logging/monitoring, use Cloud Logging/Cloud Monitoring as a base, while also incorporating tools like Datadog as needed.

2. **Single "Permission Service" Microservice**
   - Separate authentication (Auth) and authorization (Authorization), and create a common API for the authorization foundation.
   - Combine RBAC (Role-Based Access Control) and ABAC (Attribute-Based) to accommodate the complex hierarchies of large enterprises.

3. **Thorough Implementation of Infrastructure as Code and CI/CD**
   - Manage infrastructure as code using Terraform/Cloud Deployment Manager.
   - Improve operational efficiency with Cloud Build or GitHub Actions + Cloud Run/GKE auto-deployment.

4. **Standardization of Operations and Audit Logs**
   - Aggregate all permission changes and access requests in Cloud Logging.
   - Store audit data in BigQuery to facilitate reporting and analysis.

### 4. Example of Technology Stack Selection
1. **Authentication**
   - Utilize Google Identity Platform or Auth0 to achieve SSO with OIDC/OAuth2.
   - A structure that easily accommodates internal/external ID linkage.

2. **Permission Management (Authorization)**
   - Run a **custom permission service** (based on Go or Python) on GKE or Cloud Run.
   - Data Store: Manage role/policy definitions with Cloud SQL (PostgreSQL).
   - Use Memorystore (Redis) for caching for fast reference.

3. **Each SaaS Module**
   - Containerize functionality services such as attendance management, payroll calculation, and year-end adjustments, referencing the permission service.
   - API Gateway: Integrate traffic control and authorization flows with Cloud Endpoints or Istio (Service Mesh).

4. **Monitoring and Log Collection**
   - Standard use of **Cloud Logging/Cloud Monitoring** across all containers.
   - Error alerts and SLO/SLA management may also be linked to Opsgenie/PagerDuty.

### 5. Roadmap (Example: 3 Phases)
1. **Phase 1 (Up to 3 months)**
   - **Design and PoC of Permission Foundation**
     - Formulate data models for role/policy definitions.
     - Separate services for Auth (authentication) and Authorization (authorization).
   - **Build CI/CD Pipeline**
     - Demonstrate automated deployment with Cloud Build or GitHub Actions + Terraform integration.

2. **Phase 2 (Up to 6 months)**
   - **Introduce Permission Service to Major Modules**
     - Migrate existing roles starting with attendance management and user management.
     - Integrate some remaining on-premises systems (e.g., VPN connection → VPC peering).
   - **Establish Audit Logs**
     - Accumulate change history and access history in BigQuery and prototype reporting functionality.

3. **Phase 3 (Up to 12 months)**
   - **Apply to All Modules**
     - Gradually migrate microservices for payroll calculation, year-end adjustments, etc., to the permission service.
   - **Load Testing and Optimization for Large-Scale Tenants**
     - Validate Redis caching and scaling policies before operational deployment.
   - **Strengthen Governance for Long-Term Operations**
     - Fully operationalize policy management UI and automatic generation of audit reports.

## 7. Skills and Knowledge Required for Consideration
To think about and execute architecture strategy, the following skills and knowledge are important:

1. **Fundamentals of Software Design**
   - Basic design methodologies such as clean architecture, microservices, monoliths, etc.
   - Best practices like SOLID principles and DDD.

2. **Holistic Perspective of the Entire System**
   - Understanding architecture characteristics such as scalability, security, performance, availability, and operational management, along with the ability to analyze trade-offs.

3. **Understanding Business Requirements**
   - Ability to explain how a particular architecture contributes to the business.

4. **Awareness of Technology Trends**
   - Ability to keep up with the latest technologies and industry trends to broaden architecture options.

5. **Leadership and Communication**
   - Skills for communication, consensus building, and decision-making with teams and stakeholders.

6. **Understanding Organizational Strategy**
   - Ability to understand how architecture strategy connects with business strategy and formulate strategies aligned with organizational policies.

## 8. Conclusion
In this article, I introduced **architecture strategy**, its necessity, the formulation process, and specific examples. Reflecting on the key points, they can be summarized into three main points:

1. **Long-Term Perspective Connecting Business and Technology**
   - Not just technical selection, but showing policies and roadmaps linked to business goals.
2. **Visualizing Technical Debt and Trade-offs to Foster Consensus**
   - Identifying non-functional requirements and risks, and involving multiple stakeholders to find optimal solutions.
3. **Continuous Verification and Improvement**
   - Since the business environment and technology trends are always changing, architecture strategies should be reviewed regularly.

By clarifying architecture strategy, the direction of the team aligns, and in the long run, development efficiency, quality, and business promotion capabilities are expected to improve significantly. I want to assess whether a strategy is necessary based on the situation of my product or organization and incorporate appropriate approaches.

## 9. Reference Materials
### Books
- [**The Path of Staff Engineer: Architecting to Become an Excellent Technical Professional**](https://amzn.to/4aXEjd3)
  - Explains behaviors and techniques for leading teams and organizations as a software engineer.
- **[Fundamentals of Software Architecture: A Systematic Approach Based on Engineering](https://amzn.to/3CVjMcv)**
  - Systematically learn about non-functional requirements and design decisions in architecture.
- **[Design It! An Introduction to Architecture for Programmers](https://amzn.to/42UWjTm)**
  - Acquire practical practices from an architect's perspective.
- **[Decision-Making Techniques for Software Architects: Utilizing Leadership, Technology, and Product Management](https://amzn.to/4b2kRMj)**
  - Specifically explains how architects make decisions and lead teams.
- **[97 Things Every Software Architect Should Know](https://amzn.to/40RZisY)**
  - A collection of veteran architect's experiences and quotes summarized in short chapters.
- **[The Architect's Textbook: Building Software Architecture that Creates Value](https://amzn.to/3CTb3aU)**
  - Systematically understand the architect's perspective and necessary knowledge and skills.
- **[Optimal Architecture Strategy for Requirements](https://amzn.to/4gCH3hj)**
  - Introduces a way of thinking that strategically connects requirements definition and architecture design.
- **[Good Strategy, Bad Strategy](https://amzn.to/4hZTc0Q)**
  - Recommended for those who want to deeply understand the concept of "strategy." Learn the differences between good and bad strategies through examples.

### Organization and Team Related
- **[Five Dialogues to Change Organizations](https://amzn.to/41iZu6h)**
  - Dialogue techniques and facilitation methods to promote organizational change.
- **[Invitation to Engineering Organization Theory: Refactoring Thinking and Organizations to Face Uncertainty](https://amzn.to/4hD5Sei)**
  - Introduces strategies for building organizations and team growth to face uncertainty.
- **[Team Topologies: Adaptive Organization Design for Delivering Valuable Software Quickly](https://amzn.to/3Qoj2j2)**
  - A way of thinking to optimize team structures and collaboration in software development organizations.
- **[Dynamic Reteaming, 2nd Edition: Effective Team Formation through Five Patterns](https://amzn.to/4i1fGyz)**
  - Introduces practical patterns for how to proceed with team reorganization.

### Blogs and Sites
- [bmf-tech.com - Read through materials related to system design](https://bmf-tech.com/posts/%e3%82%b7%e3%82%b9%e3%83%86%e3%83%a0%e8%a8%ad%e8%a8%88%e9%96%a2%e9%80%a3%e3%81%ae%e8%b3%87%e6%96%99%e3%82%92%e8%aa%ad%e3%81%bf%e6%bc%81%e3%81%a3%e3%81%9f)
  - Summarizes reference materials on system design and architecture.
- [jlhood.com - How to Set the Technical Direction for Your Team](https://jlhood.com/how-to-set-team-technical-direction/)
  - Practical examples of setting technical policies at the team level.
- [sarahtaraporewalla.com - Defining a Tech Strategy](https://sarahtaraporewalla.com/agile/design/architecture/Defining-a-Tech-Strategy)
  - A simple introduction to the key points of architecture strategy.
- [leaddev.com - Technical strategy power chords](https://leaddev.com/technical-direction/technical-strategy-power-chords)
  - An article focusing on how to create the "core" of technical strategy.
- [staffeng.com](https://staffeng.com/guides/learning-materials/)
  - Knowledge systems and learning guides for staff engineers.
- [zenn.dev - What is the difference between Architecture, Design, and Development?](https://zenn.dev/fujishiro/scraps/1144522710389a)
  - A clear explanation of the differences between architecture and design.