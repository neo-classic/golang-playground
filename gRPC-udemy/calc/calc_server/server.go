package main

import (
	"context"
	"fmt"
	"github.com/neo-classic/golang-playground/gRPC-udemy/calc/calcpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	calcpb.UnimplementedCalcServiceServer
}

func (*server) DoCalc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	num1 := req.GetCalc().GetNum1()
	num2 := req.GetCalc().GetNum2()
	operation := req.GetCalc().GetOperation()
	var result float64

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	res := &calcpb.CalcResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumber(req *calcpb.PrimeNumberRequest, stream calcpb.CalcService_PrimeNumberServer) error {
	n := req.GetNum().GetNum()
	var k int64 = 2

	for n > 1 {
		if n%k == 0 {
			sendNumber(k, stream)
			n = n / k
		} else {
			k = k + 1
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func sendNumber(k int64, stream calcpb.CalcService_PrimeNumberServer) {
	res := &calcpb.PrimeNumberResponse{Result: k}
	stream.Send(res)
}

func main() {
	fmt.Println("Calc server is running...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
