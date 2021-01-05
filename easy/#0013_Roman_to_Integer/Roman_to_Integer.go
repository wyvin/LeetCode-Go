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

func romanToInt(s string) int {
	var (
		sum int
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