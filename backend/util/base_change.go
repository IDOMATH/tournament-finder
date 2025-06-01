package util

import (
	"fmt"
	"strconv"
	"strings"
)

const runeDifference = 87

func TenToThirtysix(in int) string {
	var reversedResult strings.Builder
	if in < 10 {
		return strconv.Itoa(in)
	}
	if in < 36 {
		return string(rune(87 + in))
	}

	for ; in > 35; in -= 36 {
		place := in % 36
		if place < 10 {
			fmt.Println("less than 10 place: ", place)
			reversedResult.WriteString(strconv.Itoa(place))
		} else {
			reversedResult.WriteString(sr(place))
		}
	}

	var result strings.Builder

	for i := len(reversedResult.String()) - 1; i >= 0; i-- {
		result.WriteString(string(reversedResult.String()[i]))
	}

	return result.String()
}

func sr(val int) string {
	return string(rune(runeDifference + val))
}

// func Example() {
// 	fmt.Println(int('a')) // 97
// 	val := 10
// 	fmt.Println(string(rune(87 + val)))
// 	fmt.Println(string(rune(10 + 87)))
// }
