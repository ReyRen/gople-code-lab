package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{Point{
			X: 8,
			Y: 8,
		}, 5},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)

	w.X = 42
	/*但是在包外部，因为circle和point没有导出，不能访问它们的成员*/

	fmt.Printf("%#v\n", w)
}
