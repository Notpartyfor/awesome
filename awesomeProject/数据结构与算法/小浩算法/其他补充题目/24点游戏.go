package main

import "fmt"

// 你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通 *，/，+，-，(，) 的运算得到 24 。

// 输入: [4, 1, 8, 7]
//输出: True
//解释: (8-4) * (7-1) = 24

// 注意:
//
//​ 1、除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
//
//​ 2、每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
//
//​ 3、你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。-

// 暴力解

// 判断两个数四种运算下来会不会为24
func judgePoint24_2_mistake(a, b float64) bool {
	return a+b == 24 || a*b == 24 || a-b == 24 || b-a == 24 || a/b == 24 || b/a == 24
}

// 但是这个方法写的正确吗？其实不对！因为在计算机中，实数在计算和存储过程中会有一些微小的误差
// 对于一些与零作比较的语句来说，有时会因误差而导致原本是等于零但结果却小于或大于零之类的情况发生
// 所以常用一个很小的数 1e-6 代替 0，进行判读！
//
//
// (1e-6：表示1乘以10的负6次方。Math.abs(x)<1e-6 其实相当于x==0。
// 1e-6(也就是0.000001)叫做epslon，用来抵消浮点运算中因为误差造成的相等无法判断的情况。这个知识点需要掌握！)

// 例如func main() {
//    var a float64
//    var b float64
//    b = 2.0
//    //math.Sqrt：开平方根
//    c := math.Sqrt(2)
//    a = b - c*c
//    fmt.Println(a == 0)                  //false
//    fmt.Println(a < 1e-6 && a > -(1e-6)) //true
//}

// 这里直接用 a==0 就会得到false，但是通过 a < 1e-6 && a > -(1e-6) 却可以进行准确的判断。

//go语言
//judgePoint24_2：判断两个数的所有操作符组合是否可以得到24
// /和-有两类，+和*实际上是同一个
func judgePoint24_2(a, b float64) bool {
	return (a+b < 24+1e-6 && a+b > 24-1e-6) || // 这里相当于a + b == 24
		(a*b < 24+1e-6 && a*b > 24-1e-6) ||
		(a-b < 24+1e-6 && a-b > 24-1e-6) ||
		(b-a < 24+1e-6 && b-a > 24-1e-6) ||
		(a/b < 24+1e-6 && a/b > 24-1e-6) ||
		(b/a < 24+1e-6 && b/a > 24-1e-6)
}

// 然后接着写三个数的是否可以得到24
//硬核代码，不服来辩！
//
func judgePoint24_3(a, b, c float64) bool {
	return judgePoint24_2(a+b, c) ||
		judgePoint24_2(a-b, c) ||
		judgePoint24_2(a*b, c) ||
		judgePoint24_2(a/b, c) ||
		judgePoint24_2(b-a, c) ||
		judgePoint24_2(b/a, c) ||
		judgePoint24_2(a+c, b) ||
		judgePoint24_2(a-c, b) ||
		judgePoint24_2(a*c, b) ||
		judgePoint24_2(a/c, b) ||
		judgePoint24_2(c-a, b) ||
		judgePoint24_2(c/a, b) ||
		judgePoint24_2(c+b, a) ||
		judgePoint24_2(c-b, a) ||
		judgePoint24_2(c*b, a) ||
		judgePoint24_2(c/b, a) ||
		judgePoint24_2(b-c, a) ||
		judgePoint24_2(b/c, a)
}

//硬核代码，不服来辩！
func judgePoint24(nums []int) bool {
	return judgePoint24_3(float64(nums[0])+float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])+float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])+float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[0]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[0]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])+float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])*float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[1])+float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])*float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[1]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[1]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])+float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])*float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[1]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[1]), float64(nums[2]), float64(nums[0]))
}

func main() {
	nums := []int{0, 0, 0, 24}
	fmt.Println(judgePoint24(nums))
}
