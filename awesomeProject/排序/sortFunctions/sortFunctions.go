package sortFunctions

import (
	"awesomeProject/common"
	"math"
)

// todo 冒泡排序
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
//
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
//
// 针对所有的元素重复以上的步骤，除了最后一个。
//
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// todo 基数排序
// 基数排序是一种 非比较型 整数排序算法，其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。
// 由于整数也可以表达字符串（比如名字或日期）和特定格式的浮点数，所以基数排序也不是只能使用于整数。

func RadixSort(arr []int) []int {
	Digit := 0
	maxDigit := 1
	maxValue := common.MaxInSlice(arr)

	// 找出最大的位数 (话说这个不就是找一个len就好了么...如果是整数
	for int(math.Pow10(maxDigit)) < maxValue {
		maxDigit += 1
	}

	for Digit < maxDigit {

		// 0-9
		temp := make([][]int, 10)
		for i := 0; i < 10; i++ {
			temp[i] = make([]int, 0)
		}
		// 求出每一个元素中的个，十，百位的值，放到对应的0 1 2 3 4 5 6 7 8 9 中
		for _, v := range arr {
			t := (v / int(math.Pow10(Digit))) % 10
			temp[t] = append(temp[t], v)
		}

		coll := make([]int, 0)

		for _, bucket := range temp {
			for _, value := range bucket {
				coll = append(coll, value)
			}
		}

		arr = coll
		Digit += 1
	}
	return arr
}

// todo 选择排序
// 选择排序是一种简单直观的排序算法，无论什么数据进去都是 O(n²) 的时间复杂度。
// 所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧

// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
//
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//
// 重复第二步，直到所有元素均排序完毕

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// todo 插入排序
// 插入排序是一种最简单直观的排序算法，它的工作原理是通过构建有序序列
// 对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
// 插入排序和冒泡排序一样，也有一种优化算法，叫做拆半插入。

//将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
//
//从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。

// （如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）

func InsertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

// todo 希尔排序
// 希尔排序，也称递减增量排序算法，是插入排序的一种更高效的改进版本。但希尔排序是非稳定排序算法。
//
// 希尔排序是基于插入排序的以下两点性质而提出改进方法的：
//
// 插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率；
// 但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位；
// 希尔排序的基本思想是：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
// 待整个序列中的记录“基本有序”时，再对全体记录进行依次直接插入排序。

// 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
//
// 按增量序列个数 k，对序列进行 k 趟排序；
//
// 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
// 仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}
