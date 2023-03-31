package models

import "time"

type Product struct {
	ProductID        int
	ProductName      string
	ProductOwnerName string
	Developers       []string
	ScrumMasterName  string
	StartDate        time.Time
	Methodology      string
}
