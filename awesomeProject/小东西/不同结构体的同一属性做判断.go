package main

import (
	"fmt"
	"reflect"
)

type Name struct {
	Name    string `json:"Name"`
	Display uint   `json:"Display,string"`
}

type Sex struct {
	Sex     string `json:"Sex"`
	Display uint   `json:"Display,string"`
}

type PersonSim struct {
	Name
	Sex
}

func IsDisplayed(D uint,ANY interface{}) interface{} {
	typ := reflect.TypeOf(ANY)
	fmt.Println(typ)
	return typ
}

func main() {
	n := &Name{
		Name:    "test",
		Display: 1,
	}

	fmt.Printf("%T",IsDisplayed(n.Display,n))

}
