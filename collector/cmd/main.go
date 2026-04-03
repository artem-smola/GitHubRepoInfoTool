package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/artem-smola/GitHubRepoInfoTool/collector/internal/adapter"
	"github.com/artem-smola/GitHubRepoInfoTool/collector/internal/controller"
	"github.com/artem-smola/GitHubRepoInfoTool/collector/internal/usecase"
	"github.com/artem-smola/GitHubRepoInfoTool/proto/gen"
	"google.golang.org/grpc"
)

const grpcServerPort = ":50051"

func main() {
	collectorAdapter := adapter.NewHTTPClient(&http.Client{Timeout: 10 * time.Second})
	collectorUsecase := usecase.NewRepoInfoProvider(collectorAdapter)
	collectorController := controller.NewGRPCServer(collectorUsecase)

	listener, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcServerPort, err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterCollectorServer(grpcServer, collectorController)

	log.Printf("collector gRPC server is listening on %s", grpcServerPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
