package main

import (
	"log"
	"net"
	"sync"

	"github.com/rosspatil/go-kit-example/endpoint"
	"github.com/rosspatil/go-kit-example/pb"
	"github.com/rosspatil/go-kit-example/service"
	"github.com/rosspatil/go-kit-example/transport"
	"google.golang.org/grpc"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	s := service.NewService()
	e := endpoint.CreateEndPoint(*s)
	go func() {
		g := transport.NewHTTP(e)
		g.Run(":8090")
		wg.Done()
	}()
	go func() {

		g1 := transport.NewGRPC(e)
		listener, err := net.Listen("tcp", ":8091")
		if err != nil {
			log.Fatal(err)
		}
		server := grpc.NewServer()
		pb.RegisterServiceServer(server, g1)
		err = server.Serve(listener)
		if err != nil {
			log.Fatal(err)

		}
		wg.Done()
	}()
	wg.Wait()
}
