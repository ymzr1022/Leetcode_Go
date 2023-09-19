package leetcodes

/**
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。

如果你能使这个正方形，则返回 true ，否则返回 false 。



示例 1:



输入: matchsticks = [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。
示例 2:

输入: matchsticks = [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。


提示:

1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 108

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/matchsticks-to-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Makesquare(matchsticks []int) bool {
	sums := 0
	n := len(matchsticks)
	for i := 0; i < n; i++ {
		sums += matchsticks[i]
	}
	if sums%4 != 0 {
		return false
	}
	target := sums / 4
	path := make([]int, 4)
	var process func(int) bool
	process = func(index int) bool {
		if index == n {
			return true
		}

		for i := 0; i < 4; i++ {
			if i > 0 && path[i] == path[i-1] {
				continue
			}
			if path[i]+matchsticks[index] > target {
				continue
			} else {
				path[i] += matchsticks[index]

				if process(index + 1) {
					return true
				}

				path[i] -= matchsticks[index]
			}

		}
		return false
	}

	return process(0)
}
