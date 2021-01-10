package _0027_Remove_Element

// 2021-01-09 双指针（前/后）
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	newLen := 0
	for i:=len(nums)-1;i>=newLen;i--{
		for nums[newLen] != val && newLen < i {
			newLen++
		}
		if nums[i] != val {
			nums[i],nums[newLen] = nums[newLen],nums[i]
			newLen++
		}
	}
	return newLen
}

// 2021-01-10 双指针（前/前）
func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}