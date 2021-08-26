package main

import (
	"fmt"
	"sync"
)

var x = 0

func inc(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go inc(&wg, ch)
	}
	wg.Wait()

	fmt.Println(x)

}
