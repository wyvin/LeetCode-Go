package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*/

func singleNumber(nums []int) int {
	var num int
	for i := 0; i <len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}

func main() {
	nums := []int{4,1,0,1,0}
	fmt.Println(singleNumber(nums))
}
