package main

import "fmt"

/*
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

/*
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hl := []rune(haystack)
	nl := []rune(needle)
	nlen := len(nl)
	for i := 0; i < len(hl) - nlen + 1; i++ {
		if hl[i] == nl[0] {
			for j, v := range hl[i:i+nlen] {
				fmt.Println(v, nl[j])
				if v != nl[j] {
					break
				}
				if j == nlen-1 {
					return i
				}
			}
		}
	}
	return -1
}
*/

// sunday
func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}
	if haystack == "" {
		return -1
	}
	hl := []rune(haystack)
	nl := []rune(needle)
	hlen := len(hl)
	nlen := len(nl)
	if nlen >= hlen {
		return -1
	}
	h := 0
	n := 0
	next := nlen

	for n < nlen {
		if hl[h] == nl[n] {
			h++
			n++
		} else {
			if next >= hlen {
				return -1
			}
			tmp := -1
			for i := nlen - 1; i >= 0; i-- {
				if nl[i] == hl[next] {
					tmp = i
					break
				}
			}
			next += nlen - tmp
			h = next - nlen
			if h > hlen - nlen {
				return -1
			}
			n = 0
		}
	}
	return h - nlen
}

func main() {
	//s := "mississippi"
	//n := "sippia"
	//n := "issip"
	//n := "pi"
	//s := "hello"
	//n := "ll"
	//s := "aaaa"
	//n := "ab"
	//s := "aaaa"
	//n := "aaaaa"
	//s := "aaaaa"
	//n := "baaaa"
	s := ""
	n := "a"
	fmt.Println(strStr(s, n))
}
