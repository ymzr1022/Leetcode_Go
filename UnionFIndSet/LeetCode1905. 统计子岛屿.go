package unionfindset

import (
	"fmt"
	"sort"
)

/*
给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。

如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。

请你返回 grid2 中 子岛屿 的 数目 。



示例 1：


输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
输出：3
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
示例 2：


输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
输出：2
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。


提示：

m == grid1.length == grid2.length
n == grid1[i].length == grid2[i].length
1 <= m, n <= 500
grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。
*/

var p []int

func findd(x int) int {
	if p[x] != x {
		p[x] = findd(p[x])
	}
	return p[x]
}

func queryy(x, y int) bool {
	return p[findd(x)] == p[findd(y)]
}

func unions(x, y int) {
	p[findd(x)] = p[findd(y)]
}

func CountSubIslands(grid1 [][]int, grid2 [][]int) (res int) {
	n := len(grid1)
	m := len(grid1[0])
	var getIndex func(i, j int) int
	getIndex = func(i, j int) int {
		return m*i + j
	}
	// var getIJ func(x int) (int, int)
	// getIJ = func(x int) (int, int) {
	// 	return x / m, x - x/m
	// }
	pp := make([][2]int, m*n)
	var process func(i, j, k int)
	process = func(i, j, k int) {
		if i >= n || i < 0 || j >= m || j < 0 || grid2[i][j] != 1 {
			return
		}
		grid2[i][j] = k
		pp[getIndex(i, j)][0] = k
		pp[getIndex(i, j)][1] = getIndex(i, j)
		process(i+1, j, k)
		process(i-1, j, k)
		process(i, j+1, k)
		process(i, j-1, k)
	}
	p = make([]int, m*n)

	k := 1
	for i, v := range grid2 {
		for j, vv := range v {
			if vv == 1 {
				k++
				process(i, j, k)
			}
		}
	}

	pp1 := make([][2]int, m*n)
	var process1 func(i, j, k int)
	process1 = func(i, j, k int) {
		if i >= n || i < 0 || j >= m || j < 0 || grid1[i][j] != 1 {
			return
		}
		grid1[i][j] = k
		pp1[getIndex(i, j)][0] = k
		pp1[getIndex(i, j)][1] = getIndex(i, j)
		process(i+1, j, k)
		process(i-1, j, k)
		process(i, j+1, k)
		process(i, j-1, k)
	}
	p = make([]int, m*n)

	k = 1
	for i, v := range grid1 {
		for j, vv := range v {
			if vv == 1 {
				k++
				process1(i, j, k)
			}
		}
	}
	fmt.Printf("p: %v\n", p)
	fmt.Printf("grid2: %v\n", grid2)
	fmt.Printf("pp: %v\n", pp)
	fmt.Printf("pp1: %v\n", pp1)
	sort.Slice(pp, func(i, j int) bool {
		return pp[i][0] > pp[j][0]
	})
	for i, v := range pp {
		index := v[1]
		num := v[0]
		if num != 0 {
			target := pp1[index][0]
			if target == 0 {
				continue
			}
			i++
			flg := true
			for i < len(pp) && pp[i][0] == num {
				if pp1[pp[i][1]][0] != target {
					flg = false
					break
				}
			}
			if flg {
				res++
			}
		}
	}
	return
}
