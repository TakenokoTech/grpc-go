# grpc-go

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
