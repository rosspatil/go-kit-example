package main

import (
	"log"
	"net"
	"sync"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkin "github.com/openzipkin/zipkin-go"
	httpReporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/rosspatil/go-kit-example/endpoint"
	"github.com/rosspatil/go-kit-example/pb"
	"github.com/rosspatil/go-kit-example/service"
	"github.com/rosspatil/go-kit-example/transport"
	"google.golang.org/grpc"
)

var (
	zipkinHTTPEndpoint = "http://localhost:9411/api/v2/spans"
)

func init() {
	reporter := httpReporter.NewReporter(zipkinHTTPEndpoint)
	ze, err := zipkin.NewEndpoint("go-kit-example", "localhost:8090")
	if err != nil {
		log.Fatalln(err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(ze))
	if err != nil {
		log.Fatalln(err)
	}
	tracer := zipkinot.Wrap(nativeTracer)
	opentracing.InitGlobalTracer(tracer)
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	s := service.NewService()
	e := endpoint.CreateEndPoint(*s)
	go func() {
		g := transport.NewHTTP(e)
		g.Run(":8080")
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
