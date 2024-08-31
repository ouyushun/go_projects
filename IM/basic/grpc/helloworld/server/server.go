package main

import (
	"IM/basic/grpc/helloworld/pb"
	"context"
)

type ProductService struct {
}

func (prd *ProductService) GetProductStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		ProdStock: req.ProductId,
	}, nil

}
