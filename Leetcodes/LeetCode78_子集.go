package leetcodes

/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Subsets(nums []int) (res [][]int) {
	// sort.Ints(nums)
	var set []int
	var process func(int)
	process = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}
		if i > len(nums) {
			return
		}
		process(i + 1)
		set = append(set, nums[i])
		process(i + 1)
		set = set[:len(set)-1]

	}
	process(0)

	return
}
