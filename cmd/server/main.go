// 编写cmd/server/main.go以设置和启动Thrift服务端。

package main

import (
  "your_project/pkg/thrift/example"
  "your_project/service"
  "github.com/apache/thrift/lib/go/thrift"
  "log"
  "net"
)

func main() {
  addr := "localhost:9090"
  transportFactory := thrift.NewTBufferedTransportFactory(8192)
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

  // 创建一个监听器
  listener, err := net.Listen("tcp", addr)
  if err != nil {
    log.Fatalf("Error creating listener: %v", err)
  }

  // 创建处理器
  handler := &service.UserServiceImpl{}
  processor := example.NewUserServiceProcessor(handler)

  // 启动Thrift服务
  server := thrift.NewTSimpleServer4(processor, listener, transportFactory, protocolFactory)
  log.Printf("Starting Thrift server on %s", addr)
  if err := server.Serve(); err != nil {
    log.Fatalf("Error starting server: %v", err)
  }
}
