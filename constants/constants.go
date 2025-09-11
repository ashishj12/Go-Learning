package main

import "fmt"

const age = 20

func main() {

	//suppose we want to group contants
	const (
		host = "localhost"
		port = 5000
	)

	fmt.Println(host, port)
	fmt.Println(age)

}
