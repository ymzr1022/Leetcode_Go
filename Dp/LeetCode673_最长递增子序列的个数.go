package dp

import "fmt"

/**
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。



示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。


提示:

1 <= nums.length <= 2000
-106 <= nums[i] <= 106

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/number-of-longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindNumberOfLIS(nums []int) (res int) {

	n := len(nums)
	timess := make([]int, n)
	var process func(int) int
	process = func(i int) int {
		res := 0
		for j, _ := range nums[:i] {
			if nums[j] < nums[i] {
				res = max(process(j), res)
				timess[res]++
			}
		}
		return res + 1
	}
	dp := make([]int, n+1)
	count := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
		count[i] = 1
	}
	count[0] = 1
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}

		}

		ans = max(ans, dp[i])

	}
	fmt.Printf("dp: %v\n", dp)
	fmt.Printf("ans: %v\n", ans)
	fmt.Printf("count: %v\n", count)
	for i := 0; i < n; i++ {
		if dp[i] == ans {
			res += count[i]
		}
	}
	return
}
