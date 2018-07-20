package parser

import (
	"io/ioutil"
	"testing"

	"../../model"
	"../../engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,  "http://album.zhenai.com/u/106518746", "好人")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 elements, but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url: "http://album.zhenai.com/u/106518746",
		Type: "zhenai",
		Id: "106518746",
		Payload: model.Profile{
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
		},
	}

	if actual != expected {
		t.Errorf("expected %v, but was %v", expected, actual)
	}
}
