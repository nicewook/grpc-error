package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	sqrtpb "github.com/nicewook/grpc-error/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type sqrtServer struct{}

func (*sqrtServer) Sqrt(ctx context.Context, req *sqrtpb.SqrtReq) (*sqrtpb.SqrtRes, error) {

	in := req.GetSqrtInput()
	fmt.Printf("-- received input value from the client: %d\n", in)
	if in > 50 {
		return nil, status.Errorf(codes.OutOfRange, "input %d: input should be smaller than 50", in)
	} else if in < 0 {
		return nil, status.Error(codes.InvalidArgument, "input should not be negative number")
	}
	result := float32(math.Sqrt(float64(in)))
	res := &sqrtpb.SqrtRes{SqrtResult: result}
	return res, nil
}

func serve() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("fail to listen tcp: %v", err)
	}

	s := grpc.NewServer()
	sqrtpb.RegisterSqrtServiceServer(s, &sqrtServer{})
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("fail to serve gRPC server: %v", err)
	}
}

func main() {
	fmt.Println("gRPC Sqrt server start...")
	serve()
}
