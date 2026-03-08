---
title: Why the Singleton Pattern is an Anti-Pattern
slug: singleton-pattern-anti-pattern
date: 2025-10-18T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Singleton Pattern
  - Golang
translation_key: singleton-pattern-anti-pattern
---

## Introduction

The Singleton pattern is one of the most well-known and widely used design patterns. However, from the perspective of clean code and maintainability, it has many issues.

In this article, we will discuss the main problems of the Singleton pattern, using specific Go code examples.

## What is the Singleton Pattern?

The Singleton pattern is a design pattern that guarantees that a class has only one instance and provides a global point of access to it.

### Basic Implementation in Go

```go
package main

import (
    "sync"
)

// Database is a database connection implemented as a singleton
type Database struct {
    connectionString string
}

var (
    instance *Database
    once     sync.Once
)

// GetInstance returns the singleton instance of Database
func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{
            connectionString: "localhost:5432",
        }
    })
    return instance
}

func (db *Database) Query(sql string) string {
    return "result from " + db.connectionString
}
```

Although this pattern seems convenient at first glance, it actually causes many problems.

## Problems with the Singleton Pattern

### 1. Lack of Bijectivity

Singletons do not correspond directly to concepts in the real world. In reality, most concepts can have multiple instances.

#### Example of the Problem

```go
// In the real world, multiple database connections can exist
func Example() {
    // Primary DB
    primaryDB := GetInstance()

    // I want to connect to a replica DB... but it's impossible with a singleton
    // replicaDB := GetReplicaInstance() // This is not possible

    primaryDB.Query("SELECT * FROM users")
}
```

In the real world, there are often needs to connect to a primary and a replica, or to multiple databases. However, the Singleton pattern takes away this flexibility.

#### Improvement

```go
// Using interfaces and dependency injection
type DBConnection interface {
    Query(sql string) string
}

type PostgresDB struct {
    connectionString string
}

func NewPostgresDB(connStr string) *PostgresDB {
    return &PostgresDB{connectionString: connStr}
}

func (db *PostgresDB) Query(sql string) string {
    return "result from " + db.connectionString
}

// Flexibly manage multiple connections
func ImprovedExample() {
    primaryDB := NewPostgresDB("primary.db.com:5432")
    replicaDB := NewPostgresDB("replica.db.com:5432")

    primaryDB.Query("INSERT INTO users VALUES (...)" )
    replicaDB.Query("SELECT * FROM users")
}
```

### 2. Tight Coupling

Singletons provide a global access point that is difficult to separate, causing the codebase to become tightly coupled.

#### Example of the Problem

```go
type UserService struct {
    // Dependency on the database is hidden
}

func (s *UserService) GetUser(id int) string {
    // Directly depends on the global singleton
    db := GetInstance()
    return db.Query("SELECT * FROM users WHERE id = " + string(rune(id)))
}

func (s *UserService) CreateUser(name string) {
    db := GetInstance()
    db.Query("INSERT INTO users (name) VALUES ('" + name + "')")
}
```

Problems with this code:
- `UserService` implicitly depends on the `Database` singleton.
- The dependency is not explicit, making it hard to understand just by reading the code.
- Difficult to replace with mocks during testing.

#### Improvement

```go
// Use dependency injection to explicitly show dependencies
type UserService struct {
    db DBConnection // Dependency is explicit
}

func NewUserService(db DBConnection) *UserService {
    return &UserService{db: db}
}

func (s *UserService) GetUser(id int) string {
    return s.db.Query("SELECT * FROM users WHERE id = " + string(rune(id)))
}

func (s *UserService) CreateUser(name string) {
    s.db.Query("INSERT INTO users (name) VALUES ('" + name + "')")
}
```

### 3. Difficult to Test

The existence of a singleton makes it very difficult to create unit tests.

#### Example of the Problem

```go
// Code to be tested
func ProcessUser(userID int) string {
    db := GetInstance() // Depends on the singleton
    result := db.Query("SELECT * FROM users WHERE id = " + string(rune(userID)))
    return "Processed: " + result
}

// Test code - cannot be replaced with a mock
func TestProcessUser(t *testing.T) {
    // Problem: The actual database is being used
    // No way to replace with a mock
    result := ProcessUser(1)

    // Since it connects to the actual DB, the test is slow and unstable
    if result == "" {
        t.Error("Expected non-empty result")
    }
}
```

