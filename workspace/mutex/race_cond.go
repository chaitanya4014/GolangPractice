package main

import (
	"fmt"
	"sync"
)

var x = 0

func proc(wg *sync.WaitGroup, mx *sync.Mutex) {
	mx.Lock()
	x = x + 1
	mx.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go proc(&wg, &mx)
	}
	wg.Wait()
	fmt.Println("final val of x is:", x)
}
