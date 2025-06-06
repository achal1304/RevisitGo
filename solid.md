package main

import "fmt"

// Open/Closed: PaymentProcessor is open for extension but closed for modification
type PaymentProcessor interface {
    ProcessPayment(amount float64)
}

type CreditCardPayment struct {}

func (c *CreditCardPayment) ProcessPayment(amount float64) {
    fmt.Println("Processing credit card payment:", amount)
}

type PaypalPayment struct {}

func (p *PaypalPayment) ProcessPayment(amount float64) {
    fmt.Println("Processing PayPal payment:", amount)
}

func ProcessPaymentMethod(p PaymentProcessor, amount float64) {
    p.ProcessPayment(amount)
}

func main() {
    cc := &CreditCardPayment{}
    pp := &PaypalPayment{}
    
    ProcessPaymentMethod(cc, 100)
    ProcessPaymentMethod(pp, 50)
}
