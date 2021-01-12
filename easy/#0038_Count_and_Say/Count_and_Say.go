package _0038_Count_and_Say

import (
	"fmt"
	"strconv"
	"strings"
)

// 2021-01-12 用字符串处理效率较低10ms+
func countAndSay(n int) string {
	var (
		res = "1"
		builder strings.Builder
		tmp string
		s string
	)
	for n > 1 {
		n--
		// 分割的字符串
		tmp = ""
		// 上一个字符
		s = string(res[0])
		builder.Reset()
		for _, item := range res {
			if string(item) == s {
				tmp += string(item)
			} else {
				s = string(item)
				builder.WriteString(fmt.Sprintf("%d%s",len(tmp), string(tmp[0])))
				tmp = s
			}
		}
		// 加上最后一串字符
		builder.WriteString(fmt.Sprintf("%d%s",len(tmp), string(tmp[0])))
		res = builder.String()
	}
	return res
}

// 用整数处理 0ms
func countAndSay2(n int) string {
	var (
		res = []int{1}
		tmp []int
		s int
		l int
		builder strings.Builder
	)
	for n > 1 {
		n--
		l = 0
		// 相同的数字
		tmp = []int{}
		// 上一个数字
		s = res[0]
		for _, item := range res {
			if item == s {
				l++
			} else {
				tmp = append(tmp, l, s)
				s = item
				l = 1
			}
		}
		// 加上最后一串字符
		tmp = append(tmp, l, s)
		res = tmp
	}
	for i, _ := range res {
		builder.WriteString(strconv.Itoa(res[i]))
	}
	return builder.String()
}