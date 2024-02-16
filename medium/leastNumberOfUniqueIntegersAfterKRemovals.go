package medium

import (
	"container/heap"
)

// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := map[int]int{}
	for _, v := range arr {
		cnt[v] += 1
	}
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range cnt {
		heap.Push(h, v)
	}
	for k > 0 {
		k -= heap.Pop(h).(int)
		if k < 0 {
			return h.Len() + 1
		}
	}
	return h.Len()
}
