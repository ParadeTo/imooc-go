package main

import (
	"../crawler/engine"
	"../crawler/scheduler"
	"../crawler/zhenai/parser"
	"./rpcsupport"
	persisClient "./persist/client"
	workerClient "./worker/client"
	"net/rpc"
	"log"
	"flag"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "","itemsaver host")
	workerHosts = flag.String("worker_hosts","","worker hosts (comma separated)")
)

func main() {
	// simple
	// e := engine.SimpleEngine{}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// concurrent
	flag.Parse()
	itemChan, err := persisClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := workerClient.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueueScheduler{},
		WorkerCount:    100,
		ItemChan:       itemChan,
		RequestProcess: processor,
		Deduplicate:    engine.NewSimpleDeDuplicate(),
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/nanyang",
		Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
