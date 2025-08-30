package domain

import "github.com/google/uuid"

type Debt struct {
	ID                int64
	UserID            uuid.UUID
	Amount            float64
	CreatedAt         string
	UpdatedAt         string
	IsPaid            bool
	LastPaymentAmount float64
	PaidAt            string
	DebtSummary       float64
	Description       string
}
