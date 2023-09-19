package binarysearch

import "fmt"

/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。



示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0


提示：

1 <= nums.length <= 104
-105 <= nums[i] <= 105


进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
*/

func FindUnsortedSubarray(nums []int) int {
	st := []int{}            // 单调栈（从左到右遍历为单调「递增」栈，找「左不满足的项」）
	leftIdx := len(nums) - 1 // 最短数组的左边界（通过「最小的栈顶下标」可得）
	for i := 0; i < len(nums); i++ {
		for len(st) > 0 && nums[st[len(st)-1]] > nums[i] {
			leftIdx = min(st[len(st)-1], leftIdx)
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	fmt.Printf("nums: %v\n", leftIdx)
	fmt.Printf("nums: %v\n", st)
	st = []int{} // 单调栈（从右到左遍历为单调「递减」栈，找「右不满足的项」）

	rightIdx := 0 // 最短数组的右边界（通过「最大的栈顶下标」可得）
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i] {
			rightIdx = max(st[len(st)-1], rightIdx)
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	fmt.Printf("nums: %v\n", rightIdx)
	fmt.Printf("nums: %v\n", st)
	if leftIdx == len(nums)-1 && rightIdx == 0 {
		return 0
	} // 因 leftIdx 和 rightIdx 仍为初值，所以本来 nums 就有序，按题意返回 0
	return rightIdx - leftIdx + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
