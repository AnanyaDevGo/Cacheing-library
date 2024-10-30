# Caching Client Library in Go

This caching client library, written in Go, enables efficient, unified interaction with popular caching backends: **Redis**, **Memcached**, and **In-Memory** cache with **Least Recently Used (LRU) eviction policy**. This library is designed to streamline cache operations with a unified API, providing developers the flexibility to use multiple backends interchangeably.

## Overview

The library provides:
- **Multi-Backend Support**: Redis, Memcached, and in-memory caching.
- **Unified API**: A single, simple API for caching operations across different backends.
- **In-Memory LRU Caching**: Automatically evicts the least recently used items when memory limits are reached.
- **Cache Policies**: Set expiration times, handle cache invalidation, and manage data freshness.
- **Performance Optimized**: Benchmarked for high-performance cache operations.

## Features

- **Simple CRUD Operations**: Set, get, update, and delete cache entries easily.
- **Flexible Expiration Control**: Custom time-to-live (TTL) for each cache entry.
- **In-Memory Cache with LRU**: Efficient memory usage through LRU eviction.
- **Detailed Error Handling**: Clear error reporting for all operations, especially for connection issues.
- **Performance Benchmarks**: Consistent performance across high-throughput scenarios.

## Installation

Install the library using Go modules:

```bash
go get github.com/AnanyaDevGo/Caching-Library-in-Go
```

## Usage

The library offers a consistent interface for Redis, Memcached, and in-memory caching. Below are examples demonstrating how to use each backend.

### 1. Import the Library

```go
import (
    "github.com/AnanyaDevGo/Caching-Library-in-Go/cache"
    "time"
    "fmt"
)
```

### 2. Initialize the Cache Client

The library allows you to choose a backend by initializing the respective cache client.

#### In-Memory Cache (with LRU Eviction)

```go
inMemoryCache := cache.NewInMemoryCache(100) // LRU cache with max 100 entries
inMemoryCache.Set("exampleKey", "exampleValue", time.Minute*5)
```

#### Redis Cache

```go
redisCache := cache.NewRedisCache("localhost:6379", "", 10) // Redis with DB 10
redisCache.Set("exampleKey", "exampleValue", time.Minute*5)
```

#### Memcached Cache

```go
memcachedCache := cache.NewMemcachedCache("localhost:11211")
memcachedCache.Set("exampleKey", "exampleValue", time.Minute*5)
```

### 3. Basic Operations

```go
// Set a cache entry with expiration
key := "exampleKey"
value := "exampleValue"
ttl := time.Minute * 5

err := inMemoryCache.Set(key, value, ttl)
if err != nil {
    log.Fatalf("Failed to set value: %v", err)
}

// Get a cache entry
retrievedValue, err := inMemoryCache.Get(key)
if err != nil {
    log.Fatalf("Failed to retrieve value: %v", err)
} else {
    fmt.Printf("Retrieved value: %s\n", retrievedValue)
}

// Delete a cache entry
err = inMemoryCache.Delete(key)
if err != nil {
    log.Fatalf("Failed to delete value: %v", err)
}
```

## Architecture

This library's architecture emphasizes flexibility and performance with its multi-backend approach.

- **Unified API**: Ensures consistent cache operations across Redis, Memcached, and in-memory LRU.
- **LRU In-Memory Caching**: Efficient memory management with auto-eviction of stale entries.
- **Redis and Memcached Integration**: Supports common cache operations like `Set`, `Get`, `Delete`.
- **Cache Policies**: Custom expiration and invalidation for data freshness.

## Benchmark Performance

Performance benchmarks for each backend:

- **In-Memory Cache**:
  - Set: 9.3M ops/sec
  - Get: 61M ops/sec
  - Delete: 4.3M ops/sec
  - Clear: 3.1M ops/sec

- **Redis Integration**:
  - Set: ~50k ops/sec
  - Get: ~60k ops/sec

- **Memcached Integration**:
  - Set: 10.6k ops/sec
  - Get: 8.8k ops/sec

These benchmarks reflect the library’s suitability for high-throughput caching.

## Example Directory Structure

```
├── cache/
│   ├── redis.go               # Redis client implementation
│   ├── memcached.go           # Memcached client implementation
│   ├── inmemory.go            # LRU in-memory cache
│   ├── cache.go               # Core interface and logic
│   └── errors.go              # Custom error handling
├── examples/
│   └── main.go                # Usage 
└── README.md                  # Project documentation
```

## Configuration Options

- **Redis**:
  - `HOST`: e.g., `localhost`
  - `PORT`: e.g., `6379`
  - `PASSWORD`: Optional for secure connections
  - `DB`: Default is `0`

- **Memcached**:
  - `HOST`: e.g., `localhost`
  - `PORT`: e.g., `11211`

## Contributing

Contributions are welcome! Please open issues or pull requests on GitHub, and follow Go conventions. Make sure to add tests for new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.