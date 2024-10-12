package types

import (
	"math"

	"github.com/IDOMATH/Convert"
)

type Tournament struct {
	Name        string
	Host        Location
	Organizer   Contact
	AgeDivision [8]bool
}

const (
	IsBoysVarsity = iota
	IsGirlsVarsity
	IsBoysJv
	IsGirlsJv
	IsBoysMs
	IsGirlsMs
	IsBoysYouth
	IsGirlsYouth
)

func AgeDivisionArrayToInt(s [8]bool) int {
	var value int = 0
	for i := range s {
		if s[i] {
			value += int(math.Pow(2, float64(i)))
		}
	}
	return value
}

// TODO: Change this to make it more general
func AgeDivisionIntToArray(val int) [8]bool {
	var arr [8]bool
	Convert.IntPow()
	for i := 7; i >= 0; i-- {
		if val >= Convert.IntPow(2, i) {
			arr[i] = true
			val -= Convert.IntPow(2, i)
		}
	}
	return arr
}

type Location struct {
	Name    string
	Address Address
}

type Address struct {
	HouseNumber string
	Street      string
	City        string
	State       string
	PostalCode  int
}

type Contact struct {
	Name  string
	Email string
}
