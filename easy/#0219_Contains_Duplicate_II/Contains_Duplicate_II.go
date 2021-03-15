package _0219_Contains_Duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)
	for i, num := range nums {
		if j, exist := hash[num]; exist {
			if i-j <= k {
				return true
			}
		}
		hash[num] = i
	}
	return false
}
