package server

import (
	"context"
	"encoding/json"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/pingcap/errors"
)

type GRPCServer struct {
	api.UnimplementedPortServer
	ports []database.Port
	Repo  *database.Repo
}

func (d *GRPCServer) GetPorts(ctx context.Context, req *api.GetPortsRequest) (*api.GetPortsResponse, error) {
	d.ports = d.Repo.GetPorts()

	w, err := json.Marshal(d.ports)
	if err != nil {
		return nil, err
	}
	return &api.GetPortsResponse{List: string(w)}, nil
}
func (d *GRPCServer) GetPort(ctx context.Context, req *api.GetPortRequest) (*api.GetPortResponse, error) {
	d.ports = d.Repo.GetPorts()

	for _, item := range d.ports {
		if item.ID == req.Id {
			w, err := json.Marshal(item)
			if err != nil {
				return nil, err
			}
			return &api.GetPortResponse{Item: string(w)}, nil
		}
	}
	return &api.GetPortResponse{Item: "Not found, check the Id"}, errors.New("Not found")
}
func (d *GRPCServer) UpsertPorts(ctx context.Context, req *api.UpsertPortsRequest) (*api.UpsertPortsResponse, error) {
	d.ports = d.Repo.GetPorts()

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
				d.Repo.UpdatePort(port)
				continue
			}
		}
		if isNotInDatabase {
			d.Repo.AddPort(port)
		}
	}
	d.ports = d.Repo.GetPorts()

	w, err := json.Marshal(d.ports)
	if err != nil {
		return nil, err
	}

	return &api.UpsertPortsResponse{List: string(w)}, nil
}
