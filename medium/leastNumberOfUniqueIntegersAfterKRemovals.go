package medium

import (
	"container/heap"
)

// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals

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
