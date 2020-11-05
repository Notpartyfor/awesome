package sortFunctions

import (
	"awesomeProject/common"
	"fmt"
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
		// 求出每一个元素中对应位数（从个位开始）的值，放到对应的0 1 2 3 4 5 6 7 8 9 中
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
	// 从第一个开始，设min
	for i := 0; i < length-1; i++ {
		min := i
		// 从i+1个开始向后找比min小的数
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 将min和i值置换
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
		// 前面的元素存在并且比当前的数大，就交换，然后继续向前走，直到当前的数比前面的数大
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

// todo 希尔排序，不明白
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

// todo 归并排序
// 算法步骤
// 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//
// 设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//
// 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//
// 重复步骤 3 直到某一指针达到序列尾；
//
// 将另一序列剩下的所有元素直接复制到合并序列尾。

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	// 分左右两边
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	// 这个看上去有点像合并两个有序链表
	// 左右两边，对比首元素
	// 把小的塞进去
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// 把剩下的也加进去
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// todo 快速排序
// 算法步骤
//从数列中挑出一个元素，称为 “基准”（pivot）;
//
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
// 在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//
// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
//
// 递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，但是这个算法总会退出，
// 因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。

// 小浩算法
func quickSortXH(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	// 将left作为基准
	pivot := left
	index := pivot + 1
	// 从pivot（基准）向右走，如果比基准小的，就跟index互换
	// index 指向的是最小的？
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	// 头尾交换？
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 百度
// 第一种写法
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		// 找出基准后，第一个比基准小的值
		for j >= p && values[j] >= temp {
			j--
		}
		// 跳出for循环，如果不是因为j >= p 的话，也就是values[j] < temp 了
		// 也就是找到了基准后的第一个比基准值小的数
		// 交换
		if j >= p {
			values[p] = values[j]
			p = j
		}
		// 找出基准前，第一个比基准大的数
		for i <= p && values[i] <= temp {
			i++
		}
		// 交换
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	// 此时p左边比p小，右边比p大
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}

// 第二种写法
func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	// mid 做基准值
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println(values)
		// 比基准值大就放到tail，然后tail指向前一位
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			// 比基准值小就放到head，然后head指向后一位
			values[i], values[head] = values[head], values[i]
			head++
			// 对比下一个数
			i++
		}
	}
	// 将基准值放到中间
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}

// 第三种写法
func Quick3Sort(a []int, left int, right int) {

	if left >= right {
		return
	}

	explodeIndex := left

	for i := left + 1; i <= right; i++ {

		if a[left] >= a[i] {

			//分割位定位++
			explodeIndex++
			a[i], a[explodeIndex] = a[explodeIndex], a[i]

		}

	}

	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex], a[left]

	Quick3Sort(a, left, explodeIndex-1)
	Quick3Sort(a, explodeIndex+1, right)

}

// todo 堆排序
// 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。
// 堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
// 堆排序可以说是一种利用堆的概念来排序的选择排序。分为两种方法：
//
//大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
//小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；

//堆排序的平均时间复杂度为 Ο(nlogn)。
//
// 算法步骤
//将待排序序列构建成一个堆 H[0……n-1]，根据（升序降序需求）选择大顶堆或小顶堆；
//
//把堆首（最大值）和堆尾互换；
//
//把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
//
//重复步骤 2，直到堆的尺寸为 1。

func heapSort(arr []int) []int {
	arrLen := len(arr)
	// 初始化能塞满arr元素进去的最大堆
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		// 排序后交换首尾，此时尾巴肯定是最大/最小
		swap(arr, 0, i)
		// 将堆的规模-1
		arrLen -= 1
		// 继续排序
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}
