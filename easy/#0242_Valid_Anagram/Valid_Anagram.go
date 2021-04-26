package _0242_Valid_Anagram

import "sort"

// 排序后判断对等
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 方法二 哈希表
