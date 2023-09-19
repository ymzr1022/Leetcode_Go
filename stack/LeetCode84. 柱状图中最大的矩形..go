package stack

import "fmt"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。



示例 1:



输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：



输入： heights = [2,4]
输出： 4


提示：

1 <= heights.length <=105
0 <= heights[i] <= 104
*/

func LargestRectangleArea(heights []int) (res int) {
	n := len(heights)
	small := make([][2]int, n)
	for i := range small {
		small[i][0] = -1
		small[i][1] = -1
	}
	stack := make([]int, 0)
	for i, v := range heights {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {

			for len(stack) > 0 && v < heights[stack[len(stack)-1]] {
				small[stack[len(stack)-1]][1] = i

				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				small[i][0] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	for i, _ := range small {
		left := small[i][0]
		if left == -1 {
			left = -1
		}
		right := small[i][1]
		if right == -1 {
			right = n
		}
		res = max(res, (right-1-left)*heights[i])
	}
	fmt.Printf("small: %v\n", small)
	return
}
