package _0204_Count_Primes

func countPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	return ans
}

// 质数：大于1，只有1和自己两个约数
// 偶数都不是质数
// 只循环奇数是否能被奇数整除，能整除则不是质数
func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	// 过滤掉偶数
	if value%2 == 0 {
		return false
	}
	for i := 3; i*i <= value; i += 2 {
		if value%i == 0 {
			return false
		}
	}
	return true
}

// 上面的方法可以根据数学性质进一步优化
