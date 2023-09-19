package trace

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/

func letterCombinations(digits string) (res []string) {
	strs := [8][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
	path := make([]byte, 0)
	var process func(int)
	process = func(i int) {
		if i == len(digits) {
			tmp := make([]byte, len(digits))
			copy(tmp, path)
			res = append(res, string(tmp))
			return
		}
		index := int(digits[i] - '2')
		for _, v := range strs[index] {
			path = append(path, []byte(v)...)
			process(i + 1)
			path = path[:len(path)-1]
		}
	}

	process(0)
	return
}
