package stack

import (
	"fmt"
	"strconv"
)

/*
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。


示例 1 ：

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
示例 2 ：

输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 ：

输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。


提示：

1 <= k <= num.length <= 105
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零
*/

func removeKdigits(num string, k int) string {
	stack := make([]int, 0)
	str := ""
	tmp := k
	// for i := range num {
	// 	var ind = num[i] - '0'
	// 	for len(stack) > 0 && k > 0 && num[stack[len(stack)-1]]-'0' > ind {
	// 		stack = stack[:len(stack)-1]
	// 		k--
	// 	}
	// 	stack = append(stack, i)
	// }
	for i := range num {
		var ind = num[i] - '0'
		for len(str) > 0 && k > 0 && str[len(str)-1]-'0' > ind {
			str = str[:len(str)-1]
			k--
		}
		str = str + string(num[i])
	}
	fmt.Printf("stack: %v\n", str)
	if len(str) <= len(num) {
		str = str[:len(stack)-tmp+(len(num)-len(stack))]
	}
	for str[0]-'0' == 0 {
		str = str[1:]
	}
	if str == "" {
		return "0"
	}
	return str
}

func removeKdigits1(num string, k int) string {
	stack := make([]int, 0)
	tmp := k
	for i := range num {
		var ind = num[i] - '0'
		for len(stack) > 0 && k > 0 && num[stack[len(stack)-1]]-'0' > ind {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, i)
	}
	res := ""

	if len(stack) <= len(num) {
		stack = stack[:len(stack)-tmp+(len(num)-len(stack))]
	}
	for _, v := range stack {
		if res == "" && num[v]-'0' == 0 {
			continue
		}
		res += strconv.Itoa(int(num[v] - '0'))
	}
	if res == "" {
		return "0"
	}
	return res
}
