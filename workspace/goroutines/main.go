package main

import "fmt"

func get_digits(n int, ch chan int) {
	for n > 0 {
		rem := n % 10
		ch <- rem
		n = n / 10
	}
	close(ch)
}

func squares(n int, sq chan int) {
	digits_ch := make(chan int)
	go get_digits(n, digits_ch)
	sqr_sum := 0
	for v := range digits_ch {

		sqr_sum += (v * v)
	}
	sq <- sqr_sum
}

func cubes(n int, cb chan int) {
	digits_ch := make(chan int)
	go get_digits(n, digits_ch)
	cube_sum := 0
	for v := range digits_ch {
		cube_sum += (v * v * v)
	}
	cb <- cube_sum
}

func main() {
	fmt.Println("Enter the num: ")
	var num int
	fmt.Scanln(&num)
	sq_ch := make(chan int)
	cb_ch := make(chan int)
	go squares(num, sq_ch)
	go cubes(num, cb_ch)
	res := <-sq_ch + <-cb_ch

	fmt.Println(res)
}

// package main

// import "fmt"

// type employee struct {
// 	fname string
// 	lname string
// }

// type contact struct {
// 	zip    int
// 	mobile int
// }

// type printName interface {
// 	printnm()
// }

// type printContact interface {
// 	printcnt()
// }

// type embedded interface {
// 	printName
// 	printContact
// }

// func (e employee) printnm() {
// 	fmt.Println("First and Last Name of the employee is: ", e.fname, e.lname)
// }

// func (e employee) printcnt() {
// 	fmt.Println("First and Last Name of the employee is: ", e.fname, e.lname)
// }

// func (c contact) printcnt() {
// 	fmt.Println("Contact details of the employee is: ", c.mobile, c.zip)
// }

// func (c contact) printnm() {
// 	fmt.Println("Contact details of the employee is: ", c.mobile, c.zip)
// }

// func main() {

// 	e := employee{
// 		fname: "raj",
// 		lname: "ram",
// 	}
// 	c := contact{
// 		zip:    12345,
// 		mobile: 123654,
// 	}
// 	var emb embedded = e
// 	var emb1 embedded = c
// 	emb.printcnt()
// 	emb.printnm()
// 	emb1.printnm()
// 	emb1.printcnt()
// }
