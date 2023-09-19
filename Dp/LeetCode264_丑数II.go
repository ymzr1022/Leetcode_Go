package dp

/**
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ugly-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func NthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp2, dp3, dp5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(min(2*dp[dp2], 3*dp[dp3]), 5*dp[dp5])
		if dp[i] == 2*dp[dp2] {
			dp2++
		}
		if dp[i] == 3*dp[dp3] {
			dp3++
		}
		if dp[i] == 5*dp[dp5] {
			dp5++
		}
	}
	return dp[n-1]
}
