package _0020_Valid_Parentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s) % 10 % 2 == 1 {
		return false
	}

	stack := make([]string, len(s)/2)
	i := -1
	for _, item := range []byte(s) {
		if p, ok := parenthesesMap[string(item)]; ok {
			i++
			if i > len(stack) - 1 {
				return false
			}
			stack[i] = p
		} else {
			if i<0 || stack[i] != string(item) {
				return false
			} else {
				i--
			}
		}
	}
	return i==-1
}

var parenthesesMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}