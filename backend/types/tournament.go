package types

import (
	"math"
	"time"

	"github.com/IDOMATH/CheetahMath/formulas"
)

type Tournament struct {
	Id              int
	Name            string
	LocationName    string
	LocationAddress string
	LocationState   string
	OrganizerId     int
	IsFull          bool
	StartDate       time.Time
	EndDate         time.Time

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
func (t *Tournament) AgeDivisionIntToArray(val int) {
	var arr [8]bool
	for i := 7; i >= 0; i-- {
		if val >= formulas.IntPow(2, i) {
			arr[i] = true
			val -= formulas.IntPow(2, i)
		}
	}
	t.AgeDivision = arr
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
