package leetcodes

import (
	"fmt"
	"math"
)

/**
你有一个凸的 n 边形，其每个顶点都有一个整数值。给定一个整数数组 values ，其中 values[i] 是第 i 个顶点的值（即 顺时针顺序 ）。

假设将多边形 剖分 为 n - 2 个三角形。对于每个三角形，该三角形的值是顶点标记的乘积，三角剖分的分数是进行三角剖分后所有 n - 2 个三角形的值之和。

返回 多边形进行三角剖分后可以得到的最低分 。


示例 1：



输入：values = [1,2,3]
输出：6
解释：多边形已经三角化，唯一三角形的分数为 6。
示例 2：



输入：values = [3,7,4,5]
输出：144
解释：有两种三角剖分，可能得分分别为：3*7*5 + 4*5*7 = 245，或 3*4*5 + 3*4*7 = 144。最低分数为 144。
示例 3：



输入：values = [1,3,1,4,1,5]
输出：13
解释：最低分数三角剖分的得分情况为 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13。


提示：

n == values.length
3 <= n <= 50
1 <= values[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-score-triangulation-of-polygon
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MinScoreTriangulation(values []int) int {
	n := len(values)

	var process func(int, int) int
	process = func(i int, j int) int {
		if i+1 == j {
			return 0
		}
		res := math.MaxInt32
		for k := i + 1; k < j; k++ {
			res = min(res, process(i, k)+values[i]*values[j]*values[k]+process(k, j))
		}
		return res
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			if j != i+1 {
				dp[i][j] = math.MaxInt32
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			res := dp[i][j]
			for k := i + 1; k < j; k++ {
				res = min(res, dp[i][k]+values[i]*values[j]*values[k]+dp[k][j])
			}
			dp[i][j] = res
		}
	}
	fmt.Printf("dp: %v\n", dp[0][n-1])

	return process(0, n-1)
}
