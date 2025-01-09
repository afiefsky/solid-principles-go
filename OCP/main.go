package main

import "fmt"

// PaymentProcessor interface defines a common behavior for all payment methods.
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// CreditCard payment method
type CreditCard struct{}

func (cc CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
}

// PayPal payment method
type PayPal struct{}

func (pp PayPal) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

func main() {
	// Using CreditCard
	var payment PaymentProcessor = CreditCard{}
	payment.ProcessPayment(50.0)

	// Using PayPal
	payment = PayPal{}
	payment.ProcessPayment(75.0)
}
