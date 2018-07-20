package parser

import (
	"regexp"

	"../../engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches { // every time, m is the same variable
		// every time, username is a new variable
		username := string(m[2])
		url := string(m[1])
		// result.Items = append(result.Items, "User "+username)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(contents []byte) engine.ParseResult {
				// fmt.Printf("Value: %d Value-Addr: %X\n", username, &username)
				// fmt.Printf("Value: %d Value-Addr: %X\n", m, &m)
				return ParseProfile(contents, url, username)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
