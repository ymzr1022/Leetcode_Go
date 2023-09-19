package dp

import "fmt"

/**
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1


提示：

1 <= envelopes.length <= 105
envelopes[i].length == 2
1 <= wi, hi <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/russian-doll-envelopes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MaxEnvelopes(envelopes [][]int) (ans int) {
	n := len(envelopes)
	var process func(int) int
	process = func(i int) int {
		res := 0
		for j := range envelopes[:i] {
			if envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
				res = max(process(j), res)
			}
		}
		ans = max(ans, res+1)
		return res + 1
	}
	process(n - 1)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		res := 0
		for j := 0; j < i; j++ {
			if envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
				res = max(dp[j], res)
			}
		}
		dp[i] = res + 1
	}
	fmt.Printf("dp: %v\n", dp)
	return
}
