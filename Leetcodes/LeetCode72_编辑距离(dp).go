package leetcodes

/**
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	// var process func(int, int) int
	// process = func(i, j int) int {
	// 	if i < 0 {
	// 		return j + 1
	// 	}
	// 	if j < 0 {
	// 		return i + 1
	// 	}
	// 	if word1[i] == word2[j] {
	// 		return process(i-1, j-1)
	// 	}
	// 	return min(min(process(i-1, j-1), process(i-1, j)), process(i, j-1)) + 1
	// }

	// fmt.Printf("process(m-1, n-1): %v\n", process(m-1, n-1))

	//------------------------------递推-------------------------------------
	// dp := make([][]int, m+1)
	// for i := range dp {
	// 	dp[i] = make([]int, n+1)
	// 	dp[i][0] = i
	// }
	// for j := range dp[0] {
	// 	dp[0][j] = j
	// }

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if word1[i] == word2[j] {
	// 			dp[i+1][j+1] = dp[i][j]
	// 		} else {
	// 			dp[i+1][j+1] = min(min(dp[i][j], dp[i][j+1]), dp[i+1][j]) + 1
	// 		}
	// 	}
	// }
	// fmt.Printf("dp: %v\n", dp)
	// fmt.Printf("dp[m][n]: %v\n", dp[m][n])

	//----------------------优化--------------------------
	dp := make([]int, n+1)
	for j := 1; j <= m; j++ {
		dp[j] = j
	}
	for i := 0; i < m; i++ {
		pre := dp[0]
		dp[0]++
		for j := 0; j < n; j++ {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = pre
			} else {
				dp[j+1] = min(min(pre, dp[j+1]), dp[j]) + 1
			}
			pre = tmp

		}
	}
	// fmt.Printf("dp: %v\n", dp)
	return dp[n]
}
