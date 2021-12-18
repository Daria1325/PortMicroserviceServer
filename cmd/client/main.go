package main

//
//import (
//	"context"
//	"flag"
//	"github.com/BurntSushi/toml"
//	api "github.com/daria/PortMicroservice/api/proto"
//	cnfg "github.com/daria/PortMicroservice/data/config"
//	"google.golang.org/grpc"
//	"io/ioutil"
//	"log"
//	"os"
//	"strconv"
//	"time"
//)
//
//var (
//	list    = flag.Bool("list", false, "To print all the ports")
//	upsert  = flag.Bool("upsert", false, "To upsert the database")
//	getPort = flag.Int("getPort", -1, "Select the Id of the port you need information about")
//	configPath ="configs/dataConfig.toml"
//
//)
//
//func printPorts(client api.PortClient, req *api.GetPortsRequest) {
//	log.Print("Looking for ports")
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	resp, err := client.GetPorts(ctx, req)
//	if err != nil {
//		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
//	}
//	log.Println(resp.GetList())
//}
//func upsertPorts(client api.PortClient, req *api.UpsertPortsRequest) {
//	log.Print("Updating ports")
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	resp, err := client.UpsertPorts(ctx, req)
//	if err != nil {
//		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
//	}
//	log.Println(resp.GetList())
//}
//
//func printPort(client api.PortClient, req *api.GetPortRequest) {
//	log.Print("Looking for ports")
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	resp, err := client.GetPort(ctx, req)
//	if err != nil {
//		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
//	}
//	log.Println(resp.GetItem())
//}
//
//func readJson(path string) ([]byte, error) {
//	jsonFile, err := os.Open(path)
//	if err != nil {
//		log.Print(err.Error())
//		return nil, err
//	}
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	defer jsonFile.Close()
//	return byteValue, nil
//}
//func config() *cnfg.Config {
//	config := cnfg.NewConfig()
//	_, err := toml.DecodeFile(configPath, config)
//	if err != nil {
//		log.Fatal(err)
//	}
//return config
//}
//
//func main() {
//	config:=config()
//	flag.Parse()
//
//	conn, err := grpc.Dial(config.BindAddr, grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	c := api.NewPortClient(conn)
//
//	if *list {
//		printPorts(c, &api.GetPortsRequest{Name: "1"})
//	}
//	if *upsert {
//		jsonData, err := readJson(config.JsonPath)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		upsertPorts(c, &api.UpsertPortsRequest{Name: string(jsonData)})
//	}
//	if *getPort != -1 {
//		printPort(c,&api.GetPortRequest{Id: strconv.Itoa(*getPort)})
//	}
//
//}
