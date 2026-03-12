---
title: "JWTとは？安全なAPI認証のためのJSON Web Token解説"
slug: jwt-introduction
date: 2020-09-11T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - JWT
translation_key: jwt-introduction
---


# 概要
JWTについて調べたことをまとめておく。
OAuthやOpen ID Connectなど実際の利用事例については触れない。

# JWTとは
JWT（JSON Web Token）は、JSONデータ構造を用いたURLセーフなClaim（JWTを用いたJSONオブジェクトはClaim Setと呼ばれる）を表現するフォーマット。
JWTでは、デジタル署名またはメッセージ認証コード（MAC）を用いたJWS（JSON Web Signature）、あるいは暗号化を用いたJWE（JSON Web Encryption）が利用される。

JWT、JWS、JWEのそれぞれのRFCは下記の通り。
[ietf.org - rfc7519 JSON Web Token](https://tools.ietf.org/html/rfc7519)
[ietf.org - rfc7516 JSON Web Encryption](https://tools.ietf.org/html/rfc7516)
[ietf.orf - rfc7515 JSON Web Signature](https://tools.ietf.org/html/rfc7515)

その他の関連するRFCとしては、以下のようなものもある。

[ietf.org - rfc7517 JSON Web Key](https://tools.ietf.org/html/rfc7517)

JWSやJWE,JWKの仕様で仕様される暗号化アルゴリズムと識別子に関する仕様。
[ietf.org - rfc7518 JSON Web Algorithm](https://tex2e.github.io/rfc-translater/html/rfc7518.html)

上記の仕様をまとめてJWxと呼ばれることがあるらしい。

# JWTのデータ構造
[jwt.io](https://jwt.io/)でJWTのエンコードとデコードをUIで体験することができる。

JWTの例が下記。
ピリオドで区切られた3つのセクションはそれぞれ、`ヘッダー.ペイロード.署名`の役割を担っている。

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

ヘッダー、ペイロード、署名の順にデコードしたものが下記。
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

## ヘッダー
ヘッダーには、署名の検証を行うためのデータ（JSONをBase64エンコードした文字列）が含まれる。

## ペイロード
ペイロードはClaim（JSONをBase64エンコードした文字列）を含む。
Claimには以下の3種類がある。

### Registered Claim Names
[www.iana.org - jwt](https://www.iana.org/assignments/jwt/jwt.xhtml)に登録済みのClaim。
必須ではなく、推奨。
Claimの種類は[ietf.org - rfc7519 JSON Web Token](https://tools.ietf.org/html/rfc7519)を参照。

### Public Claim Names
JWTを使用するユーザーが自由に定義することができるClaimだが、衝突防止のため、[www.iana.org - jwt](https://www.iana.org/assignments/jwt/jwt.xhtml)に登録するか、別途対応をする必要がある。

### Private Claim Names
JWTを使用する当事者間で自由に定義することができる。Registerd ClaimやPublic Claimで予約されているもの以外に限る。

## 署名
トークンの改ざん検証のためのデータを含む。

# JWTとセキュリティ
JWTの扱いについての注意点は以下の記事がよくまとまっているので一読しておきたい。
[auth0.com - JWT の最新ベスト プラクティスに関するドラフトを読み解く](https://auth0.com/blog/jp-a-look-at-the-latest-draft-for-jwt-bcp/)

# GolangでJWTを使ってみる
JWSを用いたJWTの実装をGolangでやってみる。

コードは[github.com - bmf-san/go-snippets](https://github.com/bmf-san/go-snippets/blob/de956c3a332ba3f39525a400a8856daa5d888284/architecture_design/auth/jwt.go)にも置いてある。

GolangでJWTを扱うために[github.com - dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)というパッケージを利用している。

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
  // Actualy, thease values comes from a form or something for getting user infomations.
  //These values are require validation.
  userName  = "bmf"
  userEmail = "foobar@example.com"
  userPass  = "password"
)

var secretKey = []byte("thisisexampleforauthjwt")

// ex.
// curl -X POST -H 'Content-Type:application/json' http://localhost:9999/login
func login(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    // Actualy, you need to get values from a form or somthing. ex. name, email, password.
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
  // Actualy, this values stored in a something storage so you need to get it from a something storage by using a something key.
  // ex. user := model.GetByEmail(email) → user.password
  // Here, hash a userPass for password verification(bcrypt.VerifyPassword).
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

# 参考
- [jwt.io](https://jwt.io/)
- [assets.ctfassets.net - jwt-handbook-jp.pdf](https://assets.ctfassets.net/2ntc334xpx65/5HColfm15cUhMmDQnupNzd/30d5913d94e79462043f6d8e3f557351/jwt-handbook-jp.pdf)
- [techblog.yahoo.co.jp - JSON Web Token（JWT）の紹介とYahoo! JAPANにおけるJWTの活用](https://techblog.yahoo.co.jp/advent-calendar-2017/jwt/)
- [hiyosi.tumblr.com - JWTについて簡単にまとめてみた](https://hiyosi.tumblr.com/post/70073770678/jwt%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E7%B0%A1%E5%8D%98%E3%81%AB%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%81%9F)
- [www.wakuwakubank.com - JWT(JSON Web Token)の「仕組み」と「注意点」](https://www.wakuwakubank.com/posts/523-it-jwt/)

