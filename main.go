package main

type CheckingAccount struct {
	owner         string
	accountNumber int
	agencyNumber  int
	balance       int
}

func (c *CheckingAccount) Withdraw(amount int) {
	if amount >= 0 && c.balance >= amount {
		c.balance -= amount
	}
}

func main() {

}
