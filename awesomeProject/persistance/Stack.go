package persistance

import "fmt"

//用切片实现Stack
type Element interface {}

type  sliceEntry struct{
	element []Element
}

func NewSliceEntry() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry)Push(e Element) {
	entry.element = append(entry.element,e)
}

func  (entry *sliceEntry)Pop() Element {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return nil
	}
	lastElement := entry.element[size-1]
	entry.element[size-1] = nil
	entry.element  = entry.element[:size-1]
	return lastElement
}

func  (entry *sliceEntry)Top() Element {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return nil
	}
	return entry.element[size-1]
}


func  (entry *sliceEntry)Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("stack is empty!")
		return false
	}
	for i :=0;i<entry.Size();i++ {
		entry.element[i] = nil
	}
	entry.element = make([]Element,0)
	return true
}

func  (entry *sliceEntry)Size() int {
	return len(entry.element)
}

func  (entry *sliceEntry)IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}
