package dp

/**
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。



示例 1:

输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
示例 2:

输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。


提示:

0 <= s1.length, s2.length <= 1000
s1 和 s2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MinimumDeleteSum(s1 string, s2 string) (res int) {
	n1 := len(s1)
	n2 := len(s2)
	for i := 0; i < n1; i++ {
		res += int(s1[i])
	}
	for i := 0; i < n2; i++ {
		res += int(s2[i])
	}
	// var process func(int, int) int
	// process = func(i1, i2 int) int {
	// 	if i1 < 0 || i2 < 0 {
	// 		return 0
	// 	}
	// 	if s1[i1] == s2[i2] {
	// 		return process(i1-1, i2-1) + int(s1[i1])
	// 	}
	// 	return max(process(i1-1, i2), process(i1, i2-1))
	// }
	dp := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		pre := 0
		for j := 0; j < n2; j++ {
			tmp := dp[i+1]
			if s1[i] == s2[j] {
				dp[j+1] = pre + int(s1[i])
			} else {
				dp[j+1] = max(dp[j+1], dp[j])
			}
			pre = tmp
		}
	}
	res -= dp[n2] * 2
	return
}
