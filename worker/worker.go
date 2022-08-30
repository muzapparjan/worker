package worker

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	maxMessageSize   = int(1 << 20)
	keepAliveTime    = 30 * time.Second
	keepAliveTimeout = 10 * time.Second
	address          = ":12345"
)

type Worker struct {
	server *grpc.Server
	api    map[string]IAPI
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
	api := InitAPI()
	worker := &Worker{
		server: server,
		api:    api,
	}
	return worker, nil
}

func (worker *Worker) Work() error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Printf("Start listening address %v", listener.Addr().String())
	return worker.server.Serve(listener)
}
