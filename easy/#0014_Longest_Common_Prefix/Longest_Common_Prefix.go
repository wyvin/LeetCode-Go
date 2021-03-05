package _0014_Longest_Common_Prefix

import (
	"math"
)

// 2021-01-05
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 默认取第一个字符串为前缀
	prefix := []byte(strs[0])
	var str []byte
	var length int
	// 应该跳过第一个字符串，从第二个开始循环对比
	for i, _ := range strs {
		if strs[i] == "" {
			return ""
		}
		str = []byte(strs[i])
		// 取两者中的最小长度
		length = int(math.Min(float64(len(str)), float64(len(prefix))))
		// 对比，发现不同的字节则跳出循环
		prefix = prefix[:length]
		for j := 0; j < length; j++ {
			if str[j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return string(prefix)
}
