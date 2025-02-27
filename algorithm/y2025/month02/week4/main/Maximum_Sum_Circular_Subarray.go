package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	max_sum := math.MinInt
	min_sum := 0
	var maxF, minF, sum int
	for _, x := range nums {
		maxF = max(maxF, 0) + x
		max_sum = max(max_sum, maxF)

		minF = min(minF, 0) + x
		min_sum = min(min_sum, minF)

		sum += x
	}

	if sum == min_sum {
		return max_sum
	}
	return max(max_sum, sum-min_sum)
}

func main() {

}
