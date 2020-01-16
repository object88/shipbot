package serve

import (
	"log"
	"net"

	"github.com/object88/shipbot/proto"
	"google.golang.org/grpc"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) ListClusters(req *proto.ListClustersRequest, stream proto.Shipbot_ListClustersServer) error {
	return nil
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterShipbotServer(grpcServer, s)
	grpcServer.Serve(lis)
}
