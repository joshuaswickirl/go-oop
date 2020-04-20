package payment

import "fmt"

type Account struct {
	amount float32
}

func (a *Account) AvailableFunds() float32 {
	println("Getting account funds...")
	return 1
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment, using method...")
	return true
}
