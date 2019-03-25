/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/

package main

import (
	"fmt"
	"strconv"
)

// 自己的实现方式
func isPalindrome(x int) bool {
	if x < 0 {
        return false
    } else if x <= 9 {
        return true
    } else if x % 10 == 0 {
        return false
    }
	xAbsStr := strconv.Itoa(x)
	xAbsSlice := []rune(xAbsStr)
	for i, j := 0, len(xAbsSlice)-1; i < j; i, j = i+1, j-1 {
		xAbsSlice[i], xAbsSlice[j] = xAbsSlice[j], xAbsSlice[i]
	}
	xAbsSliceStr := string(xAbsSlice)
	return xAbsSliceStr == xAbsStr
}

//leetcode 上更优的实现
func isPalindrome2(x int) bool {
	if x < 0 {
        return false
    } else if x <= 9 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    
    var y int
    var r int
    for x > y {
        r = x % 10
        x = x / 10
        y = y * 10 + r
        
        if x == y || x / 10 == y {
            return true
        }
    }
    return false
}

func main() {
	res := isPalindrome2(121)
	fmt.Println(res)
}
