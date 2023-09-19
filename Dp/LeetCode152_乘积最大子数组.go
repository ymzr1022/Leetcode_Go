package dp

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。



示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。


提示:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MaxProduct(nums []int) (res int) {
	n := len(nums)
	res = nums[0]
	var process func(int) (int, int)
	process = func(i int) (int, int) {
		if i < 0 {
			return 1, 0
		}
		if nums[i] == 0 {
			res = max(res, 0)
			return 0, 0
		} else if nums[i] > 0 {
			max1, min1 := process(i - 1)
			if max1 <= 0 {
				max1 = 1
			}
			res = max(res, max1*nums[i])
			return max1 * nums[i], min1 * nums[i]
		} else {
			max1, min1 := process(i - 1)
			if min1 == 0 {
				min1 = 1
			}
			if max1 <= 0 {
				max1 = 1
			}
			res = max(res, min1*nums[i])
			return min1 * nums[i], max1 * nums[i]
		}
	}
	process(n - 1)
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			process(i - 1)
		}
	}
	return
}
