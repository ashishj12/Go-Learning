// package main

// import "fmt"

// func main() {

// 	// slice -> dynamic
// 	// most used contruct in golang
// 	// useful methods

// 	//uninitized slice is equal to nil
// 	// var nums[]int
// 	// fmt.Println(nums == nil)
// 	// fmt.Println(len(nums))

// 	// var nums = make([]int, 2, 5) //third parameter refer to capcity if we don't intizlied capacity so it can automatically arrage size acc to elements

// 	// fmt.Println(cap(nums)) //capacity refers to maximum numbers of elements that can fit
// 	// fmt.Println(nums == nil)

// 	// nums = append(nums, 1)
// 	// nums = append(nums, 2)
// 	// nums = append(nums, 3)
// 	// nums = append(nums, 4)

// 	// fmt.Println(nums)
// 	// fmt.Println(cap(nums))

// 	var nums = make([]int, 0, 5)
// 	nums = append(nums, 2)
// 	var nums2 = make([]int, len(nums))

// 	fmt.Println(nums2, nums)

// }

// -> pratice slices

package main

import "fmt"

func main() {

	//an uninitialized slice
	// var s []string
	// fmt.Println("uninit slice", s, s == nil, len(s) == 0)

	s := make([]string, 3)
	fmt.Println("emp:", s, "len", len(s), "cap:", cap(s))

}
