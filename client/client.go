package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/CalculatorService/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

var (
	serverAddr = flag.String("address", "localhost:5051", "The server address in the format of host:post")
)

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorServiceClient(conn)
	GetSum(client)
	GetPrimeNumbers(client)
	GetAverage(client)
	GetMaximum(client)
}

func GetSum(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of a Unary RPC...")
	request := &pb.GetSumRequest{
		Num1: 87,
		Num2: 168,
	}

	response, err := client.GetSum(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling the GetSum unary rpc : %v", err)
	}
	fmt.Println("Response from GetSum Call : ", response.GetSum())
}

func GetPrimeNumbers(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of Server-Side Streaming RPC...")
	request := &pb.GetPrimeNumbersRequest{
		Num: 87,
	}
	responseStream, err := client.GetPrimeNumbers(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling GetPrimeNumbers server-side streaming rpc : %v", err)
	}
	fmt.Println("Response from GetPrimeNumbers Server")
	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receving server stream : %v", err)
		}
		fmt.Printf("%v ", response.GetPrime())
	}
}

func GetAverage(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of Client-Side Streaming RPC...")
	stream, err := client.GetAverage(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client-side streaming : %v", err)
	}

	requests := []*pb.GetAverageRequest{
		{
			Num: 7,
		},
		{
			Num: 37,
		},
		{
			Num: 94,
		},
		{
			Num: 753,
		},
		{
			Num: 91,
		},
		{
			Num: 71,
		},
		{
			Num: 70,
		},
		{
			Num: 48,
		},
		{
			Num: 7,
		},
	}

	for _, req := range requests {
		fmt.Println("Sending request... : ", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server : %v", err)
	}
	fmt.Println("Response from server : ", response.GetAverage())
}

func GetMaximum(client pb.CalculatorServiceClient) {
	fmt.Println("Starting Bi-directional stream by calling GreetEveryone over GRPC...")
	requests := []*pb.GetMaximumNumberRequest{
		{
			Num: 7,
		},
		{
			Num: 10,
		},
		{
			Num: 3,
		},
		{
			Num: 16,
		},
		{
			Num: 18,
		},
		{
			Num: 20,
		},
		{
			Num: 4,
		},
		{
			Num: 16,
		},
		{
			Num: 30,
		},
		{
			Num: 18,
		},
	}

	stream, err := client.GetMaximum(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client side streaming : %v", err)
	}

	waitChan := make(chan struct{})

	go func(requests []*pb.GetMaximumNumberRequest) {
		for _, req := range requests {
			fmt.Println("Sending Request... : ", req.GetNum())
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending request to GetMaximumNumber service : %v", err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(requests)

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(waitChan)
				return
			}
			if err != nil {
				log.Fatalf("error receiving response from server : %v", err)
			}
			fmt.Printf("Response From Server : %v", response.GetMaximumNum())
		}
	}()
}
