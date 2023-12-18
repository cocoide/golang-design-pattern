package mysql

import (
	"github.com/cocoide/golang-design-pattern/patterns/dependecy_injection/entity"
	"github.com/cocoide/golang-design-pattern/patterns/dependecy_injection/repository"
	"gorm.io/gorm"
)

type bankImpl struct {
	db *gorm.DB
}

func newBank(db *gorm.DB) repository.Bank {
	return &bankImpl{db: db}
}

func (b *bankImpl) Withdraw(accountID, amount int) error {
	var bank entity.Bank
	if err := transaction(b.db, func(tx *gorm.DB) error {
		if err := tx.Where("account_id = ?", accountID).First(&entity.Bank{}).Error; err != nil {
			return err
		}
		if bank.CannotWithdraw(amount) {
			// error handling
			return nil
		}
		bank.Balance -= amount
		if err := tx.Model(&bank).Update("balance", bank.Balance).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		// logging
		return err
	}
	return nil
}