#### Improvement

```go
// Mock implementation
type MockDB struct {
    queryFunc func(sql string) string
}

func (m *MockDB) Query(sql string) string {
    if m.queryFunc != nil {
        return m.queryFunc(sql)
    }
    return "mock result"
}

// Testable implementation
func ProcessUserImproved(userID int, db DBConnection) string {
    result := db.Query("SELECT * FROM users WHERE id = " + string(rune(userID)))
    return "Processed: " + result
}

// Test code - can use mocks
func TestProcessUserImproved(t *testing.T) {
    // Inject mock DB
    mockDB := &MockDB{
        queryFunc: func(sql string) string {
            return "test user data"
        },
    }

    result := ProcessUserImproved(1, mockDB)

    expected := "Processed: test user data"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
```

### 4. State Accumulation

With multiple test executions, unnecessary data can accumulate in the singleton.

#### Example of the Problem

```go
type Cache struct {
    data map[string]string
}

var cacheInstance *Cache
var cacheOnce sync.Once

func GetCache() *Cache {
    cacheOnce.Do(func() {
        cacheInstance = &Cache{
            data: make(map[string]string),
        }
    })
    return cacheInstance
}

func (c *Cache) Set(key, value string) {
    c.data[key] = value
}

func (c *Cache) Get(key string) string {
    return c.data[key]
}

// Test 1
func TestCacheSet(t *testing.T) {
    cache := GetCache()
    cache.Set("key1", "value1")

    if cache.Get("key1") != "value1" {
        t.Error("Expected value1")
    }
}

// Test 2 - State from Test 1 remains
func TestCacheGet(t *testing.T) {
    cache := GetCache()

    // Problem: Data from the previous test remains
    // "key1" already exists
    if cache.Get("key1") != "" {
        t.Error("Expected empty cache, but got data from previous test")
    }
}
```

#### Improvement

```go
// Generate instances each time
type ImprovedCache struct {
    data map[string]string
}

func NewCache() *ImprovedCache {
    return &ImprovedCache{
        data: make(map[string]string),
    }
}

func (c *ImprovedCache) Set(key, value string) {
    c.data[key] = value
}
}

func (c *ImprovedCache) Get(key string) string {
    return c.data[key]
}

// Test 1 - Independent instance
func TestImprovedCacheSet(t *testing.T) {
    cache := NewCache() // New instance
    cache.Set("key1", "value1")

    if cache.Get("key1") != "value1" {
        t.Error("Expected value1")
    }
}

// Test 2 - Independent instance
func TestImprovedCacheGet(t *testing.T) {
    cache := NewCache() // Another new instance

    // Not affected by previous tests
    if cache.Get("key1") != "" {
        t.Error("Expected empty cache")
    }
}
```

### 5. Concurrency Issues

Using the Singleton pattern requires thread-safe implementations in concurrent environments, increasing complexity.

#### Example of the Problem (Non-thread-safe Singleton)

```go
type Counter struct {
    count int
}

var counterInstance *Counter
var counterOnce sync.Once

// Implemented as a singleton (instance creation is thread-safe)
func GetCounter() *Counter {
    counterOnce.Do(func() {
        counterInstance = &Counter{count: 0}
    })
    return counterInstance
}

// Problem: Since it's a singleton, if this method is not thread-safe,
// race conditions will occur for all callers
func (c *Counter) Increment() {
    c.count++ // Race condition
    // Note: No compile error
    // However, unexpected results occur during concurrent execution
}

func (c *Counter) GetCount() int {
    return c.count // This also has a race condition
}

// Details of the race condition:
// Because it's a singleton, all goroutines access the same instance
// c.count++ breaks down into the following operations:
//   1. Read value from memory (READ)
//   2. Increment value (INCREMENT)
//   3. Write back to memory (WRITE)
//
// Example: If the current count = 5, and two goroutines call Increment() simultaneously
//   goroutine A: Reads count = 5
//   goroutine B: Reads count = 5  ← A reads the same value
//   goroutine A: Calculates 5 + 1 = 6
//   goroutine B: Calculates 5 + 1 = 6
//   goroutine A: Writes count = 6
//   goroutine B: Writes count = 6 ← Overwritten
//   Result: Incremented twice, but count is 6 (expected value is 7)

// Problems occur during concurrent execution
func ConcurrentExample() {
    var wg sync.WaitGroup

    // 1000 goroutines access the same singleton instance
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter := GetCounter() // All access the same instance
            counter.Increment()      // Race condition
        }()
    }

    wg.Wait()

    // Expected: 1000
    // Actual: Less than 1000 (e.g., 987, 934, etc.)
    // Reason: Since it's a singleton, all goroutines share the same instance,
    //       calling the non-thread-safe Increment() function
    fmt.Println("Count:", GetCounter().GetCount())
}

// How to detect race conditions
// Problems may not manifest during normal execution,
// but running go run -race main.go will show warnings:
//
// WARNING: DATA RACE
// Write at 0x... by goroutine 7:
//   main.(*Counter).Increment()
// Previous write at 0x... by goroutine 6:
//   main.(*Counter).Increment()
```

