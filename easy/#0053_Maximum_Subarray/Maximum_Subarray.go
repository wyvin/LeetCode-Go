package _0053_Maximum_Subarray

// 2020-01-13 贪心法
// 如果之前和小于0则丢弃，在当前位置开始重新算
func maxSubArray(nums []int) int {
	// 之前和
	presum := 0
	// 当前和
	sum := nums[0]
	// 最大和
	max := sum
	for i:=1;i<len(nums);i++{
		presum = sum
		if presum < 0 {
			sum = nums[i]
		} else {
			sum = presum+nums[i]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
