---
title: What is a Test Double?
slug: test-doubles-explained
date: 2025-10-18T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Golang
  - Test Double
translation_key: test-doubles-explained
---

## Introduction

When writing unit tests, if the code under test depends on external databases, API servers, file systems, etc., the following problems may arise:

- Tests run slowly
- Setting up the test environment is complex
- Test results are unstable (e.g., network errors)
- Reproducing specific states or error cases is difficult

To solve these problems, we use **Test Doubles**.

A test double is a "stand-in" that replaces a dependent component with a test-specific implementation instead of the real one, similar to a stunt double in movies.

In this article, we will explain the five types of test doubles (Dummy, Stub, Fake, Spy, Mock), their purposes, and how to differentiate them, along with Go code examples.

## Basics of Test Doubles

### Five Types of Test Doubles

There are five types of test doubles, each with different purposes and usages.

| Type   | Purpose                     | Features                          |
|--------|-----------------------------|-----------------------------------|
| **Dummy** | Fills in arguments only      | Not actually used                 |
| **Stub**  | Returns fixed values         | Used for state verification       |
| **Fake**  | Simplified implementation     | Lightweight version that works    |
| **Spy**   | Records calls                | Verifies history later            |
| **Mock**  | Predefined expectations       | Used for behavior verification     |

### Prerequisite: Code Under Test

In the following example, we will test a service that depends on a data store.

```go
package main

import "errors"

// Store is the interface for data storage

type Store interface {
    Get(key string) (string, error)
    Put(key string, value string) error
}

// UserService is a service that depends on Store

type UserService struct {
    store Store
}

func NewUserService(s Store) *UserService {
    return &UserService{store: s}
}

// FetchValue calls store.Get() internally
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

// SaveValue calls store.Put() internally
func (svc *UserService) SaveValue(key, value string) error {
    if value == "" {
        return errors.New("value cannot be empty")
    }
    return svc.store.Put(key, value)
}
```

### Types of Test Doubles and Implementation Examples

Let's look at specific code examples and use cases for each type of test double.

### 1. Dummy

**Dummy** is an object that exists only to fill in arguments and is not actually used.

#### Usage Example

```go
package main

import "testing"

// Dummy implementation

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

// Logger interface

type Logger interface {
    Info(msg string)
}

// SimpleLogger is a simple Logger implementation

type SimpleLogger struct{}

func (l *SimpleLogger) Info(msg string) {
    // Actually logs, but does nothing here
}

// ProcessData is a function with multiple dependencies (store is not used)
func ProcessData(store Store, logger Logger) error {
    // This function only uses logger, not store
    logger.Info("processing started")
    return nil
}

func TestProcessData(t *testing.T) {
    // store is not used, so Dummy is sufficient
    dummy := NewDummyStore(t)
    logger := &SimpleLogger{}

    err := ProcessData(dummy, logger)

    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    // If store's methods are called, the test will fail with t.Fatal()
}
```

#### Use Cases

- When an argument is needed to satisfy a function signature but is not actually used.
- To detect misuse by failing immediately if called.

### 2. Stub

**Stub** is a simple implementation that returns fixed values in response to calls. It is used for state verification.

#### Usage Example

```go
package main

import (
    "errors"
    "testing"
)

// Stub implementation

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

// Successful test case
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

// Error case test
func TestFetchValue_Error(t *testing.T) {
    stub := &StubStore{err: errors.New("connection failed")}
    svc := NewUserService(stub)

    _, err := svc.FetchValue("foo")
    if err == nil {
        t.Error("expected error, got nil")
    }
}

// Empty string test
func TestFetchValue_EmptyValue(t *testing.T) {
    stub := &StubStore{value: ""}
    svc := NewUserService(stub)

    _, err := svc.FetchValue("foo")
    if err == nil {
        t.Error("expected error for empty value")
    }
}
```

#### Use Cases

- When you want to return specific return values or errors in tests.
- For tests that verify state (results).
- The simplest and easiest to use test double.

### 3. Fake

