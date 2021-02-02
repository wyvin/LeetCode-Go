package _0125_Valid_Palindrome

// ASCII：
// 1-9：48～57
// A-Z：65～90
// a-z：97～122
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	left := 0
	right := len(s) - 1
	for left <= right {
		if !isLetter(s[left]) {
			left++
			continue
		}
		if !isLetter(s[right]) {
			right--
			continue
		}
		if isEqual(s[left], s[right]) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isLetter(b byte) bool {
	if 'A' <= b && b <= 'Z' {
		return true
	}
	if 'a' <= b && b <= 'z' {
		return true
	}
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}

func isEqual(a, b byte) bool {
	if a == b {
		return true
	}
	if a <= 57 || b <= 57 {
		return false
	}
	if a > b {
		return a-32 == b
	}
	return b-32 == a
}
