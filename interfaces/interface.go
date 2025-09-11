//In Go, an interface is a type that lists methods without providing their code. You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.

package main

import "fmt"

// make an interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

// create an payment gateway struct
type payment struct {
	gateway paymenter
}

// create method to integrate with payment struct
func (p payment) makePayment(amount float32) {

	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	p.gateway.pay(amount)

}

// razorpay struct to handle the payment gateway
type razorpay struct{}

// method to integrate the razor pay struct
func (r razorpay) pay(amount float32) {

	//logic to make payment
	fmt.Println("making payment using razorpay", amount)

}

// // stripe struct to handle the payement gateway
// type stripe struct{}

// // method to handle the logic of intergration of razorpay
// func (s stripe) stripePay(amount float32) {

// 	//logic to make payment
// 	fmt.Println("making payment using stripe", amount)

// }

// paypal integration for payment
type paypal struct{}

// method to handle the logic
func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}


func (p paypal) refund(amount float32, account string) {
	fmt.Println("handle refund of amount", amount, account)
}

// fake payment for testing purpose
type fakePayment struct{}

// method to handle the logic
func (f fakePayment) pay(amount float32) {
	fmt.Println("making fake payment for testing purpose")
}

func main() {
	// fakeGw := razorpay{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)

}
