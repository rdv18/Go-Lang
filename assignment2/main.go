package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	var mp = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	s, _ := reader.ReadString('\n')

	words := strings.Fields(s)

	for _, s1 := range words {

		v := mp[s1]
		v++
		mp[s1] = v

	}

	p := make(PairList, len(mp))
	i := 0
	for k, v := range mp {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))

	t := 0
	for _, r := range p {
		t++
		if t <= 10 {

			fmt.Printf("Words most frequent:  %v : Times it occured: %v \n", r.Key, r.Value)
		} else {

			return
		}

	}

}
