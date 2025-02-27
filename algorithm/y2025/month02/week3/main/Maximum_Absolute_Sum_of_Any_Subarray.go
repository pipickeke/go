package main

func maxAbsoluteSum(nums []int) int {
	var s, mx, mn int
	for _, x := range nums {
		s += x
		if s > mx {
			mx = s
		} else if s < mn {
			mn = s
		}
	}
	return mx - mn
}
