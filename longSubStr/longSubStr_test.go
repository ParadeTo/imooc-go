package longSubStr

import "testing"

func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥发发"

	for i := 0; i < 13; i++ {
		s = s + s
	}

	ans := 7

	b.Logf("len(s) = %d", len(s))

	for i := 0; i < b.N; i++ {
		_, actual := longestSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; " +
				"expected %d",
					actual, s, ans)
		}
	}
}