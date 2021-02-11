package _0169_Majority_Element

// 方法一：用哈希表记录元素出现的次数
// 方法二：排序，然后取n/2下标的值

// 进阶方法：Boyer-Moore 投票算法
// 时间复杂度O(n)，空间复杂度O(1)
func majorityElement(nums []int) int {
	m := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			m = num
		}
		if num == m {
			count++
		} else {
			count--
		}
	}
	return m
}
