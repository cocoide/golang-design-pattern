// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/handler"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository/gorm"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/usecase"
)

// Injectors from wire.go:

func initializeHandlers(config *entity.Config) (*handlers, error) {
	dbConfig := config.DB
	db := gorm.NewDBClient(dbConfig)
	gormParams := &gorm.GormParams{
		DB: db,
	}
	bank := gorm.NewBankImpl(gormParams)
	bankUsecase := usecase.NewBankUsecase(bank)
	bankHandler := handler.NewBankHandler(bankUsecase)
	mainHandlers := &handlers{
		BankHandler: bankHandler,
	}
	return mainHandlers, nil
}

// wire.go:

type handlers struct {
	BankHandler *handler.BankHandler
}