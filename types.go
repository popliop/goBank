package main

import "math/rand"

type Account struct {
	ID        int    `json:"ID"`
	Firstname string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		Firstname: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
	}
}
