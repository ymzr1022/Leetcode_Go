package contest

import (
	"fmt"
	"sort"
)

/*
给你一个字符串 s 和一个整数 k 。

k 子序列指的是 s 的一个长度为 k 的 子序列 ，且所有字符都是 唯一 的，也就是说每个字符在子序列里只出现过一次。

定义 f(c) 为字符 c 在 s 中出现的次数。

k 子序列的 美丽值 定义为这个子序列中每一个字符 c 的 f(c) 之 和 。

比方说，s = "abbbdd" 和 k = 2 ，我们有：

f('a') = 1, f('b') = 3, f('d') = 2
s 的部分 k 子序列为：
"abbbdd" -> "ab" ，美丽值为 f('a') + f('b') = 4
"abbbdd" -> "ad" ，美丽值为 f('a') + f('d') = 3
"abbbdd" -> "bd" ，美丽值为 f('b') + f('d') = 5
请你返回一个整数，表示所有 k 子序列 里面 美丽值 是 最大值 的子序列数目。由于答案可能很大，将结果对 109 + 7 取余后返回。

一个字符串的子序列指的是从原字符串里面删除一些字符（也可能一个字符也不删除），不改变剩下字符顺序连接得到的新字符串。

注意：

f(c) 指的是字符 c 在字符串 s 的出现次数，不是在 k 子序列里的出现次数。
两个 k 子序列如果有任何一个字符在原字符串中的下标不同，则它们是两个不同的子序列。所以两个不同的 k 子序列可能产生相同的字符串。


示例 1：

输入：s = "bcca", k = 2
输出：4
解释：s 中我们有 f('a') = 1 ，f('b') = 1 和 f('c') = 2 。
s 的 k 子序列为：
bcca ，美丽值为 f('b') + f('c') = 3
bcca ，美丽值为 f('b') + f('c') = 3
bcca ，美丽值为 f('b') + f('a') = 2
bcca ，美丽值为 f('c') + f('a') = 3
bcca ，美丽值为 f('c') + f('a') = 3
总共有 4 个 k 子序列美丽值为最大值 3 。
所以答案为 4 。
示例 2：

输入：s = "abbcd", k = 4
输出：2
解释：s 中我们有 f('a') = 1 ，f('b') = 2 ，f('c') = 1 和 f('d') = 1 。
s 的 k 子序列为：
abbcd ，美丽值为 f('a') + f('b') + f('c') + f('d') = 5
abbcd ，美丽值为 f('a') + f('b') + f('c') + f('d') = 5
总共有 2 个 k 子序列美丽值为最大值 5 。
所以答案为 2 。


提示：

1 <= s.length <= 2 * 105
1 <= k <= s.length
s 只包含小写英文字母。
*/

func CountKSubsequencesWithMaxBeauty1(s string, k int) int {
	ind := make([][2]int, 26)
	for i := range s {
		index := s[i] - 'a'
		ind[index][0] = int(index)
		ind[index][1]++
	}
	sort.Slice(ind, func(i, j int) bool { return ind[i][1] > ind[j][1] })
	fmt.Printf("ind: %v\n", ind)

	mp := make(map[int]int)
	for i := 0; i < 26; i++ {
		if i < k {
			mp[ind[i][0]] = ind[i][1]
		}
		if i >= k && ind[k-1][1] == ind[i][1] {
			mp[ind[i][0]] = ind[i][1]
		}
	}
	fmt.Printf("mp: %v\n", mp)
	xx := make([]byte, 0)
	for i := range s {
		index := s[i] - 'a'
		if _, ok := mp[int(index)]; ok {
			xx = append(xx, index+97)
		}
	}
	s = string(xx)

	fmt.Printf("s: %v\n", s)
	n := len(s)
	set := make(map[int]struct{})
	var process func(int) int
	process = func(i int) int {
		if len(set) == k {
			mm := make([]int, 0)
			for k := range set {
				mm = append(mm, k)
			}
			sort.Slice(mm, func(i, j int) bool { return mm[i] > mm[j] })
			flg := false
			for i := range mm {
				if ind[i][0] != mm[i] {
					flg = false
				}
			}
			if flg {
				return 1
			}
		}
		if i >= n {
			return 0
		}
		index := int(s[i] - 'a')
		if _, ok := set[index]; ok {
			return process(i + 1)
		} else {
			nget := process(i + 1)
			set[index] = struct{}{}
			get := process(i + 1)
			delete(set, index)
			return nget + get
		}
	}
	return process(0)
}
