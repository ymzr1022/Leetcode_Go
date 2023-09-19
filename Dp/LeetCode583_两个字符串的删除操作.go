package dp

/**

给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。

每步 可以删除任意一个字符串中的一个字符。



示例 1：

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
示例  2:

输入：word1 = "leetcode", word2 = "etco"
输出：4


提示：

1 <= word1.length, word2.length <= 500
word1 和 word2 只包含小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-operation-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minDistance(word1 string, word2 string) (res int) {
	n1, n2 := len(word1), len(word2)
	// var process func(int, int) int
	// process = func(i, j int) int {
	// 	if i < 0 || j < 0 {
	// 		return 0
	// 	}
	// 	if word1[i] == word2[j] {
	// 		return process(i-1, j-1) + 1
	// 	}
	// 	return Max(process(i-1, j), process(i, j-1))
	// }
	// res = process(n1-1, n2-1)

	dp := make([]int, n1+1)

	for i := 0; i < n1; i++ {
		pre := 0
		for j := 0; j < n2; j++ {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = pre + 1
			} else {
				dp[j+1] = Max(dp[j+1], dp[j])
			}
			pre = tmp
		}
	}
	// fmt.Printf("dp: %v\n", dp)
	res = dp[n2]

	res = n1 + n2 - res - res
	return
}
