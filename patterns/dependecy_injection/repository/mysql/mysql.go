package mysql

import (
	"github.com/cocoide/golang-design-pattern/patterns/dependecy_injection/repository"
	"gorm.io/gorm"
)

func NewClient(db *gorm.DB) *repository.Database {
	return &repository.Database{
		Bank: newBank(db),
	}
}
