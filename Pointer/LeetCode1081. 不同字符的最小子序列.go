package pointer

/*
返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"


提示：

1 <= s.length <= 1000
s 由小写英文字母组成


注意：该题与 316 https://leetcode.cn/problems/remove-duplicate-letters/ 相同
*/

func SmallestSubsequence(s string) string {
	count := make([]int, 26)
	// times := 0
	for _, v := range s {
		count[int(v-'a')]++
	}
	res := make([]byte, 0)
	in := make([]bool, 26)
	for _, v := range s {
		ind := byte(v - 'a')
		if !in[ind] {
			in[ind] = true
			for len(res) > 0 && count[int(res[len(res)-1])] > 0 && res[len(res)-1] > ind {
				in[res[len(res)-1]] = false
				res = res[:len(res)-1]

			}
			res = append(res, ind)
		}
		count[int(ind)]--

	}
	ans := ""
	for _, v := range res {
		ans += string(v + 'a')
	}
	return ans
}
