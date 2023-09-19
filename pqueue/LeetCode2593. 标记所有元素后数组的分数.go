package pqueue

import "container/heap"

/*
给你一个数组 nums ，它包含若干正整数。

一开始分数 score = 0 ，请你按照下面算法求出最后分数：

从数组中选择最小且没有被标记的整数。如果有相等元素，选择下标最小的一个。
将选中的整数加到 score 中。
标记 被选中元素，如果有相邻元素，则同时标记 与它相邻的两个元素 。
重复此过程直到数组中所有元素都被标记。
请你返回执行上述算法后最后的分数。



示例 1：

输入：nums = [2,1,3,4,5,2]
输出：7
解释：我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[2,1,3,4,5,2] 。
- 2 是最小未标记元素，所以标记它和左边相邻元素：[2,1,3,4,5,2] 。
- 4 是仅剩唯一未标记的元素，所以我们标记它：[2,1,3,4,5,2] 。
总得分为 1 + 2 + 4 = 7 。
示例 2：

输入：nums = [2,3,5,1,3,2]
输出：5
解释：我们按照如下步骤标记元素：
- 1 是最小未标记元素，所以标记它和相邻两个元素：[2,3,5,1,3,2] 。
- 2 是最小未标记元素，由于有两个 2 ，我们选择最左边的一个 2 ，也就是下标为 0 处的 2 ，以及它右边相邻的元素：[2,3,5,1,3,2] 。
- 2 是仅剩唯一未标记的元素，所以我们标记它：[2,3,5,1,3,2] 。
总得分为 1 + 2 + 2 = 5 。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 106
*/
type pair struct {
	num, index int
}

type hp struct {
	paris []*pair
}

func (h hp) Len() int { return len(h.paris) }
func (h hp) Less(i, j int) bool {
	return h.paris[i].num < h.paris[j].num || (h.paris[i].num == h.paris[j].num && h.paris[i].index < h.paris[j].index)
}
func (h hp) Swap(i, j int) { h.paris[i], h.paris[j] = h.paris[j], h.paris[i] }

func (h *hp) Push(x any) { h.paris = append(h.paris, x.(*pair)) }
func (h *hp) Pop() any   { m := h.paris; v := m[h.Len()-1]; h.paris = m[:h.Len()-1]; return v }

func findScore(nums []int) (res int64) {
	h := hp{}
	visit := make([]bool, len(nums))
	for i, v := range nums {
		heap.Push(&h, &pair{
			v, i,
		})
	}
	for h.Len() > 0 {
		node := heap.Pop(&h).(*pair)
		if !visit[node.index] {
			res += int64(node.num)
			if node.index-1 >= 0 {
				visit[node.index-1] = true
			}
			if node.index+1 < len(nums) {
				visit[node.index+1] = true
			}
			visit[node.index] = true
		}
	}
	return
}
