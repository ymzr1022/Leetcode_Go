package bfs

import "fmt"

/*
给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。

岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。

你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。

返回必须翻转的 0 的最小数目。

[[1,1,0,0,0]
,[1,0,0,0,0]
,[1,0,0,0,0]
,[0,0,0,1,1]
,[0,0,0,1,1]]

示例 1：

输入：grid = [[0,1],[1,0]]
输出：1
示例 2：

输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
输出：2
示例 3：

输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1


提示：

n == grid.length == grid[i].length
2 <= n <= 100
grid[i][j] 为 0 或 1
grid 中恰有两个岛
*/

func ShortestBridge(grid [][]int) (res int) {
	m := len(grid)
	n := len(grid[0])
	x := 1
	cnt2 := 0
	cnt3 := 0
	queue2 := make([][2]int, 0)
	var infect func(i, j, k int)
	infect = func(i, j, k int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return
		}
		grid[i][j] = x
		if k == 1 {
			cnt2++
			queue2 = append(queue2, [2]int{i, j})
		} else {
			cnt3++
		}
		infect(i+1, j, k)
		infect(i-1, j, k)
		infect(i, j+1, k)
		infect(i, j-1, k)
	}
	times := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				x++
				times++
				infect(i, j, times)
			}
		}
	}
	fmt.Printf("queue2: %v\n", queue2)
	for len(queue2) != 0 {
		size := len(queue2)
		for k := 0; k < size; k++ {
			i := queue2[k][0]
			j := queue2[k][1]
			if i+1 < m {
				if grid[i+1][j] == 0 {
					grid[i+1][j] = 4
					queue2 = append(queue2, [2]int{i + 1, j})
				}
				if grid[i+1][j] == 3 {
					return
				}
			}
			if i-1 >= 0 {
				if grid[i-1][j] == 0 {
					grid[i-1][j] = 4
					queue2 = append(queue2, [2]int{i - 1, j})
				}
				if grid[i-1][j] == 3 {
					return
				}
			}
			if j+1 < n {
				if grid[i][j+1] == 0 {
					grid[i][j+1] = 4
					queue2 = append(queue2, [2]int{i, j + 1})
				}
				if grid[i][j+1] == 3 {
					return
				}
			}
			if j-1 >= 0 {
				if grid[i][j-1] == 0 {
					grid[i][j-1] = 4
					queue2 = append(queue2, [2]int{i, j - 1})
				}
				if grid[i][j-1] == 3 {
					return
				}
			}

		}
		queue2 = queue2[size:]
		res++
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