**Fake** is a lightweight implementation that actually performs simplified operations. It behaves similarly to the real thing but is simplified for testing.

#### Usage Example

```go
package main

import (
    "errors"
    "testing"
)

// Fake implementation: manages data in memory

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

// Test using Fake
func TestFetchValue_Fake(t *testing.T) {
    fake := NewFakeStore()
    fake.Put("foo", "bar")
    fake.Put("hello", "world")

    svc := NewUserService(fake)

    // Fetching existing key
    got, err := svc.FetchValue("foo")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got != "bar" {
        t.Errorf("got %q, want %q", got, "bar")
    }

    // Fetching another key
    got2, err := svc.FetchValue("hello")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got2 != "world" {
        t.Errorf("got %q, want %q", got2, "world")
    }
}
```

#### Use Cases

- When multiple test cases require a shared data store.
- When tests need behavior close to actual operations.
- Intermediate level tests between integration and unit tests.
- Example: in-memory databases, in-memory file systems.

### 4. Spy

**Spy** is intended to record call history (arguments, counts, etc.) and verify it later. Unlike Mock, Spy does not set expectations in advance but checks the history after execution.

#### Usage Example

```go
package main

import "testing"

// Spy implementation

type SpyStore struct {
    GetCalls []string // List of keys called by Get
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

// Test using Spy
func TestFetchValue_Spy(t *testing.T) {
    spy := &SpyStore{value: "hello"}
    svc := NewUserService(spy)

    _, _ = svc.FetchValue("foo")

    // Verify call history
    if len(spy.GetCalls) != 1 {
        t.Errorf("expected 1 call, got %d", len(spy.GetCalls))
    }
    if spy.GetCalls[0] != "foo" {
        t.Errorf("expected Get('foo'), got Get('%s')", spy.GetCalls[0])
    }
}

// Test multiple calls
func TestFetchMultipleValues_Spy(t *testing.T) {
    spy := &SpyStore{value: "test"}
    svc := NewUserService(spy)

    svc.FetchValue("key1")
    svc.FetchValue("key2")
    svc.FetchValue("key3")

    // Verify call order and arguments
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

#### Use Cases

- When you want to confirm that a method was called with the correct arguments.
- When you want to verify call counts or order.
- Testing processes with side effects like logging or notifications.

### 5. Mock

**Mock** is specialized for behavior verification by setting expectations in advance and verifying if those expectations were met after the test. Unlike Spy, Mock explicitly states what should be called before execution.

#### Usage Example

```go
package main

import (
    "errors"
    "testing"
)

// Mock implementation

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

// Set expectations (chainable)
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

// Verify expectations
func (m *MockStore) Verify() {
    if m.callIndex != len(m.expectations) {
        m.t.Errorf("expected %d calls, got %d", len(m.expectations), m.callIndex)
    }
}

// Test using Mock
func TestFetchValue_Mock(t *testing.T) {
    mock := NewMockStore(t)
    mock.ExpectGet("foo").WillReturn("bar", nil)

    svc := NewUserService(mock)

    result, err := svc.FetchValue("foo")

    // Verify expectations
    mock.Verify()

    // Verify result
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if result != "bar" {
        t.Errorf("got %q, want %q", result, "bar")
    }
}

// Test multiple calls
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

#### Use Cases

- When you want to strictly verify that methods were called in the expected order and with the expected arguments.
- When complex behavior verification is needed.
- When testing interactions with external services.

## Conclusion

Test doubles are powerful tools that make unit tests faster and more stable, enabling testing of hard-to-test code.

1. First, abstract dependencies as interfaces.
2. Use Stubs/Fakes for state verification.
3. Use Spies/Mocks for behavior verification.
4. Keep Mocks to a minimum.
5. Keep test doubles simple.

By choosing the appropriate test double, you can write tests that are maintainable and resilient to refactoring.

## References

- [xUnit Test Patterns – Test Double](http://xunitpatterns.com/Test%20Double.html)
- [Mocks Aren't Stubs – Martin Fowler](https://martinfowler.com/articles/mocksArentStubs.html)