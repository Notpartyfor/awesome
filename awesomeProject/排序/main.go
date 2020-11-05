package main

import (
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		arr = append(arr, r.Intn(10000))
	}

	//fmt.Println(arr[20:])
	//fmt.Println(sortFunctions.BubbleSort(arr)[:20])
	//fmt.Println(sortFunctions.RadixSort(arr)[:20])
	//fmt.Println(sortFunctions.SelectionSort(arr)[:20])
	//fmt.Println(sortFunctions.ShellSort(arr)[:20])
}
