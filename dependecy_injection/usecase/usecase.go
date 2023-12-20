package usecase

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository"
)

type BankUsecaseParams struct {
	Repository repository.Bank
}

func NewBankUsecase(repo repository.Bank) *BankUsecase {
	return &BankUsecase{repo: repo}
}
