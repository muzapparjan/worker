package worker

import (
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	maxMessageSize   = int(1 << 20)
	keepAliveTime    = 30 * time.Second
	keepAliveTimeout = 10 * time.Second
)

type Worker struct {
	computeTasks []*Task
	genericTasks []*Task
	server       *grpc.Server
}

func NewWorker() (*Worker, error) {
	keepAliveParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    keepAliveTime,
		Timeout: keepAliveTimeout,
	})
	enforcementPolicy := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime: keepAliveTime,
	})
	maxSendMsgSize := grpc.MaxSendMsgSize(maxMessageSize)
	maxRecvMsgSize := grpc.MaxRecvMsgSize(maxMessageSize)
	server := grpc.NewServer(keepAliveParams, enforcementPolicy, maxSendMsgSize, maxRecvMsgSize)
	RegisterWorkerServer(server, &Server{})
	worker := &Worker{
		computeTasks: []*Task{},
		genericTasks: []*Task{},
		server:       server,
	}
	return worker, nil
}

func (worker *Worker) Work() error {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		return err
	}
	return worker.server.Serve(listener)
}
