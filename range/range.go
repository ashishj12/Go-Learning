package main

import "fmt"

// range refers to iterating over data structure
func main() {

	// classic iteration in slices
	// nums := []int{6, 7, 8}
	// for  i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//iteration using range
	// sum :=0
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	//first elememt(i) is refer to the index of slice
	// for i, num:= range nums{
	// 	fmt.Println(i,num)
	// }

	//iteration in maps
	// m := map[string]string{"fname": "GO", "lname": "Lang"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }


	//iteration in strings
	for i, c := range "golang" { // -> in string for every char it return unicode code
		fmt.Println(i, c)
	}

}
