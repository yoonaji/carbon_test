package models

import "time"

type WebhookTransaction struct {
	TransactionType   string    `json:"transaction_type"`
	BankAccountID     string    `json:"bank_account_id"`
	BankAccountNumber string    `json:"bank_account_number"`
	BankCode          string    `json:"bank_code"`
	Amount            uint      `json:"amount"`
	TransactionDate   time.Time `json:"transaction_date"`
	TransactionName   string    `json:"transaction_name"`
	Balance           uint      `json:"balance"`
	ProcessingDate    string    `json:"processing_date"`
}
