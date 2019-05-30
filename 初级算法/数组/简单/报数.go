package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:
输入: 1
输出: "1"

示例 2:
输入: 4
输出: "1211"
 */
func countAndSay(n int) string {
	a := []int{1}
	now := a[0]
	length := 0
	for n - 1> 0 {
		var b []int
		for i := 0; i < len(a); i++ {
			if a[i] == now {
				length += 1
			} else {
				b = append(b, length, now)
				length = 1
				now = a[i]
			}
			if i == len(a) - 1 {
				b = append(b, length, now)
			}
		}
		a = b
		now = a[0]
		length = 0
		n--
	}
	var c strings.Builder
	for i:=0; i<len(a); i++ {
		c.WriteString(strconv.Itoa(a[i]))
	}
	return c.String()
}
func main() {
	n := 3
	fmt.Println("result: ", countAndSay(n))
}
