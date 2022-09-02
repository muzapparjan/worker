package worker

import (
	context "context"
	"log"
)

type Server struct{}

func (server *Server) Compute(ctx context.Context, request *ComputeRequest) (*ComputeResponse, error) {
	log.Printf("Worker-Compute: start computing request:%v", request)
	result := &ComputeResponse{
		Result: []byte("Sorry this worker is stupid now."),
	}
	log.Println("Worker-Compute: finished computing.")
	//log.Printf("Worker-Compute: finished computing.\nrequest:\n\t%v\nResult:\t%v", request, result)
	return result, nil
}

func (server *Server) Run(ctx context.Context, request *RunRequest) (*RunResponse, error) {
	log.Printf("Worker-Run: start running request:%v", request)
	result := &RunResponse{
		Result: []byte("Sorry this worker is stupid now."),
	}
	log.Println("Worker-Run: finished running.")
	//log.Printf("Worker-Run: finished running.\nrequest:\n\t%v\nResult:\t%v", request, result)
	return result, nil
}

func (server *Server) computeDefault(ctx context.Context, request *ComputeRequest) (*ComputeResponse, error) {
	//TODO
	return nil, nil
}
