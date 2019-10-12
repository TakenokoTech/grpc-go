# grpc-go

[![Node+CI](https://github.com/TakenokoTech/grpc-go/workflows/Go/badge.svg)](https://github.com/TakenokoTech/grpc-go/actions)
[![codecov](https://codecov.io/gh/TakenokoTech/grpc-go/branch/master/graph/badge.svg)](https://codecov.io/gh/TakenokoTech/grpc-go)

```
# Run
go run bin/grpc-go.go

# Build
go build bin/grpc-go.go

# Test
go test ./...

# Coverage
go test -race -coverprofile="report/cover.txt" -covermode=atomic ./...
go tool cover -html="report/cover.out" -o report/cover.html

# proto
protoc -I api/ api/proto/sample.proto --go_out=plugins=grpc:api
```
