package leetcodes

/**
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400
通过次数771,946提交次数1,418,719

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Rob(nums []int) int {
	// n := len(nums)

	// var process func(int) int
	// process = func(index int) int {
	// 	if index < 0 {
	// 		return 0
	// 	}
	// 	return max(process(index-1), process(index-2)+nums[index])
	// }
	// return process(n - 1)

	// dp := make([]int, n)
	// for i, x := range nums {
	// 	if i == 1 {
	// 		dp[i] = max(dp[i-1], x)
	// 	} else if i == 0 {
	// 		dp[i] = max(0, x)
	// 	} else {
	// 		dp[i] = max(dp[i-1], dp[i-2]+x)
	// 	}
	// }
	f0, f1 := 0, 0
	for _, x := range nums {
		tmp := max(f1, f0+x)
		f0 = f1
		f1 = tmp
	}
	// fmt.Printf("dp: %v\n", dp)
	return f1
}
