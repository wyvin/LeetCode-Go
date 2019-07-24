package main

import "fmt"

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1 := m - 1
	p2 := n - 1
	c := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[c] = nums1[p1]
			p1--
		} else {
			nums1[c] = nums2[p2]
			p2--
		}
		c--
	}
	if p1 < 0 {
		for i:=c; i>=0; i--  {
			nums1[i] = nums2[i]
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	m := 3
	n := 3
	fmt.Println(nums1[:3])
	//nums1[:3] = nums2
	fmt.Println(nums1)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
