package main

import (
	"log"

	pb "github.com/AdekunleDally/grpc-client-streaming-api/averagecalculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:8082"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	c := pb.NewAverageServiceClient(conn)
}
