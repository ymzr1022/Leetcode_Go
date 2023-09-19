package dp

/*
给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有 多个 石头。

每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。

请你返回每个格子恰好有一个石头的 最少移动次数 。



示例 1：



输入：grid = [[1,1,0],[1,1,1],[1,2,1]]
输出：3
解释：让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
总共需要 3 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 3 。
示例 2：



输入：grid = [[1,3,0],[1,0,0],[1,0,3]]
输出：4
解释：让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (0,1) 移动到 (0,2) 。
2 - 将一个石头从格子 (0,1) 移动到 (1,1) 。
3 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
4 - 将一个石头从格子 (2,2) 移动到 (2,1) 。
总共需要 4 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 4 。


提示：

grid.length == grid[i].length == 3
0 <= grid[i][j] <= 9
grid 中元素之和为 9 。
*/

func MinimumMoves(grid [][]int) (res int) {
	moved := make([][]int, 0)
	end := make([][]int, 0)
	for i := range grid {
		for j, v := range grid[i] {
			if v == 0 {
				end = append(end, []int{i, j})
			}
			x := v
			for x > 1 {
				moved = append(moved, []int{i, j})
				x--
			}
		}
	}
	res = 1<<31 - 1
	var getAnswer func([][]int)
	getAnswer = func(mov [][]int) {
		count := 0
		for i, v := range mov {
			count += abs(v[0]-end[i][0]) + abs(v[1]-end[i][1])
		}
		res = min(count, res)
		return
	}
	var dfs func(mov [][]int, i int)
	dfs = func(mov [][]int, i int) {
		if i == len(moved) {
			getAnswer(mov)
			return
		}
		dfs(mov, i+1)
		for j := i + 1; j < len(moved); j++ {
			mov[i], mov[j] = mov[j], mov[i]
			dfs(mov, i+1)
			mov[i], mov[j] = mov[j], mov[i]
		}
	}

	le := make([]int, len(moved))
	for i := 0; i < len(moved); i++ {
		le[i] = i
	}
	dfs(moved, 0)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
