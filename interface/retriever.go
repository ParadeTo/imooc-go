package main

import (
	"time"
	"net/http"
	"net/http/httputil"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type PosterRetriever interface {
	Retriever
	Poster
}

func session(s PosterRetriever, url string) {
	s.Get(url)
	s.Post(url, map[string]string {
		"name": "ayou",
	})
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func post(poster Poster, url string) {
	poster.Post(url, map[string]string {
		"name": "mouse",
		"course": "golang",
	})
}

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
	return r.Contents
}

func (r MockRetriever) Post(url string, form map[string]string) string {
	fmt.Println("mock retriever post")
	return r.Contents
}

type RealRetriever struct {
	UserAgent string
	TimeOut time.Duration
}



func (r RealRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}

func main() {
	url := "http://www.baidu.com"
	var r MockRetriever
	r = MockRetriever{"I am content."}
	//r = RealRetriever{"FireFox", time.Second * 30}

	session(r, url)
	download(r, url)
	post(r, url)


	//fmt.Printf("%T %v", r, r)
	//
	//if realRetriever, ok := r.(RealRetriever); ok {
	//	fmt.Println(realRetriever.UserAgent)
	//} else {
	//	fmt.Println("not real retriever")
	//}
	//
	//
	//switch r.(type) {
	//case MockRetriever:
	//	fmt.Println("I am mock retriever")
	//case RealRetriever:
	//	fmt.Println("I am real retriever")
	//}
}
