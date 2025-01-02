package main

func totalHammingDistance(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}

	ans := 0
	for xor != 0 {
		if xor&1 == 1 {
			ans++
		}
		xor >>= 1
	}
	return ans
}
