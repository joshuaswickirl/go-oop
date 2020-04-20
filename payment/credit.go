package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// CreditAccount ...
type CreditAccount struct {
	Account
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit int
}

var cardNumberPattern = regexp.MustCompile("^[\\d{4}-]{3}\\d{4}$")

// CreateCreditAccount ...
func CreateCreditAccount(ownerName, cardNumber string,
	expirationMonth, expirationYear, securityCode int,
	chargeCh chan float32) *CreditAccount {

	creditAccount := &CreditAccount{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}

	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing a credit card payment, using channel...")
}

// OwnerName ...
func (c *CreditAccount) OwnerName() string {
	return c.ownerName
}

// SetOwnerName ...
func (c *CreditAccount) SetOwnerName(name string) error {
	if len(name) == 0 {
		return errors.New("invalid name provided")
	}
	c.ownerName = name
	return nil
}

// CardNumber ...
func (c *CreditAccount) CardNumber() string {
	return c.cardNumber
}

// SetCardNumber ...
func (c *CreditAccount) SetCardNumber(number string) error {
	if !cardNumberPattern.Match([]byte(number)) {
		return errors.New("invalid credit card number format")
	}
	c.cardNumber = number
	return nil
}

// ExpirationDate ...
func (c *CreditAccount) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate ...
func (c *CreditAccount) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() ||
		(year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("expiration date must be in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode ...
func (c *CreditAccount) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode ...
func (c *CreditAccount) SetSecurityCode(code int) error {
	if code < 100 || code > 999 {
		return errors.New("invalid security code")
	}
	c.securityCode = code
	return nil
}

// AvailableCredit ...
func (c *CreditAccount) AvailableCredit() float32 {
	// this can come from a web service, client doesn't know/care
	return 5000.
}
