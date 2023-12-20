package handler

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/usecase"
	"github.com/google/wire"
)

type BankHandlerParams struct {
	Usecase *usecase.BankUsecase
}

func NewBankHandler(uc *usecase.BankUsecase) *BankHandler {
	return &BankHandler{uc: uc}
}

var Set = wire.NewSet(
	NewBankHandler,
	wire.Struct(new(BankHandlerParams), "*"),
)
