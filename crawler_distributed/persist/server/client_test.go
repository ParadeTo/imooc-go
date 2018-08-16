package main

import (
	"testing"
	"time"

	"log"

	"../../../crawler/engine"
	"../../../crawler/model"
	"../../config"
	"../../rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/106518746",
		Type: "zhenai",
		Id:   "106518746",
		Payload: model.Profile{
			Name:       "花开",
			Gender:     "男",
			Age:        23,
			Height:     168,
			Weight:     61,
			Income:     "5001-8000元",
			Marriage:   "未婚",
			Education:  "高中及以下",
			Occupation: "保安人员",
			Hokou:      "浙江台州",
			Xinzuo:     "魔羯座",
			House:      "打算婚后购房",
			Car:        "未购车",
		},
	}
	// start ItemSaverServer
	go startRpc(host, "test1")
	time.Sleep(time.Second)

	// start client
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// call save
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	log.Println("--------------")
	log.Printf("%v", err)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
