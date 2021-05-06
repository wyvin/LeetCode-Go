package _0258_Add_Digits

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for i := num % 10; num > 0; i = num % 10 {
			sum += i
			num /= 10
		}
		num = sum
	}
	return num
}

// 进阶 O(1)
// 1. 能够被9整除的整数，各位上的数字加起来也必然能被9整除，所以，连续累加起来，最终必然就是9。
// 2. 不能被9整除的整数，各位上的数字加起来，结果对9取模，和初始数对9取摸，是一样的，所以，连续累加起来，最终必然就是初始数对9取摸。
func addDigits2(num int) int {
	if num < 10 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
