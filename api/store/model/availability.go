package model

import "time"

type Availability struct {
	Id        int       `json:"id"`
	AccountId int       `json:"accountId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
