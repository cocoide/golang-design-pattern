package entity

import "time"

type Bank struct {
	AccountID string `gorm:"uniqueIndex"`
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Bank) CannotWithdraw(amount int) bool {
	return b.Balance < amount
}
