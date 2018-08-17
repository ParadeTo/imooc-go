package main

import (
	"fmt"

	"../../worker"
	"../../config"
	"../../rpcsupport"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
