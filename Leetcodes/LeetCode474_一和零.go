package leetcodes

import "strings"

/**
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。



示例 1：

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。


提示：

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100
通过次数165,395提交次数253,519

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindMaxForm(strs []string, m int, n int) int {
	// lens := len(strs)

	//------------------------------递归-------------------------------
	// var process func(int, int, int) int
	// process = func(index, m, n int) int {
	// 	if index >= lens {
	// 		return 0
	// 	}
	// 	x := strings.Count(strs[index], "0")
	// 	y := strings.Count(strs[index], "1")
	// 	if x > m || y > n {
	// 		return process(index+1, m, n)
	// 	}
	// 	return Max(process(index+1, m, n), process(index+1, m-x, n-y)+1)
	// }
	// return process(0, m, n)

	//----------------------------递推-----------------------------
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		x := strings.Count(str, "0")
		y := strings.Count(str, "1")
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if x > j || y > k {
					dp[j][k] = dp[j][k]
				} else {
					dp[j][k] = Max(dp[j][k], dp[j-x][k-y]+1)
				}
			}
		}
	}
	return dp[m][n]
}
