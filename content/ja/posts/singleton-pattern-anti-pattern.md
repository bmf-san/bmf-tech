---
title: "Singletonパターンはなぜアンチパターンなのか"
slug: "singleton-pattern-anti-pattern"
date: 2025-10-18
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "シングルトンパターン"
  - "Golang"
draft: false
---

## はじめに

シングルトンパターンは、デザインパターンの中でも最もよく知られ、広く使われているパターンの一つである。しかし、クリーンコードや保守性の観点から見ると、多くの問題を抱えている。

本記事では、シングルトンパターンの主要な問題点について、具体的なGoのコード例を交えながら解説する。

## シングルトンパターンとは

シングルトンパターンは、クラスのインスタンスが常に1つだけ存在することを保証するデザインパターンである。

### Goでの基本的な実装

```go
package main

import (
    "sync"
)

// Database はシングルトンとして実装されたデータベース接続
type Database struct {
    connectionString string
}

var (
    instance *Database
    once     sync.Once
)

// GetInstance はDatabaseのシングルトンインスタンスを返す
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

一見便利に見えるこのパターンだが、実際には多くの問題を引き起こす。

## シングルトンパターンの問題点

### 1. 全単射性の欠如

シングルトンは現実世界に直接対応する概念ではない。現実世界では、ほとんどの概念は複数のインスタンスを持つことができる。

#### 問題の例

```go
// 現実世界では複数のデータベース接続が存在しうる
func Example() {
    // プライマリDB
    primaryDB := GetInstance()

    // レプリカDBにも接続したい...しかしシングルトンでは不可能
    // replicaDB := GetReplicaInstance() // これはできない

    primaryDB.Query("SELECT * FROM users")
}
```

現実世界では、プライマリとレプリカ、あるいは複数のデータベースに接続する必要がある場合が多い。しかし、シングルトンはこの柔軟性を奪う。

#### 改善策

```go
// インターフェースと依存性注入を使用
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

// 複数の接続を柔軟に管理できる
func ImprovedExample() {
    primaryDB := NewPostgresDB("primary.db.com:5432")
    replicaDB := NewPostgresDB("replica.db.com:5432")

    primaryDB.Query("INSERT INTO users VALUES (...)")
    replicaDB.Query("SELECT * FROM users")
}
```

### 2. 密結合

シングルトンは、分離が困難なグローバルなアクセスポイントを提供する。これにより、コード全体が強く結合される。

#### 問題の例

```go
type UserService struct {
    // データベースへの依存が隠蔽されている
}

func (s *UserService) GetUser(id int) string {
    // グローバルなシングルトンに直接依存
    db := GetInstance()
    return db.Query("SELECT * FROM users WHERE id = " + string(rune(id)))
}

