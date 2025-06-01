package util

import (
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

	numPlaces := 0

	for ; in > 0; in /= 36 {
		place := in % 36
		if place < 10 {
			reversedResult.WriteString(strconv.Itoa(place))
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

func sr(val int) string {
	return string(rune(runeDifference + val))
}
