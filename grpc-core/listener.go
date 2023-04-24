package grpc_core

import (
	"fmt"
	"net"

	"github.com/micro-it-freelance/core/config"
)

func NewGRPCListener() net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		panic(err)
	}
	return listener
}
