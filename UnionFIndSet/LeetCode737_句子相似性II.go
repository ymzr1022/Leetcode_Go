package unionfindset

/**
我们可以将一个句子表示为一个单词数组，例如，句子 I am happy with leetcode"可以表示为 arr = ["I","am",happy","with","leetcode"]

给定两个句子 sentence1 和 sentence2 分别表示为一个字符串数组，并给定一个字符串对 similarPairs ，其中 similarPairs[i] = [xi, yi] 表示两个单词 xi 和 yi 是相似的。

如果 sentence1 和 sentence2 相似则返回 true ，如果不相似则返回 false 。

两个句子是相似的，如果:

它们具有 相同的长度 (即相同的字数)
sentence1[i] 和 sentence2[i] 是相似的
请注意，一个词总是与它自己相似，也请注意，相似关系是可传递的。例如，如果单词 a 和 b 是相似的，单词 b 和 c 也是相似的，那么 a 和 c 也是 相似 的。



示例 1:

输入: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","good"],["fine","good"],["drama","acting"],["skills","talent"]]
输出: true
解释: 这两个句子长度相同，每个单词都相似。
示例 2:

输入: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","onepiece"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
输出: true
解释: "leetcode" --> "platform" --> "anime" --> "manga" --> "onepiece".
因为“leetcode”和“onepiece”相似，而且前两个单词是相同的，所以这两句话是相似的。
示例 3:

输入: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","hunterXhunter"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
输出: false
解释: “leetcode”和“onepiece”不相似。


提示:

1 <= sentence1.length, sentence2.length <= 1000
1 <= sentence1[i].length, sentence2[i].length <= 20
sentence1[i] 和 sentence2[i] 只包含大小写英文字母
0 <= similarPairs.length <= 2000
similarPairs[i].length == 2
1 <= xi.length, yi.length <= 20
xi 和 yi 只含英文字母

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sentence-similarity-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var ufmap map[string]string

func find(x string) string {
	if ufmap[x] != x {
		ufmap[x] = find(ufmap[x])
	}
	return ufmap[x]
}

func query(a, b string) bool {
	return ufmap[find(a)] == ufmap[find(b)]
}
func union(a, b string) {
	ufmap[find(a)] = ufmap[find(b)]
}

func AreSentencesSimilarTwo(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	ufmap = make(map[string]string)
	for _, v := range sentence1 {
		ufmap[v] = v
	}
	for _, v := range sentence2 {
		ufmap[v] = v
	}
	for _, v := range similarPairs {
		if ufmap[v[0]] == "" {
			ufmap[v[0]] = v[0]
		}
		if ufmap[v[1]] == "" {
			ufmap[v[1]] = v[1]
		}
		union(v[0], v[1])
	}
	n := len(sentence1)

	for i := 0; i < n; i++ {
		if !query(sentence1[i], sentence2[i]) {
			return false
		}
	}
	return true
}
