package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/status"

	sqrtpb "github.com/nicewook/grpc-error/proto"

	"google.golang.org/grpc"
)

func sendReq(c sqrtpb.SqrtServiceClient, in int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := c.Sqrt(ctx, &sqrtpb.SqrtReq{SqrtInput: in})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			fmt.Printf("gRPC error Code: %v\n", status.Code())
			fmt.Printf("gRPC error Details: %v\n", status.Details())
			fmt.Printf("gRPC error Err: %v\n", status.Err())
			fmt.Printf("gRPC error Message: %v\n\n", status.Message())
		} else {
			log.Fatalf(": %v", err)
		}
		return
	}
	fmt.Printf("%d DO square root result: %v\n\n", in, res.SqrtResult)
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial to gRPC server: %v", err)
	}

	c := sqrtpb.NewSqrtServiceClient(cc)

	sendReq(c, 10)
	sendReq(c, 100)
	sendReq(c, -2)
}
