package pointer

import "strings"

/*
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

如果不存在这样的子字符串，则返回 0。



示例 1：

输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2：

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。


提示：

1 <= s.length <= 104
s 仅由小写英文字母组成
1 <= k <= 105
*/

func longestSubstring(s string, k int) (res int) {
	mp := make(map[string]int)
	for _, v := range s {
		mp[string(v)]++
	}
	for kk, v := range mp {
		if v < k {
			for _, v := range strings.Split(s, kk) {
				res = max(longestSubstring(v, k), res)
			}
		}
	}

	return
}
