package main

import (
	"fmt"
	"log"

	"../../config"
	"../../persist"
	"../../rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(startRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func startRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
