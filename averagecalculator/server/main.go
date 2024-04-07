package main

import (
	"log"
	"net"

	pb "github.com/AdekunleDally/grpc-client-streaming-api/averagecalculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.AverageServiceServer
}

var addr string = "0.0.0.0:8082"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen on : %v\n", err)
	}
	log.Printf("Listening on :%s\n", addr)
	s := grpc.NewServer()
	pb.RegisterAverageServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v\n", err)
	}
}
