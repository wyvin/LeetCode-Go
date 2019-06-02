package main

import "fmt"

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。

示例 1:
输入: a = "11", b = "1"
输出: "100"

示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
 */
func addBinary(a string, b string) string {
	result := ""
	flag := 0       // 存储进位
	i, j := len(a)-1, len(b)-1
	for i>=0 || j>=0 {
		t1,t2 := 0,0
		if i>=0 {
			t1 = int(a[i]-'0')
		}
		if j>=0 {
			t2 = int(b[j]-'0')
		}
		sum := t1 + t2 + flag   // 计算当前位置
		switch sum {
		case 3: flag = 1; result = "1" + result
		case 2: flag = 1; result = "0" + result
		case 1: flag = 0; result = "1" + result
		case 0: flag = 0; result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 {  // 最终进位
		result = "1" + result
	}
	return result
}

func main() {
	a := "11"
	b := "11"
	fmt.Println(addBinary(a, b))
}
