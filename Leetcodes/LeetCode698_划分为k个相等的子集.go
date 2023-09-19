package leetcodes

import (
	"sort"
)

/**
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。



示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
示例 2:

输入: nums = [1,2,3,4], k = 3
输出: false


提示：

1 <= k <= len(nums) <= 16
0 < nums[i] < 10000
每个元素的频率在 [1,4] 范围内

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-to-k-equal-sum-subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CanPartitionKSubsets(nums []int, k int) bool {
	sums := 0
	for i := 0; i < len(nums); i++ {
		sums += nums[i]
	}
	if sums%k != 0 {
		return false
	}
	sum := sums / k
	sums = 0
	// var flag bool
	// sort.Ints(nums)
	// times := 0
	var bucket = make([]int, k)
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	var process func(int) bool
	process = func(index int) bool {
		if index == len(nums) {
			// for i := 0; i < k; i++ {
			// 	if bucket[i] != sum {
			// 		return false
			// 	}
			// }
			return true
		}

		for i := 0; i < k; i++ {
			if i > 0 && bucket[i] == bucket[i-1] {
				continue
			}
			if bucket[i]+nums[index] > sum {
				continue
			}
			if bucket[i]+nums[index] <= sum {
				bucket[i] += nums[index]
				if process(index + 1) {
					return true
				}
				bucket[i] -= nums[index]
			}
		}
		return false
	}

	return process(0)
}
