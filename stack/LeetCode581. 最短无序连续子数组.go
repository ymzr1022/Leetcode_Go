package stack

import (
	"fmt"
)

/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。



示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0


提示：

1 <= nums.length <= 104
-105 <= nums[i] <= 105


进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
*/

type heaps []int

func (h *heaps) Push(a int) {
	*h = append(*h, a)
}

func (h *heaps) Pop() int {
	tmp := make([]int, h.Len())
	copy(tmp, *h)
	x := tmp[h.Len()-1]
	*h = tmp[:h.Len()-1]
	return x
}

func (h heaps) Len() int {
	return len(h)
}

func (h heaps) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h heaps) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func FindUnsortedSubarray(nums []int) int {
	stack := make([]int, 0)
	ind := &heaps{}
	for i, v := range nums {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && v < nums[stack[len(stack)-1]] {
				// heap.Push(ind, i)
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}

	}
	fmt.Printf("stack: %v\n", stack)
	fmt.Printf("ind: %v\n", ind)
	maxs := -1
	mins := 1 << 30
	// for _, v := range ind {
	// 	if maxs < v {
	// 		maxs = v
	// 	}
	// 	if mins > v {
	// 		mins = v
	// 	}
	// }
	return maxs - mins + 1
}
