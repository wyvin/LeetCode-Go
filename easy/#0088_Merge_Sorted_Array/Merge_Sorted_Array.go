package _0088_Merge_Sorted_Array

// 2021-01-19 双指针（从后往前）
func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[n+m-1] = nums2[n-1]
			n--
		} else {
			nums1[n+m-1] = nums1[m-1]
			m--
		}
	}
	if n > 0 && m == 0 {
		for n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
		}
	}
	return
}
