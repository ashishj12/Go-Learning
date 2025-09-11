package main

import "fmt"

//basically variadic function refers to that when we want that an particular function receive n amount(unlimited) of argumets for that reasons we use variadic functions

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {

	nums := []int{3, 4, 5, 6, 7, 8, 9, 10}
	result := Sum(nums...)
	fmt.Println(result)

}
