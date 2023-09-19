package leetcodes

import "fmt"

/**
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。



示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1


提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	if target < 0 {
		target *= -1
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if (sum+target)%2 != 0 {
		return 0
	}
	aim := (sum + target) / 2

	// ----------------------------递归--------------------------------
	// var process func(int, int) int
	// process = func(index, a int) int {
	// 	if index < 0 {
	// 		if a == 0 {
	// 			return 1
	// 		} else {
	// 			return 0
	// 		}
	// 	}
	// 	if a < nums[index] {
	// 		return process(index-1, a)
	// 	}
	// 	return process(index-1, a) + process(index-1, a-nums[index])
	// }
	// return process(n-1, aim)

	//---------------------------------递推------------------------------
	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = make([]int, aim+1)
	// }
	// dp[0][0] = 1
	// for i, x := range nums {
	// 	for j := 0; j <= aim; j++ {
	// 		if j < x {
	// 			dp[i+1][j] = dp[i][j]
	// 		} else {
	// 			dp[i+1][j] = dp[i][j] + dp[i][j-x]
	// 		}
	// 	}
	// }
	// return dp[n][aim]

	//--------------------------------------优化------------------------------
	dp := make([]int, aim+1)
	dp[0] = 1
	for _, x := range nums {
		for j := aim; j >= x; j-- {
			dp[j] = dp[j] + dp[j-x]
		}
	}
	fmt.Printf("dp: %v\n", dp)
	return dp[aim]
}
