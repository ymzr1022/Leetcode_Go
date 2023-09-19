package dfs

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：
r,d,l,u
1,1,-1,-1
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func SpiralOrder(matrix [][]int) (res []int) {
	m := len(matrix)
	n := len(matrix[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	res = make([]int, 0)
	var dfs func(int, int, string)
	dfs = func(i, j int, s string) {
		if visit[i][j] {
			return
		}
		res = append(res, matrix[i][j])
		visit[i][j] = true
		if s == "R" {
			if j+1 >= n || visit[i][j+1] {
				dfs(i+1, j, "D")
			} else {
				dfs(i, j+1, "R")
			}
		}
		if s == "D" {
			if i+1 >= m || visit[i+1][j] {
				dfs(i, j-1, "L")
			} else {
				dfs(i+1, j, "D")
			}
		}
		if s == "L" {
			if j-1 < 0 || visit[i][j-1] {
				dfs(i-1, j, "U")
			} else {
				dfs(i, j-1, "L")
			}
		}
		if s == "U" {
			if i-1 < 0 || visit[i-1][j] {
				dfs(i, j+1, "R")
			} else {
				dfs(i-1, j, "U")
			}
		}
	}
	dfs(0, 0, "R")
	return
}
