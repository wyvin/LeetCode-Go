package main

import (
	"fmt"
	"math"
)

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
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。
请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func solution(x int) (rev int) {
	rev = 0
	// maxInt32 := 2 << 31 -1
	// minInt32 := -2 << 31
	for x != 0{
		pop := x % 10
		x /= 10

		if rev > math.MinInt32/10 || (rev == math.MinInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func main() {
	var b = 2<<29 - 2
	fmt.Println(b)
	s := solution(b)
	fmt.Println(s)
}
