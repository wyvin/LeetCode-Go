package _0058_Length_of_Last_Word

// 2021-01-14
func lengthOfLastWord(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			l++
		} else {
			if l > 0 {
				return l
			}
		}
	}
	return l
}
