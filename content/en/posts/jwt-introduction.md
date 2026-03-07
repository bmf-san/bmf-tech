---
title: What is JWT
slug: jwt-introduction
date: 2020-09-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JWT
translation_key: jwt-introduction
---

# Overview
This post summarizes what I have researched about JWT. I will not touch on actual use cases such as OAuth or Open ID Connect.

# What is JWT
JWT (JSON Web Token) is a format that represents a URL-safe Claim using a JSON data structure (the JSON object using JWT is called a Claim Set). JWT uses JWS (JSON Web Signature) with a digital signature or message authentication code (MAC), or JWE (JSON Web Encryption) with encryption.

The respective RFCs for JWT, JWS, and JWE are as follows.
[ietf.org - rfc7519 JSON Web Token](https://tools.ietf.org/html/rfc7519)
[ietf.org - rfc7516 JSON Web Encryption](https://tools.ietf.org/html/rfc7516)
[ietf.org - rfc7515 JSON Web Signature](https://tools.ietf.org/html/rfc7515)

Other related RFCs include:

[ietf.org - rfc7517 JSON Web Key](https://tools.ietf.org/html/rfc7517)

Specifications regarding encryption algorithms and identifiers specified in the JWS, JWE, and JWK specifications.
[ietf.org - rfc7518 JSON Web Algorithm](https://tex2e.github.io/rfc-translater/html/rfc7518.html)

The above specifications are sometimes collectively referred to as JWx.

# JWT Data Structure
You can experience encoding and decoding JWT in the UI at [jwt.io](https://jwt.io/).

An example of a JWT is as follows. The three sections separated by periods serve the roles of `header.payload.signature`.

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

The decoded header, payload, and signature are as follows:
```
{
  "alg": "HS256",
  "typ": "JWT",
  "alg": "HS256"
}
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  your-256-bit-secret
)
```

## Header
The header contains data (a string encoded in Base64 JSON) for verifying the signature.

## Payload
The payload contains Claims (a string encoded in Base64 JSON). There are three types of Claims:

### Registered Claim Names
Claims that are registered at [www.iana.org - jwt](https://www.iana.org/assignments/jwt/jwt.xhtml). They are not mandatory but recommended. For the types of Claims, refer to [ietf.org - rfc7519 JSON Web Token](https://tools.ietf.org/html/rfc7519).

### Public Claim Names
Claims that can be freely defined by users using JWT, but to avoid conflicts, they need to be registered at [www.iana.org - jwt](https://www.iana.org/assignments/jwt/jwt.xhtml) or handled separately.

### Private Claim Names
Claims that can be freely defined between parties using JWT, limited to those not reserved by Registered Claims or Public Claims.

## Signature
Contains data for verifying the integrity of the token.

# JWT and Security
Points to note about handling JWT are well summarized in the following article, which I recommend reading.
[auth0.com - Understanding the Latest Draft for JWT BCP](https://auth0.com/blog/jp-a-look-at-the-latest-draft-for-jwt-bcp/)

# Using JWT with Golang
I will implement JWT using JWS in Golang.

The code is also available at [github.com - bmf-san/go-snippets](https://github.com/bmf-san/go-snippets/blob/de956c3a332ba3f39525a400a8856daa5d888284/architecture_design/auth/jwt.go).

I am using the package [github.com - dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) to handle JWT in Golang.

```go
// Refered to https://github.com/EricLau1/go-api-login
package main

import (
  "encoding/json"
  "fmt"
  "html"
  "log"
  "net/http"
  "strings"
  "time"

  "github.com/dgrijalva/jwt-go"

  "golang.org/x/crypto/bcrypt"
)

const (
  // Actually, these values come from a form or something for getting user information.
  // These values require validation.
  userName  = "bmf"
  userEmail = "foobar@example.com"
  userPass  = "password"
)

var secretKey = []byte("thisisexampleforauthjwt")

// ex.
// curl -X POST -H 'Content-Type:application/json' http://localhost:9999/login
func login(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    // Actually, you need to get values from a form or something. ex. name, email, password.
    token, err := signIn(userEmail, userPass)
    if err != nil {
      toJSON(w, err.Error(), http.StatusUnauthorized)
      return
    }

    toJSON(w, token, http.StatusOK)
    return
  }

  toJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
  return
}

func toJSON(w http.ResponseWriter, data interface{}, statusCode int) {
  w.Header().Set("Content-type", "application/json; charset=UTF8")
  w.WriteHeader(statusCode)
  err := json.NewEncoder(w).Encode(data)
  if err != nil {
    log.Fatal(err)
  }
}

func signIn(userEmail string, userPass string) (string, error) {
  // Actually, these values are stored in something, so you need to get it from something storage by using something key.
  // ex. user := model.GetByEmail(email) → user.password
  // Here, hash a userPass for password verification (bcrypt.VerifyPassword).
  hashedUserPass, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }

  err = bcrypt.CompareHashAndPassword([]byte(hashedUserPass), []byte(userPass))
  if err != nil {
    return "", err
  }

  // If password verification is ok, creates and returns a jwt.
  jwt, err := generateJWT()
  if err != nil {
    return "", err
  }

  return jwt, nil
}

func generateJWT() (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["authorized"] = true
  claims["user_email"] = userEmail
  claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
  return token.SignedString(secretKey)
}

func jwtExtract(r *http.Request) (map[string]interface{}, error) {
  headerAuthorization := r.Header.Get("Authorization")
  bearerToken := strings.Split(headerAuthorization, " ")
  tokenString := html.EscapeString(bearerToken[1])
  claims := jwt.MapClaims{}
  _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return secretKey, nil
  })
  if err != nil {
    return nil, err
  }
  return claims, nil
}

// ex.
// curl -H http://localhost:9999/public
func public(w http.ResponseWriter, r *http.Request) {
  toJSON(w, "public page", http.StatusOK)
  return
}

// ex.
// curl -H 'Content-Type:application/json' -H "Authorization:Bearer <JWT>" http://localhost:9999/private
func private(w http.ResponseWriter, r *http.Request) {
  jwtParams, err := jwtExtract(r)
  if err != nil {
    toJSON(w, err.Error(), http.StatusUnauthorized)
    return
  }
  email, ok := jwtParams["user_email"].(string)
  if !ok {
    toJSON(w, "payload invalid", http.StatusUnauthorized)
    return
  }
  toJSON(w, email, http.StatusOK)
  return
}

func middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    authorizationHeader := r.Header.Get("Authorization")
    if authorizationHeader != "" {
      bearerToken := strings.Split(authorizationHeader, " ")
      if len(bearerToken) == 2 {
        token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
          if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unauthorized")
          }
          return secretKey, nil
        })
        if err != nil {
          w.WriteHeader(http.StatusUnauthorized)
          w.Write([]byte(err.Error()))
          return
        }
        if token.Valid {
          next.ServeHTTP(w, r)
        }
      } else {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Unauthorized"))
      }
    }
  }
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/login", login)
  mux.HandleFunc("/public", public)
  mux.HandleFunc("/private", middlewareAuth(private))

  if err := http.ListenAndServe(":9999", mux); err != nil {
    fmt.Println(err)
  }
}
```

# References
- [jwt.io](https://jwt.io/)
- [assets.ctfassets.net - jwt-handbook-jp.pdf](https://assets.ctfassets.net/2ntc334xpx65/5HColfm15cUhMmDQnupNzd/30d5913d94e79462043f6d8e3f557351/jwt-handbook-jp.pdf)
- [techblog.yahoo.co.jp - Introduction to JSON Web Token (JWT) and its use at Yahoo! JAPAN](https://techblog.yahoo.co.jp/advent-calendar-2017/jwt/)
- [hiyosi.tumblr.com - A Simple Summary of JWT](https://hiyosi.tumblr.com/post/70073770678/jwt%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E7%B0%A1%E5%8D%98%E3%81%AB%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%81%9F)
- [www.wakuwakubank.com - The Mechanism and Precautions of JWT (JSON Web Token)](https://www.wakuwakubank.com/posts/523-it-jwt/)