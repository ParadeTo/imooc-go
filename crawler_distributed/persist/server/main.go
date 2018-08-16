package main

import (
	"fmt"
	"log"

	"../../config"
	"../../persist"
	"../../rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal(startRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
