package types

// CreateAccountRequest represents a request to create a new account
type CreateAccountRequest struct {
	Firstname string `json:"firstName"`
	LastName  string `json:"lastName"`
}
