package types

import (
	"math"

	"github.com/idomath/CheetahMath"
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

func (t Tournament) AgeDivisionArrayToInt() int {
	var value int = 0
	for i := range t.AgeDivision {
		if t.AgeDivision[i] {
			value += int(math.Pow(2, float64(i)))
		}
	}
	return value
}

// TODO: Change this to make it more general
func (t Tournament) AgeDivisionIntToArray(val int) [8]bool {
	var arr [8]bool
	for i := 7; i >= 0; i-- {
		if val >= CheetahMath.IntPow(2, i) {
			arr[i] = true
			val -= CheetahMath.IntPow(2, i)
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
