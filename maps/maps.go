package main

import "fmt"

func main() {

	//maps is like objects key value pairs
	m := make(map[string]string)
	//setting an element
	m["name"] = "golang"
	//get an element
	fmt.Println(m["name"])

	//IMP:-> if key does not exist in the map then it return zero value
	fmt.Println(m["backend"])

}
