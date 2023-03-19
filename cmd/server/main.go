// 编写cmd/server/main.go以设置和启动Thrift服务端。

package main

import (
  "learn-thrift/pkg/thrift/example"
  "learn-thrift/service"
  "github.com/apache/thrift/lib/go/thrift"
  "log"
  "net"
)

type myTServerSocket struct {
  listener net.Listener
}

func (m *myTServerSocket) Listen() error {
  return nil
}

func (m *myTServerSocket) Accept() (thrift.TTransport, error) {
  conn, err := m.listener.Accept()
  if err != nil {
      return nil, err
  }
  return thrift.NewTSocketFromConnTimeout(conn, 0), nil
}

func (m *myTServerSocket) Close() error {
  return m.listener.Close()
}

func (m *myTServerSocket) Addr() net.Addr {
  return m.listener.Addr()
}

func (m *myTServerSocket) Interrupt() error {
  return nil
}

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

  tServerTransport := &myTServerSocket{listener: listener}

  // 启动Thrift服务
  server := thrift.NewTSimpleServer4(processor, tServerTransport, transportFactory, protocolFactory)
  log.Printf("Starting Thrift server on %s", addr)
  if err := server.Serve(); err != nil {
      log.Fatalf("Error starting server: %v", err)
  }
  
}
