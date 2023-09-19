package leetcodes

import (
	"fmt"
	"math"
)

/**
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格，和一个整型 k 。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。


提示：

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MaxProfit4(k int, prices []int) int {
	n := len(prices)

	var process func(int, int, int) int
	process = func(index, status, k int) int {
		if index < 0 {
			if status == 0 {
				return 0
			} else {
				return math.MinInt32
			}
		}
		if k < 0 {
			return math.MinInt32
		}
		if status == 0 {
			return max(process(index-1, 0, k), process(index-1, 1, k-1)+prices[index])
		} else {
			return max(process(index-1, 1, k), process(index-1, 0, k)-prices[index])
		}
	}

	// fmt.Printf("process(0, 0, k): %v\n", process(0, 0, k))

	// bp := make([][]int, n+1)
	// for i := range bp {
	// 	bp[i] = make([]int, 2)
	// }
	// bp[0][0] = 0
	// bp[0][1] = -prices[0]
	// for i := 1; i < n; i++ {
	// 	bp[i][0] = max(bp[i-1][1]+prices[i], bp[i-1][0])
	// 	for j := k; j >= 0; j-- {
	// 		bp[i][1] = max(bp[i-1][1], bp[i-1][0]-prices[i])
	// 	}
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	for j := 0; j < k+2; j++ {
		dp[0][j] = 0
	}
	for _, v := range prices {
		for j := k; j >= 0; j-- {
			dp[1][j+1] = max(dp[1][j+1], dp[0][j+1]-v)
			dp[0][j+1] = max(dp[0][j+1], dp[1][j]+v)
		}
	}
	fmt.Printf("dp: %v\n", dp)
	fmt.Printf("dp[0][k]: %v\n", dp[0][k+1])
	// }
	// fmt.Printf("bp: %v\n", bp)
	return process(n-1, 0, k)
}
