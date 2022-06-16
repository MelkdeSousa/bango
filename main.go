package main

import (
	"fmt"
	"os"

	"github.com/MelkdeSousa/bango/domain"
)

type Account interface {
	Withdraw(amount int) error
}

func PayBankBillet(account Account, amount int) {
	err := account.Withdraw(amount)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	ownerMelk := domain.Owner{
		Name:      "Melk de Sousa",
		Id:        "123456789",
		BirthDate: "08/27/2002",
	}
	ownerLemuel := domain.Owner{
		Name:      "Lemuel de Sousa",
		Id:        "987654321",
		BirthDate: "04/14/2004",
	}

	accountMelk := domain.Account{}
	accountLemuel := domain.Account{}

	accountMelk.SetOwner(ownerMelk)
	accountMelk.SetAccountNumber("123456789")
	accountMelk.SetAgencyNumber("12345")
	accountMelk.SetBalance(1000)

	accountLemuel.SetOwner(ownerLemuel)
	accountLemuel.SetAccountNumber("987654321")
	accountLemuel.SetAgencyNumber("54321")
	accountLemuel.SetBalance(2000)

	PayBankBillet(&accountMelk, 100)

	fmt.Println("Melk de Sousa:", accountMelk.GetBalance())
}
