---
title: Investigating Permission Management Design
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
  - Resource Links
translation_key: permission-management-design
---

# Overview
I researched examples of permission management design and wanted to take some notes.

# Research Notes
I organized the information I researched, but there are still things I don't fully understand, so it's not completely organized.

## Elements that Constitute Permissions
Permissions can be thought to consist of the following elements:

- Who (Principal)
- What (Resource)
- What can be done (ALLOW) or not done (DENY)

## Permission Design Methods
The following methods are generally considered for design:

- ACL (Access Control List)
  - Set permissions for each user
  - Permissions are managed in a list
- RBAC (Role Based Access Control)
  - Assign roles to users and assign permissions to roles
- ABAC (Attribute Based Access Control)
  - Set permissions based on attributes of users, resources, environment, etc.

The flexibility of permissions and the complexity of implementation are ranked as ACL < RBAC < ABAC.

## Perspectives on Permission Design
There are several perspectives to consider in permission design:

- Scope of Permission
  - How far do permissions apply?
  - Functional Scope
    - It can target a single function or a specific group of functions.
    - e.g., Whether the user can access the user information retrieval API.
  - Data Scope
    - The most primitive application range of permissions.
    - e.g., Viewing permissions for user information → Can see name and age, but not address.
  - Sometimes both need to be considered, and sometimes only one is sufficient, but flexible design is required.
- Control Targets of Permissions
  - What criteria are used to apply permissions?
    - Content
      - Apply permissions to specific functions, data, etc.
    - Context
      - Apply permissions based on whether specific conditions are met.
        - e.g., Only users with a payment amount of 1 million yen or more can access.
      - Is it even a target managed as a permission? If it is managed as a permission, how is the state managed?
    - Time
      - Apply permissions only during specific time periods.
        - e.g., Accessible only on weekdays.
      - Access permissions also fall under the scope of permissions.
- Constraints on Permissions
  - Constraints such as priority of permissions and dependencies between permissions.
  - Conflicting Permissions
    - If a user has permission to view only specific data and also has permission to view all data, which takes precedence?
    - How to express the relationship of conflicting permissions?
      - Considerations for implementation include whether to evaluate manually when adding new permissions, define priorities based on attributes or roles, or evaluate using a set-theoretic approach.
      - It may be necessary to consider from the perspective of user experience and security (principle of least privilege).
- Layers of Permission Application
  - Which layer of the system is permissions applied to?
    - Application, database, network, OS, etc.
  - Where is permission application necessary?
  - The layer is likely determined by the functional and data scope settings.
- Handling of Administrator Permissions
  - How to set administrator permissions.
  - Administrator permissions carry security risks.
    - Need to consider risk management perspectives such as the principle of least privilege, separation of duties, audit logs, and emergency operation preparations.
- Operational Flow of Permission Management
  - Need to consider an operational flow to ensure permissions are managed appropriately.
  - Principle of Least Privilege
    - Grant only the minimum necessary permissions.
  - Regular Audits and Reviews
    - Regularly check permission settings and remove unnecessary permissions.
  - Centralized Management
    - Provide an interface that makes it easy to understand types of permissions and their application status.
    - Consistent permission application flow.
      - If permission management is uniquely implemented in each system, consistency is likely to be disrupted.
    - Separation of Permissions
      - Ensure flexible permission settings.
      - The control target of a single permission should not be too broad.

## Required System Characteristics
I considered the system characteristics required for a system that manages permissions.

- Extensibility
  - Permissions can be added with arbitrary flexibility.
- Scalability
  - A highly flexible permission design is likely to anticipate increased system complexity and data volume.
  - Consider capacity when users and permissions increase linearly.
- Reliability
  - Ensure that adding or changing permissions does not adversely affect existing permissions.
- Security
  - Adhere to the principle of least privilege, minimize blast radius, and prepare for emergency operations in case of failures.

# Thoughts
Depending on the industry and domain of the service, permission management seems to be one of the functions that is particularly required for B2B services.

I felt that there isn't a well-established body of information or best practices in this area.

It seems like a field that could fill an entire book, but there don't seem to be many related publications.

