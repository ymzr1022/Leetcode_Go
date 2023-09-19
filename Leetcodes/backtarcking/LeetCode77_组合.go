package leetcodes

func Combine(n int, k int) (res [][]int) {
	var nums []int
	var process func(int)
	process = func(i int) {
		if len(nums)+(n-i+1) < k {
			return
		}
		if len(nums) == k {
			comb := make([]int, k)
			copy(comb, nums)
			res = append(res, comb)
			return
		}
		if i > n {
			return
		}
		process(i + 1)
		nums = append(nums, i)
		process(i + 1)
		nums = nums[:len(nums)-1]
	}
	process(1)
	return
}
