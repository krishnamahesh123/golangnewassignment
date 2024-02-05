// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type sample interface {
	area() int
	perimeter() int
}

type Rect struct {
	width  int
	length int
}

func (r Rect) area() int {
	return r.width * r.length
}

func (r Rect) perimeter() int {
	return 2*r.width + 2*r.length
}

//func modify(s sample) {
//fmt.Println("Testing")
//s.width = 100
//fmt.Println(s)
//}

func main() {
	var s sample = Rect{width: 20, length: 10}

	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	//modify("check")
}
