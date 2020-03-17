package main

import (
	"github.com/rosspatil/go-kit-example/endpoint"
	"github.com/rosspatil/go-kit-example/service"
	"github.com/rosspatil/go-kit-example/transport"
)

func main() {
	s := service.NewService()
	e := endpoint.CreateEndPoint(*s)
	g := transport.NewHTTP(e)
	g.Run(":8090")
}
