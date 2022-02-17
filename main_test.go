package main

import (
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/dsrvlabs/vatz-plugin-matic/plugin"
)

func TestGrpc(t *testing.T) {
	c, err := net.Listen("tcp", "0.0.0.0:9091")
	if err != nil {
		log.Println(err)
	}

	s := grpc.NewServer()

	serv := pluginServer{}
	plugin.RegisterManagerPluginServer(s, &serv)

	if err := s.Serve(c); err != nil {
		log.Panic(err)
	}
}
