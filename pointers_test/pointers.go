package main

// import "fmt"

// type Wallet struct {
// 	balance int
// }

// func (w Wallet) Deposit(amount int) {
// 	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
// 	w.balance += amount
// }

// func (w Wallet) Balance() int {
// 	return w.balance
// }

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
