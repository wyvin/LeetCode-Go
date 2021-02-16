package _0171_Excel_Sheet_Column_Number

import "math"

// 从后开始算
func titleToNumber(s string) int {
	bs := []byte(s)
	sum := 0
	bsl := len(bs)
	for i := bsl - 2; i >= 0; i-- {
		sum += (int(bs[i]-'A') + 1) * int(math.Pow(26, float64(bsl-1-i)))
	}
	sum += int(bs[bsl-1]-'A') + 1
	return sum
}

// 从前开始算
func titleToNumber1(s string) int {
	sum := 0
	bs := []byte(s)
	for _, b := range bs {
		sum = 26*sum + (int(b-'A') + 1)
	}
	return sum
}
