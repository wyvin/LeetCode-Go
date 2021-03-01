package _0191_Number_of_1_Bits

// 经过#0190题后，这道看上起就简单多了
func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		ans += int(num & 1)
		num >>= 1
	}
	return ans
}
