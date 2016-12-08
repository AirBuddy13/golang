package main

import (
	"fmt"
	"sort"
)


type people [] string

func (p people) Len() int{ return len(p) }

func (p people) Swap(i, j int){ p[i], p[j] = p[j], p[i]	}

func (p people) Less(i, j int) bool{ return p[i] < p[j] }




func main() {
	p := people{"D","A", "F"}
	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
	sort.Sort(sort.Reverse(p))
	fmt.Println(p)
	//s := []string{"F", "D"}


	i := []int {5,2,3}
	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)
	sort.Sort(sort.IntSlice(i))

	sort.Search("D", p)

	fmt.Println(i)


}