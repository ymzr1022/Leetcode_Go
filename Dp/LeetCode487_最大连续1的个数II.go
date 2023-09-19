package dp

/**
给定一个二进制数组 nums ，如果最多可以翻转一个 0 ，则返回数组中连续 1 的最大个数。



示例 1：

输入：nums = [1,0,1,1,0]
输出：4
解释：翻转第一个 0 可以得到最长的连续 1。
     当翻转以后，最大连续 1 的个数为 4。
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：4


提示:

1 <= nums.length <= 105
nums[i] 不是 0 就是 1.

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/max-consecutive-ones-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMaxConsecutiveOnes(nums []int) (res int) {
	start, end, count0, end1 := 0, 0, 0, 0
	n := len(nums)
	for start < n && end < n {
		if nums[end] == 1 {
			end++
			res = Max(res, end-start)
		} else {
			if count0 < 1 {
				count0++
				end1 = end
				end++
				res = Max(res, end-start)
			} else if count0 == 1 {
				count0--
				res = Max(res, end-start)

				start = end1 + 1
				end = start
			}
		}
	}
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
