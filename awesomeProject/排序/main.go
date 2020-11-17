package main

import (
	"awesomeProject/排序/sortFunctions"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		arr = append(arr, r.Intn(10000))
	}

	fmt.Println(arr[:20])
	fmt.Println(sortFunctions.BubbleSort(arr)[:20])
	fmt.Println(sortFunctions.BubbleSortMine(arr)[:20])
	fmt.Println(sortFunctions.RadixSort(arr)[:20])
	fmt.Println(sortFunctions.RadixSortMine(arr)[:20])
	fmt.Println(sortFunctions.SelectionSort(arr)[:20])
	fmt.Println(sortFunctions.SelectionSortMine(arr)[:20])
	fmt.Println(sortFunctions.InsertionSort(arr)[:20])
	fmt.Println(sortFunctions.InsertionSortMine(arr)[:20])
	//fmt.Println(sortFunctions.ShellSort(arr)[:20])
	//fmt.Println(sortFunctions.MergeSort(arr)[:20])
	//fmt.Println(sortFunctions.QuickSortXH(arr)[:20])
	//fmt.Println(sortFunctions.HeapSort(arr)[:20])
	//fmt.Println(sortFunctions.CountingSort(arr)[:20])
	//fmt.Println(sortFunctions.BucketSort(arr,500)[:20])
}
