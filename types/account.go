package types

import (
	"math/rand"
	"time"
)

// Account represents a user's bank account
type Account struct {
	ID          int       `json:"ID"`
	Firstname   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Number      int64     `json:"number"`
	Balance     int64     `json:"balance"`
	CreatedTime time.Time `json:"createdTime"`
}

// NewAccount creates a new Account instance
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Firstname:   firstName,
		LastName:    lastName,
		Number:      int64(rand.Intn(1000000)), // Generate a random account number
		Balance:     0,                         // Initialize balance to 0
		CreatedTime: time.Now().UTC(),
	}
}
