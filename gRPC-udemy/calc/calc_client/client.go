package main

import (
	"context"
	"fmt"
	"github.com/neo-classic/golang-playground/gRPC-udemy/calc/calcpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Calc client in running...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect to Calc server: %v\n", err)
	}
	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)
	//callCalc(c)

	callPrimeNumber(c)
}

func callCalc(c calcpb.CalcServiceClient) {
	var num1 float64 = 10
	var num2 float64 = 2
	var operation string = "-"

	req := &calcpb.CalcRequest{
		Calc: &calcpb.Calc{
			Num1:      num1,
			Num2:      num2,
			Operation: operation,
		},
	}

	res, err := c.DoCalc(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while connecting Calc gRPC: %v\n", err)
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operation, num2, res.Result)
}

func callPrimeNumber(c calcpb.CalcServiceClient) {
	var num int64 = 120

	req := &calcpb.PrimeNumberRequest{
		Num: &calcpb.Number{
			Num: num,
		},
	}

	resStream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while connecting to PrimeNumber: %v", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumbers: %v", res.GetResult())
	}
}
