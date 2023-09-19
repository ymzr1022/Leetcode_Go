package leetcodes

import "fmt"

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。


提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CanPartition(nums []int) bool {
	sums := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sums += nums[i]
	}
	if sums%2 != 0 {
		return false
	}
	sum := sums / 2
	dp := make([]int, sum+1)
	// for i := range dp {
	// 	dp[i] = make([]int, sum+1)
	// }
	// var process func(int, int) int
	// process = func(index, target int) int {
	// 	if index < 0 {
	// 		if target == 0 {
	// 			return 1
	// 		} else {
	// 			return 0
	// 		}
	// 	}
	// 	if target < nums[index] {
	// 		return process(index-1, target)
	// 	}
	// 	tmp := process(index-1, target) + process(index-1, target-nums[index])
	// 	fmt.Printf("tmp: %v  %v   %v\n", tmp, index, target)

	// 	return tmp
	// }

	// i := process(n-1, sum)
	// fmt.Printf("dp: %v\n", dp)
	// fmt.Printf("i: %v\n", i)
	dp[0] = 1
	// for i := 0; i < n; i++ {
	// 	for j := sum; j >= nums[i]; j-- {
	// 		if j >= nums[i] {
	// 			dp[j] = dp[j] + dp[j-nums[i]]
	// 		}
	// 	}
	// }
	for _, x := range nums {
		for i := sum; i >= x; i-- {
			dp[i] = dp[i] + dp[i-x]
		}
	}
	fmt.Printf("dp: %v\n", dp)
	// fmt.Printf("dp: %v\n", dp)
	// if dp[n-1][sum] > 0 {
	// 	return true
	// }
	if dp[sum] != 0 {
		return true
	}
	return false
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
