package worker

import context "context"

type Server struct{}

func (server *Server) Compute(context.Context, *ComputeRequest) (*ComputeResponse, error) {
	return &ComputeResponse{
		Result: "Sorry this worker is stupid now.",
	}, nil
}

func (server *Server) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return &RunResponse{
		Result: "Sorry this worker is stupid now.",
	}, nil
}
