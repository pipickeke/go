package main

func main() {

}

func countPairs(nums []int, low, high int) (ans int) {
	t := &trie{&trieNode{}}
	t.put(nums[0])
	for _, v := range nums[1:] {
		ans += t.countLimit(v, high+1) - t.countLimit(v, low)
		t.put(v)
	}
	return
}

const HIGH_BIT = 14

type trieNode struct {
	son [2]*trieNode
	cnt int
}

type trie struct {
	root *trieNode
}

func (t *trie) put(v int) *trieNode {
	o := t.root
	for i := HIGH_BIT; i >= 0; i-- {
		b := (v >> i) & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie) countLimit(v, limit int) (cnt int) {
	o := t.root
	for i := HIGH_BIT; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				cnt += o.son[b].cnt
			}
			b ^= 1
		}
		if o.son[b] == nil {
			return
		}
		o = o.son[b]
	}
	return
}
