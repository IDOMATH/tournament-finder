package util

import (
	"strconv"
	"strings"
)

const runeDifference = 87

// TenToThirtysix takes a numner in base 10 and returns a string representation
// of that number in base 36.
func TenToThirtysix(in int64) string {
	var reversedResult strings.Builder
	if in < 10 {
		return strconv.Itoa(int(in))
	}
	if in < 36 {
		return string(rune(87 + in))
	}

	numPlaces := 0

	for ; in > 0; in /= 36 {
		place := in % 36
		if place < 10 {
			reversedResult.WriteString(strconv.Itoa(int(place)))
		} else {
			reversedResult.WriteString(sr(place))
		}
		numPlaces++
	}

	if len(reversedResult.String()) < numPlaces {
		reversedResult.WriteString("1")
	}

	var result strings.Builder

	for i := len(reversedResult.String()) - 1; i >= 0; i-- {
		result.WriteString(string(reversedResult.String()[i]))
	}

	return result.String()
}

func ThirtysixToTen(bts string) int64 {
	return 0
}

// sr takes an int value that should be between 10 and 35
// I'm not about to put checks in for that, just use it properly
func sr(val int64) string {
	return string(rune(runeDifference + val))
}
