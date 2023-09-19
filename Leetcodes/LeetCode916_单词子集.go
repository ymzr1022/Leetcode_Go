package leetcodes

import "strings"

/**
给你两个字符串数组 words1 和 words2。

现在，如果 b 中的每个字母都出现在 a 中，包括重复出现的字母，那么称字符串 b 是字符串 a 的 子集 。

例如，"wrr" 是 "warrior" 的子集，但不是 "world" 的子集。
如果对 words2 中的每一个单词 b，b 都是 a 的子集，那么我们称 words1 中的单词 a 是 通用单词 。

以数组形式返回 words1 中所有的通用单词。你可以按 任意顺序 返回答案。



示例 1：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
输出：["facebook","google","leetcode"]
示例 2：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
输出：["apple","google","leetcode"]
示例 3：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","oo"]
输出：["facebook","google"]
示例 4：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["lo","eo"]
输出：["google","leetcode"]
示例 5：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["ec","oc","ceo"]
输出：["facebook","leetcode"]


提示：

1 <= words1.length, words2.length <= 104
1 <= words1[i].length, words2[i].length <= 10
words1[i] 和 words2[i] 仅由小写英文字母组成
words1 中的所有字符串 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/word-subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func WordSubsets(words1 []string, words2 []string) (res []string) {
	map2 := make(map[string]int)
	map3 := make(map[string]int)
	var process2 func(string)
	process2 = func(s string) {
		for i := 0; i < len(s); i++ {
			map2[string(s[i])] = map2[string(s[i])] + 1
		}
		for key, v := range map2 {
			if map3[key] < v {
				map3[key] = v
			}
		}
	}

	for i := 0; i < len(words2); i++ {
		map2 = make(map[string]int)
		process2(words2[i])
	}
	for i := 0; i < len(words1); i++ {
		flg := true
		for key, v := range map3 {
			if v > strings.Count(words1[i], key) {
				flg = false
				break
			}
		}
		if flg {
			res = append(res, words1[i])
		}
	}

	return
}
