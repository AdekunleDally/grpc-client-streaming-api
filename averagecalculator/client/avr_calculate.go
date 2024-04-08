package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/AdekunleDally/grpc-client-streaming-api/averagecalculator/proto"
)

func calculateAverage(c pb.AverageServiceClient) {
	log.Println("Welcome to the Average Calculator App. Please enter the numbers you wish to find the average")

	scanner := bufio.NewScanner(os.Stdin)
	time.Sleep(3 * time.Second)

	fmt.Println("Enter a stream of Integers")
	time.Sleep(3 * time.Second)

	fmt.Println("Press 'Enter' After each integer")
	time.Sleep(3 * time.Second)

	fmt.Println("Then type 'done' when you are done entering all the integers")
	numbs := []*pb.NumberRequest{}

	stream, err := c.AverageCalc(context.Background())
	if err != nil {
		log.Fatalf("Error while calling AverageCalculator  %v\n", err)
	}
	for scanner.Scan() {
		input := scanner.Text()

		if input == "done" {
			break
		}

		numb, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please Enter a valid integer")
			continue
		}
		numbs = append(numbs, &pb.NumberRequest{InputNumber: numb})

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
	log.Println("Sending the numbers below to the server")

	for _, numb := range numbs {
		log.Println(numb)
		stream.Send(numb)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from longGreet %v\n", err)
	}
	log.Println("Waiting for the answer from the server")
	time.Sleep(5 * time.Second)
	log.Println("The Average is : ", res.AvrResult)

}
