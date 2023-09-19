package dp

import "fmt"

/**
一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：

"AAJF" ，将消息分组为 (1 1 10 6)
"KJF" ，将消息分组为 (11 10 6)
注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。

给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。

题目数据保证答案肯定是一个 32 位 的整数。



示例 1：

输入：s = "12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2：

输入：s = "226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
示例 3：

输入：s = "06"
输出：0
解释："06" 无法映射到 "F" ，因为存在前导零（"6" 和 "06" 并不等价）。


提示：

1 <= s.length <= 100
s 只包含数字，并且可能包含前导零。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func NumDecodings(s string) int {
	n := len(s)
	var process func(int) int
	process = func(i int) int {
		if i > n {
			return 0
		}
		if i == n {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		if s[i] == '1' || s[i] == '2' {
			if s[i] == '2' {
				if i+1 < n && s[i+1] <= '6' {
					return process(i+1) + process(i+2)
				} else {
					return process(i + 1)
				}
			}
			return process(i+1) + process(i+2)
		}
		return process(i + 1)
	}

	dp := make([]int, n+2)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {

			if s[i] == '1' || s[i] == '2' {
				if s[i] == '2' {
					if i+1 < n && s[i+1] <= '6' {
						dp[i] = dp[i+1] + dp[i+2]
					} else {
						dp[i] = dp[i+1]
					}
				} else {
					dp[i] = dp[i+1] + dp[i+2]
				}
			} else {
				dp[i] = dp[i+1]
			}
		}
	}
	fmt.Printf("dp: %v\n", dp)
	return process(0)
}
