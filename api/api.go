package api

import (
	"fmt"
	pb "go_grpc/proto"
	"net"
	"time"

	"go_grpc/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type go_grpcAPI struct {
	pb.UnimplementedGoGrpcServiceServer
	address       string
	app           app.AppSvc
	tracerEnabled bool
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 10 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

// initialized API
func (f *go_grpcAPI) Start() {
	listener, err := net.Listen("tcp", f.address)
	if err != nil {
		fmt.Printf("failed to listen address: %v", err.Error())
	}
	server := grpc.NewServer()
	if f.tracerEnabled {
		server = grpc.NewServer(
			grpc.KeepaliveEnforcementPolicy(kaep),
			grpc.KeepaliveParams(kasp),
		)
	}
	pb.RegisterGoGrpcServiceServer(server, f)

	fmt.Printf("go_grpc proto API start listening request %v\n", f.address)
	if err := server.Serve(listener); err != nil {
		fmt.Printf("failed to server proto API: %v", err.Error())

	}
}

func InitApi(host string, port int, app app.AppSvc, tracerEnabled bool) *go_grpcAPI {
	address := fmt.Sprintf("%s:%d", host, port)
	return &go_grpcAPI{
		address:       address,
		app:           app,
		tracerEnabled: tracerEnabled,
	}
}
