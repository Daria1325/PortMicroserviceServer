package server

import (
	"context"
	"encoding/json"
	"fmt"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/pingcap/errors"
	"log"
	"strconv"
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
		if strconv.Itoa(item.ID) == req.Id {
			w, err := json.Marshal(item)
			if err != nil {
				return nil, err
			}
			w = []byte(fmt.Sprintf("[%s]", w))
			return &api.GetPortResponse{Item: string(w)}, nil
		}
	}
	return &api.GetPortResponse{Item: "Not found, check the Id"}, errors.New("Not found")
}
func (d *GRPCServer) UpsertPorts(ctx context.Context, req *api.UpsertPortsRequest) (*api.UpsertPortsResponse, error) {
	d.ports = d.Repo.GetPorts()

	var portArray []database.Port

	err := json.Unmarshal([]byte(req.Name), &portArray)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, port := range portArray {
		isNotInDatabase := true
		if len(d.ports) != 0 {
			for _, item := range d.ports {
				if item.ID == port.ID {
					isNotInDatabase = false
					d.Repo.UpdatePort(port)
					continue
				}
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
