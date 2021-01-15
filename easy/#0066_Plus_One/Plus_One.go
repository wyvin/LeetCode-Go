package _0066_Plus_One

// 20201-01-15
// 在原数组上修改，最高位需要进位时移动数组
func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
		i--
	}
	return append([]int{1}, digits...)
}
