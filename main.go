package main

type CheckingAccount struct {
	owner         string
	accountNumber int
	agencyNumber  int
	balance       int
}

func (c *CheckingAccount) Withdraw(amount int) {
	if c.balance >= amount {
		c.balance -= amount
	}
}

func main() {

}
