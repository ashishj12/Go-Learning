package main

import "fmt"

// when we create an function in golang then it accept an argument with the type and also add an return type of value if we did'n then it shows error of too many returns(by default)
//for e.g.

// func Add(a int, b int) int {
// 	return a + b //-> in this case if we don't include the int type after function it shows an error of too many attempts
// }

func Languages() (string, string, string) {
	return "golang", "javascript", "java"
}

func main() {

	lang1, lang2, lang3 := Languages()
	fmt.Println(lang1, lang2, lang3)
}
