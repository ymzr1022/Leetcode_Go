package leetcodes

import (
	"fmt"
	"sort"
)

func CombinationSum(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var nums []int
	var sum = 0
	var process func(int)
	var dif = 0
	var nu = 0
	process = func(i int) {
		if sum == target {

			comb := make([]int, len(nums))
			copy(comb, nums)
			res = append(res, comb)
			return
		}
		if i >= len(candidates) {
			return
		}
		if candidates[i] > target || sum > target {
			return
		}
		process(i + 1)
		tmp := candidates[i]
		time := target / tmp
		for k := 1; k <= time; k++ {
			sum += tmp * k
			dif++
			numss := make([]int, k)
			for x := 0; x < k; x++ {
				numss[x] = tmp
			}

			nums = append(nums, numss...)
			process(i + 1)
			sum -= tmp * k
			nums = nums[:len(nums)-k]
			dif--
		}

	}

	process(0)
	fmt.Printf("nu: %v\n", nu)
	return
}
