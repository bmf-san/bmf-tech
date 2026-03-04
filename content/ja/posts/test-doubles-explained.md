---
title: "テストダブルとは何か"
slug: "test-doubles-explained"
date: 2025-10-18
author: bmf-san
categories:
  - "テスト"
tags:
  - "Golang"
  - "テストダブル"
draft: false
---

## はじめに

単体テストを書く際、テスト対象のコードが外部のデータベース、APIサーバー、ファイルシステムなどに依存していると、以下のような問題が発生する：

- テストの実行が遅い
- テスト環境の準備が複雑
- テスト結果が不安定（ネットワークエラーなど）
- 特定の状態やエラーケースの再現が困難

これらの問題を解決するために使われるのが**テストダブル（Test Double）**である。

テストダブルとは、テストにおいて依存先のコンポーネントを本物の代わりに置き換える「代役」のことである。映画のスタントダブルのように、本物の代わりにテスト専用の実装を使用する。

本記事では、テストダブルの5つの種類（Dummy、Stub、Fake、Spy、Mock）について、それぞれの目的と使い分けを、Goのコード例を交えて解説する。

## テストダブルの基礎知識

### テストダブルの5つの種類

テストダブルには5つの種類がある。それぞれ目的と使い方が異なる。

| 種類 | 目的 | 特徴 |
|------|------|------|
| **Dummy** | 引数を埋めるだけ | 実際には使用されない |
| **Stub** | 決まった値を返す | 状態検証に使用 |
| **Fake** | 簡易的な実装 | 実際に動作する軽量版 |
| **Spy** | 呼び出しを記録 | 履歴を後で検証 |
| **Mock** | 期待を事前設定 | 振る舞い検証に使用 |

### 前提：テスト対象のコード

以下の例では、データストアに依存するサービスをテストする。

```go
package main

import "errors"

// Store はデータストレージのインターフェース
type Store interface {
    Get(key string) (string, error)
    Put(key string, value string) error
}

// UserService はStoreに依存するサービス
type UserService struct {
    store Store
}

func NewUserService(s Store) *UserService {
    return &UserService{store: s}
}

// FetchValue は内部でstore.Get()を呼ぶ
func (svc *UserService) FetchValue(key string) (string, error) {
    v, err := svc.store.Get(key)
    if err != nil {
        return "", err
    }
    if v == "" {
        return "", errors.New("value not found")
    }
    return v, nil
}

// SaveValue は内部でstore.Put()を呼ぶ
func (svc *UserService) SaveValue(key, value string) error {
    if value == "" {
        return errors.New("value cannot be empty")
    }
    return svc.store.Put(key, value)
}
```

### テストダブルの種類と実装例

それぞれのテストダブルについて、具体的なコード例と使いどころを見ていく。

### 1. Dummy

**Dummy**は、引数を埋めるためだけに存在し、実際には使用されないオブジェクトである。

#### 使用例

```go
package main

import "testing"

// Dummy実装
type DummyStore struct {
    t *testing.T
}

func NewDummyStore(t *testing.T) *DummyStore {
    return &DummyStore{t: t}
}

func (d *DummyStore) Get(key string) (string, error) {
    d.t.Fatal("Get should not be called")
    return "", nil
}

func (d *DummyStore) Put(key, value string) error {
    d.t.Fatal("Put should not be called")
    return nil
}

// Logger インターフェース
type Logger interface {
    Info(msg string)
}

// SimpleLogger は簡易的なLogger実装
type SimpleLogger struct{}

func (l *SimpleLogger) Info(msg string) {
    // 実際にはログ出力するが、ここでは何もしない
}

// ProcessData は複数の依存を持つ関数（storeは使わない）
func ProcessData(store Store, logger Logger) error {
    // この関数ではloggerだけを使い、storeは使わない
    logger.Info("processing started")
    return nil
}

func TestProcessData(t *testing.T) {
    // storeは使われないので、Dummyで十分
    dummy := NewDummyStore(t)
    logger := &SimpleLogger{}

    err := ProcessData(dummy, logger)

    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    // もしstoreのメソッドが呼ばれたら、t.Fatal()でテストが失敗する
}
```

#### 使いどころ

- 関数のシグネチャを満たすために引数が必要だが、実際には使われない場合
- 呼ばれたら即座に失敗させることで、誤用を検出できる

### 2. Stub

**Stub**は、呼び出しに対して決まった値を返すだけの単純な実装である。状態検証に使われる。

#### 使用例

