package contest

import "fmt"

func minimumRightShifts(nums []int) (res int) {
	next := make([]int, 0)
	next = append(next, nums...)
	next = append(next, nums...)
	fmt.Printf("res: %v\n", next)
	n := len(nums)
	res = -1
	for i := 0; i < n; i++ {
		flg := true
		for j := i + 1; j < i+n; j++ {
			if flg && next[j] > next[j-1] {
				flg = false
			}
		}
		if flg {
			if i == 0 {
				return 0
			}
			res = n - i
			break
		}
	}
	return
}
