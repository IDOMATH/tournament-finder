package util

import (
	"github.com/IDOMATH/CheetahMath/formulas"
	"strconv"
	"strings"
)

const runeDifference = 87

// TenToThirtysix takes a number in base 10 and returns a string representation
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
	var baseTen int64 = 0
	for i := 0; i < len(bts); i++ {
		char := rune(bts[len(bts)-1-i])
		if char-'0' > 9 {
			char = char - runeDifference
		} else {
			char = char - 48
		}
		baseTen = baseTen + (int64(char) * formulas.IntPow(36, int64(i)))
	}
	return baseTen
}

// sr takes an int value that should be between 10 and 35
// I'm not about to put checks in for that, just use it properly
func sr(val int64) string {
	return string(rune(runeDifference + val))
}
