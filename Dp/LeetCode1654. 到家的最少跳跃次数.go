package dp

import "fmt"

/*
有一只跳蚤的家在数轴上的位置 x 处。请你帮助它从位置 0 出发，到达它的家。

跳蚤跳跃的规则如下：

它可以 往前 跳恰好 a 个位置（即往右跳）。
它可以 往后 跳恰好 b 个位置（即往左跳）。
它不能 连续 往后跳 2 次。
它不能跳到任何 forbidden 数组中的位置。
跳蚤可以往前跳 超过 它的家的位置，但是它 不能跳到负整数 的位置。

给你一个整数数组 forbidden ，其中 forbidden[i] 是跳蚤不能跳到的位置，同时给你整数 a， b 和 x ，请你返回跳蚤到家的最少跳跃次数。如果没有恰好到达 x 的可行方案，请你返回 -1 。



示例 1：

输入：forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
输出：3
解释：往前跳 3 次（0 -> 3 -> 6 -> 9），跳蚤就到家了。
示例 2：

输入：forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
输出：-1
示例 3：

输入：forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
输出：2
解释：往前跳一次（0 -> 16），然后往回跳一次（16 -> 7），跳蚤就到家了。


提示：

1 <= forbidden.length <= 1000
1 <= a, b, forbidden[i] <= 2000
0 <= x <= 2000
forbidden 中所有位置互不相同。
位置 x 不在 forbidden 中。
*/

func MinimumJumps(forbidden []int, a int, b int, x int) (res int) {
	visit := make(map[int]bool)
	for _, v := range forbidden {
		visit[v] = true
	}
	type info struct {
		index int
		left  bool
	}
	queue := [][2]int{[2]int{0, 0}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			in, left := queue[i][0], queue[i][1]
			if in == x {
				return
			}
			new_pos := 0
			if left == 0 {
				new_pos = in - b
				if new_pos >= 0 && !visit[new_pos] {
					queue = append(queue, [2]int{new_pos, 1})
					visit[new_pos] = true
				}
			}
			new_pos = in + a
			if new_pos <= 6000 && !visit[new_pos] {
				queue = append(queue, [2]int{new_pos, 0})
				visit[new_pos] = true
			}

		}
		res++
		queue = queue[size:]
	}
	return -1
}

func MinimumJumps2(forbidden []int, a int, b int, x int) int {

	visited := make([]bool, 2001+a+b)
	for i := range forbidden {
		visited[forbidden[i]] = true
	}

	queue := [][2]int{[2]int{0, 0}}
	res := -1
	for len(queue) > 0 {

		length := len(queue)
		res++
		for i := 0; i < length; i++ {
			cur, isBack := queue[i][0], queue[i][1]
			if cur == x {
				return res
			}
			if isBack == 0 && cur-b > 0 && !visited[cur-b] {
				visited[cur-b] = true
				queue = append(queue, [2]int{cur - b, 1})
			}
			if cur+a < len(visited) && !visited[cur+a] {
				visited[cur+a] = true
				queue = append(queue, [2]int{cur + a, 0})
			}
		}

		queue = queue[length:]
		fmt.Printf("queue: %v\n", queue)
	}

	return -1
}

// for i := 0; i < x+a; {
// 	if i == x {
// 		break
// 	}
// 	_, ok := set[i]
// 	if ok {
// 		res = -1
// 		break
// 	}
// 	if flg {
// 		res++
// 		flg = false
// 		i += a
// 	} else {
// 		if (i > x && i-b >= 0) || i > b {
// 			res++
// 			flg = true
// 			i -= b
// 		} else {
// 			res++
// 			flg = false
// 			i += a
// 		}
// 	}
// }
