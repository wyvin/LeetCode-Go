package _0035_Search_Insert_Position

// 2021-01-11
func searchInsert(nums []int, target int) int {
	for i, _ := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 二分查找
func searchInsert2(nums []int, target int) int {
	n := len(nums)
	l, h := 0, n-1
	res, mid := n, 0
	for l <= h {
		mid = (l + h) >> 1
		// 不能用>=比较，因为会出现h+1=n的答案
		// 循环的条件内得不到h+1
		if target <= nums[mid] {
			res = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
