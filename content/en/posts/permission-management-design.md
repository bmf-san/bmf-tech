---
title: Exploring Permission Management Design
description: 'Understand ACL, RBAC, and ABAC permission design methods with functional/data scope, control targets, and operational constraints.'
slug: permission-management-design
date: 2024-05-22T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ABAC
  - ACL
  - RBAC
  - Permission Management
  - System Design
  - Link Collection
translation_key: permission-management-design
---



# Overview
I researched case studies on permission management design and took notes.

# Research Notes
I organized the information I researched, but there are still things I don't fully understand, so it's not completely organized.

## Elements Composing Permissions
Permissions seem to be composed of the following elements:

- Who (Principal)
- What (Resource)
- What actions (Action) are allowed (ALLOW) or not allowed (DENY)


## Permission Design Methods
The following methods are generally considered for design:

- ACL (Access Control List)
  - Set permissions for each user
  - Permissions are managed in a list
- RBAC (Role Based Access Control)
  - Assign roles to users and assign permissions to roles
- ABAC (Attribute Based Access Control)
  - Set permissions based on attributes of users, resources, environments, etc.

The flexibility of permissions and the complexity of implementation are ACL < RBAC < ABAC.

## Perspectives on Permission Design
The following perspectives seem relevant for permission design:

- Scope of Permission Application
  - How far do permissions apply?
  - Functional scope
    - Can target a single function or a specific group of functions
    - ex. Whether the user information retrieval API can be used
  - Data scope
    - The most primitive scope of permission application
    - ex. Viewing permissions for user information → Can see name and age but not address
  - Sometimes both need to be considered, while other times only one is sufficient, but flexible design is required
- Control Targets of Permissions
  - On what basis are permissions applied?
    - Content
      - Apply permissions to specific functions, data, etc.
    - Context
      - Apply permissions based on whether specific conditions are met
        - ex. Only users with a transaction amount of 1 million yen or more can access
      - Is it even a target managed as a permission? If managed as a permission, how is the state managed?
    - Time
      - Apply permissions only during specific time periods
        - ex. Accessible only on weekdays
      - Access permissions are also within the scope of permissions
- Constraints of Permissions
  - Constraints such as priority of permissions and dependencies between permissions
  - Conflicting permissions
    - If a user has permissions to view only specific data but also has permissions to view all data, which takes precedence?
    - How to express the relationship of conflicting permissions?
      - Considerations for implementation include evaluating manually when adding new permissions, defining priorities based on attributes or roles, or using a set-theoretic approach
      - May need to consider from the perspective of user experience and security (principle of least privilege)
- Layers of Permission Application
  - Which layer of the system is permission applied to?
    - Application, database, network, OS, etc.
  - From where is permission application necessary?
  - The layer seems to be determined by the functional and data scope settings
- Handling of Administrator Permissions
  - How to set administrator permissions
  - Administrator permissions pose a security risk
    - Need to consider risk management perspectives such as the principle of least privilege, separation of duties, audit logs, and emergency operations
- Operational Flow of Permission Management
  - Need to consider an operational flow for proper permission management
  - Principle of least privilege
    - Grant only the minimum necessary permissions
  - Regular audits and reviews
    - Regularly check permission settings and remove unnecessary permissions
  - Centralized management
     - Provide an interface that makes it easy to understand the types of permissions and their application status
     - Consistent permission application flow
       - If permission management is uniquely implemented in each system, consistency is likely to be disrupted (I think)
     - Separation of permissions
       - Flexible permission settings should be possible
       - The control target of a single permission should not be too broad

## Required System Characteristics
Considered the system characteristics required for a system that manages permissions.

- Extensibility
  - Can add permissions with any flexibility
- Scalability
  - A highly flexible permission design is likely to increase system complexity and data volume
  - Consider capacity when users and permissions increase linearly
- Reliability
  - Ensure that adding or changing permissions does not adversely affect existing permissions
- Security
  - Adhere to the principle of least privilege, minimize blast radius, and prepare emergency operations for failures

# Impressions
Although it's a general impression as it depends on the industry and business domain of the service, I think permission management is particularly demanded in B2B services.

It seems like there isn't a well-established body of information or best practices.

