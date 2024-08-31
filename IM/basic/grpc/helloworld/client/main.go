package main

import (
	"IM/basic/grpc/helloworld/client/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//无认证
	conn, err := grpc.NewClient(":6666", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	prodClient := pb.NewProductServiceClient(conn)
	stock, err := prodClient.GetProductStock(context.Background(), &pb.ProductRequest{
		ProductId: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("查询结构:", stock)
}
