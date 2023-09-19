package bfs

/*
在一个 8x8 的棋盘上，放置着若干「黑皇后」和一个「白国王」。

给定一个由整数坐标组成的数组 queens ，表示黑皇后的位置；以及一对坐标 king ，表示白国王的位置，返回所有可以攻击国王的皇后的坐标(任意顺序)。



示例 1：



输入：queens = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], king = [0,0]
输出：[[0,1],[1,0],[3,3]]
解释：
[0,1] 的皇后可以攻击到国王，因为他们在同一行上。
[1,0] 的皇后可以攻击到国王，因为他们在同一列上。
[3,3] 的皇后可以攻击到国王，因为他们在同一条对角线上。
[0,4] 的皇后无法攻击到国王，因为她被位于 [0,1] 的皇后挡住了。
[4,0] 的皇后无法攻击到国王，因为她被位于 [1,0] 的皇后挡住了。
[2,4] 的皇后无法攻击到国王，因为她和国王不在同一行/列/对角线上。
示例 2：



输入：queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]
输出：[[2,2],[3,4],[4,4]]
示例 3：



输入：queens = [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]], king = [3,4]
输出：[[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]


提示：

1 <= queens.length <= 63
queens[i].length == 2
0 <= queens[i][j] < 8
king.length == 2
0 <= king[0], king[1] < 8
一个棋盘格上最多只能放置一枚棋子。
*/

func QueensAttacktheKing(queens [][]int, king []int) (res [][]int) {
	grid := make([][8]int, 8)
	for _, queen := range queens {
		grid[queen[0]][queen[1]] = 1
	}
	visit := make([][8]bool, 8)
	queue := make([][]int, 0)
	queue = append(queue, []int{king[0], king[1], 1})
	queue = append(queue, []int{king[0], king[1], 2})
	queue = append(queue, []int{king[0], king[1], 3})
	queue = append(queue, []int{king[0], king[1], 4})
	queue = append(queue, []int{king[0], king[1], 5})
	queue = append(queue, []int{king[0], king[1], 6})
	queue = append(queue, []int{king[0], king[1], 7})
	queue = append(queue, []int{king[0], king[1], 8})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y, to := queue[i][0], queue[i][1], queue[i][2]
			if to == 1 {
				if x+1 < 8 {
					if !visit[x+1][y] {
						if grid[x+1][y] == 1 {
							res = append(res, []int{x + 1, y})
						} else {
							queue = append(queue, []int{x + 1, y, 1})
						}
						visit[x+1][y] = true
					}
				}
			}
			if to == 2 {
				if x-1 >= 0 {
					if !visit[x-1][y] {
						if grid[x-1][y] == 1 {
							res = append(res, []int{x - 1, y})
						} else {
							queue = append(queue, []int{x - 1, y, 2})
						}
						visit[x-1][y] = true
					}
				}
			}
			if to == 3 {
				if x+1 < 8 && y+1 < 8 {
					if !visit[x+1][y+1] {
						if grid[x+1][y+1] == 1 {
							res = append(res, []int{x + 1, y + 1})
						} else {
							queue = append(queue, []int{x + 1, y + 1, 3})
						}
						visit[x+1][y+1] = true
					}
				}
			}
			if to == 4 {
				if x+1 < 8 && y-1 >= 0 {
					if !visit[x+1][y-1] {
						if grid[x+1][y-1] == 1 {
							res = append(res, []int{x + 1, y - 1})
						} else {
							queue = append(queue, []int{x + 1, y - 1, 4})
						}
						visit[x+1][y-1] = true
					}
				}
			}
			if to == 5 {
				if x-1 >= 0 && y+1 < 8 {
					if !visit[x-1][y+1] {
						if grid[x-1][y+1] == 1 {
							res = append(res, []int{x - 1, y + 1})
						} else {
							queue = append(queue, []int{x - 1, y + 1, 5})
						}
						visit[x-1][y+1] = true
					}
				}
			}
			if to == 6 {
				if x-1 >= 0 && y-1 >= 0 {
					if !visit[x-1][y-1] {
						if grid[x-1][y-1] == 1 {
							res = append(res, []int{x - 1, y - 1})
						} else {
							queue = append(queue, []int{x - 1, y - 1, 6})
						}
						visit[x-1][y-1] = true
					}
				}
			}
			if to == 7 {
				if y+1 < 8 {
					if !visit[x][y+1] {
						if grid[x][y+1] == 1 {
							res = append(res, []int{x, y + 1})
						} else {
							queue = append(queue, []int{x, y + 1, 7})
						}
						visit[x][y+1] = true
					}
				}
			}
			if to == 8 {
				if y-1 >= 0 {
					if !visit[x][y-1] {
						if grid[x][y-1] == 1 {
							res = append(res, []int{x, y - 1})
						} else {
							queue = append(queue, []int{x, y - 1, 8})
						}
						visit[x][y-1] = true
					}
				}
			}
		}
		queue = queue[size:]
	}
	return
}
