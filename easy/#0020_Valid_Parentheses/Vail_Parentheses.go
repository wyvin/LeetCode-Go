package _0020_Valid_Parentheses

// 利用栈，如果在map的括号则入栈
// 如果不在则退栈并对比，不正确则返回false
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	// 最大的深度为一半的长度
	// 也可以使用append动态增加
	stack := make([]string, len(s)/2)
	i := -1
	for _, item := range []byte(s) {
		if p, ok := parenthesesMap[string(item)]; ok {
			i++
			if i > len(stack)-1 {
				return false
			}
			stack[i] = p
		} else {
			if i < 0 || stack[i] != string(item) {
				return false
			} else {
				i--
			}
		}
	}
	return i == -1
}

var parenthesesMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}
