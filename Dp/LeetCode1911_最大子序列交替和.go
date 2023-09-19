package dp

import "fmt"

/**
一个下标从 0 开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和 。

比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4 。
给你一个数组 nums ，请你返回 nums 中任意子序列的 最大交替和 （子序列的下标 重新 从 0 开始编号）。

一个数组的 子序列 是从原数组中删除一些元素后（也可能一个也不删除）剩余元素不改变顺序组成的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4] 的一个子序列（加粗元素），但是 [2,4,2] 不是。



示例 1：

输入：nums = [4,2,5,3]
输出：7
解释：最优子序列为 [4,2,5] ，交替和为 (4 + 5) - 2 = 7 。
示例 2：

输入：nums = [5,6,7,8]
输出：8
解释：最优子序列为 [8] ，交替和为 8 。
示例 3：

输入：nums = [6,2,1,2,4,5]
输出：10
解释：最优子序列为 [6,1,5] ，交替和为 (6 + 5) - 1 = 10 。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-alternating-subsequence-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MaxAlternatingSum(nums []int) int64 {
	n := len(nums)

	var process func(int, int) int
	process = func(i1, i2 int) int {
		if i1 < 0 {
			return 0
		}
		if i2 == 0 {
			return max(process(i1-1, 1)+nums[i1], process(i1-1, 0))
		} else {
			return process(i1-1, 0) - nums[i1]
		}
	}
	fmt.Printf("process(n-1, 0): %v\n", process(n-1, 0))
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, 2)
	}
	for i := 0; i < n; i++ {
		dp[i+1][0] = maxx(dp[i][1]+int64(nums[i]), dp[i][0])
		dp[i+1][1] = maxx(dp[i][0]-int64(nums[i]), dp[i][1])
	}
	fmt.Printf("dp: %v\n", dp)
	var dp0 int64 = -1 << 31
	var dp1 int64 = 0
	for _, v := range nums {
		dp0, dp1 = maxx(dp1+int64(v), dp0), dp0-int64(v)
	}

	var dp00 int64 = 0
	var dp10 int64 = 0
	for _, v := range nums {
		dp00, dp10 = maxx(dp10+int64(v), dp00), maxx(dp00-int64(v), dp10)
	}
	fmt.Printf("dp00: %v\n", dp00)
	fmt.Printf("dp10: %v\n", dp10)
	// fmt.Printf("dp: %v\n", dp)
	return maxx(dp0, dp1)
}

func maxx(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
