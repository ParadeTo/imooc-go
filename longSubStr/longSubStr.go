package longSubStr

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

//func main() {
//	fmt.Print(longestSubStr("黑化肥挥发发灰会花飞灰化肥发发"))
//}
