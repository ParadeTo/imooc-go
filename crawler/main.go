package main

import (
	"./engine"
	"./persist"
	"./scheduler"
	"./zhenai/parser"
)

func main() {
	// simple
	// e := engine.SimpleEngine{}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// concurrent
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueueScheduler{},
		WorkerCount:    100,
		ItemChan:       itemChan,
		RequestProcess: engine.Worker,
		Deduplicate:    engine.NewSimpleDeDuplicate(),
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
