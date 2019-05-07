package main

import "fmt"

/*
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
func longestCommonPrefix(strs []string) string {
	prefix := ""
	length := len(strs)
	if length == 1 {
		return strs[0]
	}
	if length == 0 || len(strs[0]) <= 0 {
		return prefix
	}
	i := 0
	var start string
	lump:
	for {
		start = string(strs[0][i])
		for j := 0; j < length; j++ {
			if len(strs[j]) - 1 < i {
				break lump
			}
			if start == string(strs[j][i]) {
				if j == length - 1 {
					prefix += start
				}
				continue
			} else {
				break lump
			}
		}
		if i == len(strs[0]) - 1 {
			break lump
		}
		i++
	}
	return prefix
}

func main() {
	test := []string{"a"}
	fmt.Println(longestCommonPrefix(test))
}