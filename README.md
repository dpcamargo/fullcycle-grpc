# Steps

1 - Install protoc v3+
```brew install protobuf```

2 - Install plugins
```$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```

3 - Create proto file

4 - Generate proto
```protoc --go_out=. --go-grpc_out=. proto/course_category.proto```

5 - Run evans
```evans -r repl```

6 - Exec:
```package pb
service CategoryService
call <func>
```
