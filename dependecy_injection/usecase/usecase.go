package usecase

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository"
	"github.com/google/wire"
)

type BankUsecaseParams struct {
	Repository repository.Bank
}

func NewBankUsecase(repo repository.Bank) *BankUsecase {
	return &BankUsecase{repo: repo}
}

var Set = wire.NewSet(
	NewBankUsecase,
	wire.Struct(new(BankUsecaseParams), "*"),
)
