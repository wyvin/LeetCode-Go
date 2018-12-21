package main

import (
	"time"
	"fmt"
)

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 */

func rotate(nums []int, k int)  {
	if k % len(nums) == 0 {
		return
	}
	if k > len(nums) {
		k = k - len(nums)
	}

	var tmp []int
	tmp = append(tmp, nums[len(nums) - k:]...)

	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for j := k-1; j >= 0; j-- {
		nums[j] = tmp[j]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 5
	a := time.Now()
	rotate(nums, k)
	fmt.Println(nums)
	time.Sleep(time.Second*10)
	elapsed := time.Since(a)
	fmt.Println("app elapsed:", elapsed)
}
