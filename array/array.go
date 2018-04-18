package main

import "fmt"

func changeArr(arr *[5]int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	// s2 = [5, 6, *7]
	s3 := append(s2, 10) // capacity of s2 is 3 and length is 2, so can append item
	fmt.Println(s3, arr)
	s4 := append(s3, 11) // capacity and length of s3 is 3, can not append item, to append must create a new array
	fmt.Println(s4, cap(s4), arr) // now the capacity of s4 is 6 ( 2 times of s3)
	s5 := append(s4, 12)
	fmt.Println(s5, cap(s5), arr)

	var s []int
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	s6 := make([]int, 2, 4)
	s6[0] = 1
	fmt.Println(len(s6), cap(s6), s6)

	var s7 []int = make([]int, 1, 4)
	copy(s7, s6)
	fmt.Println(s7)
	s8 := append(s7, 2)
	fmt.Println(s6, s7, s8)

	s2 = append(s2[:3], s2[4:]...)
}
