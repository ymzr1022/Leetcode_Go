package dp

/**
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: rowIndex = 3
输出: [1,3,3,1]
示例 2:

输入: rowIndex = 0
输出: [1]
示例 3:

输入: rowIndex = 1
输出: [1,1]


提示:

0 <= rowIndex <= 33


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func GetRow(rowIndex int) (res []int) {

	if rowIndex == 0 {
		res = append(res, 1)
		return
	}
	dp := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		tmp := 1
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[j] = 1
			} else if j == i {
				dp[j-1] = tmp
				dp[j] = 1
			} else {
				tmp1 := dp[j-1] + dp[j]
				dp[j-1] = tmp
				tmp = tmp1
			}
		}
	}
	res = append(res, dp...)
	return
}
