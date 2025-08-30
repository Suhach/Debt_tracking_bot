package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateDebtRequest struct {
		UserID      uuid.UUID `json:"user_id" validate:"required"`
		Amount      float64   `json:"amount" validate:"required,gt=0"`
		Description string    `json:"description" validate:"max=255"`
	}
	CreateDebtResponse struct {
		ID                int64     `json:"id"`
		UserID            uuid.UUID `json:"user_id"`
		Amount            float64   `json:"amount"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		IsPaid            bool      `json:"is_paid"`
		LastPaymentAmount float64   `json:"last_payment_amount,omitempty"`
		PaidAt            time.Time `json:"paid_at,omitempty"`
		Description       string    `json:"description,omitempty"`
	}

	GetDebtsResponse struct {
		ID                int64     `json:"id"`
		UserID            uuid.UUID `json:"user_id"`
		Amount            float64   `json:"amount"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		IsPaid            bool      `json:"is_paid"`
		LastPaymentAmount float64   `json:"last_payment_amount,omitempty"`
		PaidAt            time.Time `json:"paid_at,omitempty"`
	}

	PaymentRequest struct {
		DebtID        int64   `json:"debt_id" validate:"required"`
		PaymentAmount float64 `json:"payment_amount" validate:"required,gt=0"`
	}

	PaymentResponse struct {
		DebtID        int64     `json:"debt_id"`
		UserID        uuid.UUID `json:"user_id"`
		PaymentAmount float64   `json:"payment_amount"`
		PaidAt        time.Time `json:"paid_at"`
		IsFullyPaid   bool      `json:"is_fully_paid"`
		Remaining     float64   `json:"remaining_amount,omitempty"`
	}

	DebtSummaryResponse struct {
		UserID           uuid.UUID `json:"user_id"`
		Username         string    `json:"username"`
		TotalDebt        float64   `json:"total_debt"`
		TotalPaid        float64   `json:"total_paid"`
		RemainingAmount  float64   `json:"remaining_amount"`
		ActiveDebtsCount int       `json:"active_debts_count"`
		PaidDebtsCount   int       `json:"paid_debts_count"`
	}
)
