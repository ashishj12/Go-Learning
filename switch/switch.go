package main

import (
	"fmt"
	// "time"
)

func main() {

	// //simple switch
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	//muiliple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's Weekend")
	// default:
	// 	fmt.Println("It's Workday")
	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's an string")
		case bool:
			fmt.Println("It's an boolean")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI(22.12)
}

