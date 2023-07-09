package pkg

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	ddospb "github.com/cloverway/schema/pbgo/v2/ddos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// RunServer runs gRPC service to publish DDoS service
func RunGrpcServer(ctx context.Context) error {
	appService := NewDdosServer()
	port := "8080"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	ddospb.RegisterDdoSServiceServer(server, appService)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Authorization" || headerName == "Token" {
		return headerName, true
	}
	return "", false
}

// RunServer runs HTTP/REST gateway
func RunRestServer(ctx context.Context) error {
	//var grpcPort, httpPort string
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//mux := runtime.NewServeMux()
	mux := runtime.NewServeMux()
	port := "8081"
	fmt.Println("Running rest server on", port)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := ddospb.RegisterDdoSServiceHandlerFromEndpoint(ctx, mux, "localhost:"+port, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

/*
func RunRestServer(ctx context.Context) error {
	// Create a gRPC server.
	s := grpc.NewServer()

	// Register the gRPC service with the server.
	ddospb.RegisterService(s, &MyService{})

	// Create a gRPC gateway client.
	c, err := client.NewClient(
		context.Background(),
		"localhost:50051",
		&gatewayutil.DefaultClientOptions{},
	)
	if err != nil {
		panic(err)
	}

	// Register the gRPC service with the gateway client.
	gateway.RegisterService(c, &MyService{})

	// Listen on port 50051.
	err = s.Serve(context.Background())
	if err != nil {
		panic(err)
	}
}

*/
