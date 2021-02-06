package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	// fmt.Println(s[2:], s[3:])
	// fmt.Println(remove(s, 2))
	fmt.Println(pointerToArrayReverse(&s))
}

func remove(slice []int, i int) []int {
	// fmt.Println(slice)
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice) -1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func pointerToArrayReverse(ps *[]int) []int {
	s := *ps
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}



