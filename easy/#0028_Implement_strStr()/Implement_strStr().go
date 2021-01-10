package _0028_Implement_strStr__

// 2021-01-10
func strStr(haystack string, needle string) int {
	length := len(needle)
	for i:=0;i<=len(haystack)-length;i++{
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}