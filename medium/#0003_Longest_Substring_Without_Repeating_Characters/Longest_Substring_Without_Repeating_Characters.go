package _0003_Longest_Substring_Without_Repeating_Characters

import "strings"

// 滑动窗口 - map
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, len(s))
	left, right := 0, 0
	maxlen := 0
	for right < len(s) {
		rb := s[right]
		window[rb]++
		right++
		for window[rb] > 1 {
			lb := s[left]
			window[lb]--
			left++
		}
		maxlen = max(maxlen, right-left)
	}
	return maxlen
}

// 滑动窗口 - 迭代
func lengthOfLongestSubstring2(s string) int {
	l, r := 0, 0
	ans := 0
	for i := range s {
		index := strings.IndexByte(s[l:i], s[i])
		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}
		ans = max(len(s[l:r]), ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
