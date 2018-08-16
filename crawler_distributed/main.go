package main

import (
	"fmt"

	"../crawler/engine"
	"../crawler/scheduler"
	"../crawler/zhenai/parser"
	"./config"
	"./persist/client"
)

func main() {
	// simple
	// e := engine.SimpleEngine{}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// concurrent
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		Deduplicate: engine.NewSimpleDeDuplicate(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
