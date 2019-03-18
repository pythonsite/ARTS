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

// 其他人大答案的耗时和内存占用都比上面两种方法小
func reverse3(x int) int {
	var r int
	var limit int
	if x >= 0 {
		limit = math.MaxInt32 / 10 // 2的32次方前9位数字
		for x != 0 {
			remainder := x % 10 //这里每次获取x的各位数字
			// Int32 占4个字节（-2147483648~2147483647），所以翻转后的最后一位不能大于7
			if r > limit || (r == limit && remainder > 7) {
				return 0
			}
			r = r*10 + remainder
			x /= 10
		}
	} else {
		limit = math.MinInt32 / 10
		for x != 0 {
			remainder := x % 10
			// Int32 占4个字节（-2147483648~2147483647）,所以翻转后最后一位不能小于-8
			if r < limit || (r == limit && remainder < -8) {
				return 0
			}
			r = r*10 + remainder
			x /= 10
		}
	}
	return r
}

func main() {
	num := 1534236469
	res := reverse2(num)
	fmt.Println(res)
}
