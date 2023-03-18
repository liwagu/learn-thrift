// 编写cmd/client/main.go以设置和使用Thrift客户端。
package main

import (
  "learn-thrift/pkg/thrift/example"
  "github.com/apache/thrift/lib/go/thrift"
  "log"
  "context"
)

func main() {
  addr := "localhost:9090"
  transportFactory := thrift.NewTBufferedTransportFactory(8192)
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

  // 创建一个连接
  transport, err := thrift.NewTSocket(addr)
  if err != nil{
    log.Fatalf("Error creating socket: %v", err)
}

// 使用TransportFactory和ProtocolFactory包装连接
transport, err = transportFactory.GetTransport(transport)
if err != nil {
    log.Fatalf("Error wrapping transport: %v", err)
}
defer transport.Close()

if err := transport.Open(); err != nil {
    log.Fatalf("Error opening transport: %v", err)
}

// 创建Thrift客户端
protocol := protocolFactory.GetProtocol(transport)
client := example.NewUserServiceClientProtocol(transport, protocol, protocol)

// 调用远程方法
ctx := context.Background()
userProfile, err := client.GetUserProfile(ctx, 1)
if err != nil {
    log.Fatalf("Error calling GetUserProfile: %v", err)
}

log.Printf("Received UserProfile: %+v", userProfile)
}
