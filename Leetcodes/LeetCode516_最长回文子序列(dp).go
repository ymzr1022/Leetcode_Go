package leetcodes

import "fmt"

/**
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。



示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。


提示：

1 <= s.length <= 1000
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func LongestPalindromeSubseq(s string) int {

	n := len(s)
	var process func(int, int) int
	process = func(l, r int) int {
		if l == r {
			return 1
		}
		if l > r {
			return 0
		}
		if s[l] == s[r] {
			return process(l+1, r-1) + 2
		}
		return max(process(l+1, r), process(l, r-1))
	}

	fmt.Printf("process(0, n-1): %v\n", process(0, n-1))
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Printf("dp: %v\n", dp[0][n-1])
	return 0
}
