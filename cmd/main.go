package main

import (
	"fmt"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/daria/PortMicroservice/cmd/server"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"google.golang.org/grpc"
	"net"
)

var (
	ConfigPath = "configs/dataConfig.toml"
)

func main() {
	config, _ := cnfg.NewConfigPath(ConfigPath)
	grpcServer := grpc.NewServer()
	portService := server.GRPCServer{}
	portService.Repo = database.Init(config)
	defer portService.Repo.Close()

	api.RegisterPortServer(grpcServer, &portService)

	lis, err := net.Listen("tcp", config.BindAddr)
	if err != nil {
		fmt.Errorf("failed to listen: %v", err)
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Errorf("failed to start gRPC server: %v", err)
		return
	}
}
