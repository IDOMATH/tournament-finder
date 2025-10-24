package types

import (
	"time"
)

type Tournament struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	LocationName  string    `json:"locationName"`
	StreetAddress string    `json:"streedAddress"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	OrganizerId   int       `json:"organizerId"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	BoysVarsity   int       `json:"boysVarsity"`
	GirlsVarsity  int       `json:"girlsVarsity"`
	BoysJv        int       `json:"boysJv"`
	GirlsJv       int       `json:"girlsJv"`
	BoysMs        int       `json:"boysMs"`
	GirlsMs       int       `json:"girlsMs"`
	BoysYouth     int       `json:"boysYouth"`
	GirlsYouth    int       `json:"girlsYouth"`
	IsFull        bool      `json:"isFull"`
}
