# WTYPE

[![zread](https://img.shields.io/badge/Ask_Zread-_.svg?style=flat&color=00b0aa&labelColor=000000&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIHZpZXdCb3g9IjAgMCAxNiAxNiIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTQuOTYxNTYgMS42MDAxSDIuMjQxNTZDMS44ODgxIDEuNjAwMSAxLjYwMTU2IDEuODg2NjQgMS42MDE1NiAyLjI0MDFWNC45NjAxQzEuNjAxNTYgNS4zMTM1NiAxLjg4ODEgNS42MDAxIDIuMjQxNTYgNS42MDAxSDQuOTYxNTZDNS4zMTUwMiA1LjYwMDEgNS42MDE1NiA1LjMxMzU2IDUuNjAxNTYgNC45NjAxVjIuMjQwMUM1LjYwMTU2IDEuODg2NjQgNS4zMTUwMiAxLjYwMDEgNC45NjE1NiAxLjYwMDFaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00Ljk2MTU2IDEwLjM5OTlIMi4yNDE1NkMxLjg4ODEgMTAuMzk5OSAxLjYwMTU2IDEwLjY4NjQgMS42MDE1NiAxMS4wMzk5VjEzLjc1OTlDMS42MDE1NiAxNC4xMTM0IDEuODg4MSAxNC4zOTk5IDIuMjQxNTYgMTQuMzk5OUg0Ljk2MTU2QzUuMzE1MDIgMTQuMzk5OSA1LjYwMTU2IDE0LjExMzQgNS42MDE1NiAxMy43NTk5VjExLjAzOTlDNS42MDE1NiAxMC42ODY0IDUuMzE1MDIgMTAuMzk5OSA0Ljk2MTU2IDEwLjM5OTlaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik0xMy43NTg0IDEuNjAwMUgxMS4wMzg0QzEwLjY4NSAxLjYwMDEgMTAuMzk4NCAxLjg4NjY0IDEwLjM5ODQgMi4yNDAxVjQuOTYwMUMxMC4zOTg0IDUuMzEzNTYgMTAuNjg1IDUuNjAwMSAxMS4wMzg0IDUuNjAwMUgxMy43NTg0QzE0LjExMTkgNS42MDAxIDE0LjM5ODQgNS4zMTM1NiAxNC4zOTg0IDQuOTYwMVYyLjI0MDFDMTQuMzk4NCAxLjg4NjY0IDE0LjExMTkgMS42MDAxIDEzLjc1ODQgMS42MDAxWiIgZmlsbD0iI2ZmZiIvPgo8cGF0aCBkPSJNNCAxMkwxMiA0TDQgMTJaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00IDEyTDEyIDQiIHN0cm9rZT0iI2ZmZiIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIvPgo8L3N2Zz4K&logoColor=ffffff)](https://zread.ai/Wuchieh/wtype)

A collection of generic data structures, concurrency-safe variants, and utility helpers for Go.

## Requirements

- Go 1.27 or later

## Installation

```bash
go get github.com/wuchieh/wtype
```

## Documentation

API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/wuchieh/wtype).

## Overview

| Type / function | Description |
| --- | --- |
| `Cache[T]`, `SafeCache[T]`, `CustomCache[T]` | TTL caches with concurrency-safe and custom storage variants |
| `Set[T]`, `SafeSet[T]` | Set implementations with JSON marshal/unmarshal support |
| `Context[T]` | Handler-chain context that also implements `context.Context` |
| `GormSlice[T]`, `SqlJSON`, `SqlJSON2[T]` | GORM and SQL JSON column helpers |
| `SyncMap[K, V]` | Type-safe wrapper around `sync.Map` |
| `String`, `SliceString` | Chainable string and string-slice helpers |
| `EventCenter` | Event subscription and dispatch with one-time handlers |
| `WaitGroup` | `sync.WaitGroup` with `Go` and `WaitContext` helpers |
| `ErrorGroup` | Concurrent task runner with error collection and fail-fast cancellation |
| `NewCtxReader` | `io.Reader` that checks context cancellation before each read |
| Utility functions | `UseOptions`, `SliceConvert`, `MapConvert`, `DoShared`, and more |

## Quick Start

### ErrorGroup

```go
group := wtype.NewErrorGroup(
	wtype.ErrorGroupWithContext(ctx),
	wtype.ErrorGroupErrorAbort(true),
)

group.Go(func(ctx context.Context) error {
	return doWork(ctx)
})

if err := group.Wait(); err != nil {
	log.Fatal(err)
}
```

### WaitGroup

```go
var wg wtype.WaitGroup
wg.Go(func() {
	// work
})

if err := wg.WaitContext(ctx); err != nil {
	log.Fatal(err)
}
```

### Context-aware reader

```go
reader := wtype.NewCtxReader(ctx, file)
_, err := io.Copy(dst, reader)
```

### Functional options

```go
cfg := wtype.UseOptions(Config{}, WithTimeout(time.Second), WithRetry(3))
```

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for the release history. The latest release is
[v1.12.0](https://github.com/Wuchieh/wtype/releases/tag/v1.12.0).

## TODO

- [x] add lint
- [x] add godoc comments

### Cache

- [x] add example
- [x] add test

### SafeCache

- [x] add example
- [x] add test

### Set

- [x] add example
- [x] add test
- [x] json.Marshal
- [x] json.Unmarshal

### SafeSet

- [x] add example
- [x] add test
- [x] json.Marshal (need use *SafeSet)
- [x] json.Unmarshal

### Context

- [x] add example
- [x] add test

### GormSlice

- [ ] add example
- [x] add test

### SyncMap

- [x] add example
- [x] add test
- [x] json.Marshal (need use *SyncMap)
- [x] json.Unmarshal

### String & StringSlice

- [x] add example
- [x] add test

### EventCenter

- [x] add example
- [x] add test
- [x] add Once

### WaitGroup

- [ ] add example
- [x] add test

### ErrorGroup

- [ ] add example
- [x] add test
- [x] add options

### CtxReader

- [ ] add example
- [x] add test

### UseOptions

- [ ] add example
- [ ] add test
