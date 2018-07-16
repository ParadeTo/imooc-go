package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const URL = "http://www.shougolf.com/"

var cityNameRe = regexp.MustCompile(`<h3><a [^<]+</a>([^<]+)</h3>`)
var clubRe = regexp.MustCompile(`<a href="/club/(\d+)" [^<]+</a>`)
var clubNameRe = regexp.MustCompile(`<div class="tit">([\s\S]+)</div>[\s]*<ul>`)
var clubPhoneRe = regexp.MustCompile(`<li><span>联系电话</span>([^<]+)</li>`)
var clubAreaRe = regexp.MustCompile(`<li><span>球场面积</span>([^<]+)</li>`)
var clubBriefRe = regexp.MustCompile(`<div id="profile">([\S]+)</div>`)
var clubAddressRe = regexp.MustCompile(`<li style="width: 850px;"><span>联系地址</span>([^<]+)</li>`)
var blankRe = regexp.MustCompile("\\s+")

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Panicf("Fetcher error: %s", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func writer(filename string, dataChan chan string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	file.WriteString("球会名,联系方式,面积,地址,介绍\n")

	for {
		data, ok := <-dataChan
		if !ok {
			fmt.Println("writer return")
			return
		}
		file.WriteString(data)
	}
}

func worker(cityId int, wg *sync.WaitGroup) {
	cityUrl := URL + "clubmore?cityid=" + fmt.Sprintf("%d", cityId)
	content, _ := fetch(cityUrl)

	// 提取城市名
	matches := cityNameRe.FindAllStringSubmatch(string(content), -1)
	var cityName string
	if len(matches) == 1 && len(matches[0]) == 2 {
		cityName = (matches[0][1])
	}

	c := make(chan string)
	go writer("./city/"+cityName+".csv", c)

	// 所有的俱乐部
	var wgClub sync.WaitGroup
	matches = clubRe.FindAllStringSubmatch(string(content), -1)
	wgClub.Add(len(matches))
	for _, m := range matches {
		go fetchClub(m[1], c, &wgClub)
	}
	wgClub.Wait()
	// fmt.Println(clubId)
	wg.Done()
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return blankRe.ReplaceAllString(strings.TrimSpace(src), "")
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

func fetchClub(clubId string, cityFileChan chan string, wg *sync.WaitGroup) {
	clubUrl := URL + "club/" + clubId
	content, _ := fetch(clubUrl)

	clubName := trimHtml(extractString(content, clubNameRe))
	// clubBrief := trimHtml(extractString(content, clubBriefRe))
	clubPhone := extractString(content, clubPhoneRe)
	clubArea := extractString(content, clubAreaRe)
	clubAddress := extractString(content, clubAddressRe)

	// brief
	var clubBrief string
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	dom.Find("#profile").Each(func(i int, selection *goquery.Selection) {
		clubBrief = blankRe.ReplaceAllString(selection.Text(), "")
	})

	// game
	// var clubGame string
	// dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	// dom.Find(".mathtimelineh1").Each(func(i int, selection *goquery.Selection) {
	// 	clubGame = blankRe.ReplaceAllString(selection.Text(), "")
	// })
	// fmt.Println(clubGame)

	cityFileChan <- clubName + "," + clubPhone + "," + clubArea + "," + clubAddress + "," + clubBrief + "\n"
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(30)

	for i := 1; i <= 30; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	time.Sleep(time.Second * 5)
}
