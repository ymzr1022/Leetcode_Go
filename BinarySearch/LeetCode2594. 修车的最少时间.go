package binarysearch

import "math"

/*
给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * n2 分钟内修好 n 辆车。

同时给你一个整数 cars ，表示总共需要修理的汽车数目。

请你返回修理所有汽车 最少 需要多少时间。

注意：所有机械工可以同时修理汽车。



示例 1：

输入：ranks = [4,2,3,1], cars = 10
输出：16
解释：
- 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
- 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
- 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
- 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
16 分钟是修理完所有车需要的最少时间。
示例 2：

输入：ranks = [5,1,8], cars = 6
输出：16
解释：
- 第一位机械工修 1 辆车，需要 5 * 1 * 1 = 5 分钟。
- 第二位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
- 第三位机械工修 1 辆车，需要 8 * 1 * 1 = 8 分钟。
16 分钟时修理完所有车需要的最少时间。

a * n1 * n1 = b * n2 * n2
b / a = (n2 / n1) ^ 2
提示：
[0,1,0],
[0,0,0],
[0,0,1]
[1,1,1,1,1],
[1,0,0,0,1],
[1,0,1,0,1],
[1,0,0,0,1],
[1,1,1,1,1]
1 <= ranks.length <= 105
1 <= ranks[i] <= 100
1 <= cars <= 106
*/

func RepairCars(ranks []int, cars int) int64 {

	minr := ranks[0]
	cnts := make([]int, 102)
	for _, v := range ranks {
		cnts[v]++
		if v < minr {
			minr = v
		}
	}
	l, r := 0, cars*cars*minr
	for l < r {
		mid := (l + r) >> 1
		s := 0
		for i := minr; i < 101 && s < cars; i++ {
			s += cnts[i] * int(math.Sqrt(float64(mid/i)))
		}
		if s >= cars {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
