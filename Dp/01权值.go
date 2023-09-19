package dp

import (
	"fmt"
	"strconv"
)

func GetNum(str string) int {
	n := len(str)
	var process func(int, int) int
	process = func(i, next int) int {
		if i < 0 {
			return 0
		}
		if i == n-1 {
			if str[i] == '0' {
				return min(process(i-1, 0), process(i-1, 1)+1)
			} else {
				return min(process(i-1, 1)+1, process(i-1, 1))
			}
		}
		i2, _ := strconv.Atoi(str[i : i+1])
		if i2 != next {
			return process(i-1, i2)
		} else {
			if i2 == 0 {
				return process(i-1, 1) + 1
			} else {
				return process(i-1, 0) + 1
			}
		}
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		i2, _ := strconv.Atoi(str[i : i+1])
		if i2 == 0 {
			dp[i+1][0] = dp[i][1]
			dp[i+1][1] = dp[i][0] + 1
		} else {
			dp[i+1][0] = dp[i][1] + 1
			dp[i+1][1] = dp[i][0]
		}
	}
	fmt.Printf("dp: %v\n", dp)
	dp0, dp1 := 0, 0
	for i := 0; i < n; i++ {
		i2, _ := strconv.Atoi(str[i : i+1])
		if i2 == 0 {
			dp0, dp1 = dp1, dp0+1
		} else {
			dp0, dp1 = dp1+1, dp0
		}
	}
	return min(dp0, dp1)

}

var p []int

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func query(a, b int) bool {
	return p[find(a)] == p[find(b)]
}

func union(a, b int) {
	p[find(a)] = p[find(b)]
}
