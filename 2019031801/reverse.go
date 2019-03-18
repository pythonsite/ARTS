/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

// 刚开始自己的实现方式
func reverse(x int) int {
	var b bool
	if x < 0 {
		x = int(math.Abs(float64(x)))
		b = true
	}
	xStr := strconv.Itoa(x)
	xSlice := []rune(xStr)
	for i, j := 0, len(xSlice)-1; i < j; i, j = i+1, j-1 {
		xSlice[i], xSlice[j] = xSlice[j], xSlice[i]
	}
	xresStr := string(xSlice)
	res, _ := strconv.Atoi(xresStr)
	if b == true {
		if res > math.MaxInt32 {
			return 0
		}
		return -res
	}
	if res > math.MaxInt32 {
		return 0
	} else {
		return res
	}
}

// 对代码进行改善
func reverse2(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	xAbs := int(math.Abs(float64(x)))
	xAbsStr := strconv.Itoa(xAbs)
	xAbsSlice := []rune(xAbsStr)
	for i, j := 0, len(xAbsSlice)-1; i < j; i, j = i+1, j-1 {
		xAbsSlice[i], xAbsSlice[j] = xAbsSlice[j], xAbsSlice[i]
	}
	xAbsSliceStr := string(xAbsSlice)
	res, _ := strconv.Atoi(xAbsSliceStr)
	if res > math.MaxInt32 {
		return 0
	}
	if x < 0 {
		return -res
	}
	return res
}

func main() {
	num := 1534236469
	res := reverse2(num)
	fmt.Println(res)
}