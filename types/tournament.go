package types

import "math"

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
	for i := 7; i >= 0; i-- {
		if val >= 128 {
			arr[i] = true
			val -= 128
		}
		if val >= 64 {
			arr[i] = true
			val -= 64
		}
		if val >= 32 {
			arr[i] = true
			val -= 32
		}
		if val >= 16 {
			arr[i] = true
			val -= 16
		}
		if val >= 8 {
			arr[i] = true
			val -= 8
		}
		if val >= 4 {
			arr[i] = true
			val -= 4
		}
		if val >= 2 {
			arr[i] = true
			val -= 2
		}
		if val >= 1 {
			arr[i] = true
			val -= 1
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
