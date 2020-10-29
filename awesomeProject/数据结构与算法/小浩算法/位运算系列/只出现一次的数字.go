package main

import "fmt"

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// todo 用异或处理，自己异或自己为0，0异或自己为自己，全部异或完就剩下自己了
func main() {
	nums1 := []int{1, 2, 3, 4, 5, 2, 4, 3, 1}
	fmt.Println(singleNumber1(nums1))
	fmt.Println(singleNumber11(nums1))

	nums2 := []int{1, 3, 2, 3, 2, 1, 3, 2, 1, 5, 5, 4, 5, 6, 4, 4}
	fmt.Println(singleNumber2(nums2))
	fmt.Println(singleNumber3(nums2))
}

func singleNumber1(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

// 改变题目为其余元素出现三次以上，就无法使用异或了兄弟，那看看有什么办法可以让相同的多个数相抵消呢？
// 用map就等被很遗憾吧

// 额，但是三次以上，我们直接乘以2，也是可以的吧，但是这样1次也变2次了

// 其余元素都出现三次的话：就需要骚操作了

// todo  异或本身的运算可以理解成二进制的加法，然后对2取模（砍掉所有进位）
// 两个相同的数相抵消 = a ^ a = 2a % 2
//  101
//  101
// 1010
// 0000 = 异或后 = 对2取模后
// 三个相同的数相抵消 = a ? a ? a = 3a % 3?
//  1001
//  1001
//  1001
// 11011
// 00000 = ？ = 对3取模后

// 其实很好理解...任何东西乘以3再对3取模那肯定为0 =-=
// 如果取余会把最后的答案也取了余，拿不到
func singleNumber11(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res % 2
}

// 所以其实是要对每一位去进行
func singleNumber2(nums []int) int {
	number, res := 0, 0
	for i := 0; i < 64; i++ {
		//初始化每一位1的个数为0
		number = 0
		for _, k := range nums {
			//通过右移i位的方式，计算每一位1的个数
			number += (k >> i) & 1
		}
		//最终将抵消后剩余的1放到对应的位数上
		res |= (number) % 3 << i
	}
	return res
}

// 用三进制去弄，巨佬答案

func singleNumber3(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		b = b ^ v & ^a
		a = a ^ v & ^b
	}
	return b
}
