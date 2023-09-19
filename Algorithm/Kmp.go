package algorithm

func GetIndex(str string, target string) int {
	n := len(str)
	i := 0
	j := 0
	next := GetNext(target)
	for i < n && j < len(target) {
		if str[i] == target[j] {
			i++
			j++
		} else if next[j] >= 0 {
			j = next[j]
		} else {
			j = 0
			i++
		}
	}
	if j == len(target) {
		return i - j
	}
	return -1
}

func GetNext(str string) []int {
	n := len(str)
	dp := make([]int, n)
	dp[0] = -1
	dp[1] = 0
	cn := 0
	for i := 2; i < n; i++ {
		for cn >= 0 {
			if str[cn] == str[i-1] {
				cn++
				dp[i] = cn
				break
			} else if dp[cn] != -1 {
				cn = dp[cn]
			} else {
				dp[i] = 0
				cn = 0
				break
			}
		}
	}
	return dp
}
