package payment

import "fmt"

type CheckingAccount struct{}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds...")
	return 125
}

func (c *CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with checking account")
	return true
}
