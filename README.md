# grpc-go

[![Node+CI](https://github.com/TakenokoTech/grpc-go/workflows/Go/badge.svg)](https://github.com/TakenokoTech/grpc-go/actions)
[![codecov](https://codecov.io/gh/TakenokoTech/grpc-go/branch/master/graph/badge.svg)](https://codecov.io/gh/TakenokoTech/grpc-go)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/TakenokoTech/grpc-go/blob/master/LICENSE)

```
# Run
go run bin/grpc-go.go

# Build
go build bin/grpc-go.go

# Test
go test ./...

# Coverage
go test -coverprofile="report/cover.out" -covermode=atomic ./...
go tool cover -html="report/cover.out" -o report/cover.html

# Lint
golint ./...

# proto
protoc -I api/ api/proto/sample.proto --go_out=plugins=grpc:api
```
