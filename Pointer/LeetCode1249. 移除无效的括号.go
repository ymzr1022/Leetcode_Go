package pointer

import "fmt"

/*
给你一个由 '('、')' 和小写字母组成的字符串 s。

你需要从字符串中删除最少数目的 '(' 或者 ')' （可以删除任意位置的括号)，使得剩下的「括号字符串」有效。

请返回任意一个合法字符串。

有效「括号字符串」应当符合以下 任意一条 要求：

空字符串或只包含小写字母的字符串
可以被写作 AB（A 连接 B）的字符串，其中 A 和 B 都是有效「括号字符串」
可以被写作 (A) 的字符串，其中 A 是一个有效的「括号字符串」


示例 1：

输入：s = "lee(t(c)o)de)"
输出："lee(t(c)o)de"
解释："lee(t(co)de)" , "lee(t(c)ode)" 也是一个可行答案。
示例 2：

输入：s = "a)b(c)d"
输出："ab(c)d"
示例 3：

输入：s = "))(("
输出：""
解释：空字符串也是有效的
*/

func minRemoveToMakeValid(s string) (res string) {
	del := make([]int, 0)
	cntl := 0
	lIndex := make([]int, 0)

	for i, v := range s {
		if v == ')' {
			if cntl == 0 {
				del = append(del, i)
			} else {
				cntl--
				lIndex = lIndex[1:]
			}
		}
		if v == '(' {
			cntl++
			lIndex = append(lIndex, i)
		}

	}
	if cntl > 0 {
		del = append(del, lIndex[:cntl]...)
		lIndex = lIndex[cntl:]
	}
	fmt.Printf("del: %v\n", del)
	if len(del) == 0 {
		return s
	}
	d := 0
	for i, v := range s {
		if i == del[d] {
			if d+1 < len(del) {
				d++
			}
			continue
		}
		res += string(v)
	}
	return
}