func (s *UserService) CreateUser(name string) {
    db := GetInstance()
    db.Query("INSERT INTO users (name) VALUES ('" + name + "')")
}
```

このコードの問題点：
- `UserService`が`Database`のシングルトンに暗黙的に依存している
- 依存関係が明示的でなく、コードを読むだけでは分からない
- テスト時にモックに置き換えることが困難

#### 改善策

```go
// 依存性注入を使用して明示的に依存関係を示す
type UserService struct {
    db DBConnection // 依存関係が明示的
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

### 3. テストが困難

シングルトンの存在により、ユニットテストの作成が非常に困難になる。

#### 問題の例

```go
// テストしたいコード
func ProcessUser(userID int) string {
    db := GetInstance() // シングルトンに依存
    result := db.Query("SELECT * FROM users WHERE id = " + string(rune(userID)))
    return "Processed: " + result
}

// テストコード - モックに置き換えられない
func TestProcessUser(t *testing.T) {
    // 問題: 実際のデータベースが使われてしまう
    // モックに置き換える方法がない
    result := ProcessUser(1)

    // 実際のDBに接続してしまうため、テストが遅く、不安定
    if result == "" {
        t.Error("Expected non-empty result")
    }
}
```

#### 改善策

```go
// モック実装
type MockDB struct {
    queryFunc func(sql string) string
}

func (m *MockDB) Query(sql string) string {
    if m.queryFunc != nil {
        return m.queryFunc(sql)
    }
    return "mock result"
}

// テスト可能な実装
func ProcessUserImproved(userID int, db DBConnection) string {
    result := db.Query("SELECT * FROM users WHERE id = " + string(rune(userID)))
    return "Processed: " + result
}

// テストコード - モックを使用可能
func TestProcessUserImproved(t *testing.T) {
    // モックDBを注入
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

### 4. 状態の蓄積

複数のテスト実行により、シングルトンに不要なデータが蓄積される。

#### 問題の例

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

// テスト1
func TestCacheSet(t *testing.T) {
    cache := GetCache()
    cache.Set("key1", "value1")

    if cache.Get("key1") != "value1" {
        t.Error("Expected value1")
    }
}

// テスト2 - テスト1の状態が残っている
func TestCacheGet(t *testing.T) {
    cache := GetCache()

    // 問題: 前のテストのデータが残っている
    // "key1"がすでに存在してしまう
    if cache.Get("key1") != "" {
        t.Error("Expected empty cache, but got data from previous test")
    }
}
```

#### 改善策

```go
// インスタンスを都度生成
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

func (c *ImprovedCache) Get(key string) string {
    return c.data[key]
}

// テスト1 - 独立したインスタンス
func TestImprovedCacheSet(t *testing.T) {
    cache := NewCache() // 新しいインスタンス
    cache.Set("key1", "value1")

    if cache.Get("key1") != "value1" {
        t.Error("Expected value1")
    }
}

// テスト2 - 独立したインスタンス
func TestImprovedCacheGet(t *testing.T) {
    cache := NewCache() // 別の新しいインスタンス

    // 前のテストの影響を受けない
    if cache.Get("key1") != "" {
        t.Error("Expected empty cache")
    }
}
```

### 5. 並行処理の問題

シングルトンパターンを使うと、並行環境でスレッドセーフな実装が必要になり、複雑さが増す。

#### 問題の例（スレッドセーフではないシングルトン）

```go
type Counter struct {
    count int
}

var counterInstance *Counter
var counterOnce sync.Once

// シングルトンとして実装（インスタンス生成はスレッドセーフ）
func GetCounter() *Counter {
    counterOnce.Do(func() {
        counterInstance = &Counter{count: 0}
    })
    return counterInstance
}

// 問題: シングルトンなので、このメソッドがスレッドセーフでないと
// すべての呼び出し元でレースコンディションが発生する
func (c *Counter) Increment() {
    c.count++ // レースコンディション
    // 注意: コンパイルエラーにはならない
    // しかし、並行実行時に予期しない結果になる
}

func (c *Counter) GetCount() int {
    return c.count // これもレースコンディション
}

// レースコンディションの詳細:
// シングルトンのため、すべてのgoroutineが同じインスタンスにアクセスする
// c.count++ は以下の処理に分解される:
//   1. メモリから値を読み取る (READ)
//   2. 値を1増やす (INCREMENT)
//   3. メモリに書き戻す (WRITE)
//
// 例: 現在のcount = 5 の状態で、2つのgoroutineが同時にIncrement()を呼ぶ
//   goroutine A: count = 5 を読む
//   goroutine B: count = 5 を読む  ← Aと同じ値を読む
//   goroutine A: 5 + 1 = 6 を計算
//   goroutine B: 5 + 1 = 6 を計算
//   goroutine A: count = 6 を書く
//   goroutine B: count = 6 を書く ← 上書きされる
//   結果: 2回インクリメントしたのに、countは6（期待値は7）

// 並行実行時に問題が発生
func ConcurrentExample() {
    var wg sync.WaitGroup

    // 1000個のgoroutineがシングルトンの同じインスタンスにアクセス
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter := GetCounter() // すべて同じインスタンス
            counter.Increment()      // レースコンディション
        }()
    }

    wg.Wait()

    // 期待値: 1000
    // 実際: 1000より小さい値（例: 987, 934など）
    // 理由: シングルトンなので全goroutineが同じインスタンスを共有し、
    //       スレッドセーフでないIncrement()を呼ぶため
    fmt.Println("Count:", GetCounter().GetCount())
}

// レースコンディションの検出方法
// 通常の実行では問題が顕在化しないこともあるが、
// go run -race main.go を実行すると警告が表示される:
//
// WARNING: DATA RACE
// Write at 0x... by goroutine 7:
//   main.(*Counter).Increment()
// Previous write at 0x... by goroutine 6:
//   main.(*Counter).Increment()
```

#### 改善策1: スレッドセーフなシングルトン実装

シングルトンを使う場合、すべてのメソッドをスレッドセーフにする必要がある：

```go
type Counter struct {
    count int
    mu    sync.Mutex // すべてのメソッドでロックが必要
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

// すべてのメソッドでmutexによる排他制御が必要
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

問題点：
- シングルトンなので、すべてのgoroutineが同じmutexで競合する
- パフォーマンスのボトルネックになる
- デッドロックのリスクが高まる
- シングルトンの他の問題（テスタビリティ、密結合など）も残る

#### 改善策2: シングルトンを使わない実装

```go
// シングルトンを使わず、必要に応じて複数のインスタンスを作成
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

// 改善例: 各goroutineグループで独立したカウンターを使用
func ImprovedConcurrentExample() {
    // 10個の独立したカウンターを作成
    counters := make([]*SafeCounter, 10)
    for i := range counters {
        counters[i] = NewSafeCounter()
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(idx int) {
            defer wg.Done()
            // 各goroutineは異なるカウンターにアクセス
            // → mutex競合が分散される
            counters[idx%10].Increment()
        }(i)
    }

    wg.Wait()

    // 最後に集計
    total := 0
    for _, c := range counters {
        total += c.GetCount()
    }
    fmt.Println("Total count:", total) // 確実に1000
}
```

メリット：
- 複数のインスタンスを使うことで、mutex競合が分散される
- パフォーマンスが向上する
- テストが容易（各カウンターを独立してテスト可能）
- シングルトンの制約から解放される

### 6. 単一責任原則の違反

シングルトンクラスは、本来の責任に加えて「インスタンス管理」という責任も持つことになる。

#### 問題の例

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

// Loggerは2つの責任を持つ:
// 1. ログを書く（本来の責任）
// 2. 自身のインスタンスを管理する（シングルトンの責任）
func (l *Logger) Log(message string) {
    // ログ処理
    fmt.Println("Log to", l.logFile, ":", message)
}
```

#### 改善策

```go
// Loggerは本来の責任（ログ記録）のみを持つ
type ImprovedLogger struct {
    logFile string
}

func NewLogger(logFile string) *ImprovedLogger {
    return &ImprovedLogger{logFile: logFile}
}

func (l *ImprovedLogger) Log(message string) {
    fmt.Println("Log to", l.logFile, ":", message)
}

// インスタンス管理は別の場所（main関数やDIコンテナ）で行う
func main() {
    logger := NewLogger("/var/log/app.log")

    // 必要なら複数のロガーも作成可能
    errorLogger := NewLogger("/var/log/error.log")

    logger.Log("Application started")
    errorLogger.Log("An error occurred")
}
```

### 7. 依存性注入の阻害

シングルトンは、依存性注入のパターンを阻害し、コンポーネント間の分離を困難にする。

#### 問題の例

```go
type EmailService struct {
    // 依存関係が隠蔽されている
}

func (s *EmailService) SendEmail(to, message string) {
    // グローバルなシングルトンに依存
    logger := GetLogger()
    logger.Log("Sending email to " + to)

    // 実際のメール送信処理
    fmt.Println("Email sent to", to)
}

// 使用側
func NotifyUser(userEmail string) {
    service := &EmailService{} // 依存関係が見えない
    service.SendEmail(userEmail, "Hello!")
}
```

#### 改善策

```go
// インターフェースを定義
type LogWriter interface {
    Log(message string)
}

// EmailServiceは依存関係を明示的に受け取る
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

// 使用側 - 依存関係が明示的
func ImprovedNotifyUser(userEmail string, logger LogWriter) {
    service := NewEmailService(logger) // 依存関係が明確
    service.SendEmail(userEmail, "Hello!")
}

// テストでモックを注入可能
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

### 8. 柔軟性の欠如

一度作成されたシングルトンオブジェクトの変更や置換が困難である。

#### 問題の例

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
    // 問題: テスト環境でも本番環境のURLが使われてしまう
    return "Calling " + config.apiURL
}
```

#### 改善策

```go
// 環境ごとに異なる設定を使用可能
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

// 環境に応じて柔軟に設定を切り替え可能
func FlexibleAPICall(config *FlexibleConfig) string {
    return "Calling " + config.apiURL
}

func Example() {
    // 開発環境
    devConfig := NewConfig(Development)
    FlexibleAPICall(devConfig)

    // 本番環境
    prodConfig := NewConfig(Production)
    FlexibleAPICall(prodConfig)
}
```

### 9. 文脈依存の一意性

一意なオブジェクトであるという概念は、一定のスコープ内に依存するべきであり、グローバルに適用すべきではない。

#### 問題の例

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

// 問題: 複数ユーザーのセッションを管理できない
func HandleRequest(userID int) {
    session := GetSession()
    // すべてのユーザーが同じセッションを共有してしまう
    session.userID = userID
}
```

#### 改善策

```go
// コンテキストごとにセッションを管理
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

// 使用例
func ImprovedHandleRequest(userID int, sm *SessionManager) {
    session := sm.GetSession(userID)
    // 各ユーザーが独立したセッションを持つ
    fmt.Println("User", session.userID, "logged in at", session.loginTime)
}
```

### 10. 非効率なメモリ使用

現代のGC（ガベージコレクタ）は、永続的なオブジェクトよりも一時的なオブジェクトを効率的に管理する。

#### 問題の例

```go
type DataProcessor struct {
    cache map[string][]byte // 永続的にメモリを占有
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
    // 問題点:
    // 1. キャッシュが無制限に蓄積され、メモリ使用量が増大し続ける
    // 2. シングルトンなのでプログラム終了までメモリが解放されない
    // 3. GCが効率的にメモリを回収できない（長寿命オブジェクトのため）
    dp.cache[data] = []byte(data)
}

// 使用例
func ProcessLargeDataset() {
    processor := GetDataProcessor()

    // 100万件のデータを処理
    for i := 0; i < 1000000; i++ {
        data := fmt.Sprintf("data-%d", i)
        processor.Process(data)
    }
    // キャッシュに100万件のデータが残り続ける
    // 他の処理で同じプロセッサを使うと、メモリがさらに増える
}
```

#### 改善策

```go
// 改善案1: 短命なオブジェクトを使用
type ImprovedDataProcessor struct {
    // 状態を持たない、またはローカルスコープで管理
}

func NewDataProcessor() *ImprovedDataProcessor {
    return &ImprovedDataProcessor{}
}

func (dp *ImprovedDataProcessor) Process(data string) []byte {
    // 処理ごとに一時的なデータを作成
    result := []byte(data)
    // 関数終了後、resultへの参照がなくなればGCが回収できる
    return result
}

func ProcessImprovedDataset() {
    // 処理ごとに新しいプロセッサを作成
    for i := 0; i < 1000000; i++ {
        processor := NewDataProcessor()
        data := fmt.Sprintf("data-%d", i)
        result := processor.Process(data)

        // 使用後、processorとresultへの参照がなくなる
        // GCが次のサイクルで効率的に回収できる
        _ = result
    }
    // メモリ使用量は一定に保たれる
}
```

## まとめ

シングルトンパターンは、一見便利に見えるが、以下の重大な問題を引き起こす：

1. テスタビリティの低下: モックへの置き換えが困難
2. 密結合: コンポーネント間の分離が困難
3. 柔軟性の欠如: 実行時の動作変更が困難
4. 並行処理の問題: スレッドセーフな実装が複雑
5. 原則違反: 単一責任原則など、SOLIDの原則に反する

### 代替案

シングルトンの代わりに、以下のアプローチを推奨する：

1. 依存性注入（DI）: 依存関係を明示的に注入
2. ファクトリーパターン: インスタンス生成を制御
3. コンテキスト管理: スコープごとにインスタンスを管理
4. 関数型アプローチ: 状態を持たない関数を使用

```go
// 推奨される構造
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
    // 依存関係を明示的に構築
    db := NewPostgresDB("localhost:5432")
    logger := NewLogger("/var/log/app.log")
    config := NewConfig(Production)

    app := NewApplication(db, logger, config)

    // アプリケーションを実行
    _ = app
}
```

シングルトンパターンは、特別な理由がない限り避けるべきである。代わりに、依存性注入やコンテキスト管理を使用することで、テスタブルで保守性の高いコードを書くことができる。

## 参考

- [クリーンコードクックブック ―コードの設計と品質を改善するためのレシピ集](https://amzn.to/47uvc3g)
