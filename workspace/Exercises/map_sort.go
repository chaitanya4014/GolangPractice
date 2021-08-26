package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	key   string
	value string
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(a, b int)      { p[a], p[b] = p[b], p[a] }
func (p PairList) Less(a, b int) bool { return p[a].value < p[b].value }

func main() {
	mp := map[string]string{
		"Apple": "cc",
		"Dog":   "pp",
		"Cat":   "ff",
		"Ball":  "aa",
	}

	pl := make(PairList, len(mp))
	i := 0
	for k, v := range mp {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	for _, v := range pl {
		fmt.Println(v)
	}

}
