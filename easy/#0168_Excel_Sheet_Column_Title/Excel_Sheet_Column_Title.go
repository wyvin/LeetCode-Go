package _0168_Excel_Sheet_Column_Title

import "strings"

func convertToTitle(n int) string {
	if n <= 26 {
		return Base26[n]
	}
	var buf strings.Builder
	var list []string
	for n > 26 {
		tmp := n % 26
		if tmp == 0 {
			tmp = 26
		}
		list = append(list, Base26[tmp])
		n = (n - tmp) / 26
	}
	if n != 0 {
		list = append(list, Base26[n])
	}
	for i := len(list) - 1; i >= 0; i-- {
		buf.WriteString(list[i])
	}
	return buf.String()
}

var Base26 = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

// 最优解
func convertToTitle2(n int) string {
	var rst string
	for n != 0 {
		n--
		rst = string(n%26+'A') + rst
		n = n / 26
	}
	return rst
}
