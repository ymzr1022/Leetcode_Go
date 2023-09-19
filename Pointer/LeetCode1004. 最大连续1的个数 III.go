package pointer

import "fmt"

/*
给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。



示例 1：

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。


提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1
0 <= k <= nums.length
*/

func LongestOnes(nums []int, k int) (res int) {
	cnt := 0
	left := 0
	for i, v := range nums {
		fmt.Printf("i: %v\n", i)
		if v == 1 {
			cnt++
			res = max(cnt, res)
		} else {
			if k > 0 {
				cnt++
				res = max(cnt, res)
				k--
			} else {
				for nums[left] != 0 {
					left++
					cnt--

				}
				left++
				res = max(cnt, res)
			}
		}
	}
	return
}
