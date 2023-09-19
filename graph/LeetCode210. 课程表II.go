package graph

/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。



示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2：

输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
示例 3：

输入：numCourses = 1, prerequisites = []
输出：[0]

0,1,0,0,-1,1,0,0,0,-1,0
提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
所有[ai, bi] 互不相同
*/

func FindOrder(numCourses int, prerequisites [][]int) []int {
	gra := make([][]int, numCourses)
	ins := make([]int, numCourses)
	for _, v := range prerequisites {
		gra[v[1]] = append(gra[v[1]], v[0])
		ins[v[0]]++
	}
	visit := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		if onPath[i] {
			hasCycle = true
		}
		if visit[i] || hasCycle {
			return
		}
		visit[i] = true
		onPath[i] = true
		for _, v := range gra[i] {
			dfs(v)
		}
		path = append(path, i)
		onPath[i] = false
	}

	// for i := 0; i < numCourses; i++ {
	// 	dfs(i)
	// }
	// if hasCycle {
	// 	return []int{}
	// } else {
	// 	reverse(path)
	// 	return path
	// }
	queue := make([]int, 0)
	count := 0
	for i, v := range ins {
		if v == 0 {
			queue = append(queue, i)
			count++
			path = append(path, i)
		}
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			ind := queue[i]
			for _, v := range gra[ind] {
				ins[v]--
				if ins[v] == 0 {
					queue = append(queue, v)
					count++
					path = append(path, v)
				}
			}
		}
		queue = queue[size:]
	}
	if count == numCourses {
		return path
	} else {
		return []int{}
	}
}

// func MakeGraph(Grids [][]int) (res [][]int) {

// return
// }
func Reverse(path []int) {
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
}
