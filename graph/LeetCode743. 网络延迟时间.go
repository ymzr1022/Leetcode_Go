package graph

import (
	"container/heap"
	"fmt"
)

/*
有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。



示例 1：



输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2
示例 2：

输入：times = [[1,2,1]], n = 2, k = 1
输出：1
示例 3：

输入：times = [[1,2,1]], n = 2, k = 2
输出：-1


提示：

1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
所有 (ui, vi) 对都 互不相同（即，不含重复边）
*/

func NetworkDelayTime(times [][]int, n int, k int) (res int) {
	gra := make([][][]int, n+1)
	for _, v := range times {
		gra[v[0]] = append(gra[v[0]], []int{v[1], v[2]})
	}

	var h hp
	dir := make([]int, n+1)
	for i := range dir {
		dir[i] = 1<<31 - 1
	}
	visit := make([]bool, n+1)
	dir[k] = 0
	heap.Push(&h, &pair{k, 0})
	visit[k] = true
	for h.Len() > 0 {

		p := heap.Pop(&h).(*pair)
		if dir[p.node] < p.val {
			continue
		}
		for _, v := range gra[p.node] {
			val := v[1]
			node := v[0]

			if dir[p.node]+val < dir[node] {
				dir[node] = dir[p.node] + val
				heap.Push(&h, &pair{node: node, val: dir[node]})
			}
		}

	}
	fmt.Printf("dir: %v\n", dir)
	cnt := 0
	for _, v := range dir {
		if v != 1<<31-1 {
			res = max(res, v)
		} else {
			cnt++
		}
	}
	if cnt > 1 {
		return -1
	}
	return
}

type pair struct {
	node int
	val  int
}

type hp struct {
	pairs []*pair
}

func (h hp) Len() int { return len(h.pairs) }

func (h hp) Less(i, j int) bool {
	return h.pairs[i].val < h.pairs[j].val
}

func (h hp) Swap(i, j int) {
	h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i]
}

func (h *hp) Push(pairt any) {
	h.pairs = append(h.pairs, pairt.(*pair))
}

func (h *hp) Pop() any {
	tmp := h.pairs[0]
	old := h.pairs
	h.pairs = old[1:]
	heap.Init(h)
	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
