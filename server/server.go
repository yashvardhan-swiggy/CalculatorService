package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/CalculatorService/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math"
	"net"
	"time"
)

var (
	port = flag.Int("port", 5051, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (*server) GetSum(ctx context.Context, request *pb.GetSumRequest) (*pb.GetSumResponse, error) {
	fmt.Println("Sum function was invoked :- Unary Streaming")
	sum := request.GetNum1() + request.GetNum2()
	return &pb.GetSumResponse{
		Sum: sum,
	}, nil
}

func (*server) GetPrimeNumbers(request *pb.GetPrimeNumbersRequest, response pb.CalculatorService_GetPrimeNumbersServer) error {
	fmt.Println("Prime Numbers Function was invoked :- Server Side Streaming")
	for i := 2; i < int(request.GetNum()); i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result := &pb.GetPrimeNumbersResponse{
				Prime: int64(i),
			}
			time.Sleep(100 * time.Millisecond)
			response.Send(result)
		}
	}
	return nil
}

func (*server) GetAverage(stream pb.CalculatorService_GetAverageServer) error {
	fmt.Println("GetAverage Function was invoked :- Client Side Streaming")
	result := 0
	count := 0
	for {
		received, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GetAverageResponse{
				Average: int64(result / count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading from client stream : %v", err)
			return err
		}
		result += int(received.GetNum())
		count++
	}
	return nil
}

func (*server) GetMaximum(stream pb.CalculatorService_GetMaximumServer) error {
	fmt.Println("GetMaximum Function was invoked :- Bi-Directional Streaming")

	slice := make([]int64, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while receiving data from client : %v", err)
			return err
		}
		slice = append(slice, req.GetNum())
		max := slice[0]
		for _, value := range slice {
			if value > max {
				max = value
			}
		}
		sendErr := stream.Send(&pb.GetMaximumNumberResponse{
			MaximumNum: max,
		})
		if sendErr != nil {
			log.Fatalf("error while sending response to client : %v", err)
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	fmt.Println("\nServer started...")
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0: %d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &server{})
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