```go
package main

import (
    "errors"
    "testing"
)

// Stub実装
type StubStore struct {
    value string
    err   error
}

func (s *StubStore) Get(key string) (string, error) {
    return s.value, s.err
}

func (s *StubStore) Put(key, value string) error {
    return nil
}

// 正常系のテスト
func TestFetchValue_Success(t *testing.T) {
    stub := &StubStore{value: "hello"}
    svc := NewUserService(stub)

    got, err := svc.FetchValue("foo")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got != "hello" {
        t.Errorf("got %q, want %q", got, "hello")
    }
}

// エラー系のテスト
func TestFetchValue_Error(t *testing.T) {
    stub := &StubStore{err: errors.New("connection failed")}
    svc := NewUserService(stub)

    _, err := svc.FetchValue("foo")
    if err == nil {
        t.Error("expected error, got nil")
    }
}

// 空文字列のテスト
func TestFetchValue_EmptyValue(t *testing.T) {
    stub := &StubStore{value: ""}
    svc := NewUserService(stub)

    _, err := svc.FetchValue("foo")
    if err == nil {
        t.Error("expected error for empty value")
    }
}
```

#### 使いどころ

- テストで特定の返り値やエラーを返したい場合
- 状態（結果）を検証するテスト
- 最もシンプルで使いやすいテストダブル

### 3. Fake

**Fake**は、実際に簡易的な動作をする軽量実装である。本物に近い振る舞いをするが、テスト用に簡略化されている。

#### 使用例

```go
package main

import (
    "errors"
    "testing"
)

// Fake実装：メモリ内でデータを管理
type FakeStore struct {
    data map[string]string
}

func NewFakeStore() *FakeStore {
    return &FakeStore{data: make(map[string]string)}
}

func (f *FakeStore) Get(key string) (string, error) {
    value, exists := f.data[key]
    if !exists {
        return "", errors.New("key not found")
    }
    return value, nil
}

func (f *FakeStore) Put(key, value string) error {
    f.data[key] = value
    return nil
}

// Fakeを使ったテスト
func TestFetchValue_Fake(t *testing.T) {
    fake := NewFakeStore()
    fake.Put("foo", "bar")
    fake.Put("hello", "world")

    svc := NewUserService(fake)

    // 存在するキーの取得
    got, err := svc.FetchValue("foo")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got != "bar" {
        t.Errorf("got %q, want %q", got, "bar")
    }

    // 別のキーの取得
    got2, err := svc.FetchValue("hello")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got2 != "world" {
        t.Errorf("got %q, want %q", got2, "world")
    }
}
```

#### 使いどころ

- 複数のテストケースで共通のデータストアが必要な場合
- 実際の動作に近いテストが必要な場合
- 統合テストとユニットテストの中間レベルのテスト
- 例：メモリ内データベース、インメモリファイルシステム

### 4. Spy

**Spy**は、呼び出し履歴（引数、回数など）を記録し、後で検証することを目的とする。Mockとの違いは、Spyは事前に期待値を設定せず、実行後に履歴を確認する点である。

#### 使用例

```go
package main

import "testing"

// Spy実装
type SpyStore struct {
    GetCalls []string // Getが呼ばれたキーのリスト
    PutCalls []struct {
        Key   string
        Value string
    }
    value string
}

func (s *SpyStore) Get(key string) (string, error) {
    s.GetCalls = append(s.GetCalls, key)
    return s.value, nil
}

func (s *SpyStore) Put(key, value string) error {
    s.PutCalls = append(s.PutCalls, struct {
        Key   string
        Value string
    }{key, value})
    return nil
}

// Spyを使ったテスト
func TestFetchValue_Spy(t *testing.T) {
    spy := &SpyStore{value: "hello"}
    svc := NewUserService(spy)

    _, _ = svc.FetchValue("foo")

    // 呼び出し履歴を検証
    if len(spy.GetCalls) != 1 {
        t.Errorf("expected 1 call, got %d", len(spy.GetCalls))
    }
    if spy.GetCalls[0] != "foo" {
        t.Errorf("expected Get('foo'), got Get('%s')", spy.GetCalls[0])
    }
}

// 複数回呼び出しのテスト
func TestFetchMultipleValues_Spy(t *testing.T) {
    spy := &SpyStore{value: "test"}
    svc := NewUserService(spy)

    svc.FetchValue("key1")
    svc.FetchValue("key2")
    svc.FetchValue("key3")

    // 呼び出し順序と引数を検証
    expected := []string{"key1", "key2", "key3"}
    if len(spy.GetCalls) != len(expected) {
        t.Fatalf("expected %d calls, got %d", len(expected), len(spy.GetCalls))
    }
    for i, want := range expected {
        if spy.GetCalls[i] != want {
            t.Errorf("call %d: expected %q, got %q", i, want, spy.GetCalls[i])
        }
    }
}
```

