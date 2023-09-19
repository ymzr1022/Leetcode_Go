package contest

import (
	"fmt"
	"sort"
)

/*
给你一个下标从 0 开始的 非递减 整数数组 nums 。

你可以执行以下操作任意次：

选择 两个 下标 i 和 j ，满足 i < j 且 nums[i] < nums[j] 。
将 nums 中下标在 i 和 j 处的元素删除。剩余元素按照原来的顺序组成新的数组，下标也重新从 0 开始编号。
请你返回一个整数，表示执行以上操作任意次后（可以执行 0 次），nums 数组的 最小 数组长度。



示例 1：

输入：nums = [1,3,4,9]
输出：0
解释：一开始，nums = [1, 3, 4, 9] 。
第一次操作，我们选择下标 0 和 1 ，满足 nums[0] < nums[1] <=> 1 < 3 。
删除下标 0 和 1 处的元素，nums 变成 [4, 9] 。
下一次操作，我们选择下标 0 和 1 ，满足 nums[0] < nums[1] <=> 4 < 9 。
删除下标 0 和 1 处的元素，nums 变成空数组 [] 。
所以，可以得到的最小数组长度为 0 。
示例 2：

输入：nums = [2,3,6,9]
输出：0
解释：一开始，nums = [2, 3, 6, 9] 。
第一次操作，我们选择下标 0 和 2 ，满足 nums[0] < nums[2] <=> 2 < 6 。
删除下标 0 和 2 处的元素，nums 变成 [3, 9] 。
下一次操作，我们选择下标 0 和 1 ，满足 nums[0] < nums[1] <=> 3 < 9 。
删除下标 0 和 1 处的元素，nums 变成空数组 [] 。
所以，可以得到的最小数组长度为 0 。
示例 3：

输入：nums = [1,1,2]
输出：1
解释：一开始，nums = [1, 1, 2] 。
第一次操作，我们选择下标 0 和 2 ，满足 nums[0] < nums[2] <=> 1 < 2 。
删除下标 0 和 2 处的元素，nums 变成 [1] 。
无法对数组再执行操作。
所以，可以得到的最小数组长度为 1 。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 109
nums 是 非递减 数组。
*/

func MinLengthAfterRemovals(nums []int) int {
	mp := make(map[int]int, 0)
	for _, v := range nums {
		mp[v]++
	}
	tmp := make([]int, 0)
	for _, v := range mp {
		tmp = append(tmp, v)
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	fmt.Printf("tmp: %v\n", tmp)
	need := tmp[0]
	flg := false
	for i := 1; i < len(tmp); i++ {
		need = need - tmp[i]
		if need < 0 {
			flg = true
			need = -need
		}
	}
	if !flg {
		return need
	} else {
		return need % 2
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
