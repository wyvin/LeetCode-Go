package _0202_Happy_Number

// n不会无穷大
// 9999999999999 -> 1053; 9999 -> 324; 999 -> 243
// n只会从最大值掉到243以下，然后在243内循环

// 使用哈希检查循环
func isHappy(n int) bool {
	tag := make(map[int]struct{})
	for n != 1 {
		tmp := 0
		for n != 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		if _, ok := tag[tmp]; ok {
			return false
		} else {
			tag[tmp] = struct{}{}
		}
		n = tmp
	}
	return n == 1
}

// 快慢指针 时间复杂度虽然一样，但时间内的工作变多了
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 硬编码，循环243内每个数，得出循环链，在链上的都返回false
