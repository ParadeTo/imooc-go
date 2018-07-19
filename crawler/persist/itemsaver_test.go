package persist

import (
	"testing"
	"../model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "好人",
		Gender:     "男",
		Age:        36,
		Height:     178,
		Weight:     84,
		Income:     "3000元以下",
		Marriage:   "未婚",
		Education:  "高中及以下",
		Occupation: "--",
		Hokou:      "广东深圳",
		Xinzuo:     "--",
		House:      "已购房",
		Car:        "未购车",
	}

	id, err := save(expected)

	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s, %s", id, *resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}