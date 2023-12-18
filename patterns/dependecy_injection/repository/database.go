package repository

type Database struct {
	Bank Bank
}

type Bank interface {
	Withdraw(accountID, amount int) error
}
