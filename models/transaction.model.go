package models

type TransactionModel struct {
	TransactionID     string  `gorm:"column:transaction_id;primaryKey"` // 별도 생성한 PK
	TransactionType   string  `gorm:"column:transaction_type;type:varchar(20);not null"`
	BankAccountID     string  `gorm:"column:bank_account_id;type:varchar(100);not null"`
	BankAccountNumber string  `gorm:"column:bank_account_number;type:varchar(30);not null"`
	BankCode          string  `gorm:"column:bank_code;type:char(3);not null"`
	Amount            int     `gorm:"column:amount;not null"`
	TransactionDate   string  `gorm:"column:transaction_date;type:varchar(50);not null"` // string 대신 time.Time 추천!
	TransactionName   string  `gorm:"column:transaction_name;type:varchar(100)"`
	Category          string  `gorm:"column:category;type:varchar(50)"`
	CarbonScore       float64 `gorm:"column:carbon_score;type:float"`
	UserID            string  `gorm:"column:user_id;type:varchar(100);not null"` // 꼭 추가해야 함!
}

type CreateTransactionRequest struct {
	TransactionType   string `json:"transaction_type" binding:"required"`
	BankAccountID     string `json:"bank_account_id" binding:"required"`
	BankAccountNumber string `json:"bank_account_number" binding:"required"`
	BankCode          string `json:"bank_code" binding:"required"`
	Amount            int    `json:"amount" binding:"required"`
	TransactionDate   string `json:"transaction_date" binding:"required"` // 또는 time.Time
	TransactionName   string `json:"transaction_name" binding:"required"`
	UserID            string `json:"user_id" binding:"required"`
}

type Webhook struct {
	TransactionType   string `json:"transaction_type" binding:"required"`
	BankAccountID     string `json:"bank_account_id" binding:"required"`
	BankAccountNumber string `json:"bank_account_number" binding:"required"`
	BankCode          string `json:"bank_code" binding:"required"`
	Amount            int    `json:"amount" binding:"required"`
	TransactionDate   string `json:"transaction_date" binding:"required"` // 또는 time.Time
	TransactionName   string `json:"transaction_name" binding:"required"`
}

type UpdateTransaction struct {
	TransactionType   string  `json:"transaction_type"`
	BankAccountID     string  `json:"bank_account_id"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankCode          string  `json:"bank_code"`
	Amount            int     `json:"amount"`
	TransactionDate   string  `json:"transaction_date"` // 또는 time.Time
	TransactionName   string  `json:"transaction_name"`
	Category          string  `json:"category"`
	CarbonScore       float64 `json:"carbon_score"`
}
