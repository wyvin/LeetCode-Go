package _0028_Implement_strStr__

// 2021-01-10
func strStr(haystack string, needle string) int {
	length := len(needle)
	if length == 0 {
		return 0
	}
	for i:=0;i<=len(haystack)-length;i++{
		if haystack[i] == needle[0] && haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}