package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	bmap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	sl := []rune(s)
	stack := []string{}

	Loop:
	for i := 0; i <= len(sl) - 1; i++ {
		for k, v := range bmap {
			if string(sl[i]) == k {
				if len(stack) == 0 {
					return false
				}
				var pop = stack[len(stack) - 1]
				if pop != v {
					return false
				}
				stack = stack[:len(stack) - 1]
				continue Loop
			}
		}
		stack = append(stack, string(sl[i]))
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[}"))
}
