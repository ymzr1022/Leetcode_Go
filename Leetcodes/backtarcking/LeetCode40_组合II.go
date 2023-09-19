package leetcodes

import "sort"

/**
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。



示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CombinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	sum := 0
	var nums []int
	var process func(int)
	var flgs = make([]bool, len(candidates))

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
		if i > 0 && candidates[i] == candidates[i-1] && !flgs[i-1] {
			return
		}
		tmp := candidates[i]
		nums = append(nums, tmp)
		sum += tmp
		flgs[i] = true
		process(i + 1)
		sum -= tmp
		nums = nums[:len(nums)-1]
		flgs[i] = false
	}
	process(0)
	return
}
