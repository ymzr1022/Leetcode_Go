package dp

import (
	"strconv"
)

func DiffWaysToCompute(expression string) (res []int) {

	if isNum(expression) {
		i, _ := strconv.Atoi(expression)
		res = append(res, i)
		return
	}

	for i, v := range expression {

		if v == '+' || v == '-' || v == '*' {
			left := DiffWaysToCompute(expression[:i])
			right := DiffWaysToCompute(expression[i+1:])

			for _, lNum := range left {
				for _, rNum := range right {
					ans := 0
					if v == '+' {
						ans = lNum + rNum
					}
					if v == '-' {
						ans = lNum - rNum
					}
					if v == '*' {
						ans = lNum * rNum
					}
					res = append(res, ans)
				}
			}
		}

	}
	return
}

func isNum(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