#### Improvement 1: Thread-safe Singleton Implementation

If using a singleton, all methods must be thread-safe:

```go
type Counter struct {
    count int
    mu    sync.Mutex // Lock needed for all methods
}

var (
    counterInstance *Counter
    counterOnce     sync.Once
)

func GetCounter() *Counter {
    counterOnce.Do(func() {
        counterInstance = &Counter{count: 0}
    })
    return counterInstance
}

// All methods need mutex for exclusive control
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

Problems:
- Since it's a singleton, all goroutines compete for the same mutex.
- Becomes a performance bottleneck.
- Increases the risk of deadlocks.
- Other issues with the singleton (testability, tight coupling, etc.) remain.

#### Improvement 2: Implementation without Singleton

```go
// Create multiple instances as needed without using a singleton
type SafeCounter struct {
    count int
    mu    sync.Mutex
}

func NewSafeCounter() *SafeCounter {
    return &SafeCounter{count: 0}
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// Improved example: Use independent counters for each goroutine group
func ImprovedConcurrentExample() {
    // Create 10 independent counters
    counters := make([]*SafeCounter, 10)
    for i := range counters {
        counters[i] = NewSafeCounter()
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(idx int) {
            defer wg.Done()
            // Each goroutine accesses a different counter
            // → Mutex contention is distributed
            counters[idx%10].Increment()
        }(i)
    }

    wg.Wait()

    // Aggregate at the end
    total := 0
    for _, c := range counters {
        total += c.GetCount()
    }
    fmt.Println("Total count:", total) // Definitely 1000
}
```

Benefits:
- By using multiple instances, mutex contention is distributed.
- Performance improves.
- Testing is easier (each counter can be tested independently).
- Freed from the constraints of a singleton.

### 6. Violation of the Single Responsibility Principle

Singleton classes take on the responsibility of "instance management" in addition to their primary responsibility.

#### Example of the Problem

```go
type Logger struct {
    logFile string
}

var loggerInstance *Logger
var loggerOnce sync.Once

func GetLogger() *Logger {
    loggerOnce.Do(func() {
        loggerInstance = &Logger{
            logFile: "/var/log/app.log",
        }
    })
    return loggerInstance
}

// Logger has two responsibilities:
// 1. Writing logs (its primary responsibility)
// 2. Managing its own instance (singleton responsibility)
func (l *Logger) Log(message string) {
    // Log processing
    fmt.Println("Log to", l.logFile, ":", message)
}
```

#### Improvement

```go
// Logger only holds its primary responsibility (logging)
type ImprovedLogger struct {
    logFile string
}

func NewLogger(logFile string) *ImprovedLogger {
    return &ImprovedLogger{logFile: logFile}
}

func (l *ImprovedLogger) Log(message string) {
    fmt.Println("Log to", l.logFile, ":", message)
}

// Instance management is handled elsewhere (main function or DI container)
func main() {
    logger := NewLogger("/var/log/app.log")

    // Can create multiple loggers if needed
    errorLogger := NewLogger("/var/log/error.log")

    logger.Log("Application started")
    errorLogger.Log("An error occurred")
}
```

### 7. Hindrance to Dependency Injection

Singletons hinder the pattern of dependency injection, making it difficult to separate components.

#### Example of the Problem

```go
type EmailService struct {
    // Dependency is hidden
}

func (s *EmailService) SendEmail(to, message string) {
    // Depends on a global singleton
    logger := GetLogger()
    logger.Log("Sending email to " + to)

    // Actual email sending process
    fmt.Println("Email sent to", to)
}

// Usage
func NotifyUser(userEmail string) {
    service := &EmailService{} // Dependency is not visible
    service.SendEmail(userEmail, "Hello!")
}
```

#### Improvement

```go
// Define an interface

type LogWriter interface {
    Log(message string)
}

// EmailService explicitly receives its dependencies

type ImprovedEmailService struct {
    logger LogWriter
}

func NewEmailService(logger LogWriter) *ImprovedEmailService {
    return &ImprovedEmailService{logger: logger}
}

func (s *ImprovedEmailService) SendEmail(to, message string) {
    s.logger.Log("Sending email to " + to)
    fmt.Println("Email sent to", to)
}

// Usage - dependencies are explicit
func ImprovedNotifyUser(userEmail string, logger LogWriter) {
    service := NewEmailService(logger) // Dependencies are clear
    service.SendEmail(userEmail, "Hello!")
}

// Can inject mocks for testing

type MockLogger struct{}

func (m *MockLogger) Log(message string) {
    fmt.Println("Mock log:", message)
}

func TestEmailService(t *testing.T) {
    mockLogger := &MockLogger{}
    service := NewEmailService(mockLogger)
    service.SendEmail("test@example.com", "Test message")
}
```

### 8. Lack of Flexibility

Once a singleton object is created, it becomes difficult to change or replace it.

#### Example of the Problem

```go
type Config struct {
    apiURL string
    apiKey string
}

var configInstance *Config
var configOnce sync.Once

func GetConfig() *Config {
    configOnce.Do(func() {
        configInstance = &Config{
            apiURL: "https://api.production.com",
            apiKey: "prod-key-123",
        }
    })
    return configInstance
}

func MakeAPICall() string {
    config := GetConfig()
    // Problem: Even in the test environment, the production URL is used
    return "Calling " + config.apiURL
}
```

#### Improvement

```go
// Use different settings for different environments
type Environment string

const (
    Development Environment = "development"
    Staging     Environment = "staging"
    Production  Environment = "production"
)

type FlexibleConfig struct {
    apiURL string
    apiKey string
}

func NewConfig(env Environment) *FlexibleConfig {
    configs := map[Environment]FlexibleConfig{
        Development: {
            apiURL: "https://api.dev.com",
            apiKey: "dev-key-123",
        },
        Staging: {
            apiURL: "https://api.staging.com",
            apiKey: "staging-key-123",
        },
        Production: {
            apiURL: "https://api.production.com",
            apiKey: "prod-key-123",
        },
    }

    cfg := configs[env]
    return &cfg
}

// Flexibly switch settings according to the environment
func FlexibleAPICall(config *FlexibleConfig) string {
    return "Calling " + config.apiURL
}

func Example() {
    // Development environment
    devConfig := NewConfig(Development)
    FlexibleAPICall(devConfig)

    // Production environment
    prodConfig := NewConfig(Production)
    FlexibleAPICall(prodConfig)
}
```

### 9. Context-Dependent Uniqueness

The concept of being a unique object should depend on a certain scope and should not be applied globally.

#### Example of the Problem

```go
type Session struct {
    userID    int
    loginTime string
}

var sessionInstance *Session
var sessionOnce sync.Once

func GetSession() *Session {
    sessionOnce.Do(func() {
        sessionInstance = &Session{
            userID:    0,
            loginTime: "",
        }
    })
    return sessionInstance
}

// Problem: Cannot manage sessions for multiple users
func HandleRequest(userID int) {
    session := GetSession()
    // All users share the same session
    session.userID = userID
}
```

#### Improvement

```go
// Manage sessions per context
type SessionManager struct {
    sessions map[int]*Session
    mu       sync.RWMutex
}

func NewSessionManager() *SessionManager {
    return &SessionManager{
        sessions: make(map[int]*Session),
    }
}

func (sm *SessionManager) GetSession(userID int) *Session {
    sm.mu.RLock()
    session, exists := sm.sessions[userID]
    sm.mu.RUnlock()

    if !exists {
        sm.mu.Lock()
        session = &Session{
            userID:    userID,
            loginTime: "2024-01-01 00:00:00",
        }
        sm.sessions[userID] = session
        sm.mu.Unlock()
    }

    return session
}

func (sm *SessionManager) RemoveSession(userID int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    delete(sm.sessions, userID)
}

// Usage example
func ImprovedHandleRequest(userID int, sm *SessionManager) {
    session := sm.GetSession(userID)
    // Each user has an independent session
    fmt.Println("User", session.userID, "logged in at", session.loginTime)
}
```

### 10. Inefficient Memory Usage

Modern GC (Garbage Collector) manages temporary objects more efficiently than persistent objects.

#### Example of the Problem

```go
type DataProcessor struct {
    cache map[string][]byte // Permanently occupies memory
}

var processorInstance *DataProcessor
var processorOnce sync.Once

func GetDataProcessor() *DataProcessor {
    processorOnce.Do(func() {
        processorInstance = &DataProcessor{
            cache: make(map[string][]byte),
        }
    })
    return processorInstance
}

func (dp *DataProcessor) Process(data string) {
    // Issues:
    // 1. Cache accumulates indefinitely, increasing memory usage
    // 2. Since it's a singleton, memory is not released until program termination
    // 3. GC cannot efficiently reclaim memory (due to long-lived objects)
    dp.cache[data] = []byte(data)
}

// Usage example
func ProcessLargeDataset() {
    processor := GetDataProcessor()

    // Process 1 million data entries
    for i := 0; i < 1000000; i++ {
        data := fmt.Sprintf("data-%d", i)
        processor.Process(data)
    }
    // Cache retains 1 million data entries
    // Using the same processor for other processes increases memory further
}
```

#### Improvement

```go
// Improvement 1: Use short-lived objects
type ImprovedDataProcessor struct {
    // No state, or managed in local scope
}

func NewDataProcessor() *ImprovedDataProcessor {
    return &ImprovedDataProcessor{}
}

func (dp *ImprovedDataProcessor) Process(data string) []byte {
    // Create temporary data for each process
    result := []byte(data)
    // After function ends, references to result are gone, allowing GC to reclaim
    return result
}

func ProcessImprovedDataset() {
    // Create a new processor for each process
    for i := 0; i < 1000000; i++ {
        processor := NewDataProcessor()
        data := fmt.Sprintf("data-%d", i)
        result := processor.Process(data)

        // After use, references to processor and result are gone
        // GC can efficiently reclaim in the next cycle
        _ = result
    }
    // Memory usage remains constant
}
```

## Conclusion

The Singleton pattern, while seemingly convenient, causes the following significant problems:

1. Decreased testability: Difficult to replace with mocks
2. Tight coupling: Difficult to separate components
3. Lack of flexibility: Difficult to change behavior at runtime
4. Concurrency issues: Complex to implement thread-safe
5. Principle violations: Violates the Single Responsibility Principle and other SOLID principles

### Alternatives

Instead of using a singleton, the following approaches are recommended:

1. Dependency Injection (DI): Explicitly inject dependencies
2. Factory Pattern: Control instance creation
3. Context Management: Manage instances per scope
4. Functional Approach: Use stateless functions

```go
// Recommended structure
type Application struct {
    db     DBConnection
    logger LogWriter
    config *Config
}

func NewApplication(db DBConnection, logger LogWriter, config *Config) *Application {
    return &Application{
        db:     db,
        logger: logger,
        config: config,
    }
}

func main() {
    // Build dependencies explicitly
    db := NewPostgresDB("localhost:5432")
    logger := NewLogger("/var/log/app.log")
    config := NewConfig(Production)

    app := NewApplication(db, logger, config)

    // Run the application
    _ = app
}
```

The Singleton pattern should be avoided unless there are special reasons. Instead, using dependency injection or context management allows for writing testable and maintainable code.

## References

- [Clean Code Cookbook: A Collection of Recipes for Improving Code Design and Quality](https://amzn.to/47uvc3g)