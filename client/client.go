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
	fmt.Println("\nServer Invoked...For,\n\n-->GetSumRPC Input : 1\n\n-->GetPrimeNumbersRPC Input : 2\n\n-->GetAverageRPC Input : 3\n\n-->GetMaximumRPC Input : 4\n")
	var input int
	fmt.Scanln(&input, "\n")
	switch input {
	case 1:
		fmt.Println("GetSum RPC called...\n")
		GetSum(client)
	case 2:
		fmt.Println("GetPrimeNumbers RPC called...\n")
		GetPrimeNumbers(client)
	case 3:
		fmt.Println("GetAverage RPC called...\n")
		GetAverage(client)
	case 4:
		fmt.Println("GetMaximum RPC called...\n")
		GetMaximum(client)
	default:
		fmt.Println("Incorrect input")
	}
}

func GetSum(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of a Unary RPC...\n")
	request := &pb.GetSumRequest{
		Num1: 87,
		Num2: 168,
	}
	fmt.Println("Print sum of two numbers : ", request.GetNum1(), " ", request.GetNum2(), "\n")
	response, err := client.GetSum(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling the GetSum unary rpc : %v", err)
	}
	fmt.Println("Response from GetSum Call, Sum : ", response.GetSum(), "\n")
}

func GetPrimeNumbers(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of Server-Side Streaming RPC...\n")
	request := &pb.GetPrimeNumbersRequest{
		Num: 87,
	}
	fmt.Println("Print prime numbers less than ", request.GetNum(), "\n")
	responseStream, err := client.GetPrimeNumbers(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling GetPrimeNumbers server-side streaming rpc : %v", err)
	}
	fmt.Println("Response from GetPrimeNumbers Server : \n")
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
	fmt.Printf("\n\n")
}

func GetAverage(client pb.CalculatorServiceClient) {
	fmt.Println("Starting use of Client-Side Streaming RPC...\n")
	stream, err := client.GetAverage(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client-side streaming : %v", err)
	}
	fmt.Println("Return Average of following numbers :\n")
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

	for i, req := range requests {
		fmt.Println("Num -", i+1, " : ", req.GetNum())
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server : %v", err)
	}
	fmt.Printf("\n")
	fmt.Println("Response from server, Average Received : ", response.GetAverage(), "\n")
}

func GetMaximum(client pb.CalculatorServiceClient) {
	fmt.Println("Starting Bi-directional stream by calling GreetEveryone over GRPC...\n")
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
	fmt.Println("Print Maximum of stream of numbers :\n")
	stream, err := client.GetMaximum(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client side streaming : %v", err)
	}

	waitChan := make(chan struct{})

	go func(requests []*pb.GetMaximumNumberRequest) {
		for i, req := range requests {
			fmt.Println("Num -", i+1, " : ", req.GetNum(), "\n")
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending request to GetMaximumNumber service : %v", err)
			}
			time.Sleep(100 * time.Millisecond)
		}
		stream.CloseSend()
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
			fmt.Println("Maximum : ", response.GetMaximumNum(), "\n")
		}
	}()

	<-waitChan
	fmt.Println()
}
