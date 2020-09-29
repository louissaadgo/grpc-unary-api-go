package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"projects/grpcunaryapi/sumpb/sumpb"
	"time"
)

const port = "localhost:50051"

func main(){
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := sumpb.NewSumServiceClient(conn)
	res, err := c.Sum(ctx, &sumpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Something Went Wrong: %v", err)
	}
	log.Printf("The sum is: %v", res.GetSum())
}