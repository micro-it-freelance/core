package core_grpc

import (
	"fmt"
	"net"

	config "github.com/micro-it-freelance/core/config"
	"google.golang.org/grpc"
)

func Serve(server *grpc.Server, listener net.Listener) {
	fmt.Printf("Listen on :%d\n", config.GRPC.Port)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
