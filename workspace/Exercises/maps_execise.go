package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	mp := make(map[string]int)
	str_slice := strings.Fields(s)
	for _, str := range str_slice {
		cnt := 0
		for _, i := range str_slice {
			if str == i {
				cnt += 1
			} else {
				continue
			}
			mp[str] = cnt
		}
	}
	return mp
}

func main() {
	wc.Test(WordCount)
	// mp := WordCount("I ate a donut. Then I ate another donut.")
	// for k, v := range mp {
	// 	fmt.Println(k, ":", v)
	// }
}
