package main

import "fmt"

func main()  {
	var i int
	var b bool
	var ip *int
	var s string
	var slice []string
	var m map[string]string

	fmt.Printf("i = %v b = %v s = [%s] ip = %v slice = %v m = %v\n",i,b,s,ip, slice,m)
	fmt.Printf("slice and map is nil = %v\n", (slice == nil && m == nil))
}