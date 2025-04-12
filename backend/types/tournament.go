package types

import (
	"time"
)

type Tournament struct {
	Id             int
	Name           string
	LocationName   string
	StreetAddress  string
	City           string
	State          string
	OrganizerId    int
	IsFull         bool
	StartDate      time.Time
	EndDate        time.Time
	IsBoysVarsity  bool
	IsGirlsVarsity bool
	IsBoysJv       bool
	IsGirlsJv      bool
	IsBoysMs       bool
	IsGirlsMs      bool
	IsBoysYouth    bool
	IsGirlsYouth   bool
}
