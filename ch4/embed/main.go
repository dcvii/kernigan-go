package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func (w Wheel) String() string {
	return fmt.Sprintf("Wheel{Circle:Center{X:%v, Y:%v}, Radius:%v}, Spokes:%v}", w.Circle.Center.X, w.Circle.Center.Y, w.Circle.Radius, w.Spokes)
}

func main() {

	w1 := Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)

	w2 := Wheel{Circle{Point{9, 9}, 6}, 25}
	fmt.Printf(w2.String())
}
