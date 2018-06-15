package parser

import (
	"io/ioutil"
	"testing"

	"../../model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "好人")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 elements, but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

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

	if profile != expected {
		t.Errorf("expected %v, but war %v", expected, profile)
	}
}
