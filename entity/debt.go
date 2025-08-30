package entity

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	ID                int64     `db:"id"`
	UserID            uuid.UUID `db:"user_id"`
	Amount            float64   `db:"amount"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	IsPaid            bool      `db:"is_paid"`
	LastPaymentAmount float64   `db:"last_payment_amount"`
	PaidAt            time.Time `db:"paid_at"`
}
