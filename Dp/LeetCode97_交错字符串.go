package dp

/**
给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。



示例 1：


输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
示例 3：

输入：s1 = "", s2 = "", s3 = ""
输出：true


提示：

0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1、s2、和 s3 都由小写英文字母组成


进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/interleaving-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func IsInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1+n2 != len(s3) {
		return false
	}
	var process func(int, int, int) bool
	process = func(index1, index2, index3 int) bool {
		if index3 < 0 {
			return true
		}
		if index1 >= 0 && s1[index1] == s3[index3] {
			if process(index1-1, index2, index3-1) {
				return true
			}
		}
		if index2 >= 0 && s2[index2] == s3[index3] {
			if process(index1, index2-1, index3-1) {
				return true
			}
		}
		return false
	}
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)

	}
	dp[0][0] = true

	for i := 1; i <= n2; i++ {
		dp[0][i] = s2[:i] == s3[:i]
	}

	for i := 1; i <= n1; i++ {
		dp[i][0] = s1[:i] == s3[:i]
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			k := i + j - 1
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[k]) || (dp[i][j-1] && s2[j-1] == s3[k])
		}
	}
	return dp[n1][n2]
}
