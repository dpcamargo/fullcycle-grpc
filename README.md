# Steps

1 - Install protoc v3+
```brew install protobuf```

2 - Install plugins
```$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```

3 - Create proto file

4 - Generate proto
```protoc --go_out=. --go-grpc_out=. proto/course_category.proto```