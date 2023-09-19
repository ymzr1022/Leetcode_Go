package dp

/**
子数组是以0下标开始的数组的连续非空子序列，从 i 到 j（0 <= i <= j < nums.length）的 子数组交替和 被定义为 nums[i] - nums[i+1] + nums[i+2] - ... +/- nums[j] 。

给定一个以0下标开始的整数数组nums，返回它所有可能的交替子数组和的最大值。



示例 1：

输入：nums = [3,-1,1,2]
输出：5
解释：
子数组 [3,-1,1]有最大的交替子数组和3 - (-1) + 1 = 5.
示例 2：

输入：nums = [2,2,2,2,2]
输出：2
解释：
子数组 [2], [2,2,2]和 [2,2,2,2,2]有相同的最大交替子数组和为2
[2]: 2.
[2,2,2]: 2 - 2 + 2 = 2.
[2,2,2,2,2]: 2 - 2 + 2 - 2 + 2 = 2.
示例 3：

输入：nums = [1]
输出：1
解释：
仅有一个非空子数组，为 [1]，它的交替子数组和为 1


提示：

1 <= nums.length <= 105
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-alternating-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func AlternatingSubarrays(nums []int) (res int64) {
	// n := len(nums)

	// var process func(int, int) int
	// process = func(i1, i2 int) int {
	// 	if i1 >= n {
	// 		return 0
	// 	}
	// 	if i2 == 0 {
	// 		return max(process(i1+1, 1)+nums[i1], nums[i1])
	// 	} else {
	// 		return max(process(i1+1, 0)-nums[i1], 0)
	// 	}
	// }
	// dp := make([][]int64, n)
	// for i := range dp {
	// 	dp[i] = make([]int64, 2)
	// 	dp[i][0] = -1 << 31
	// 	dp[i][1] = -1 << 31
	// }
	// dp[0][0] = int64(nums[0])
	// dp[0][1] = 0
	// ress := int64(nums[0])
	// for i := 1; i < n; i++ {
	// 	dp[i][0] = dp[i-1][1] - int64(nums[i])
	// 	dp[i][1] = maxx(dp[i-1][0]+int64(nums[i]), int64(nums[i]))
	// 	ress = maxx(ress, maxx(dp[i][0], dp[i][1]))
	// }
	// fmt.Printf("dp: %v\n", dp)
	// fmt.Printf("process(n-1, 0): %v\n", process(0, 0))
	var dp0 int64 = int64(nums[0])
	var dp1 int64 = 0
	for i := 0; i < len(nums); i++ {
		dp0, dp1 = maxx(dp1+int64(nums[i]), int64(nums[i])), dp0-int64(nums[i])
		res = maxx(res, maxx(dp0, dp1))
	}

	return
}
