package main

import "fmt"


type Shape interface {
	area() float64
}

type Race interface {
	Shape
	test()
}
type circle struct {

}

func (c circle) area() float64{
	return 0.0
}

func (c circle) test(){

}


func printArea(s Shape){
	fmt.Println(s.area())
}

func printTest(s Race){
	fmt.Println(s.area())
}

func main() {
	c := circle{}
	fmt.Println(c.area())
	printTest(c)

}
