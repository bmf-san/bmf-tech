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
Since I've been involved in the development of authentication services recently, I wanted to revisit the basics and summarize them.

I referred to the authentication and authorization feature in the [Software Design November 2020 issue](https://gihyo.jp/magazine/SD/archive/2020/202011).

# Relationship Between ID and Authentication/Authorization
- ID
    - Identity
        - Management unit for system usage
            - ex. User, Object, Organization etc...
    - Identifier
        - Unit managed as data
            - ex. User IDs managed by SaaS, SNS, other services
    - Attribute
        - Each piece of information that makes up an Identifier
            - ex. Identity is a person, Identifier is taro_yamada, Attribute is gender, date of birth, address, affiliation etc...

# Process of Identification, Authentication, and Authorization
- Identification, Authentication, Authorization
    - Identification
        - Uniquely identify an Identifier
    - Authentication
        - Verify the legitimacy of an Identifier (≒ whether it is the person) using credentials
    - Authorization
        - Decide how to assign permissions to users

# Embedding Authentication and Authorization
- ID and Permissions
    - Role
        - A system that defines patterns of service or permission combinations and assigns users to them
    - RBAC (Role-Based Access Control)
        - A system that consolidates access restrictions to specific resources as a role and assigns users to it
    - ABAC (Attribute-Based Access Control)
        - A system that imposes access restrictions based on specific attributes
- Authentication Methods
    - Characteristics of Credentials
        - Something you know
            - Based on user memory
        - Something you have
            - Based on what the user possesses
        - Something you are
            - Based on the user's physical characteristics
    - Two-Factor Authentication (Multi-Factor Authentication)
        - Authentication combining two or more credentials with different characteristics
    - Two-Step Authentication (Multi-Step Authentication)
        - Authentication requiring the process to be performed more than once

# Authentication and Authorization for Web Services
- Own ID Management Database
- Social Login
    - Reduces the number of authentications
    - Reduces management burden
- FIDO Authentication
    - First IDentity Online
    - Technology created by the non-profit FIDO Alliance in 2012
    - Online authentication centered on biometric authentication
    - Uses public key cryptography
    - Stores authentication information in an authenticator and performs authentication
- OAuth Authorization Framework
    - Open Authorization
    - A system that allows apps (OAuth clients) to access APIs on behalf of users
        - Authorizes API access
- OAuth2.0
    - OAuth1.0 mainly targets web services
    - OAuth2.0 also targets mobile apps
    - HTTPS is mandatory

# Differences Between Authentication and Authorization for Web Services and APIs
- Differences
    - Web Services
        - Maintains login status in HTTP Cookie after authentication and authorization
    - API
        - Determines execution permission with a token
- OpenID Connect Protocol
    - An extension of OAuth2.0 that allows identity information, including authentication results, to be passed in an ID token
    - OAuth2.0 is specialized in authorization and does not implement a mechanism to pass identity information including authentication results
- Types of OpenID Connect1.0 Authentication Flows
    - Authorization Code Flow
        - Exchanges authorization code and ID token (and access token)
    - Implicit Flow
        - Signature verification is mandatory when passing ID token
    - Hybrid Flow
        - A fusion of the above two
- Differences Between OAuth2.0 and OpenID Connect1.0
    - OAuth2.0 does not define authentication. OpenID Connect1.0 does
    - OAuth2.0 does not define the format of access tokens, but OpenID Connect1.0 defines the format of ID tokens
    - The flow until token issuance is the same, but OpenID Connect1.0 requires the implementation of a Userinfo endpoint API to obtain user information

# Enterprise Authentication and Authorization
- In systems targeting enterprises, access management and access control are important
- IAM (Identity and Access Management)
    - A concept that manages user and member ID information, authentication, authorization, and grants access rights
    - For consumer IAM, improving UX is important, while for enterprises, corporate governance is a key issue
- Local Authentication
    - Authentication that manages ID, authentication, and authorization for each system
    - Management becomes difficult as users and systems increase
- Directory Services
    - A service that collectively records and manages the location, attributes, and configuration information of resources (systems, servers, applications, etc.) connected to the network
    - LDAP (Lightweight Directory Access Protocol)
        - A communication protocol for accessing directory services
    - Kerberos Authentication
        - A protocol for mutual authentication between server and client, confirming identity. One of the technologies that enables single sign-on
- SAML (Security Assertion Markup Language)
    - A protocol for authentication between different cloud services

# References
The authentication and authorization feature in the [Software Design November 2020 issue](https://gihyo.jp/magazine/SD/archive/2020/202011) was easy to understand.
