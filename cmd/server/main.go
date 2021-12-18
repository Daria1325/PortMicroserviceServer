package main

import (
	"context"
	"encoding/json"
	"github.com/BurntSushi/toml"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	configPath = "configs/dataConfig.toml"
)

type GRPCServer struct {
	api.UnimplementedPortServer
	ports []database.Port
	repo  *database.Repo
}

func (d *GRPCServer) GetPorts(ctx context.Context, req *api.GetPortsRequest) (*api.GetPortsResponse, error) {
	d.repo = database.Init()
	defer d.repo.Close()

	d.ports = d.repo.GetPorts()

	w, err := json.Marshal(d.ports)
	if err != nil {
		log.Print(err.Error())
	}
	return &api.GetPortsResponse{List: string(w)}, nil
}
func (d *GRPCServer) GetPort(ctx context.Context, req *api.GetPortRequest) (*api.GetPortResponse, error) {
	d.repo = database.Init()
	defer d.repo.Close()

	d.ports = d.repo.GetPorts()

	for _, item := range d.ports {
		if item.ID == req.Id {
			w, err := json.Marshal(item)
			if err != nil {
				log.Print(err.Error())
			}
			return &api.GetPortResponse{Item: string(w)}, nil
		}
	}
	return &api.GetPortResponse{Item: "Not found, please check the Id"}, nil
}
func (d *GRPCServer) UpsertPorts(ctx context.Context, req *api.UpsertPortsRequest) (*api.UpsertPortsResponse, error) {
	d.repo = database.Init()
	defer d.repo.Close()
	d.ports = d.repo.GetPorts()

	isNotInDatabase := true
	var portArray []database.Port

	err := json.Unmarshal([]byte(req.Name), &portArray)
	if err != nil {
		return nil, err
	}

	for _, port := range portArray {
		for _, item := range d.ports {
			if item.ID == port.ID {
				isNotInDatabase = false
				d.repo.UpdatePort(port)
				continue
			}
		}
		if isNotInDatabase {
			d.repo.AddPort(port)
		}
	}
	d.ports = d.repo.GetPorts()

	w, err := json.Marshal(d.ports)
	if err != nil {
		log.Print(err.Error())
	}

	return &api.UpsertPortsResponse{List: string(w)}, nil
}

func config() *cnfg.Config {
	config := cnfg.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	config := config()
	grpcServer := grpc.NewServer()
	portService := GRPCServer{}
	api.RegisterPortServer(grpcServer, &portService)

	lis, err := net.Listen("tcp", config.BindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
