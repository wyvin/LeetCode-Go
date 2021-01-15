package _0014_Longest_Common_Prefix

import (
	"math"
)

// 2021-01-05
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := []byte(strs[0])
	var str []byte
	var length int
	for i, _ := range strs {
		if strs[i] == "" {
			return ""
		}
		str = []byte(strs[i])
		length = int(math.Min(float64(len(str)), float64(len(prefix))))
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
