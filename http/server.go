package http

import (
	"net/http"
	"net"
)

func Run() {
	server := &http.Server{}
	server.Handler = Handler()

	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
}
