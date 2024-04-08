package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/AdekunleDally/grpc-client-streaming-api/averagecalculator/proto"
)

func (s *Server) AverageCalc(stream pb.AverageService_AverageCalcServer) error {
	var total_value, num_requests float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if num_requests == 0 {
				return errors.New("no new numbers received")
			}
			average_value := float64(total_value) / float64(num_requests)
			return stream.SendAndClose(&pb.AverageResponse{
				AvrResult: average_value,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		num_requests++
		total_value += req.InputNumber
	}
}
