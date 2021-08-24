package main

import (
	"fmt"
	"sync"
)

func proc(i int, wg *sync.WaitGroup) {
	fmt.Println("Started routine ", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go proc(i, &wg)
	}
	wg.Wait()
	fmt.Println("Done")

}
