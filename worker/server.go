package worker

import (
	context "context"
	"log"
)

type Server struct{}

func (server *Server) Compute(ctx context.Context, request *ComputeRequest) (*ComputeResponse, error) {
	log.Printf("Worker-Compute: start computing request:%v", request)
	result := &ComputeResponse{
		Result: "Sorry this worker is stupid now.",
	}
	log.Printf("Worker-Compute: finished computing request:%vResult:%v", request, result)
	return result, nil
}

func (server *Server) Run(ctx context.Context, request *RunRequest) (*RunResponse, error) {
	log.Printf("Worker-Run: start running request:%v", request)
	result := &RunResponse{
		Result: "Sorry this worker is stupid now.",
	}
	log.Printf("Worker-Run: finished running request:%vResult:%v", request, result)
	return result, nil
}
