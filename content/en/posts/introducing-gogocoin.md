---
title: Introducing gogocoin — A Self-Hosted Crypto Trading Bot
description: 'A deep dive into gogocoin, a Go-based self-hosted Bitcoin trading bot for bitFlyer. Covers the pluggable strategy architecture, layered design and dependency rules, trading flow, data model, and balance cache.'
slug: introducing-gogocoin
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - Golang
  - Bitcoin
  - Trading
  - Infrastructure
translation_key: introducing-gogocoin
---

# Introducing gogocoin — A Self-Hosted Crypto Trading Bot

## Why I Built It

Open-source crypto bots and automated trading services are everywhere. I built gogocoin anyway because I wanted the hands-on experience of implementing something that works exactly as I intend and actually earning returns with my own money. I had built a similar bot once before; this time I rebuilt it from scratch with the help of AI. Running it in production has been a continuous source of learning — it has become a hobby as much as a software project.

[gogocoin](https://github.com/bmf-san/gogocoin)

## Getting Started

**gogocoin is bitFlyer-only.** All exchange communication goes through the author's own [`go-bitflyer-api-client`](https://github.com/bmf-san/go-bitflyer-api-client) library; no other exchange works. The bot places orders via bitFlyer's spot-only endpoint (`/v1/me/sendchildorder`), so **margin / futures trading (e.g. FX\_BTC\_JPY) does not work**.

There are two ways to use gogocoin.

**A. Use as a library (recommended)**

`example/` is a fully working sample and a starting point for your own repo.

```bash
git clone https://github.com/bmf-san/gogocoin.git && cd gogocoin/example

# Create config and set API credentials via environment variables
cp configs/config.example.yaml configs/config.yaml
export BITFLYER_API_KEY=your_key
export BITFLYER_API_SECRET=your_secret

make run
# or: go run ./cmd/

# → Dashboard at http://localhost:8080
```

Using `example/configs/config.example.yaml` as-is, the bot runs an XRP/JPY scalping strategy with a 1000 JPY order size. Adjust `trading.symbols` and `strategy_params.scalping.order_notional` to trade different pairs or sizes. The bot stores trade data in SQLite (no external database needed).

You can also integrate gogocoin into your own module via `go get github.com/bmf-san/gogocoin@latest`.

**B. Docker for quick testing**

`example/` includes a `Dockerfile` and `docker-compose.yml` that build a fully working binary with the same EMA+RSI scalping strategy registered.

```bash
git clone https://github.com/bmf-san/gogocoin.git && cd gogocoin/example
cp configs/config.example.yaml configs/config.yaml
# Edit configs/config.yaml and set your API credentials
make up

# → Dashboard at http://localhost:8080
```

The Dockerfile build context is the repo root, so run `make up` from the `example/` directory.

## Architecture

The codebase follows a four-layer architecture. `internal/` houses domain logic, use cases, and external adapters (bitFlyer client, SQLite repository, HTTP handlers, etc.); `pkg/strategy` is a public package providing the Strategy interface definition and a scalping reference implementation. The Composition Root (wiring all services together) lives in the caller's repository — `example/cmd/main.go` is a working sample.

### C4 Context — System Overview

```mermaid
C4Context
    Person(operator, "Operator", "Administrator of the trading bot")
    System(gogocoin, "gogocoin", "Automated scalping trading bot")
    System_Ext(bitflyer, "bitFlyer", "Crypto exchange REST / WebSocket API")
    System_Ext(sqlite, "SQLite", "Local database")

    Rel(operator, gogocoin, "Control trading and check status via REST API")
    Rel(gogocoin, bitflyer, "Place orders, fetch balances, receive market data")
    Rel(gogocoin, sqlite, "Save trades, positions, and performance records")
```

### C4 Container — Main Containers

```mermaid
C4Container
    Person(operator, "Operator")

    System_Boundary(gogocoin, "gogocoin") {
        Container(example, "example/cmd", "Go", "Composition Root — startup/shutdown (caller-provided)")
        Container(http, "adapter/http", "Go net/http", "REST API server")
        Container(worker, "adapter/worker", "Go goroutine", "Background worker group")
        Container(usecase, "usecase/", "Go", "Business logic (trading / strategy / risk / analytics)")
        Container(domain, "domain/", "Go", "Domain models and interface definitions")
        Container(infra_bf, "infra/exchange/bitflyer", "Go", "bitFlyer API client")
        Container(infra_db, "infra/persistence", "Go + SQLite", "SQLite persistence")
    }

    System_Ext(bitflyer_api, "bitFlyer API", "REST / WebSocket")
    SystemDb_Ext(sqlite, "SQLite")

    Rel(operator, http, "HTTP/JSON")
    Rel(example, http, "starts")
    Rel(example, worker, "starts")
    Rel(http, usecase, "uses")
    Rel(worker, usecase, "uses")
    Rel(usecase, domain, "uses")
    Rel(infra_bf, domain, "implements IFs")
    Rel(infra_db, domain, "implements IFs")
    Rel(infra_bf, bitflyer_api, "HTTPS / WSS")
    Rel(infra_db, sqlite, "SQL")
```

### C4 Component — usecase/trading

```mermaid
C4Component
    Container_Boundary(trading, "usecase/trading") {
        Component(trader, "BitflyerTrader", "Go", "Place/cancel orders, fetch balance")
        Component(monitor, "OrderMonitor", "Go goroutine", "Poll and monitor order status")
        Component(pnl, "PnLCalculator", "Go", "Post-fill P&L calculation and persistence")
        Component(balance, "BalanceService", "Go", "Balance retrieval and TTL cache")
        Component(order, "OrderService", "Go", "Order validation and placement")
        Component(validator, "OrderValidator", "Go", "Order size validation and balance check")
    }

    ComponentDb(tradeRepo, "TradeRepository", "domain.TradeRepository")
    ComponentDb(positionRepo, "PositionRepository", "domain.PositionRepository")
    ComponentDb(balanceRepo, "BalanceRepository", "domain.BalanceRepository")

    Rel(trader, monitor, "starts / watches")
    Rel(trader, order, "delegates PlaceOrder")
    Rel(trader, validator, "ValidateOrder / CheckBalance")
    Rel(monitor, order, "GetOrders (OrderGetter IF)")
    Rel(monitor, pnl, "saveTradeToDB → CalculateAndSave")
    Rel(monitor, balance, "UpdateBalanceToDB after fill")
    Rel(pnl, tradeRepo, "SaveTrade")
    Rel(pnl, positionRepo, "GetOpenPositions / UpdatePosition / SavePosition")
    Rel(trader, balance, "GetBalance")
    Rel(balance, balanceRepo, "SaveBalance / GetLatestBalances")
```

### Dependency Graph

```mermaid
graph LR
    cmd([cmd])

    cmd --> adp_http[adapter/http]
    cmd --> adp_worker[adapter/worker]
    cmd --> infra_bf[infra/exchange/bitflyer]
    cmd --> infra_db[infra/persistence]
    cmd --> domain([domain])

    adp_http --> uc_trading[usecase/trading]
    adp_http --> uc_analytics[usecase/analytics]
    adp_http --> domain
    adp_worker --> uc_trading
    adp_worker --> uc_strategy[usecase/strategy]
    adp_worker --> uc_risk[usecase/risk]
    adp_worker --> domain

    uc_trading --> domain
    uc_strategy --> domain
    uc_risk --> domain
    uc_analytics[usecase/analytics] --> domain

    adp_worker --> uc_analytics

    infra_bf --> domain
    infra_db --> domain
    logger --> domain
    cmd --> config[config]
    cmd --> logger
```

Dependency rules are enforced in CI:

| Rule | Detail |
|---|---|
| `domain/` has zero internal imports | stdlib only; knows nothing of infra or usecase |
| `usecase/` does not import `infra/` | depends only on `domain/` interfaces |
| `adapter/` holds no concrete infra types | uses `domain/` interfaces only |
| `infra/` implements `domain/` | knows nothing of `usecase/` or `adapter/` |
| Composition Root lives in the caller's repository | `internal/` needs no wiring |

The public API (subject to semantic versioning) lives under `pkg/`. `pkg/engine` is the engine entry point; `pkg/strategy` provides the Strategy interface and registry.

## Use Cases

```mermaid
graph LR
    OP(["👤 Operator"])
    BF(["🏦 bitFlyer"])
    SYS(["⚙️ System"])

    subgraph sys["gogocoin system boundary"]
        UC1(Start trading)
        UC2(Stop trading)
        UC3(Check trading status)
        UC4(Check positions)
        UC5(Check performance)
        UC6(Check market data)
        UC7(Check balance)
        UC8(View trade history)
        UC9(View order list)
        UC10(View logs)
        UC11(View config)
        UC12(Reset strategy)
        UC13(Detect signal with scalping strategy)
        UC14(Check risk)
        UC15(Place order)
        UC16(Monitor order status)
        UC17(Calculate and record P&L)
        UC18(Maintenance: clean up old data)
        UC19(Monitor and update strategy parameters)
        UC20(Handle order timeout / cancellation)
    end

    OP --> UC1
    OP --> UC2
    OP --> UC3
    OP --> UC4
    OP --> UC5
    OP --> UC6
    OP --> UC7
    OP --> UC8
    OP --> UC9
    OP --> UC10
    OP --> UC11
    OP --> UC12
    UC13 --> UC14
    UC14 --> UC15
    UC15 --> UC16
    UC16 --> UC17
    BF -.->|"fill notification (polling)"| UC16
    SYS --> UC13
    SYS --> UC18
    SYS --> UC19
    SYS --> UC20
```

The operator controls and monitors the bot via the HTTP API (including the web dashboard). Signal generation, order placement, P&L calculation, and data cleanup run autonomously.

## Trading Flow

The following shows the main path from receiving a WebSocket tick to filling an order and recording P&L.

### 6.1 Scalping Trading Flow

```mermaid
sequenceDiagram
    participant SW as StrategyWorker
    participant ST as Strategy
    participant SigW as SignalWorker
    participant TC as TradingController
    participant RM as risk.Manager
    participant BP as balanceProvider
    participant TR as BitflyerTrader
    participant BF as bitFlyer API
    participant OM as OrderMonitor
    participant PNL as PnLCalculator
    participant BS as BalanceService
    participant DB as persistence
    participant CB as callback

    note over RM,BP: risk.Manager depends on the balanceProvider local IF.<br/>BP is implemented by BitflyerTrader. BitflyerTrader.GetBalance()<br/>delegates internally to BalanceService (TTL cache).
    note over SW,SigW: StrategyWorker writes signals to a channel.<br/>SignalWorker reads from the channel and performs risk check + order placement.

    SW->>ST: Analyze(history []MarketData)
    ST-->>SW: Signal(BUY)
    SW-)SigW: signalCh <- signal (channel send)
    SigW->>TC: IsTradingEnabled()
    TC-->>SigW: true
    SigW->>RM: CheckRiskManagement(ctx, signal)
    RM->>BP: GetBalance(ctx)
    BP->>BS: GetBalance(ctx)
    note over BS: Check TTL cache (10s).<br/>On cache hit, skip BF call.
    alt cache miss
        BS->>BF: GET /v1/me/getbalance
        BF-->>BS: balance
    end
    BS-->>BP: balance
    BP-->>RM: balance
    alt risk violation (insufficient balance / excess position)
        RM-->>SigW: non-nil error (insufficient funds, limit exceeded, etc.)
        SigW->>SigW: skip (wait for next tick)
    else risk OK
        RM-->>SigW: nil
        note over SigW: createOrderFromSignal() builds domain.OrderRequest
        SigW->>TR: PlaceOrder(ctx, order)
        TR->>BF: POST /v1/me/sendchildorder
        BF-->>TR: order_id
        note over TR,OM: MonitorExecution is started as a goroutine.<br/>PlaceOrder returns immediately (async).
        TR-)OM: go MonitorExecution(ctx, result)

        loop Polling up to 90s every 15s
            OM->>BF: GET /v1/me/getchildorders
            BF-->>OM: status=ACTIVE
        end

        BF-->>OM: status=COMPLETED
        note over OM,PNL: OrderMonitor.saveTradeToDB() calls PnL directly<br/>(before onOrderCompleted callback).
        OM->>PNL: CalculateAndSave(result)
        note over PNL,DB: For SELL, GetOpenPositions reads outside the tx (pre-read).<br/>SQLite's default isolation is effectively serializable,<br/>so phantom read risk is minimal. Pre-reading minimizes<br/>work inside the tx and reduces deadlock risk.
        PNL->>DB: GetOpenPositions() [SELL only, outside tx]
        PNL->>DB: BeginTx()
        PNL->>DB: SavePosition() [BUY] / UpdatePosition() [SELL]
        PNL->>DB: SaveTrade()
        PNL->>DB: Commit()
        PNL-->>OM: (pnl float64)
        OM->>BS: InvalidateBalanceCache()
        OM->>BS: UpdateBalanceToDB(ctx)
        BS->>BF: GET /v1/me/getbalance
        BS->>DB: SaveBalance(balance)
        OM->>CB: onOrderCompleted(result)
    end
```

`StrategyWorker` and `SignalWorker` are connected asynchronously via a Go channel. `PlaceOrder()` returns immediately after placing the order; `OrderMonitor` handles fill monitoring in a goroutine. `PnLCalculator` saves position and trade records within the same transaction; `OrderMonitor` appends the balance snapshot separately after that completes.

### 6.2 REST API Trading Control Flow

```mermaid
sequenceDiagram
    participant C as HTTP Client
    participant H as adapter/http
    participant TC as TradingController
    participant DB as AppStateRepository

    C->>H: POST /api/trading/start
    H->>TC: SetTradingEnabled(ctx, true)
    TC->>DB: SaveAppState("trading_enabled", "true")
    DB-->>TC: nil
    TC-->>H: nil
    H-->>C: 200 OK

    C->>H: POST /api/trading/stop
    H->>TC: SetTradingEnabled(ctx, false)
    TC->>DB: SaveAppState("trading_enabled", "false")
    DB-->>TC: nil
    TC-->>H: nil
    H-->>C: 200 OK
```

### 6.3 Market Data Collection Flow

```mermaid
sequenceDiagram
    participant BS as bootstrap
    participant WM as WorkerManager
    participant WS as bitflyer WebSocket
    participant MW as MarketDataWorker
    participant DB as MarketDataRepository

    BS->>WS: Connect()
    BS->>WM: StartAll(ctx)
    WM-)MW: Run(ctx)

    loop Receive tick data
        WS-->>MW: Tick(price, volume, ...)
        MW->>DB: SaveMarketData(tick)
    end

    note over BS,WS: On disconnect, bootstrap reconnects (independent of WorkerManager lifecycle)
```

### 6.4 Order Timeout / CANCELED•EXPIRED Flow

```mermaid
sequenceDiagram
    participant TR as BitflyerTrader
    participant BF as bitFlyer API
    participant OM as OrderMonitor
    participant PNL as PnLCalculator
    participant DB as persistence
    participant LOG as Logger

    TR->>BF: POST /v1/me/sendchildorder
    BF-->>TR: order_id
    note over TR,OM: MonitorExecution is started as a goroutine (no return value).<br/>Results are delivered via the onOrderCompleted callback.
    TR-)OM: go MonitorExecution(ctx, result)

    alt Timeout (90s elapsed)
        loop Polling continues
            OM->>BF: GET /v1/me/getchildorders
            BF-->>OM: status=ACTIVE
        end
        OM->>BF: GET /v1/me/getchildorders (saveFinalOrderState)
        BF-->>OM: confirm final status
        OM->>LOG: Warn("Order monitoring timeout", order_id)
        note over OM: goroutine exits. No return to PlaceOrder.
    else Terminal status (CANCELED / EXPIRED / REJECTED)
        OM->>BF: GET /v1/me/getchildorders
        BF-->>OM: status=CANCELED
        OM->>LOG: Warn("order terminal", status, order_id)
        note over OM,PNL: saveTradeToDB is called even on CANCELED to record the trade.<br/>Balance update and onOrderCompleted callback are NOT called.
        OM->>PNL: CalculateAndSave(result) [cancel record]
        PNL->>DB: BeginTx()
        PNL->>DB: SaveTrade() [status=CANCELED]
        PNL->>DB: Commit()
    end
```

### 6.5 Rate Limit Retry Flow

```mermaid
sequenceDiagram
    participant UC as usecase
    participant BF as infra/exchange/bitflyer
    participant API as bitFlyer API

    UC->>BF: PlaceOrder(req)
    note over BF: Client.WithRetry() manages retries.<br/>The usecase layer is unaware of retry logic.
    BF->>API: POST /v1/me/sendchildorder
    API-->>BF: 429 Too Many Requests
    loop Up to MaxRetries (exponential backoff)
        BF->>BF: exponential backoff wait
        BF->>API: POST /v1/me/sendchildorder (retry)
    end
    alt Retry succeeded
        API-->>BF: 200 OK
        BF-->>UC: order_id
    else Retry limit exceeded
        BF-->>UC: domain.ErrRateLimitExceeded
        note over UC: Cast with errors.As(err, &apiErr) to *domain.Error,<br/>check apiErr.Type == domain.ErrTypeRateLimit, then propagate
    end
```

### 6.6 MaintenanceWorker Flow

```mermaid
sequenceDiagram
    participant BS as bootstrap
    participant WM as WorkerManager
    participant MW as MaintenanceWorker
    participant DB as MaintenanceRepository
    participant LOG as Logger

    BS->>WM: StartAll(ctx)
    WM-)MW: Run(ctx)

    loop Periodic execution (nightly)
        MW->>DB: GetDatabaseSize()
        DB-->>MW: size bytes
        MW->>DB: CleanupOldData(retentionDays)
        DB-->>MW: deleted rows
        MW->>DB: GetTableStats()
        DB-->>MW: stats
        MW->>LOG: Info("maintenance done", stats)
    end

    note over MW: Exits immediately on ctx.Done()
```

## Strategy Interface

Every trading strategy follows the `Strategy` interface defined in `pkg/strategy/strategy.go`. This keeps the engine decoupled from any specific algorithm:

```go
// AutoScaleConfig holds the order-size auto-scaling parameters returned by
// Strategy.GetAutoScaleConfig. The engine uses this to compute buy notional
// without reading strategy-specific config keys directly.
type AutoScaleConfig struct {
    Enabled     bool
    BalancePct  float64 // % of available JPY balance to use (0-100)
    MaxNotional float64 // hard cap in JPY; 0 = unlimited
    FeeRate     float64
}

type Strategy interface {
    // GenerateSignal generates a signal from the latest market data point and
    // the historical series for the same symbol.
    GenerateSignal(ctx context.Context, data *MarketData, history []MarketData) (*Signal, error)

    // Analyze generates a signal from a batch of historical data.
    Analyze(data []MarketData) (*Signal, error)

    // Lifecycle
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
    IsRunning() bool
    GetStatus() StrategyStatus
    Reset() error

    // Metrics & trade accounting
    GetMetrics() StrategyMetrics
    RecordTrade()
    InitializeDailyTradeCount(count int)

    // Configuration
    Name() string
    Description() string
    Version() string
    Initialize(config map[string]interface{}) error
    UpdateConfig(config map[string]interface{}) error
    GetConfig() map[string]interface{}

    // Order sizing — each strategy owns this logic so the engine never reads
    // strategy-specific config keys directly.
    GetStopLossPrice(entry float64) float64   // 0 = no stop-loss
    GetTakeProfitPrice(entry float64) float64 // 0 = no take-profit
    GetBaseNotional(symbol string) float64
    GetAutoScaleConfig() AutoScaleConfig
}
```

`Initialize()` receives the `strategy_params.<name>` block from `config.yaml` as a `map[string]interface{}`. `UpdateConfig()` allows live parameter updates via the HTTP API without restarting the bot.

Strategies self-register via the global registry using the same mechanism as `database/sql` driver registration. A `register.go` file in each strategy package calls `strategy.Register("name", constructor)` inside `init()`, and `main.go` pulls the strategy in with a blank import (`_ "github.com/bmf-san/gogocoin/pkg/strategy/scalping"`). Adding a new strategy is an `import` change in `main.go` — no engine code needs to change.

## Engine Risk Management

The engine (`StrategyWorker`) enforces stop-loss and take-profit on every market tick, independently of any signal. It calls `GetStopLossPrice` / `GetTakeProfitPrice` on each tick and closes the position immediately when the price crosses the threshold — no signal required.

A `max_open_positions_per_symbol: 1` guard in `config.yaml` prevents position stacking. Without it, consecutive BUY signals during a downtrend accumulate multiple open positions on the same symbol, and when stop-loss fires all of them close simultaneously, multiplying the loss. With the guard set to 1, any BUY is rejected if the symbol already has an open position.

The engine also supports balance-proportional order sizing as a framework feature via `GetAutoScaleConfig()`. To enable it, override `GetAutoScaleConfig()` in your own strategy to return `Enabled: true` with a `BalancePct` and optional `MaxNotional`.

## Balance Cache — Double-Checked Locking

The trading loop polls account balance frequently. Calling the bitFlyer REST API on every tick would quickly exhaust the rate limit of 50 requests per minute. `BalanceService` caches the result with a 60-second TTL and uses a double-checked locking pattern to prevent thundering-herd API calls when the cache expires:

```go
func (s *BalanceService) GetBalance(ctx context.Context) ([]domain.Balance, error) {
    // First check: read without write lock
    s.cache.mu.RLock()
    cacheTimestamp := s.cache.timestamp
    cacheData := s.cache.data
    s.cache.mu.RUnlock()

    if time.Since(cacheTimestamp) < CacheDuration && len(cacheData) > 0 {
        result := make([]domain.Balance, len(cacheData))
        copy(result, cacheData)
        return result, nil
    }

    // Serialize fetches: only one goroutine calls the API at a time
    s.fetchMu.Lock()
    defer s.fetchMu.Unlock()

    // Second check: re-verify after acquiring the lock
    s.cache.mu.RLock()
    cacheTimestamp = s.cache.timestamp
    cacheData = s.cache.data
    s.cache.mu.RUnlock()
    if time.Since(cacheTimestamp) < CacheDuration && len(cacheData) > 0 {
        result := make([]domain.Balance, len(cacheData))
        copy(result, cacheData)
        return result, nil
    }

    // ... call API and update cache
}
```

The outer `cache.mu` (a `sync.RWMutex`) allows concurrent reads of fresh cache data. The inner `fetchMu` (a `sync.Mutex`) serialises API calls so that exactly one goroutine fetches when the cache is stale.

## Data Model

Trade data is persisted in SQLite. The table-to-domain-model mapping:

| Table | Content | Notes |
|---|---|---|
| `trades` | Filled order records | `order_id UNIQUE` for idempotency. Immutable (no UPDATE) |
| `positions` | FIFO positions | 3 states: OPEN / PARTIAL / CLOSED. Created on BUY, updated on SELL |
| `balances` | Balance snapshots | Append-only (no overwrites). Latest row per currency |
| `market_data` | WebSocket tick data | UNIQUE(symbol, timestamp) |
| `performance_metrics` | Daily performance metrics | Snapshot appended after each fill |
| `logs` | Structured log entries | `fields` column is JSON |
| `app_state` | Runtime flag KV store | e.g. `trading_enabled` |

```mermaid
erDiagram
    TRADES {
        INTEGER id PK
        TEXT    symbol
        TEXT    side
        TEXT    type
        REAL    size
        REAL    price
        REAL    fee
        TEXT    status
        TEXT    order_id "UNIQUE"
        DATETIME executed_at
        DATETIME created_at
        DATETIME updated_at
        TEXT    strategy_name
        REAL    pnl
    }

    POSITIONS {
        INTEGER id PK
        TEXT    symbol
        TEXT    side
        REAL    size
        REAL    used_size
        REAL    remaining_size
        REAL    entry_price
        REAL    current_price
        REAL    unrealized_pl
        REAL    pnl
        TEXT    status
        TEXT    order_id
        DATETIME created_at
        DATETIME updated_at
    }

    BALANCES {
        INTEGER  id PK
        TEXT     currency
        REAL     available
        REAL     amount
        DATETIME timestamp
    }

    MARKET_DATA {
        INTEGER  id PK
        TEXT     symbol
        DATETIME timestamp
        REAL     open
        REAL     high
        REAL     low
        REAL     close
        REAL     volume
        DATETIME created_at
    }

    PERFORMANCE_METRICS {
        INTEGER  id PK
        DATETIME date
        REAL     total_return
        REAL     daily_return
        REAL     win_rate
        REAL     max_drawdown
        REAL     sharpe_ratio
        REAL     profit_factor
        INTEGER  total_trades
        INTEGER  winning_trades
        INTEGER  losing_trades
        REAL     average_win
        REAL     average_loss
        REAL     largest_win
        REAL     largest_loss
        INTEGER  consecutive_wins
        INTEGER  consecutive_loss
        REAL     total_pnl
    }

    LOGS {
        INTEGER  id PK
        TEXT     level
        TEXT     category
        TEXT     message
        TEXT     fields
        DATETIME timestamp
    }

    APP_STATE {
        TEXT     key PK
        TEXT     value
        DATETIME updated_at
    }

    POSITIONS ||--o{ TRADES : "symbol (FIFO, logical join)"
```

No foreign key constraints are defined. The only cross-table logical reference is between `positions` and `trades`, but `PnLCalculator` writes both within the **same transaction** (BeginTx → SavePosition/UpdatePosition → SaveTrade → Commit). Transaction atomicity guarantees consistency, making DB-level FK constraints unnecessary.

## Web Dashboard

![Dashboard](/assets/images/posts/introducing-gogocoin/01_dashboard.png)

The embedded web UI at `http://localhost:8080` has four pages, navigable via the sidebar.

- **Dashboard** — Four summary cards (total P&L, today's P&L, win rate, daily trade count) plus a system status panel showing connection state, active strategy, uptime, and live prices for monitored currency pairs.
- **Performance** — Per-currency balance table (total / available), three risk metrics (Sharpe ratio, profit factor, max drawdown), and a daily P&L history table.
- **Trade History** — Full trade log with timestamp, currency pair, side, price, size, fee, and P&L columns.
- **System Logs** — Log viewer filterable by level (DEBUG / INFO / WARN / ERROR) and category (system / trading / API / strategy / UI / data).

Start and stop buttons in the top bar control the bot in real time. Configuration (API credentials, trading parameters) lives in `config.yaml`. The bot writes data to SQLite — no external database required.

## Running in Production

gogocoin ships as a single statically-linked binary. [gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template) is a sample reference for running it on ConoHa VPS, covering systemd configuration and deployment steps.

Initial VPS setup (systemd service installation, etc.) uses `make setup`. Ongoing deployment is automated via the included GitHub Actions workflow (`workflow_dispatch`), which builds for linux/amd64 and transfers the binary to the VPS with `rsync`.

## Summary

gogocoin is a minimal self-hosted trading bot that lets you freely implement your own trading strategy and run it with real money. Having your own code directly tied to real P&L is genuinely exciting — and there is no end to how deep you can go tuning strategies and building new features. If any of this sounds interesting, feel free to give it a try.

- **GitHub**: [bmf-san/gogocoin](https://github.com/bmf-san/gogocoin)
- **VPS Template**: [bmf-san/gogocoin-vps-template](https://github.com/bmf-san/gogocoin-vps-template)
