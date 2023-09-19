package leetcodes

import "math"

/**
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func NumSquares(n int) int {
	// f := int(math.Sqrt(float64(n)))

	// //------------------递归-----------------------
	// var process func(int, int) int
	// process = func(index, target int) int {
	// 	if index <= 0 {
	// 		if target == 0 {
	// 			return 0
	// 		} else {
	// 			return math.MaxInt32
	// 		}
	// 	}
	// 	if pow(index) > target {
	// 		return process(index-1, target)
	// 	}
	// 	return min(process(index-1, target), process(index, target-pow(index))+1)
	// }
	// return process(f, n)

	//---------------------递推------------------------
	// dp := make([][]int, f+2)
	// for i := range dp {
	// 	dp[i] = make([]int, n+1)
	// 	for j := range dp[i] {
	// 		dp[i][j] = math.MaxInt32
	// 	}
	// }
	// dp[0][0] = 0

	// for i := 0; i <= f; i++ {
	// 	for j := 0; j <= n; j++ {
	// 		if pow(i) > j {
	// 			dp[i+1][j] = dp[i][j]
	// 		} else {
	// 			dp[i+1][j] = min(dp[i][j], dp[i+1][j-pow(i)]+1)
	// 		}
	// 	}
	// }

	//--------------------------优化-----------------------
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; pow(i) <= n; i++ {
		for j := pow(i); j <= n; j++ {
			dp[j] = min(dp[j], dp[j-pow(i)]+1)
		}
	}
	return dp[n]
}

func pow(a int) int {
	return a * a
}
