package grpc_core

import (
	"fmt"
	"net"

	"github.com/micro-it-freelance/core/config"
	"google.golang.org/grpc"
)

func Serve(server *grpc.Server, listener net.Listener) {
	fmt.Printf("[%s] Listen on :%d\n", config.APP.Name, config.GRPC.Port)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