I feel like I've identified some perspectives to consider, but the particularly challenging part is determining how much flexibility in permissions to allow. I believe it's necessary to design with future business requirements in mind to some extent.

# References
- [ja.wikipedia.org - Role-Based Access Control](https://ja.wikipedia.org/wiki/%E3%83%AD%E3%83%BC%E3%83%AB%E3%83%99%E3%83%BC%E3%82%B9%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1)
- [ja.wikpedia.org - Access Control List](https://ja.wikipedia.org/wiki/%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1%E3%83%AA%E3%82%B9%E3%83%88)
- [kenfdev.hateblo.jp - Challenges in Permission Design in Applications](https://kenfdev.hateblo.jp/entry/2020/01/13/115032)
- [knooto.info - Designing Access Control in Systems](https://knooto.info/software-design-access-control/#top)
- [waterlow2013.hatenablog.com - Key Points for Database Design in Permission Management](https://waterlow2013.hatenablog.com/entry/2017/01/27/233405)
- [www.lyricrime.com - About Access Control Methods in Systems](https://www.lyricrime.com/posts/access-control/)
- [zenn.dev/she_techblog - Considerations on Authorization Architecture](https://zenn.dev/she_techblog/articles/6eff1f28d107be?redirected=1)
- [www.osohq.com - Authorization Academy](https://www.osohq.com/academy)
- [dzone.com - Access Control Acronyms: ACL, RBAC, ABAC, PBAC, RAdAC, and a Dash of CBAC](https://dzone.com/articles/acl-rbac-abac-pbac-radac-and-a-dash-of-cbac)
- [www.onelogin.com - RBAC vs. ABAC: Making the Right Decision](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [csrc.nist.gov - The NIST Model for Role-Based Access Control](https://csrc.nist.gov/CSRC/media/Publications/conference-paper/2000/07/26/the-nist-model-for-role-based-access-control-towards-a-unified-/documents/sandhu-ferraiolo-kuhn-00.pdf)
- [www.internetacademy.jp - Six Types of Access Control Methods](https://www.internetacademy.jp/it/management/security/six-types-of-access-control-method.html)
- [butterflymx.com - Effective Access Control Design & Planning](https://butterflymx.com/blog/access-control-design/)
- [satoricyber.com - A Comprehensive Guide to Role-Based Access Control Design](https://satoricyber.com/data-access-control/a-comprehensive-guide-to-role-based-access-control-design/)
- [tsapps.nist.gov - Role Engineering: Methods and Standards](https://tsapps.nist.gov/publication/get_pdf.cfm?pub_id=909664#:~:text=Even%20after%20a%20complete%20picture,known%20as%20%22role%20engineering%22.)
- [medium.muz.li - How to Design Access Control System for SaaS Application](https://medium.muz.li/how-to-design-access-control-system-for-saas-application-b6455c944186)
- [uxdesign.cc - Designing Permissions for a SaaS App](https://uxdesign.cc/design-permissions-for-a-saas-app-db6c1825f20e)
- [applis.io - Considerations for Designing Permission Management in Systems](https://applis.io/posts/how-to-manage-authorization)
- [link-and-motivation.hatenablog - A Bitter Memory of Permission Management Transformed into a New Service](https://link-and-motivation.hatenablog.com/entry/20220401-authorization)
- [www.okta.com - RBAC vs. ABAC: Definitions and Usage](https://www.okta.com/jp/identity-101/role-based-access-control-vs-attribute-based-access-control/)
- [www.onelogin.com - RBAC vs. ABAC: Making the Right Decision](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [zenn.dev - How to Implement User Functionality Access](https://zenn.dev/dove/articles/bc6933dbb39509)
- [zenn.dev - Designing Around Permissions](https://zenn.dev/dove/articles/8bed47a7a839ad)
- [qiita.com - Role-Based Access Control in Business Systems](https://qiita.com/kawasima/items/8dd7eda743f2fdcad78e)
- [note.com - Considering Role-Based Permission Design for SaaS](https://note.com/tumsat/n/nfbf88bfcbc29)