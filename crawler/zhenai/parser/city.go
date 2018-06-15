package parser

import (
	"regexp"

	"../../engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches { // every time, m is the same variable
		// every time, username is a new variable
		username := string(m[2])
		result.Items = append(result.Items, "User "+username)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				// fmt.Printf("Value: %d Value-Addr: %X\n", username, &username)
				// fmt.Printf("Value: %d Value-Addr: %X\n", m, &m)
				return ParseProfile(contents, username)
			},
		})
	}

	return result
}
