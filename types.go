package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Firstname string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID          int       `json:"ID"`
	Firstname   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Number      int64     `json:"number"`
	Balance     int64     `json:"balance"`
	Createdtime time.Time `json:"CreatedTime"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Firstname:   firstName,
		LastName:    lastName,
		Number:      int64(rand.Intn(1000000)),
		Createdtime: time.Now().UTC(),
	}
}
