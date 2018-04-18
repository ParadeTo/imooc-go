package main

import (
	"fmt"
	"strings"
)

func longestSubStr(s string) (int, int) {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			maxStart = start
		}
		lastOccured[ch] = i
	}
	return maxStart, maxLength
}

func main() {
	//m := map[string]string {
	//	"name": "ccmouse",
	//	"course": "golang",
	//}
	//// or
	//
	//delete(m, "name")
	//fmt.Println(m["name"])

	fmt.Println(longestSubStr("中啥的dd快感卡萨见到过"))

	s := "abc萨克国际"
	//
	//for _, b := range []byte(s) {
	//	fmt.Printf("%X ", b)
	//}
	//fmt.Println()
	//// 61 62 63 E8 90 A8 E5 85 8B E5 9B BD E9 99 85
	//
	//for i, ch := range s { // ch is a rune
	//	fmt.Printf("(%d %x)", i, ch)
	//}
	//fmt.Println()
	//// (0 61)(1 62)(2 63)(3 8428)(6 514b)(9 56fd)(12 9645)
	//
	//fmt.Println("Rune count:", utf8.RuneCountInString(s))
	//// Rune count: 7
	//
	//bytes := []byte(s)
	//for len(bytes) > 0 {
	//	ch, size := utf8.DecodeRune(bytes)
	//	bytes = bytes[size:]
	//	fmt.Printf("%c ", ch)
	//}
	//fmt.Println()
	//// a b c 萨 克 国 际
	//
	//for i, ch := range []rune(s) {
	//	fmt.Printf("(%d %c %X)", i, ch, ch)
	//}
	//fmt.Println()
	// (0 a 61)(1 b 62)(2 c 63)(3 萨 8428)(4 克 514B)(5 国 56FD)(6 际 9645)

	fmt.Println(strings.Index(s, "克"))
	//strings.Fields()
}
