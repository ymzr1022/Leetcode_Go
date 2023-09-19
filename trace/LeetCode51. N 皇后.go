package trace

/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9
*/

func SolveNQueens(n int) [][]string {

	var check func([][2]int) bool
	check = func(strs [][2]int) bool {
		str := strs[len(strs)-1]
		i2 := str[0]
		j2 := str[1]
		for i := 0; i < len(strs)-1; i++ {
			strtmp := strs[i]
			i1 := strtmp[0]
			j1 := strtmp[1]
			if i1 == i2 || j1 == j2 || abs(i2-i1) == abs(j1-j2) {
				return false
			}
		}
		return true
	}
	res := make([][][2]int, 0)
	indeies := make([][2]int, 0)

	var process func(i int)
	process = func(i int) {
		if i >= n {
			cp := make([][2]int, n)
			copy(cp, indeies)
			res = append(res, cp)
			return
		}
		for j := 0; j < n; j++ {
			indeies = append(indeies, [2]int{i, j})
			if check(indeies) {
				process(i + 1)
			}
			indeies = indeies[:len(indeies)-1]
		}
	}
	for i := 0; i < n; i++ {
		indeies = append(indeies, [2]int{0, i})
		process(1)
		indeies = indeies[:len(indeies)-1]
	}
	ans := make([][]string, 0)
	for _, v := range res {
		tmp := make([]string, 0)
		for _, x := range v {
			str := ""
			index := x[1]
			for i := 0; i < n; i++ {
				if i == index {
					str += "Q"
				} else {
					str += "."
				}
			}
			tmp = append(tmp, str)
		}
		ans = append(ans, tmp)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
