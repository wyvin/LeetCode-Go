package _0053_Maximum_Subarray

// 2020-01-13 贪心法
// 如果之前和小于0则丢弃，在当前位置开始重新算
func maxSubArray(nums []int) int {
	// 和
	sum := nums[0]
	// 最大和
	max := sum
	for i:=1;i<len(nums);i++{
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = sum+nums[i]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

// 动态规划
// 若前一个元素大于0，则将其加到当前元素上
func maxSubArray2(nums []int) int {
	max := nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
