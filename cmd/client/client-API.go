package main

import (
	"context"
	"github.com/BurntSushi/toml"
	api "github.com/daria/PortMicroservice/api/proto"
	cnfg "github.com/daria/PortMicroservice/data/config"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Path       string
	configPath = "configs/dataConfig.toml"
)

//
//type Client struct {
//	client *api.PortClient
//}
//var cl *Client
//
//
//func (cl *Client) Init() *grpc.ClientConn {
//	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	client := api.NewPortClient(conn)
//	cl.client=&client
//	return conn
//}
//func CloseCon(conn *grpc.ClientConn)  {
//	conn.Close()
//}

func InitConn() (*grpc.ClientConn, api.PortClient) {
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewPortClient(conn)

	return conn, client
}

func getPorts(w http.ResponseWriter, r *http.Request) {
	conn, client := InitConn()
	defer conn.Close()
	log.Print("Looking for ports")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetPorts(ctx, &api.GetPortsRequest{Name: "1"})
	if err != nil {
		log.Fatalf("%v.ListPorts(_) = _, %v", client, err)
	}
	log.Println(resp.GetList())
}
func getPort(w http.ResponseWriter, r *http.Request) {
	conn, client := InitConn()
	defer conn.Close()
	params := mux.Vars(r)
	log.Print("Looking for port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetPort(ctx, &api.GetPortRequest{Id: params["id"]})
	if err != nil {
		log.Fatalf("%v.Port(_) = _, %v", client, err)
	}
	log.Println(resp.GetItem())
}

func readJson(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	return byteValue, nil
}

func upsertPorts(w http.ResponseWriter, r *http.Request) {
	conn, client := InitConn()
	defer conn.Close()
	jsonData, err := readJson(Path)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Updating ports")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UpsertPorts(ctx, &api.UpsertPortsRequest{Name: string(jsonData)})
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	log.Println(resp.GetList())
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
	Path = config.JsonPath

	r := mux.NewRouter()

	r.HandleFunc("/ports", getPorts).Methods("GET")
	r.HandleFunc("/ports/{id}", getPort).Methods("GET")
	r.HandleFunc("/ports", upsertPorts).Methods("POST")

	http.ListenAndServe(":3000", r)
}
