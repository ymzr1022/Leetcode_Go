package contest

/*
给你一个整数数组 nums 和两个正整数 m 和 k 。

请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。

如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。

子数组指的是一个数组中一段连续 非空 的元素序列。



示例 1：

输入：nums = [2,6,7,3,1,7], m = 3, k = 4
输出：18
解释：总共有 3 个长度为 k = 4 的几乎唯一子数组。分别为 [2, 6, 7, 3] ，[6, 7, 3, 1] 和 [7, 3, 1, 7] 。这些子数组中，和最大的是 [2, 6, 7, 3] ，和为 18 。
示例 2：

输入：nums = [5,9,9,2,4,5,4], m = 1, k = 3
输出：23
解释：总共有 5 个长度为 k = 3 的几乎唯一子数组。分别为 [5, 9, 9] ，[9, 9, 2] ，[9, 2, 4] ，[2, 4, 5] 和 [4, 5, 4] 。这些子数组中，和最大的是 [5, 9, 9] ，和为 23 。
示例 3：

输入：nums = [1,2,1,2,1,2,1], m = 3, k = 3
输出：0
解释：输入数组中不存在长度为 k = 3 的子数组含有至少  m = 3 个互不相同元素的子数组。所以不存在几乎唯一子数组，最大和为 0 。
 3,3,2,3,4
 m = 4
 len(mp) = 2
 r - l + 1 = 4 k -
提示：

1 <= nums.length <= 2 * 104
1 <= m <= k <= nums.length
1 <= nums[i] <= 109
*/

func MaxSum(nums []int, m int, k int) (res int64) {
	mp := make(map[int]int)
	left := 0
	var sums int64 = 0
	for right, v := range nums {
		mp[v]++
		sums += int64(v)
		if right-left+1 == k && len(mp) >= m {
			if res < sums {
				res = sums
			}
		}
		if right-left+1 > k || (k-right+left-1) < (m-len(mp)) {
			sums -= int64(nums[left])
			mp[nums[left]]--
			if mp[nums[left]] == 0 {
				delete(mp, nums[left])
			}
			left++
		}
		if right-left+1 == k && len(mp) >= m {
			if res < sums {
				res = sums
			}
		}
	}
	return
}
