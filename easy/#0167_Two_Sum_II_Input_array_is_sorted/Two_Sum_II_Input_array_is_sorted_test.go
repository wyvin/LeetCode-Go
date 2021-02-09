package _0167_Two_Sum_II_Input_array_is_sorted

import (
	"fmt"
	"testing"
)

type input struct {
	X      []int
	Target int

	Answer []int
}

var output []int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = twoSum2(input.X, input.Target)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{2, 7, 11, 15},
		Target: 9,
		Answer: []int{1, 2},
	})
}

// 老办法，使用map记录出现过的
func twoSum(numbers []int, target int) []int {
	tmap := make(map[int]int)

	for i := range numbers {
		tmp := target - numbers[i]
		if j, ok := tmap[tmp]; ok {
			return []int{j, i + 1}
		}
		tmap[numbers[i]] = i + 1
	}
	return []int{}
}

// 最优解：双指针
// 由于一定存在唯一解，可用前后两个指针循环
func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tmp := numbers[left] + numbers[right]
		if tmp > target {
			right--
		} else if tmp < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

// 方法三：暴力法，两个循环，可用二分查找来优化循环
