package types

type Tournament struct {
	Divisions []Division
	Ages      []AgeGroup
	Host      Location
}

type Division int
type AgeGroup int

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
