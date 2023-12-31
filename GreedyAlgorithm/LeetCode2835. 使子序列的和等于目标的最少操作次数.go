package greedyalgorithm

import (
	"fmt"
	"strconv"
)

/*
给你一个下标从 0 开始的数组 nums ，它包含 非负 整数，且全部为 2 的幂，同时给你一个整数 target 。

一次操作中，你必须对数组做以下修改：

选择数组中一个元素 nums[i] ，满足 nums[i] > 1 。
将 nums[i] 从数组中删除。
在 nums 的 末尾 添加 两个 数，值都为 nums[i] / 2 。
你的目标是让 nums 的一个 子序列 的元素和等于 target ，请你返回达成这一目标的 最少操作次数 。如果无法得到这样的子序列，请你返回 -1 。

数组中一个 子序列 是通过删除原数组中一些元素，并且不改变剩余元素顺序得到的剩余数组。



示例 1：

输入：nums = [1,2,8], target = 7
输出：1
解释：第一次操作中，我们选择元素 nums[2] 。数组变为 nums = [1,2,4,4] 。
这时候，nums 包含子序列 [1,2,4] ，和为 7 。
无法通过更少的操作得到和为 7 的子序列。
示例 2：

输入：nums = [1,32,1,2], target = 12
输出：2
解释：第一次操作中，我们选择元素 nums[1] 。数组变为 nums = [1,1,2,16,16] 。
第二次操作中，我们选择元素 nums[3] 。数组变为 nums = [1,1,2,16,8,8] 。
这时候，nums 包含子序列 [1,1,2,8] ，和为 12 。
无法通过更少的操作得到和为 12 的子序列。
示例 3：

输入：nums = [1,32,1], target = 35
输出：-1
解释：无法得到和为 35 的子序列。


提示：

1 <= nums.length <= 1000
1 <= nums[i] <= 230
nums 只包含非负整数，且均为 2 的幂。
1 <= target < 231
*/

func MinOperations(nums []int, target int) (ans int) {
	sums := 0
	ind := make([]int, 31)
	for _, v := range nums {
		ind[len(transfer(v))-1]++
		sums += v
	}
	if sums < target {
		return -1
	}
	fmt.Printf("ind: %v\n", ind)
	s := 0
	bt := transfer(target)
	fmt.Printf("bt: %v\n", bt)
	i := 0
	for 1<<i <= target {
		s += ind[i] * (1 << i)
		if target&(1<<i) == 0 {
			i++
			continue
		}
		mask := 1<<(i+1) - 1
		if s >= (mask & target) {
			i++
			continue
		} else {
			j := i + 1
			ans++
			for ind[j] == 0 {
				j++
				ans++
			}
			i = j
		}
	}
	return ans
}

func transfer(a int) string {
	res := ""
	for a > 0 {
		x := a % 2
		a /= 2
		res = res + strconv.Itoa(x)
	}
	return res
}
