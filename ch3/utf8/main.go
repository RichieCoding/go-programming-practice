package main

import (
	"fmt"
)

func main() {
	str := "界世"
	fmt.Println([]rune(str))

	fmt.Println(string([]rune(str)))
	// fmt.Println(len(str))
	// for i := 0; i < utf8.RuneCountInString(str); i++ {

	// }

	for _, letter := range str {
		fmt.Println(letter)
	}

	utf := 30028
	fmt.Println(string(utf))
}