#### 使いどころ

- メソッドが正しい引数で呼ばれたか確認したい場合
- 呼び出し回数や順序を検証したい場合
- ログ記録、通知送信などの副作用を持つ処理のテスト

### 5. Mock

**Mock**は、事前に期待（expectation）を設定し、テスト終了後にその期待が満たされたかを検証する。振る舞い検証に特化している。Spyとの違いは、Mockはテスト実行前に「こう呼ばれるべき」という期待を明示する点である。

#### 使用例

```go
package main

import (
    "errors"
    "testing"
)

// Mock実装
type MockStore struct {
    expectations []struct {
        key   string
        value string
        err   error
    }
    callIndex int
    t         *testing.T
}

func NewMockStore(t *testing.T) *MockStore {
    return &MockStore{t: t}
}

// 期待値を設定（チェーン可能）
func (m *MockStore) ExpectGet(key string) *MockStore {
    m.expectations = append(m.expectations, struct {
        key   string
        value string
        err   error
    }{key: key})
    return m
}

func (m *MockStore) WillReturn(value string, err error) *MockStore {
    if len(m.expectations) > 0 {
        idx := len(m.expectations) - 1
        m.expectations[idx].value = value
        m.expectations[idx].err = err
    }
    return m
}

func (m *MockStore) Get(key string) (string, error) {
    if m.callIndex >= len(m.expectations) {
        m.t.Errorf("unexpected call to Get(%q)", key)
        return "", errors.New("unexpected call")
    }

    expected := m.expectations[m.callIndex]
    if key != expected.key {
        m.t.Errorf("call %d: expected Get(%q), got Get(%q)",
            m.callIndex, expected.key, key)
    }

    m.callIndex++
    return expected.value, expected.err
}

func (m *MockStore) Put(key, value string) error {
    return nil
}

// 期待が満たされたか検証
func (m *MockStore) Verify() {
    if m.callIndex != len(m.expectations) {
        m.t.Errorf("expected %d calls, got %d", len(m.expectations), m.callIndex)
    }
}

// Mockを使ったテスト
func TestFetchValue_Mock(t *testing.T) {
    mock := NewMockStore(t)
    mock.ExpectGet("foo").WillReturn("bar", nil)

    svc := NewUserService(mock)

    result, err := svc.FetchValue("foo")

    // 期待通りに呼ばれたか検証
    mock.Verify()

    // 結果も検証
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if result != "bar" {
        t.Errorf("got %q, want %q", result, "bar")
    }
}

// 複数回の呼び出しをテスト
func TestFetchMultipleValues_Mock(t *testing.T) {
    mock := NewMockStore(t)
    mock.ExpectGet("key1").WillReturn("value1", nil)
    mock.ExpectGet("key2").WillReturn("value2", nil)

    svc := NewUserService(mock)

    result1, _ := svc.FetchValue("key1")
    result2, _ := svc.FetchValue("key2")

    mock.Verify()

    if result1 != "value1" {
        t.Errorf("got %q, want %q", result1, "value1")
    }
    if result2 != "value2" {
        t.Errorf("got %q, want %q", result2, "value2")
    }
}
```

#### 使いどころ

- メソッドが期待通りの順序・引数で呼ばれたか厳密に検証したい場合
- 複雑な振る舞いの検証が必要な場合
- 外部サービスとのインタラクションをテストする場合

## まとめ

テストダブルは、単体テストを高速で安定させ、テストしづらいコードをテスト可能にする強力なツールである。

1. まず依存関係をインターフェースとして抽象化する
2. 状態検証にはStub/Fakeを使う
3. 振る舞い検証にはSpy/Mockを使う
4. Mockは必要最小限に抑える
5. テストダブルはシンプルに保つ

適切なテストダブルを選択することで、保守性が高く、リファクタリング耐性のあるテストを書くことができる。

## 参考

- [xUnit Test Patterns – Test Double](http://xunitpatterns.com/Test%20Double.html)
- [Mocks Aren't Stubs – Martin Fowler](https://martinfowler.com/articles/mocksArentStubs.html)
