package main

import (
	"fmt"

	"../crawler/engine"
	"../crawler/scheduler"
	"../crawler/zhenai/parser"
	"./config"
	persisClient "./persist/client"
	workerClient "./worker/client"
)

func main() {
	// simple
	// e := engine.SimpleEngine{}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// concurrent
	itemChan, err := persisClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := workerClient.CreateProcessor()
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
