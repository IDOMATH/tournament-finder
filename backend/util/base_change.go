package util

import (
	"fmt"
	"strconv"
)

func tenToThirtysix(in int) string {
	if in < 10 {
		return strconv.Itoa(in)
	}
	_ = in % 100

	return ""
}

func Example() {
	fmt.Println(int('a')) // 97
	//TODO: get this 97 to be an 'a' again
	fmt.Println(rune(10 + 87))
}
