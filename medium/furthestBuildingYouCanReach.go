package medium

import (
	"container/heap"
)

// https://leetcode.com/problems/furthest-building-you-can-reach
func furthestBuilding(heights []int, bricks int, ladders int) int {
	d := &IntHeap{}
	for i := 1; i < len(heights); i++ {
		h := heights[i] - heights[i-1]
		if h <= 0 {
			continue
		}
		heap.Push(d, h)
		if d.Len() > ladders {
			bricks -= heap.Pop(d).(int)
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}
