package leetcodes

import "fmt"

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func coinChange(coins []int, amount int) int {

	n := len(coins)
	var process func(int, int) int
	process = func(index, target int) int {
		if index < 0 {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}
		if coins[index] > target {
			return process(index-1, target)
		}
		return (process(index-1, target) + process(index, target-coins[index]))
	}

	dp := make([]int, amount+1)
	// for i := range dp {
	// 	dp[i] = make([]int, amount+1)
	// 	dp[i][amount] = 1
	// }
	dp[0] = 1
	for _, x := range coins {
		for j := x; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-x]
		}
	}
	fmt.Printf("process(0, amount): %v\n", process(n-1, amount))
	fmt.Printf("dp: %v\n", dp)
	return dp[amount]
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
