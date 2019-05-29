package main

import (
	"fmt"
)

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
 */
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		index :=  (end + start) / 2
		if start == end {
			if target <= nums[start] {
				return start
			} else {
				return start + 1
			}
		}
		if nums[index] == target {
			return index
		} else if target < nums[index] {
			end = index
		} else {
			start = index + 1
		}
	}
}


func main()  {
	//l := []int{1, 3, 5, 6, 8}
	l := []int{1, 3, 5, 6}
	n := 7
	fmt.Println(searchInsert(l, n))

}
