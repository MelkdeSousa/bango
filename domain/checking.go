package domain

import "errors"

type Account struct {
	owner                       Owner
	accountNumber, agencyNumber string
	balance                     int
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) GetOwner() Owner {
	return account.owner
}

func (account *Account) GetAccountNumber() string {
	return account.accountNumber
}

func (account *Account) GetAgencyNumber() string {
	return account.agencyNumber
}

func (account *Account) SetOwner(owner Owner) {
	account.owner = owner
}

func (account *Account) SetAccountNumber(accountNumber string) {
	account.accountNumber = accountNumber
}

func (account *Account) SetAgencyNumber(agencyNumber string) {
	account.agencyNumber = agencyNumber
}

func (account *Account) SetBalance(balance int) {
	account.balance = balance
}

func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errors.New("insufficient funds")
	}

	if amount < 0 {
		return errors.New("amount must be positive")
	}

	account.SetBalance(account.balance - amount)

	return nil
}

func (account *Account) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("amount must be positive")
	}

	account.SetBalance(account.balance + amount)

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
