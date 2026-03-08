---
title: Summary of OAuth 2.0 Specifications
slug: oauth-2-0-spec-summary
date: 2026-01-23T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - OAuth
  - Authentication
  - Authorization
description: Summary based on RFC 6749 (OAuth 2.0 Authorization Framework) and RFC 6750 (Bearer Token Usage).
translation_key: oauth-2-0-spec-summary
---



## RFC Terminology (RFC 2119)

| Term | Meaning |
|-----|------|
| **MUST** / **REQUIRED** / **SHALL** | Absolute requirement |
| **MUST NOT** / **SHALL NOT** | Absolute prohibition |
| **SHOULD** / **RECOMMENDED** | Recommended (should be followed unless there is a specific reason not to) |
| **SHOULD NOT** / **NOT RECOMMENDED** | Not recommended (should be avoided unless there is a specific reason not to) |
| **MAY** / **OPTIONAL** | Optional (may or may not be implemented) |

---

## Overview

OAuth 2.0 is an authorization framework that allows third-party applications to obtain **limited access** to HTTP services.

In the traditional client-server authentication model, third parties needed to use the resource owner's credentials directly. OAuth 2.0 solves this issue by introducing **access tokens**.

---

## Four Roles

| Role | Description |
|-------|------|
| **Resource Owner** | Entity that grants access to the protected resource (usually the end-user) |
| **Resource Server** | Server hosting the protected resources and accepting requests with access tokens |
| **Client** | Application accessing the protected resources on behalf of the resource owner with authorization |
| **Authorization Server** | Server that authenticates the resource owner and issues access tokens after obtaining authorization |

```mermaid
sequenceDiagram
    participant RO as Resource Owner
    participant C as Client
    participant AS as Authorization Server
    participant RS as Resource Server

    C->>RO: (A) Authorization Request
    RO-->>C: (B) Authorization Grant
    C->>AS: (C) Authorization Grant
    AS-->>C: (D) Access Token
    C->>RS: (E) Access Token
    RS-->>C: (F) Protected Resource
```

---

## Four Grant Types

### 1. Authorization Code Grant

**The most recommended flow. Suitable for web applications.**

```mermaid
sequenceDiagram
    participant RO as Resource Owner
    participant UA as User Agent<br/>(Browser)
    participant C as Client
    participant AS as Authorization Server

    C->>UA: (A) Redirect to Authorization Endpoint
    UA->>AS: (A) Authorization Request<br/>(client_id, redirect_uri, scope, state)
    AS->>RO: (B) Authenticate & Authorize
    RO-->>AS: (B) Grant Access
    AS-->>UA: (C) Authorization Code
    UA-->>C: (C) Redirect with Code
    C->>AS: (D) Token Request<br/>(code, client_id, client_secret)
    AS-->>C: (E) Access Token<br/>(+ Refresh Token)
```

| Step | Description |
|---------|------|
| (A) | Client redirects the user agent to the authorization endpoint |
| (B) | Authorization server authenticates the resource owner and grants/denies access |
| (C) | Authorization server issues an authorization code and redirects |
| (D) | Client exchanges the authorization code for an access token at the token endpoint |
| (E) | Authorization server issues an access token (and refresh token) |

**Features**:
- Authorization code is short-lived (RECOMMENDED: within 10 minutes)
- Refresh tokens can be issued
- Client authentication is performed

### 2. Implicit Grant

**For JavaScript applications running in the browser.**

```mermaid
sequenceDiagram
    participant RO as Resource Owner
    participant UA as User Agent<br/>(Browser)
    participant C as Client<br/>(JavaScript)
    participant AS as Authorization Server

    C->>UA: (A) Redirect to Authorization Endpoint
    UA->>AS: (A) Authorization Request<br/>(response_type=token)
    AS->>RO: (B) Authenticate & Authorize
    RO-->>AS: (B) Grant Access
    AS-->>UA: (C) Access Token in URL Fragment<br/>(#access_token=xxx)
    UA-->>C: (D) Extract Token from Fragment
```

| Feature | Description |
|-----|------|
| Token Acquisition | Access token is obtained directly from the authorization endpoint |
| Refresh Token | Not issued |
| Client Authentication | Not performed |
| Security | Lower than authorization code grant (token exposed in URL fragment) |

