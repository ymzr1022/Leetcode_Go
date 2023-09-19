package sudoku

func Sudokus(girds [][]int) (res [][][]int) {
	n := 9
	dp := make([][]int, 9)
	for i := range dp {
		dp[i] = make([]int, 9)
	}

	var valid func(int, int) bool
	valid = func(i1, j1 int) bool {
		tmp := make([]int, 10)
		for i := i1; i < i1+3; i++ {
			for j := j1; j < j1+3; j++ {
				tmp[dp[i][j]]++
			}
		}
		for i := 1; i < 10; i++ {
			if tmp[i] > 1 {
				return false
			}
		}
		return true
	}

	var process func(int, int)

	process = func(i, j int) {
		if i == n-1 && j == n-1 {
			res = append(res, append([][]int{}, dp...))
			return
		}
		if i >= n || j >= n {
			return
		}
		if girds[i][j] == 0 {
			for k := 1; k <= 9; k++ {
				dp[i][j] = k
				if !valid(i, j) {
					dp[i][k] = 0
				} else {
					process(i+1, j)
					process(i, j+1)
				}

			}

		} else {
			dp[i][j] = girds[i][j]
		}
	}

	return
}
