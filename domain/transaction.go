package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	transaction := &Transaction{}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	return transaction
}
