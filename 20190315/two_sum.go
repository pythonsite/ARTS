/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	//  自己的实现方式
	var array = []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				array = append(array, i, j)
				return array
			}
		}
	}
	return array
}

func twoSum2(nums []int, target int) []int {
	// 别人的实现方式
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := numMap[v]; ok {
			return []int{j, i}
		} else {
			numMap[target-v] = i
		}
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 3, 15}
	array := twoSum2(nums, 9)
	fmt.Printf("%v", array)
}

/*
小结： 自己开始做的时候首先有的想法就是通过循环嵌套的方式，当实现之后看了其他人的答案，第一个感觉是人家的方法比我的简单了很多，也扩展了一下自己的思维
不能总是局限于之前的思维方式
*/
