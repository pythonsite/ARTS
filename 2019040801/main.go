/*\
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
package main

import (
	"fmt"
)


func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 0 {
		return ""
	}
	var flag = true
	var prefix string
	var minStr = strs[0]
	for i :=1;i<len(strs);i++ {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}

	for i:=0;i<len(minStr);i++ {
		if !flag {
			break
		}
		prefix = minStr[:i+1]
		for j :=0; j < len(strs);j++ {
			if prefix != strs[j][:i+1] {
				flag = false
				prefix = minStr[:i]
				break
			}
		}
		
	}
	return prefix

}


func main() {
	array := []string{"flower","flow","flight"}
	res := longestCommonPrefix(array)
	fmt.Println(res)
}

