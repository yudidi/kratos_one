https://go-kratos.dev/en/docs/getting-started/start/

参考下面的
https://github.com/go-kratos/kratos-layout

```md
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs

# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## kratos proto client 命令用于 根据 proto 文件生成客户端 Go 语言代码
** 生成的代码通常包括：
 - .pb.go 文件：包含了 helloworld.proto 文件中定义的所有 message 消息结构的 Go 语言实现代码。
 - _grpc.pb.go 文件：包含了 gRPC 客户端代码，用于通过 gRPC 协议调用 helloworld.proto 中定义的服务接口。
 - _http.pb.go 文件 (如果 proto 文件中定义了 HTTP 路由)：包含了 HTTP 客户端代码，用于通过 HTTP 协议调用 helloworld.proto 中定义的服务接口。 需要注意的是，只有当你的 .proto 文件中使用了 google.api.http 选项定义了 HTTP 路由规则时，才会生成 HTTP 客户端代码。

### 区别:kratos proto client api/server/server_http.proto 和 api/server/server.proto

执行kratos proto client api/server/server_http.proto
生成go和gprc+http的客户端
openapi.yaml也会跟着变