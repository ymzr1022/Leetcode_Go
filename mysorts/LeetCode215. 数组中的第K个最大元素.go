package mysorts

import (
	"container/heap"
	"sort"
)

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4


提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
通过次数
922.2K
提交次数
1.5M
通过率
63.2%
*/

func FindKthLargest(nums []int, k int) int {
	var partition func(num []int, left, right int) int
	partition = func(num []int, left, right int) int {
		target := num[left]
		i, j := left, right
		for i < j {
			for i < j && num[j] >= target {
				j--
			}
			for i < j && num[i] <= target {
				i++
			}
			num[i], num[j] = num[j], num[i]
		}
		num[i], num[left] = num[left], num[i]
		return i
	}
	var topK func(num []int, k, left, right int)
	topK = func(num []int, k, left, right int) {
		if left < right {
			i := partition(nums, left, right)
			if i == k {
				return
			} else if i > k {
				topK(num, k, left, i-1)
			} else {
				topK(num, k, i+1, right)
			}
		}
	}
	topK(nums, k-1, 0, len(nums)-1)

	return nums[k-1]
}

func f2(nums []int, k int) int {
	h := hp{}
	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] < h.IntSlice[0] {
			h.replace(nums[i])
		}
	}
	return h.IntSlice[0]
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any   { m := h.IntSlice; v := m[h.Len()-1]; h.IntSlice = m[:h.Len()-1]; return v }
func (h *hp) replace(x any) {
	h.IntSlice[0] = x.(int)
	heap.Fix(h, 0)
}
