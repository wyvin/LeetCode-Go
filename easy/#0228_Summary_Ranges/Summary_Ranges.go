package _0228_Summary_Ranges

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	x := nums[0]
	y := x
	ans := make([]string, 0)
	helper := func() {
		if x == y {
			ans = append(ans, strconv.Itoa(x))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", y, x))
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == x {
			x++
			continue
		}
		helper()
		x = nums[i]
		y = x
	}
	helper()
	return ans
}

// 官方
func summaryRanges2(nums []int) []string {
	i := 0
	n := len(nums)
	ans := make([]string, 0)
	for i < n {
		left := i
		for i = i + 1; i < n && nums[i-1] == nums[i]-1; i++ {
		}
		if left == i-1 {
			ans = append(ans, strconv.Itoa(nums[left]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[left], nums[i-1]))
		}
	}
	return ans
}
