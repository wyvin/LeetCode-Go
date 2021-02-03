package _0136_Single_Number

// 2021-02-03
// 异或运算：之前遇到过类似的问题，当时知道用异或的时候觉得很神奇，印象深刻
func singleNumber(nums []int) int {
	var ans int
	for i, _ := range nums {
		ans ^= nums[i]
	}
	return ans
}
