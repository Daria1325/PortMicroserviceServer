package tests

import (
	"context"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/daria/PortMicroservice/cmd/server"
	"github.com/daria/PortMicroservice/data/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var portServer server.GRPCServer

func TestMain(m *testing.M) {
	portServer = server.GRPCServer{}
	portServer.Repo = database.Init()
	defer portServer.Repo.Close()
	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestGetPorts(t *testing.T) {
	_, err := portServer.GetPorts(context.Background(), &api.GetPortsRequest{Name: ""})
	assert.NoError(t, err)
}
func TestGetPort(t *testing.T) {
	_, err := portServer.GetPort(context.Background(), &api.GetPortRequest{Id: "1"})
	assert.NoError(t, err)

	resp, err := portServer.GetPort(context.Background(), &api.GetPortRequest{Id: "0"})
	assert.Error(t, err)
	assert.Equal(t, "Not found, check the Id", resp.GetItem())
}
func TestUpsertPorts(t *testing.T) {
	input := "[\n  {\n    \"id\": \"23\",\n    \"name\": \"Port 23\",\n    \"city\": \"Ismail\"\n  }\n]"
	_, err := portServer.UpsertPorts(context.Background(), &api.UpsertPortsRequest{Name: input})
	assert.NoError(t, err)
}
func TestNewConfigPath(t *testing.T) {
	input := "some.toml"
	_, err := config.NewConfigPath(input)
	assert.Error(t, err)
}
