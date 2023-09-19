package contest

import (
	"fmt"
	"sort"
)

/*
给你一个下标从 0 开始长度为 n 的整数数组 nums 。

从 0 到 n - 1 的数字被分为编号从 1 到 3 的三个组，数字 i 属于组 nums[i] 。注意，有的组可能是 空的 。

你可以执行以下操作任意次：

选择数字 x 并改变它的组。更正式的，你可以将 nums[x] 改为数字 1 到 3 中的任意一个。
你将按照以下过程构建一个新的数组 res ：

将每个组中的数字分别排序。
将组 1 ，2 和 3 中的元素 依次 连接以得到 res 。
如果得到的 res 是 非递减顺序的，那么我们称数组 nums 是 美丽数组 。

请你返回将 nums 变为 美丽数组 需要的最少步数。



示例 1：
			2,2,2,2,1
输入：nums = [2,1,3,2,1]
			 1,4,0,3,2
			 4,0,1,2,3
输出：3
解释：以下三步操作是最优方案：
1. 将 nums[0] 变为 1 。
2. 将 nums[2] 变为 1 。
3. 将 nums[3] 变为 1 。
执行以上操作后，将每组中的数字排序，组 1 为 [0,1,2,3,4] ，组 2 和组 3 都为空。所以 res 等于 [0,1,2,3,4] ，它是非递减顺序的。
三步操作是最少需要的步数。
示例 2：

输入：nums = [1,3,2,1,3,3]
			0, 3,2,1,4,5
			1,2,2,1,2,1
			0,3,5,1,2,4
			0,4,3,2,1,5
输出：2
解释：以下两步操作是最优方案：
1. 将 nums[1] 变为 1 。
2. 将 nums[2] 变为 1 。
执行以上操作后，将每组中的数字排序，组 1 为 [0,1,2,3] ，组 2 为空，组 3 为 [4,5] 。所以 res 等于 [0,1,2,3,4,5] ，它是非递减顺序的。
两步操作是最少需要的步数。
示例 3：

输入：nums = [2,2,2,2,3,3]
输出：0
解释：不需要执行任何操作。
组 1 为空，组 2 为 [0,1,2,3] ，组 3 为 [4,5] 。所以 res 等于 [0,1,2,3,4,5] ，它是非递减顺序的。


提示：

1 <= nums.length <= 100
1 <= nums[i] <= 3
*/

func MinimumOperations(nums []int) (res int) {
	n := len(nums)

	tmp := make([][]int, n)
	for i := range nums {
		tmp[i] = make([]int, 2)
		tmp[i][0] = nums[i]
		tmp[i][1] = i
	}

	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][0] == tmp[j][0] {
			return tmp[i][1] < tmp[j][1]
		} else {
			return tmp[i][0] < tmp[j][0]
		}
	})
	fmt.Printf("tmp: %v\n", tmp)
	for i := 0; i < n; i++ {
		nums[i] = tmp[i][1]
	}
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 0; i < n; i++ {
		ans := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				ans = max(ans, dp[j])
			}
		}
		dp[i] = ans + 1
		res = max(dp[i], res)
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
