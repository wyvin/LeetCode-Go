package _0190_Reverse_Bits

import (
	"strconv"
)

// 字符串操作
func reverseBits(num uint32) uint32 {
	numStr := strconv.FormatUint(uint64(num), 2)
	numB := []byte(numStr)
	for i := len(numB) - 1; i >= len(numB)/2; i-- {
		numB[i], numB[len(numB)-1-i] = numB[len(numB)-1-i], numB[i]
	}
	for len(numB) < 32 {
		numB = append(numB, '0')
	}
	numUint64, _ := strconv.ParseUint(string(numB), 2, 0)
	return uint32(numUint64)
}

// 位运算
// num & 1 可以求出最后一位的二进制：0&1=0；1&1=1
// num >>= 1 右移一位可以去掉最后一位二进制
func reverseBits2(num uint32) uint32 {
	ans := uint32(0)
	offset := 31
	for num != 0 {
		ans += (num & 1) << offset
		num >>= 1
		offset--
	}
	return ans
}

// 位运算 - 不循环
// 一共只有32位，每次折半互换位置
func reverseBits3(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}
