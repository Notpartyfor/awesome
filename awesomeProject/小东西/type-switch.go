package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

type Shaper2 interface {
	Area() float32
	Circumference() float32
}

func main() {
	var areaIntf Shaper2
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t.Area())
	//case *Circle:
	//	fmt.Printf("Type Circle %T with value %v\n", t, t.Area())
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	if _, ok := areaIntf.(Shaper); ok {
		fmt.Println(areaIntf.Area())
	}

}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square) Circumference() float32 {
	return 4 * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func GetShaperArea(s Shaper) float32 {
	return s.Area()
}
