package main

import "fmt"

func main() {
	var ch chan string
	select {
	case s := <-ch:
		fmt.Println(s)
	default:
		fmt.Println("Default Case")
	}
}

// func main() {
// 	ch := make(chan string)

// 	select {
// 	case <-ch:
// 	default:
// 		fmt.Println("Default Case")
// 	}
// }

// func proc(ch chan string) {
// 	time.Sleep(10500 * time.Millisecond)
// 	ch <- "From proc"
// }

// func main() {
// 	ch := make(chan string)

// 	go proc(ch)

// 	for {
// 		time.Sleep(1000 * time.Millisecond)
// 		select {
// 		case s := <-ch:
// 			fmt.Println(s)
// 			return
// 		default:
// 			fmt.Println("Default Case")
// 		}
// 	}
// }
