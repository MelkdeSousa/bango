package domain

import "errors"

type Account struct {
	Owner         string
	AccountNumber int
	AgencyNumber  int
	Balance       int
}

func (account *Account) GetBalance() int {
	return account.Balance
}

func (account *Account) Withdraw(amount int) error {
	if account.Balance < amount {
		return errors.New("insufficient funds")
	}

	if amount < 0 {
		return errors.New("amount must be positive")
	}

	account.Balance -= amount

	return nil
}

func (account *Account) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("amount must be positive")
	}

	account.Balance += amount

	return nil
}

func (account *Account) Transfer(to *Account, amount int) error {
	error := account.Withdraw(amount)

	if error != nil {
		return error
	}

	error = to.Deposit(amount)

	if error != nil {
		return error
	}

	return nil
}
