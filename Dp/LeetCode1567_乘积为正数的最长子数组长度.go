package dp

/**
给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。

一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。

请你返回乘积为正数的最长子数组长度。



示例  1：

输入：nums = [1,-2,-3,4]
输出：4
解释：数组本身乘积就是正数，值为 24 。
示例 2：

输入：nums = [0,1,-2,-3,-4]
输出：3
解释：最长乘积为正数的子数组为 [1,-2,-3] ，乘积为 6 。
注意，我们不能把 0 也包括到子数组中，因为这样乘积为 0 ，不是正数。
示例 3：

输入：nums = [-1,-2,-3,0,1]
输出：2
解释：乘积为正数的最长子数组是 [-1,-2] 或者 [-2,-3] 。


提示：

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func GetMaxLen(nums []int) (res int) {
	n := len(nums)

	dp0, dp1 := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			if dp1 > 0 {
				dp0, dp1 = dp0+1, dp1+1
			} else {
				dp0, dp1 = dp0+1, 0
			}
		} else if nums[i] < 0 {
			if dp1 > 0 {
				dp0, dp1 = dp1+1, dp0+1
			} else {
				dp0, dp1 = 0, dp0+1
			}
		} else {
			dp0, dp1 = 0, 0
		}
		res = max(res, dp0)
	}
	return
}
