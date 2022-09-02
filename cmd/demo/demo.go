package main

import (
	"context"
	"log"
	"worker/worker"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial the server:\n%v", err)
	}
	defer conn.Close()

	client := worker.NewWorkerClient(conn)
	request := &worker.ComputeRequest{
		Option: &worker.ComputeOption{
			Api:      worker.Flag_Any,
			Platform: worker.Flag_Any,
			Device:   worker.Flag_Any,
		},
		Code: nil,
		Data: nil,
	}
	ctx := context.Background()
	response, err := client.Compute(ctx, request)
	if err != nil {
		log.Fatalf("Failed to Compute:\n%v", err)
	}
	log.Printf("Compute result:\n%v", string(response.GetResult()))
}
