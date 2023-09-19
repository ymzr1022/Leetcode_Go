package contest

/*
给你一个 二维 整数数组 coordinates 和一个整数 k ，其中 coordinates[i] = [xi, yi] 是第 i 个点在二维平面里的坐标。

我们定义两个点 (x1, y1) 和 (x2, y2) 的 距离 为 (x1 XOR x2) + (y1 XOR y2) ，XOR 指的是按位异或运算。

请你返回满足 i < j 且点 i 和点 j之间距离为 k 的点对数目。



示例 1：

输入：coordinates = [[1,2],[4,2],[1,3],[5,2]], k = 5
输出：2
解释：以下点对距离为 k ：
- (0, 1)：(1 XOR 4) + (2 XOR 2) = 5 。
- (2, 3)：(1 XOR 5) + (3 XOR 2) = 5 。
示例 2：

输入：coordinates = [[1,3],[1,3],[1,3],[1,3],[1,3]], k = 0
输出：10
解释：任何两个点之间的距离都为 0 ，所以总共有 10 组点对。
*/

func CountPairs1(coordinates [][]int, k int) (res int) {
	type pair struct {
		y, t int
	}
	mp := make(map[pair]int, 0)

	for _, v := range coordinates {

		for i := 0; i <= k; i++ {
			xx := i ^ v[0]
			yy := (k - i) ^ v[1]
			res += mp[pair{xx, yy}]

		}
		mp[pair{v[0], v[1]}]++
	}
	return
}