**Note**: In OAuth 2.1 (draft), **removed** for security reasons. Use authorization code grant with PKCE instead.

### 3. Resource Owner Password Credentials Grant

**Flow where the client directly receives the user's ID/password.**

```mermaid
sequenceDiagram
    participant RO as Resource Owner
    participant C as Client
    participant AS as Authorization Server

    RO->>C: (A) Username & Password
    C->>AS: (B) Token Request<br/>(grant_type=password,<br/>username, password)
    AS-->>C: (C) Access Token<br/>(+ Optional Refresh Token)
```

| Feature | Description |
|-----|------|
| Use Case | Migration from legacy systems, highly trusted clients |
| Risk | Client has access to credentials, risk of misuse |
| Recommendation | Use only if other grant types are not available |

**Note**: In OAuth 2.1 (draft), **removed** for security reasons.

### 4. Client Credentials Grant

**Flow for accessing resources with the client's own authority.**

```mermaid
sequenceDiagram
    participant C as Client
    participant AS as Authorization Server

    C->>AS: (A) Client Authentication<br/>(client_id, client_secret)
    AS-->>C: (B) Access Token
```

| Feature | Description |
|-----|------|
| Use Case | Machine-to-machine (M2M) communication, batch processing |
| Resource Owner | Client itself is the resource owner |
| User Involvement | Not required |

---

## Tokens

### Access Token

| Item | Description |
|-----|------|
| Role | Credential for accessing protected resources |
| Expiry | Short-lived (SHOULD: less than 1 hour, RFC 6750) |
| Format | Not specified by the specification (implementation-dependent, JWT is common) |
| Scope | Limits the range of access |

### Refresh Token

| Item | Description |
|-----|------|
| Role | Used to refresh the access token |
| Expiry | Long-lived (days to weeks) |
| Usage Location | Only at the token endpoint (not sent to the resource server) |
| Issuance | Can be issued in authorization code grant, not in implicit |

```mermaid
sequenceDiagram
    participant C as Client
    participant AS as Authorization Server
    participant RS as Resource Server

    C->>AS: (A) Authorization Grant
    AS-->>C: (B) Access Token + Refresh Token

    C->>RS: (C) Access Token
    RS-->>C: (D) Protected Resource

    Note over C,RS: Time passes, token expires

    C->>RS: (E) Access Token (expired)
    RS-->>C: (F) Invalid Token Error

    C->>AS: (G) Refresh Token
    AS-->>C: (H) New Access Token<br/>(+ Optional New Refresh Token)
```

---

## Endpoints

### Authorization Endpoint

| Item | Description |
|-----|------|
| Role | Obtains authentication and authorization from the resource owner |
| Communication | TLS (MUST) |
| Used Grants | Authorization code, implicit |
| HTTP Method | GET (SHOULD), POST (MAY) |

**Request Parameters**:

| Parameter | Requirement | Description |
|-----------|:----:|------|
| `response_type` | REQUIRED | `code` (authorization code) or `token` (implicit) |
| `client_id` | REQUIRED | Client identifier |
| `redirect_uri` | OPTIONAL | Redirect URI |
| `scope` | OPTIONAL | Access scope |
| `state` | RECOMMENDED | Random value for CSRF protection |

### Token Endpoint

| Item | Description |
|-----|------|
| Role | Exchanges authorization grant for access token |
| Communication | TLS (MUST) |
| HTTP Method | POST (MUST) |
| Client Authentication | MUST (for confidential clients) |

**Request Parameters (Authorization Code Grant)**:

| Parameter | Requirement | Description |
|-----------|:----:|------|
| `grant_type` | REQUIRED | `authorization_code` |
| `code` | REQUIRED | Authorization code |
| `redirect_uri` | REQUIRED* | Same URI as in authorization request (*if specified during authorization) |
| `client_id` | REQUIRED* | Client identifier (*if not authenticated) |

**Response Example**:
```json
{
  "access_token": "2YotnFZFEjr1zCsicMWpAA",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "tGzv3JOkF0XG5Qx2TlKWIA",
  "scope": "read write"
}
```

---

## Scope

