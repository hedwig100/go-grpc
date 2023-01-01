package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptr 1: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] my unary server interceptor 1: ", res)
	return res, err
}

func myUnaryServerInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptr 2: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] my unary server interceptor 2: ", res)
	return res, err
}
