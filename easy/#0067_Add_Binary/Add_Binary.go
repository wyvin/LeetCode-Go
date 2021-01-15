package _0067_Add_Binary

import (
	"bytes"
	"strconv"
)

// 官方
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 2021-01-15
// 使用补0+判断，内存消耗比上面少- -，缺点太长
const one = "1"
const zero = "0"

func addBinary2(a string, b string) string {
	al := len(a) - 1
	bl := len(b) - 1
	if al != bl {
		for al > bl {
			b = "0" + b
			bl++
		}
		for bl > al {
			a = "0" + a
			al++
		}
	}

	res := make([]byte, 0)
	carry := zero
	tmp := 0
	for al >= 0 {
		if string(a[al]) == one {
			tmp++
		}
		if string(b[al]) == one {
			tmp++
		}
		if carry == one {
			tmp++
		}
		switch tmp {
		case 0:
			res = append(res, '0')
			carry = zero
		case 1:
			res = append(res, '1')
			carry = zero
		case 2:
			res = append(res, '0')
			carry = one
		case 3:
			res = append(res, '1')
			carry = one
		}
		tmp = 0
		al--
	}
	if carry == one {
		res = append(res, '1')
	}
	var builder = bytes.NewBuffer([]byte{})
	for i := len(res) - 1; i >= 0; i-- {
		builder.WriteByte(res[i])
	}
	return builder.String()
}
