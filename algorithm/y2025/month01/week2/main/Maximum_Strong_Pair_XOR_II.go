package main

import "slices"

/**
题目：2935. 找出强数对的最大异或值 II
标签：字典树

给你一个下标从 0 开始的整数数组 nums 。如果一对整数 x 和 y 满足以下条件，则称其为 强数对 ：

|x - y| <= min(x, y)
你需要从 nums 中选出两个整数，且满足：这两个整数可以形成一个强数对，
并且它们的按位异或（XOR）值是在该数组所有强数对中的 最大值 。

返回数组 nums 所有可能的强数对中的 最大 异或值。

注意，你可以选择同一个整数两次来形成一个强数对。


*/

const highbit = 19

type node struct {
	children [2]*node
	cnt      int
}

type trie struct {
	root *node
}

func (t trie) insert(val int) *node {
	cur := t.root
	for i := highbit; i >= 0; i-- {
		bit := (val >> i) & 1
		if cur.children[bit] == nil {
			cur.children[bit] = &node{}
		}
		cur = cur.children[bit]
		cur.cnt++
	}
	return cur
}

func (t trie) remove(val int) *node {
	cur := t.root
	for i := highbit; i >= 0; i-- {
		cur = cur.children[(val>>i)&1]
		cur.cnt--
	}
	return cur
}

func (t trie) maxxor(val int) int {
	cur := t.root
	ans := 0
	for i := highbit; i >= 0; i-- {
		bit := (val >> i) & 1
		if cur.children[bit^1] != nil && cur.children[bit^1].cnt > 0 {
			ans |= (1 << i)
			bit ^= 1
		}
		cur = cur.children[bit]
	}
	return ans
}

func maximumStrongPairXor(nums []int) int {
	slices.Sort(nums)
	t := trie{&node{}}
	left := 0
	ans := 0
	for _, y := range nums {
		t.insert(y)
		for nums[left]*2 < y {
			t.remove(nums[left])
			left++
		}
		ans = max(ans, t.maxxor(y))
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	print(maximumStrongPairXor(nums))

}
