---
title: Basics of Authentication and Authorization
slug: authentication-authorization-basics
date: 2020-11-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - FIDO
  - IAM
  - LDAP
  - OAuth
  - OpenID Connect
  - SAML
  - SSO
  - Authorization
  - Authentication
translation_key: authentication-authorization-basics
---

# Overview
Since I have been involved in the development of authentication services recently, I wanted to summarize the basic concepts once again.

I am referencing the authentication and authorization feature from the [Software Design November 2020 issue](https://gihyo.jp/magazine/SD/archive/2020/202011).

# Relationship Between ID and Authentication/Authorization
- ID
    - Identity
        - Management unit for system usage
            - ex. User, Object, Organization, etc...
    - Identifier
        - Unit managed as data
            - ex. User ID managed in SaaS, SNS, and other services
    - Attribute
        - Each piece of information that constitutes the Identifier
            - ex. Identity is a person, Identifier is taro_yamada, Attributes are gender, date of birth, address, affiliation, etc...

# Identification, Authentication, and Authorization Process
- Identification, Authentication, Authorization
    - Identification
        - Uniquely identifies the Identifier
    - Authentication
        - Verifies the legitimacy of the Identifier (≈ whether it is the person) using credentials
    - Authorization
        - Decides how to assign permissions to the user

# Incorporating Authentication and Authorization
- ID and Permissions
    - Role
        - Defines patterns of combinations of services or permissions and assigns users to them
    - RBAC (Role-Based Access Control)
        - Consolidates access restrictions for specific resources into roles and assigns users to them
    - ABAC (Attribute-Based Access Control)
        - Implements access restrictions based on specific attributes
- Authentication Methods
    - Credential Characteristics
        - Something you know
            - Based on user memory
        - Something you have
            - Something the user possesses
        - Something you are
            - Based on the user's physical characteristics
    - Two-Factor Authentication (Multi-Factor Authentication)
        - Combines two or more credentials with different characteristics for authentication
    - Two-Step Authentication (Multi-Step Authentication)
        - Requires the authentication process to be performed two or more times

# Authentication and Authorization for Web Services
- In-house ID management database
- Social Login
    - Reduces the number of authentications
    - Reduces management load
- FIDO Authentication
    - First IDentity Online
    - A technology created in 2012 by the FIDO Alliance, a non-profit organization
    - Online authentication centered on biometric authentication
    - Uses public key cryptography
    - Stores authentication information in an Authenticator and performs authentication
- OAuth Authorization Framework
    - Open Authorization
    - A mechanism that allows an app (OAuth client) to access APIs on behalf of the user
        - Authorizes API access
- OAuth 2.0
    - OAuth 1.0 primarily targeted web services
    - OAuth 2.0 also targets mobile applications
    - HTTPS is mandatory

# Differences Between Authentication/Authorization for Web Services and APIs
- Differences
    - Web Services
        - After authentication and authorization, maintains login state in HTTP Cookies
    - APIs
        - Determines execution rights based on tokens
- OpenID Connect Protocol
    - Extends OAuth 2.0 to allow the transfer of identity information by including authentication results in the ID token
    - OAuth 2.0 is focused on authorization and does not implement a mechanism to transfer identity information containing authentication results
- Types of Authentication Flows in OpenID Connect 1.0
    - Authorization Code Flow
        - Transfers by exchanging authorization code and ID token (and access token)
    - Implicit Flow
        - Signature verification is required when transferring ID tokens
    - Hybrid Flow
        - A fusion of the above two
- Differences Between OAuth 2.0 and OpenID Connect 1.0
    - OAuth 2.0 is undefined regarding authentication. OpenID Connect 1.0 is defined.
    - OAuth 2.0 does not define the format of access tokens, but OpenID Connect 1.0 defines the format of ID tokens.
    - The flow until token issuance is the same, but OpenID Connect 1.0 requires the implementation of an API for obtaining user information called the Userinfo endpoint.

# Enterprise Authentication and Authorization
- In systems targeting enterprises, access management and access control are crucial.
- IAM (Identity and Access Management)
    - The concept of managing user and member ID information, granting authentication, authorization, and access rights.
    - For consumer IAM, improving UX is important, while for enterprises, corporate governance is a critical issue.
- Local Authentication
    - Authentication that manages ID, authentication, and authorization for each system.
    - Becomes difficult to manage as the number of users and systems increases.
- Directory Services
    - A service that records and manages the location, attributes, and configuration information of resources (systems, servers, applications, etc.) connected to the network.
    - LDAP (Lightweight Directory Access Protocol)
        - A communication protocol for accessing directory services.
    - Kerberos Authentication
        - A protocol for mutual authentication between server and client, confirming identity. One of the technologies that achieve single sign-on.
- SAML (Security Assertion Markup Language)
    - A protocol for authentication between different cloud services.

# References
The authentication and authorization feature from the [Software Design November 2020 issue](https://gihyo.jp/magazine/SD/archive/2020/202011) was easy to understand.