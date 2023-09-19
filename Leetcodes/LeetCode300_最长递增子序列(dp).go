package leetcodes

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1


提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLIS(nums []int) int {
	n := len(nums)

	// var process func(int) int
	// process = func(i int) int {
	// 	res := 0
	// 	for j, x := range nums[:i] {
	// 		if x < nums[i] {
	// 			res = max(process(j), res)
	// 		}
	// 	}
	// 	return res + 1
	// }
	// ans := 0
	// for i, _ := range nums {
	// 	ans = max(ans, process(i))
	// }
	// fmt.Printf("process(1, 1): %v\n", ans)

	// var p2 func(int) int
	// p2 = func(i int) int {
	// 	if i >= n {
	// 		return 1
	// 	}
	// 	if nums[i] <= nums[i-1] {
	// 		return p2(i + 1)
	// 	}
	// 	return p2(i+1) + 1
	// }
	// fmt.Printf("p2(1): %v\n", p2(1))

	dp := make([]int, n+1)
	ans := 0
	for i := 0; i < n; i++ {
		res := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res = max(res, dp[j])
			}
		}
		dp[i] = res + 1
		ans = max(ans, dp[i])
	}
	return ans
}
