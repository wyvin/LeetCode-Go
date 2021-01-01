package __Two_Sum

// 2021-01-01
func twoSum1(nums []int, target int) []int {
	var diff int
	for i, num := range nums {
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
