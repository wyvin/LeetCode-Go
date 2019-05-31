package main

import (
	"fmt"
)

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:
输入: "Hello World"
输出: 5
 */
func lengthOfLastWord(s string) int {
	sl := []rune(s)
	length := 0
	last := 0
	for i:=0; i<len(sl); i++ {
		if string(sl[i]) == " " {
			length = 0
		} else {
			length++
		}
		if length != 0 {
			last = length
		}
	}
	return last
}

func main() {
	s := ""
	fmt.Println(lengthOfLastWord(s))
}
