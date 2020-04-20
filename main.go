package main

import (
	"github.com/joshuaswickirl/go-oop/payment"
)

type Option interface {
	ProcessPayment(float32) bool
}

func main() {
	var option Option
	chargeCh := make(chan float32)
	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5, 2021, 123, chargeCh)
	// fmt.Printf("Owner name: %v\n", option.OwnerName())
	// fmt.Printf("Card number: %v\n", option.CardNumber())
	// fmt.Println("Attempting to change card number...")
	// err := option.SetCardNumber("invalid")
	// if err != nil {
	// 	fmt.Printf("Nope: %v\n", err)
	// }
	// fmt.Printf("Available credit: %v\n", option.AvailableCredit())

	chargeCh <- 500

	option.ProcessPayment(500)
	option = payment.CreateCashAccount()
	option.ProcessPayment(500)

	ha := &payment.HybridAccount{}
	println(ha.AvailableFunds())
}
