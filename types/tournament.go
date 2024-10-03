package types

type Tournament struct {
	IsBoysVarsity  bool
	IsGirlsVarsity bool
	IsBoysJv       bool
	IsGirlsJv      bool
	IsBoysMs       bool
	IsGirlsMs      bool
	IsBoysYouth    bool
	IsGirlsYouth   bool
	Host           Location
	Organizer      Contact
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
