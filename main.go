package main

import (
	"context"
	"log"
	"net"

	"github.com/dsrvlabs/vatz-plugin-matic/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type pluginServer struct {
	plugin.UnimplementedManagerPluginServer
}

func (s *pluginServer) Init(context.Context, *emptypb.Empty) (*plugin.PluginInfo, error) {
	return nil, nil
}

func (s *pluginServer) Verify(context.Context, *emptypb.Empty) (*plugin.VerifyInfo, error) {
	return nil, nil
}

func (s *pluginServer) Execute(context.Context, *plugin.ExecuteRequest) (*plugin.ExecuteResponse, error) {
	return nil, nil
}

func main() {
	startServer()
}

func startServer() {
	c, err := net.Listen("tcp", "0.0.0.0:9091")
	if err != nil {
		log.Println(err)
	}

	s := grpc.NewServer()

	serv := pluginServer{}
	plugin.RegisterManagerPluginServer(s, &serv)

	reflection.Register(s)

	if err := s.Serve(c); err != nil {
		log.Panic(err)
	}
}