| Item | Description |
|-----|------|
| Format | Space-separated list of strings |
| Role | Limits the range of access |
| Definition | Defined by the authorization server (not specified by the specification) |
| Principle | Clients should request the minimum necessary scope |

**Example**:
```
scope=read write profile email
```

The authorization server may ignore the requested scope entirely, partially allow it, or grant additional scopes.

---

## Bearer Token (RFC 6750)

### What is a Bearer Token

A security token that any party in possession of the token can use without proving possession of a cryptographic key.

### Token Transmission Methods

| Method | Requirement | Example |
|-----|:------:|-----|
| **Authorization Header** | MUST (server) / SHOULD (client) | `Authorization: Bearer mF_9.B5f-4.1JqM` |
| Form-Encoded Body | MAY | `access_token=mF_9.B5f-4.1JqM` (POST body) |
| URI Query Parameter | SHOULD NOT | `?access_token=mF_9.B5f-4.1JqM` |

The resource server must support the Authorization Header method (MUST).

### Security Requirements

| Requirement | Description | Strength |
|-----|------|:----:|
| **TLS Required** | Always send via HTTPS | MUST |
| **Certificate Verification** | Verify TLS certificate chain | MUST |
| **No Cookie Storage** | Must not be stored in cookies that can be sent in plaintext | MUST NOT |
| **Short Expiry** | Recommended to be less than 1 hour (RFC 6750 Section 5.3) | SHOULD |
| **Avoid URL Transmission** | Tokens should not be included in page URLs | SHOULD NOT |

---

## Security Considerations (from RFC 6819)

### Major Threats and Countermeasures

| Threat | Countermeasure |
|-----|------|
| **CSRF Attacks** | Verification using `state` parameter |
| **Authorization Code Interception** | Use PKCE (Proof Key for Code Exchange) |
| **Token Leakage** | TLS required, short expiry, scope limitation |
| **Phishing** | Use only legitimate authorization servers, certificate verification |
| **Client Impersonation** | Client authentication, redirect URI verification |

### PKCE (RFC 7636)

An extension to safely use authorization code grant in public clients (SPAs and mobile apps) instead of implicit grant.

```mermaid
sequenceDiagram
    participant C as Client
    participant AS as Authorization Server

    Note over C: 1. Generate code_verifier<br/>(random string)
    Note over C: 2. Calculate code_challenge<br/>= SHA256(code_verifier)

    C->>AS: 3. Authorization Request<br/>(+ code_challenge)
    AS-->>C: Authorization Code

    C->>AS: 4. Token Request<br/>(code + code_verifier)
    Note over AS: 5. Verify:<br/>SHA256(code_verifier) == code_challenge
    AS-->>C: Access Token
```

---

## Modern Recommendations (OAuth 2.1)

OAuth 2.1 is a draft specification that integrates best practices from OAuth 2.0. Key changes:

| Change | OAuth 2.0 | OAuth 2.1 |
|-------|-----------|-----------|
| Implicit Grant | MAY | Removed |
| Password Grant | MAY | Removed |
| PKCE | OPTIONAL | REQUIRED |
| Strict Redirect URI Matching | SHOULD | MUST |
| Bearer Token (Query Parameter) | SHOULD NOT | MUST NOT |
| Refresh Token Rotation | - | SHOULD |

### Recommended Practices

| Item | Recommendation |
|-----|------|
| Grant Type | Authorization Code + PKCE |
| Implicit | Do not use |
| Password | Do not use |
| Token Format | JWT (signed) |
| Token Expiry | Short for access tokens (SHOULD: less than 1 hour), refresh with refresh tokens |
| Scope | Principle of least privilege |

---

## References

- [RFC 6749: The OAuth 2.0 Authorization Framework](https://openid-foundation-japan.github.io/rfc6749.ja.html)
- [RFC 6750: Bearer Token Usage](https://openid-foundation-japan.github.io/rfc6750.ja.html)
- [RFC 6819: OAuth 2.0 Threat Model and Security Considerations](https://openid-foundation-japan.github.io/rfc6819.ja.html)
- [RFC 7636: Proof Key for Code Exchange (PKCE)](https://datatracker.ietf.org/doc/html/rfc7636)
