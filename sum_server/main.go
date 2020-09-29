package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"projects/grpcunaryapi/sumpb/sumpb"
)

const port = ":50051"

type server struct {
	sumpb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, in *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	log.Printf("First number received: %v \n", in.GetFirstNumber())
	log.Printf("Second number received: %v " + "\n", in.GetSecondNumber())
	sum := in.GetSecondNumber() + in.GetFirstNumber()
	log.Printf("Calculation ended, sending the response")
	return &sumpb.SumResponse{Sum: sum}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to the port: %v", err)
	}
	log.Println("Listening to the port")
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}