package main

import "fmt"

// basically in this code we have two function one for integer items and one is for names (string) and for these two function we have to separaly create function, suppose we work on production so there is an changes of duplication error to avoid the error we use generics.

//to use generics in code we create an function an add syntax [T any] this refer to we pass any type of data in this, we don't need to create seprate functions for different type of functionality,

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(names []string) {
// 	for _, name := range names {
// 		fmt.Println(name)
// 	}
// }

func main() {

	nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}

	printSlice(nums)
	printSlice(names)
	// printStringSlice(names)

}
