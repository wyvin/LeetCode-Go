package _0026_Remove_Duplicates_from_Sorted_Array

// 2021-01-08
func removeDuplicates(nums []int) int {
	newLen := 0
	for i, _ := range nums {
		if nums[i]>nums[newLen] {
			newLen++
			nums[newLen] = nums[i]
		}
	}
	return newLen+1
}
