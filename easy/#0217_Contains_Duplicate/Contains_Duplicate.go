package _0217_Contains_Duplicate

import "sort"

// 排序后再循环
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 哈希表
func containsDuplicate2(nums []int) bool {
	hash := make(map[int]struct{})
	for _, i := range nums {
		if _, ok := hash[i]; ok {
			return true
		}
		hash[i] = struct{}{}
	}
	return false
}
