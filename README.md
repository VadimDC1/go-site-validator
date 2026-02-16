# Site Health Checker

Concurrent website health checker written in Go.

## Features
- Checks website availability using HTTP HEAD requests
- Concurrent checking with worker pool (goroutines + channels)
- Reports status (alive/dead), HTTP code, response time, and errors
- Compares sequential vs parallel execution time

## How to use
```go
go run main.go
