package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is cccc@gmail.com
salkdgj is askjg@dasg.com
askgjl isadg s    kk@sa.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
