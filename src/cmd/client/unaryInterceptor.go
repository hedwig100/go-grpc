package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func myUnaryClientInterceptor1(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("[pre] my unary client interceptor 1", method, req)
	err := invoker(ctx, method, req, res, cc, opts...)
	fmt.Println("[post] my unary client interceptor 1", res)
	return err
}
