package _0013_Roman_to_Integer

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// 右边加法，左边减法
// 从右边开始循环，对比后一位，比后一位小则加，大则减
// 例如："IV" = 5 - 1
//      "VI" = 1 + 5
func romanToInt(s string) int {
	var (
		sum  int
		last int
	)
	for i := len(s) - 1; i >= 0; i-- {
		if romanMap[string(s[i])] >= last {
			sum += romanMap[string(s[i])]
		} else {
			sum -= romanMap[string(s[i])]
		}
		last = romanMap[string(s[i])]
	}
	return sum
}
