package dfs

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20
*/

func GenerateMatrix(n int) (res [][]int) {
	res = make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	var dfs func(i, j, num int, s string)
	dfs = func(i, j, num int, s string) {
		if i >= n || i < 0 || j < 0 || j >= n || res[i][j] != 0 {
			return
		}
		res[i][j] = num
		if s == "R" {
			if j+1 >= n || res[i][j+1] != 0 {
				dfs(i+1, j, num+1, "D")
			} else {
				dfs(i, j+1, num+1, "R")
			}
		}
		if s == "D" {
			if i+1 >= n || res[i+1][j] != 0 {
				dfs(i, j-1, num+1, "L")
			} else {
				dfs(i+1, j, num+1, "D")
			}
		}
		if s == "L" {
			if j-1 < 0 || res[i][j-1] != 0 {
				dfs(i-1, j, num+1, "U")
			} else {
				dfs(i, j-1, num+1, "L")
			}
		}
		if s == "U" {
			if i-1 < 0 || res[i-1][j] != 0 {
				dfs(i, j+1, num+1, "R")
			} else {
				dfs(i-1, j, num+1, "U")
			}
		}
	}
	dfs(0, 0, 1, "R")
	return
}
