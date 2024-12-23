package types

// TransferRequest represents a request to transfer money between accounts
type TransferRequest struct {
	ToAccountID int `json:"toAccountID"`
	Amount      int `json:"amount"`
}
