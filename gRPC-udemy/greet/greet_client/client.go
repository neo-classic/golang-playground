package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/gRPC-udemy/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello from client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

}
