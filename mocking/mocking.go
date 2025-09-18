package main

// // import (
// // 	"fmt"
// // 	"io"
// // 	"os"
// // )

// // func Countdown(out io.Writer) {
// // 	for i := 3; i > 0; i-- {
// // 		fmt.Fprintln(out, i)
// // 	}
// // 	fmt.Fprint(out, "Go!")
// // }

// // func main() {
// // 	Countdown(os.Stdout)
// // }

// package main

// import (
// 	"fmt"
// 	"io"
// )

// type Sleeper interface{
// 	Sleep()
// }

// type SpySleep struct{
// 	Calls int
// }

// func (s *SpySleep) Sleep(){
// 	s.Calls++
// }


// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	fmt.Fprint(out, finalWord)
// 