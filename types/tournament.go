package types

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
