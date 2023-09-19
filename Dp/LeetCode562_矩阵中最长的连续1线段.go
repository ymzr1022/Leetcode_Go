package dp

/**
给定一个 m x n 的二进制矩阵 mat ，返回矩阵中最长的连续1线段。

这条线段可以是水平的、垂直的、对角线的或者反对角线的。



示例 1:



输入: mat = [[0,1,1,0],[0,1,1,0],[0,0,0,1]]
输出: 3
示例 2:



输入: mat = [[1,1,1,1],[0,1,1,0],[0,0,0,1]]
输出: 4


提示:

m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
mat[i][j] 不是 0 就是 1.

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-line-of-consecutive-one-in-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func LongestLine(mat [][]int) (res int) {
	row := len(mat)
	col := len(mat[0])

	// var process func(int, int) (int, int, int, int)
	// process = func(i, j int) (int, int, int, int) {
	// 	if i >= row || j >= col || i < 0 || j < 0 {
	// 		return 0, 0, 0, 0
	// 	}
	// 	if mat[i][j] == 1 {
	// 		down, _, _, _ := process(i+1, j)
	// 		_, right, _, _ := process(i, j+1)
	// 		_, _, cross, _ := process(i+1, j+1)
	// 		_, _, _, xcross := process(i+1, j-1)
	// 		x := max(max(max(down, right), cross), xcross) + 1
	// 		res = max(res, x)
	// 		return down + 1, right + 1, cross + 1, xcross + 1
	// 	}
	// 	return 0, 0, 0, 0
	// }
	// process(0, 0)
	dp := make([][][]int, 2)
	for i := range dp {
		dp[i] = make([][]int, col+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 1 {
				dp[(i+1)%2][j+1][0] = dp[i%2][j+1][0] + 1
				dp[(i+1)%2][j+1][1] = dp[(i+1)%2][j][1] + 1
				dp[(i+1)%2][j+1][2] = dp[i%2][j][2] + 1
				dp[(i+1)%2][j+1][3] = dp[(i)%2][j+2][3] + 1
				res = max(res, dp[(i+1)%2][j+1][0])
				res = max(res, dp[(i+1)%2][j+1][1])
				res = max(res, dp[(i+1)%2][j+1][2])
				res = max(res, dp[(i+1)%2][j+1][3])
			}
		}
	}

	return

}
