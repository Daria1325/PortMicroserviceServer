package main

import (
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/daria/PortMicroservice/cmd/server"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	ConfigPath = "configs/dataConfig.toml"
)

func main() {
	config, err := cnfg.NewConfigPath(ConfigPath)
	grpcServer := grpc.NewServer()
	portService := server.GRPCServer{}
	portService.Repo = database.Init()
	defer portService.Repo.Close()

	api.RegisterPortServer(grpcServer, &portService)

	lis, err := net.Listen("tcp", config.BindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
