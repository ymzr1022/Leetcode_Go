package leetcodes

import "fmt"

/**
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。



示例 1：

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：

输入：nums = [9], target = 3
输出：0


提示：

1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CombinationSum4(nums []int, target int) (res int) {
	n := len(nums)
	// sort.Ints(nums)
	// var sum = 0
	var process func(int, int) int
	process = func(index, target int) int {
		if index < 0 {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}
		if nums[index] > target {
			return process(index-1, target)
		}
		return process(index-1, target) + process(n-1, target-nums[index])
	}
	res = process(n-1, target)

	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		for _, x := range nums {
			if j >= x {
				dp[j] += dp[j-x]
			}
		}
	}
	fmt.Printf("dp: %v\n", dp)
	res = dp[target]
	return
}

func Get(a int, sums []int) int {
	y := 1
	for i := a; i > sums[0]; i-- {
		y *= i
	}

	for i := 1; i < len(sums); i++ {
		y /= Gets(sums[i])
	}
	return y
}

func Gets(a int) int {
	y := 1
	for i := 1; i <= a; i++ {
		y *= i
	}
	return y
}
