package _0027_Remove_Element

// 2021-01-09 双指针
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	newLen := 0
	for i:=len(nums)-1;i>=newLen;i--{
		for nums[newLen] != val && newLen < i {
			newLen++
		}
		println(newLen, i)
		if nums[i] != val {
			nums[i],nums[newLen] = nums[newLen],nums[i]
		}
	}
	if nums[newLen] != val {
		newLen++
	}
	return newLen
}