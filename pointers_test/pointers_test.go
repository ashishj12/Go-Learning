package main

import (
	// "fmt"
	"testing"
)

// create an wallet function which lets us deposit bitcoin
// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}
// 	wallet.Deposit(10)
// 	got := wallet.Balance()

// 	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
// 	want := 10
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}

// }

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
