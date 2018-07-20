package persist

import (
	"testing"
	"../model"
	"../engine"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/106518746",
		Type: "zhenai",
		Id: "106518746",
		Payload: model.Profile{
			Name: "花开",
			Gender: "男",
			Age: 23,
			Height: 168,
			Weight: 61,
			Income: "5001-8000元",
			Marriage: "未婚",
			Education: "高中及以下",
			Occupation: "保安人员",
			Hokou: "浙江台州",
			Xinzuo: "魔羯座",
			House: "打算婚后购房",
			Car: "未购车",
		},
	}

	// TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s, %s", expected.Id, *resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %+v; expected %+v", actual, expected)
	}
}