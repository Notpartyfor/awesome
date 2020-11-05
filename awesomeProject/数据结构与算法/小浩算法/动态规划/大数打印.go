package main

import "fmt"

func main() {
	nums := 90
	printNumbers(nums)
}

func printNumbers(num int) {
	//声明字符数组,用来存放一个大数
	number := make([]byte, num)
	for i, _ := range number {
		number[i] = '0'
	}
	for !incrementNumber(number) {
		saveNumber(number)
	}

}

func incrementNumber(number []byte) bool {
	// 循环体退出标识
	var isBreak bool = false
	// 进位标识
	var carryFlag byte = 0
	l := len(number)

	for i := l - 1; i >= 0; i-- {
		//取第i位的数字转化
		nSum := number[i] - '0' + carryFlag
		if i == l-1 {
			// 最低位加一
			nSum++
		}
		if nSum >= 10 {
			if i == 0 {
				isBreak = true
			} else {
				//进位之后减10，并把进位标识设置为1
				nSum -= 10
				carryFlag = 1
				number[i] = '0' + nSum
			}
		} else {
			number[i] = '0' + nSum
			break
		}
	}
	return isBreak
}

func saveNumber(number []byte) {
	var isBegin0 = true
	for i, b := range number {
		if isBegin0 && b != '0' {
			isBegin0 = false
		}
		if !isBegin0 {
			fmt.Printf("%s ", number[i:]) // 这里有东西的啊，如果只是单个只会输出对应的uint8如10就会变成 49和48，要一起输出成字符串
			break
		}
	}
}
