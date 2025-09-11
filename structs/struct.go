package main

import (
	"fmt"
	"time"
)

// struct embedding
type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// reciver type ->
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

func main() {

	myOrder := order{
		id:     "1",
		amount: 30,
		status: "failed",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	fmt.Println(myOrder.customer)

	//create instace
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 20.00,
	// 	status: "recieved",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	// //insert values
	// myOrder.createdAt = time.Now()

	// //get the values
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:     "2",
	// 	amount: 100,
	// 	status: "delivered",
	// }
	// fmt.Println("order struct", myOrder2)

}
