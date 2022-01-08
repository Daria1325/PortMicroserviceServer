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
	Repo *database.Repo
}

func (d *GRPCServer) GetPorts(ctx context.Context, _ *api.GetPortsRequest) (*api.GetPortsResponse, error) {
	ports, err := d.Repo.GetPorts()
	if err != nil {
		return nil, err
	}
	w, err := json.Marshal(ports)
	if err != nil {
		return nil, err
	}
	return &api.GetPortsResponse{List: string(w)}, nil
}
func (d *GRPCServer) GetPort(ctx context.Context, req *api.GetPortRequest) (*api.GetPortResponse, error) {
	ports, err := d.Repo.GetPorts()
	if err != nil {
		return nil, err
	}

	for _, item := range ports {
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
	updatedId := "Id: "
	ports, err := d.Repo.GetPorts()
	if err != nil {
		return nil, err
	}

	var portArray []database.Port

	err = json.Unmarshal([]byte(req.Name), &portArray)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, port := range portArray {
		isNotInDatabase := true
		if len(ports) != 0 {
			for _, item := range ports {
				if item.ID == port.ID {
					updatedId = fmt.Sprintf("%s %d", updatedId, item.ID)
					isNotInDatabase = false
					err := d.Repo.UpdatePort(port)
					if err != nil {
						return nil, err
					}
					continue
				}
			}
		}

		if isNotInDatabase {
			err := d.Repo.AddPort(port)
			if err != nil {
				return nil, err
			}
		}
	}

	return &api.UpsertPortsResponse{List: updatedId}, nil
}
