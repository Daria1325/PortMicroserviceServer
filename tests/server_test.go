package tests

import (
	"context"
	api "github.com/daria/PortMicroservice/api/proto"
	"github.com/daria/PortMicroservice/cmd/database"
	"github.com/daria/PortMicroservice/cmd/server"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var portServer server.GRPCServer

func TestMain(m *testing.M) {
	config, _ := cnfg.NewConfigPath("../configs/dataConfig.toml")
	portServer = server.GRPCServer{}
	portServer.Repo = database.Init(config)
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

	resp, err := portServer.GetPort(context.Background(), &api.GetPortRequest{Id: "100"})
	assert.Error(t, err)
	assert.Equal(t, "Not found, check the Id", resp.GetItem())
}
func TestUpsertPorts(t *testing.T) {
	input := "[{\n    \"id\": 0,\n    \"name\": \"consequat id officia\",\n    \"isActive\": true,\n    \"company\": \"GEOFARM\",\n    \"email\": \"maymclean@geofarm.com\",\n    \"phone\": \"+1 (997) 434-3843\",\n    \"address\": \"907 National Drive, Foscoe, Oregon, 5061\",\n    \"about\": \"Id laborum labore irure nisi mollit. Exercitation dolor ad nisi veniam tempor laboris Lorem nisi incididunt do reprehenderit veniam dolor consequat. Mollit deserunt occaecat tempor fugiat consequat culpa eu eu deserunt minim qui. Dolore magna ipsum nisi est occaecat deserunt aliquip laboris ex cillum veniam do minim.\\r\\n\",\n    \"registered\": \"2021-02-14T01:46:57 -02:00\",\n    \"latitude\": 70.822864,\n    \"longitude\": 156.088083\n  }]"
	_, err := portServer.UpsertPorts(context.Background(), &api.UpsertPortsRequest{Name: input})
	assert.NoError(t, err)
}
func TestNewConfigPath(t *testing.T) {
	input := "some.toml"
	_, err := cnfg.NewConfigPath(input)
	assert.Error(t, err)
}
