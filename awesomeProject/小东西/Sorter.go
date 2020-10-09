package main

import (
	"awesomeProject/排序"
	"fmt"
	"log"
	"strings"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data) //conversion to type IntArray
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

//func strings() {
//	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
//	a := sort.StringArray(data)
//	sort.Sort(a)
//	if !sort.IsSorted(a) {
//		panic("fail")
//	}
//	fmt.Printf("The sorted array is: %v\n", a)
//}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }
func (p *dayArray) List() string {
	var str []string
	for _, d := range p.data {
		str = append(str, fmt.Sprintf("%s ", d.longName))
	}
	return strings.Join(str, "")
}

// 注意调用这方法的时候因为Stringer接口定义的是值所以这里不能用指针去定义方法
func (p dayArray) String() string {
	log.Print("string")
	return p.List()
}

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}
	fmt.Println(a)
}

type Person struct {
	firstName string
	lastName  string
}

type PersonArray struct {
	data []*Person
}

func (p *PersonArray) Len() int           { return len(p.data) }
func (p *PersonArray) Less(i, j int) bool { return p.data[i].firstName < p.data[j].firstName }
func (p *PersonArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i]}

func Persons() {
	zhangsan := Person{firstName:"张", lastName:"三"}
	lisi := Person{firstName:"李", lastName:"四"}
	wangwu := Person{firstName:"王", lastName:"五"}
	data := []*Person{&zhangsan, &wangwu ,&lisi}
	a := PersonArray{data:data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range a.data {
		fmt.Printf("%s%s    ", d.firstName,d.lastName)
	}
}


func main() {
	ints()
	//strings()
	days()
	Persons()
}
