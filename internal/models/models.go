package models

import "time"

type Paycheck struct {
	Amount float64
}

type Expense struct {
	Name   string
	Amount float64
}

type Session struct {
	Date     time.Time `json:"date"`
	Paycheck Paycheck  `json:"paycheck"`
	Expenses []Expense `json:"expenses"`
}
