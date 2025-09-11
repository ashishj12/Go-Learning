package main

import "fmt"

// in other languages we use for, while and do while loop but in
// golang there is only one looping constart which is For
func main() {

	// while loop using for
	// i := 1 //-> short hand syntax of declaration
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinte loop
	// for {
	// 	println("1")
	// }

	//classic for loop
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

}
