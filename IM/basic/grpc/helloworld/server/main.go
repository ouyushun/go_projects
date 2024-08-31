package main

import (
	"IM/basic/grpc/helloworld/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

var ProDevice = &ProductService{}

func main() {
	rpcServer := grpc.NewServer()
	//指针类型的receiver 方法实现接口时，只有指针类型的对象实现了该接口。
	//ProductService = &server.ProductService{} 是指针类型，
	//值类型的对象只有（t T) 结构的方法，虽然值类型的对象也可以调用(t *T) 方法，但这实际上是Golang编译器自动转化成了&t的形式来调用方法，并不是表明值类型的对象拥有该方法。
	pb.RegisterProductServiceServer(rpcServer, ProDevice)

	//net.Listen能够监听本地端口，接收特定协议建立的连接，如果成功接收，则返回一个Listener接口
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("127.0.0.1:6666 error")
		return
	}

	err = rpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
	fmt.Println("rpc服务启动成功")
}