It seems like a field where you could write a whole book, but there don't seem to be many related books.

I feel like I've identified the perspectives to consider, but the particularly challenging part seems to be how much flexibility to allow in permissions. I thought it would be necessary to expand the design while considering future business requirements to some extent.

# References
- [ja.wikipedia.org - Role-Based Access Control](https://ja.wikipedia.org/wiki/%E3%83%AD%E3%83%BC%E3%83%AB%E3%83%99%E3%83%BC%E3%82%B9%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1)
- [ja.wikpedia.org - Access Control List](https://ja.wikipedia.org/wiki/%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1%E3%83%AA%E3%82%B9%E3%83%88)
- [kenfdev.hateblo.jp - Challenges in Application Permission Design](https://kenfdev.hateblo.jp/entry/2020/01/13/115032)
- [knooto.info - Access Control Design in Systems](https://knooto.info/software-design-access-control/#top)
- [waterlow2013.hatenablog.com - Key Points for Permission Management DB Design Patterns](https://waterlow2013.hatenablog.com/entry/2017/01/27/233405)
- [www.lyricrime.com - About System Permission Methods](https://www.lyricrime.com/posts/access-control/)
- [zenn.dev/she_techblog - Considerations on Authorization Architecture (Reading Authorization Academy II)](https://zenn.dev/she_techblog/articles/6eff1f28d107be?redirected=1)
- [www.osohq.com - Authorization Academy](https://www.osohq.com/academy)
- [dzone.com - Access Control Acronyms: ACL, RBAC, ABAC, PBAC, RAdAC, and a Dash of CBAC](https://dzone.com/articles/acl-rbac-abac-pbac-radac-and-a-dash-of-cbac)
- [www.onelogin.com - RBAC vs. ABAC: Making the Right Decision](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [csrc.nist.gov - The NIST Model for Role-Based Access Control: Towards A Unified Standard](https://csrc.nist.gov/CSRC/media/Publications/conference-paper/2000/07/26/the-nist-model-for-role-based-access-control-towards-a-unified-/documents/sandhu-ferraiolo-kuhn-00.pdf)
- [www.internetacademy.jp - Six Types of Access Control Methods](https://www.internetacademy.jp/it/management/security/six-types-of-access-control-method.html)
- [butterflymx.com - Effective Access Control Design & Access Control System Planning](https://butterflymx.com/blog/access-control-design/)
- [satoricyber.com - A Comprehensive Guide to Role-Based Access Control Design](https://satoricyber.com/data-access-control/a-comprehensive-guide-to-role-based-access-control-design/)
- [tsapps.nist.gov - Role Engineering: Methods and Standards](https://tsapps.nist.gov/publication/get_pdf.cfm?pub_id=909664#:~:text=Even%20after%20a%20complete%20picture,known%20as%20%22role%20engineering%22.)
- [medium.muz.li - How to design access control system for Saas application](https://medium.muz.li/how-to-design-access-control-system-for-saas-application-b6455c944186)
- [uxdesign.cc - Designing permissions for a SaaS app](https://uxdesign.cc/design-permissions-for-a-saas-app-db6c1825f20e)
- [applis.io - Thoughts on Designing System Permission Management](https://applis.io/posts/how-to-manage-authorization)
- [link-and-motivation.hatenablog - Transforming Bitter Memories of Permission Management into a New Service](https://link-and-motivation.hatenablog.com/entry/20220401-authorization)
- [www.okta.com - RBAC vs. ABAC: Definitions and Usage](https://www.okta.com/jp/identity-101/role-based-access-control-vs-attribute-based-access-control/)
- [www.onelogin.com - RBAC vs. ABAC: Making the Right Decision](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [zenn.dev - How to Implement "This User Can Use This Feature"](https://zenn.dev/dove/articles/bc6933dbb39509)
- [zenn.dev - Designing Around Permissions](https://zenn.dev/dove/articles/8bed47a7a839ad)
- [qiita.com - Role-Based Access Control in Business Systems](https://qiita.com/kawasima/items/8dd7eda743f2fdcad78e)
- [note.com - Considering Role-Based Permission Design for SaaS](https://note.com/tumsat/n/nfbf88bfcbc29)