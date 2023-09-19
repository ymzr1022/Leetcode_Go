package dp

import (
	"strings"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func LengthOfLongestSubstring(s string) (res int) {

	n := len(s)
	if n == 0 {
		return
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		if s[0] != s[1] {
			return 2
		}
		return 1
	}
	// var process func(int) int
	// process = func(i int) int {
	// 	if i < 0 {
	// 		return 0
	// 	}
	// 	if i == 0 {
	// 		res = max(res, 1)
	// 		return 1
	// 	}
	// 	lens := process(i - 1)
	// 	start := i - lens
	// 	str := s[start:i]
	// 	if strings.Contains(str, string(s[i])) {
	// 		res = max(res, 1)
	// 		return 1
	// 	}
	// 	res = max(res, lens+1)
	// 	return lens + 1
	// }

	// process(n - 1)
	i := 0
	j := 1
	len := 1
	res = max(res, len)
	for i < n && j < n {
		if strings.Contains(s[i:j], string(s[j])) {
			i++
			len = j - i
			res = max(res, len)
		} else {
			len++
			res = max(res, len)
			j++
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
