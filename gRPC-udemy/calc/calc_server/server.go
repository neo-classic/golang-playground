package main

import (
	"context"
	"fmt"
	"github.com/neo-classic/golang-playground/gRPC-udemy/calc/calcpb"
	"google.golang.org/grpc"
	"log"
	"net"
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
