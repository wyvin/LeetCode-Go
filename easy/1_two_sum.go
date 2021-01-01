package main

/*
两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/


// 2021-01-01
func twoSum1(nums []int, target int) []int {
	var diff int
	for i, num := range nums {
		if num > target {
			continue
		}
		diff = target - num
		for j:=i+1;j<len(nums);j++ {
			if diff == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// best version
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int,len(nums))
	for k,v := range nums{
		t := target - v
		_,ok := m[t]
		if ok {
			return []int{k,m[t]}
		}else {
			m[v] = k
		}
	}

	return []int{}
}
//func main() {
//	nums := []int{2, 7, 11, 15}
//	target := 9
//	fmt.Println(twoSum1(nums, target))
//}
