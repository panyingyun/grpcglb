# grpcglb
grpcglb, a example for learn grpc glb structure

ReadMe form https://github.com/wwcd/grpc-lb

# 说明

[gRPC服务发现&负载均衡](https://segmentfault.com/a/1190000008672912)中的例子, 修订如下问题

- register中重复PUT, watch时没有释放导致的内存泄漏
- 退出时不能正常unregister

## 启动测试程序

    # 分别启动服务端
    go run cmd/svr/svr.go - port 50001
    go run cmd/svr/svr.go - port 50002
    go run cmd/svr/svr.go - port 50003

    # 启动客户端
    go run cmd/cli/cli.go
	