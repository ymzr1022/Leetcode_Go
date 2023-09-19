package main

import (
	"fmt"

	unionfindset "ymzr1022.com/leetcode/UnionFIndSet"
	"ymzr1022.com/leetcode/linkedlist"
)

func main() {

	// 6, 3, 3, 5,7,2,8,3
	// 6  ， 5，  4， 3 ， 7
	/*
		4，5，1，2，3，6，0
	*/
	// fmt.Printf("leetcodes.NumRescueBoats(nums, 4): %v\n", leetcodes.NumRescueBoats(nums, 5))
	// fmt.Printf("leetcodes.FindContentChildren(nums, nu): %v\n", leetcodes.FindContentChildren(nums, nu))
	// flgs := make([]bool, 2)
	// fmt.Printf("flgs: %v\n", flgs)
	// fmt.Printf("leetcodes.Combine(4, 2): %v\n", leetcodes.Combine(4, 2))
	// fmt.Printf("leetcodes.CombinationSum(nums, 8): %v\n", leetcodes.CombinationSum4(nums, 24))
	// fmt.Printf("leetcodes.CombinationSum(nums, 6): %v\n", leetcodes.CombinationSum(nums, 5))
	// fmt.Printf("leetcodes.CombinationSum3(3, 7): %v\n", leetcodes.CombinationSum3(2, 18))
	// fmt.Printf("leetcodes.CombinationSum2(nums, 8): %v\n", leetcodes.CombinationSum2(nums, 5))
	// fmt.Printf("leetcodes.Get(4, 2): %v\n", leetcodes.Get(24, []int{20, 2, 2}))
	// fmt.Printf("leetcodes.Check(\"aba\"): %v\n", leetcodes.Check("aba"))
	// fmt.Printf("leetcodes.Partition(\"aaab\"): %v\n", leetcodes.Partition("aaabb"))
	// s := "123"
	// if s[:] > "122" {
	// 	fmt.Print("1")
	// 	fmt.Print(s[2])
	// }
	// fmt.Printf("leetcodes.RestoreIpAddresses(\"101023\"): %v\n", leetcodes.RestoreIpAddresses("101023"))
	// fmt.Printf("leetcodes.Check(\"\", 3, 5): %v\n", leetcodes.Check("101023", 3, 5))
	// fmt.Printf("leetcodes.Subsets([]int{1, 2, 3}): %v\n", leetcodes.SubsetsWithDup([]int{1, 2, 2}))
	// fmt.Printf("leetcodes.DistributeCookies(nums, 2): %v\n", leetcodes.CuttingRope(12))
	// str := "abc"
	// bys := []byte(str)
	// fmt.Printf("int(bys[0]): %v\n", int(bys[0]))
	// strs := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// words2 := []string{"le", "eo"}
	// dpp := [][]int{{0, 1}, {1, 0}}
	// fmt.Printf("leetcodes.WordSubsets(strs, words2): %v\n", leetcodes.LemonadeChange(nums))
	// maps := make(map[string]string)
	// fmt.Printf("maps[\"0\"]: %v\n", maps["0"] == "")
	// sentence1 := []string{"I", "love", "leetcode"}
	// sentence2 := []string{"I", "love", "onepiece"}
	// similarPairs := [][]string{{"manga", "onepiece"}, {"platform", "anime"}, {"leetcode", "platform"}, {"anime", "manga"}}

	// unionfindset.AreSentencesSimilarTwo(sentence1, sentence2, similarPairs)
	// grid := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	// fmt.Printf("dp.MinPathSum(grid): %v\n", )
	// ss := 1
	// for i := 0; i < 9; i++ {
	// 	ss *= 10
	// }
	// ss += 7

	// fmt.Printf("ss: %v\n", ss)
	// mat := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	// fmt.Printf("algorithm.GetNext(\"aab\"): %v\n", dp.LengthOfLIS(nums, 3))
	// seg := segmenttree.NewSegTree(nums)
	// seg.Build(1, 2, 1)
	// fmt.Printf("seg.Sum: %v\n", seg.Sum)
	// seg.Updates(1, 2, 3, 1, 1, 1)
	// fmt.Printf("seg.Sum: %v\n", seg.Sum)
	// fmt.Printf("seg.Query(1, 2, 1, 2, 1): %v\n", seg.Query(1, 4, 1, 2, 1))
	// fmt.Printf("seg.Sum: %v\n", seg.Sum)
	// root := tree.TreeNode{
	// 	Val: 1,
	// }
	// root.Left = &tree.TreeNode{
	// 	Val: 2,
	// }
	// root.Right = &tree.TreeNode{
	// 	Val: 2,
	// }
	// root.Left.Right = &tree.TreeNode{
	// 	Val: 3,
	// }
	// root.Left.Right.Left = &tree.TreeNode{
	// 	Val: 4,
	// }
	// root.Left.Right.Right = &tree.TreeNode{
	// 	Val: 4,
	// }
	// root.Left.Right.Left.Left = &tree.TreeNode{
	// 	Val: 5,
	// }
	// fmt.Printf("tree.LongestZigZag(&root): %v\n", tree.LongestZigZag(&root))
	// fmt.Printf("dp.NumOfSubarrays(nums): %v\n", dp.NthSuperUglyNumber(12, nums))

	// mm := myMap.NewMyConMap()
	// mm.Put(1, 2)
	// fmt.Println(mm.Get(1, 2))
	// fmt.Println(mm.Get(2, 2000))
	// go func() {
	// 	time.Sleep(1000)
	// 	mm.Put(2, 5)
	// }()
	// fmt.Printf("binarysearch.FindPeakElement(nums): %v\n", contest.LongestEqualSubarray(nums, 3))
	// fmt.Printf("dp.GetNum(\"01010101\"): %v\n", dp.GetNum("111000"))
	// na := segmenttree.Constructor1(nums)
	// fmt.Printf("na.Sums: %v\n", na.Sums)
	// na.Update(0, 3)
	// fmt.Printf("na.Sums: %v\n", na.Sums)
	// seg.Updates(1, 2, 3, 1, 1, 1)
	// fmt.Printf("seg.Sum: %v\n", seg.Sum)
	// fmt.Printf("seg.Query(2, 2, 1, 2, 1): %v\n", seg.Query(1, 2, 2, 2, 1))
	// fmt.Printf("na.Sums: %v\n", na.Sums)
	// fmt.Printf("na.SumRange(1, 1): %v\n", na.SumRange(1, 1))
	// rfq := segmenttree.Constructor(nums)
	// fmt.Printf("rfq.Times: %v\n", rfq)
	// rfq.Query(0, 11, 33)
	// fmt.Printf("dp.FindNumberOfLIS(nums): %v\n", dp.FindNumberOfLIS(nums))

	// mc := segmenttree.ConstructorC()

	// fmt.Printf("mc.Book(10, 20): %v\n", mc.Book(15, 23))
	// fmt.Printf("mc.Book(10, 20): %v\n", mc.Book(17, 26))
	// fmt.Printf("mc.Book(10, 20): %v\n", mc.Book(30, 36))
	// fmt.Printf("mc.Book(10, 20): %v\n", mc.Book(32, 40))
	// // // fmt.Printf("greedyalgorithm.MinOperations(nums, 10): %v\n", stack.FindUnsortedSubarray(nums))
	// // // fmt.Println(1 << 10)
	// // // set := make(map[int]struct{})
	// // // set[0] = struct{}{}
	// // // set[1] = struct{}{}
	// // // fmt.Printf("set: %v\n", set)
	// // // fmt.Printf("set[0]: %v\n", set[0])
	// // // fmt.Printf("set[2]: %v\n", set[2])
	// // // _, ok := set[1]
	// // // fmt.Printf("set: %v\n", ok)
	// // fmt.Printf("dp.MinimumJumps(nums, 3, 15, 9): %v\n", stack.LargestRectangleArea(nums))
	// // fmt.Printf("stringss.MyAtoi(\" -0999 sa\"): %v\n", stringss.MyAtoi(" -0999 sa"))
	// // var x int = 91283472332
	// // x = 1<<31 - 1
	// // x = -1 << 31 abcdf
	// //          fbadc
	// // fmt.Printf("x: %v\n", x)
	// // fmt.Printf("contest.CanBeEqual(\"ifjz\", \"jzfi\"): %v\n", contest.CanBeEqual("ifjz", "jzfi"))
	// fmt.Printf("contest.MaxSum(nums, 3, 3): %v\n", contest.MaxSum(nums, 3, 4))

	// grid := [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}}
	// fmt.Printf("dp.MinFallingPathSums(grid): %v\n", dp.MinFallingPathSums(grid))
	// s := skiplists.NewSkiplist()
	// s.Put(1, 1)
	// fmt.Println(s.Get(1))
	// s.Put(2, 4)
	// s.Put(2, 3)
	// s.Put(4, 6)
	// fmt.Println(s.Get(4))
	// fmt.Println(s.Get(5))

	// t := trie.Constructor()
	// t.Insert("hotdog")

	// t.Search("app")
	// t.StartsWith("app")

	// l := lru.Constructor(2)
	// l.Put(2, 1)
	// l.Put(1, 1)
	// l.Put(2, 3)96 + 480 576
	// fmt.Printf("l.Get(1): %v\n", l.Get(1))
	// l.Put(3, 5)
	// fmt.Printf("l.Get(2): %v\n", l.Get(2))
	// nums := []int{4, 5, 1, 9, 2}
	// fmt.Printf("greedyalgorithm.RepairCars(nums, 6): %v\n", greedyalgorithm.RepairCars(nums, 74))
	// nu := []int{2, 107, 109, 113, 127, 131, 137, 3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 47, 53}

	// mat := [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}
	// fmt.Printf("dfs.SpiralOrder(mat): %v\n", dfs.SpiralOrder(mat))
	// head := makeLink(nums)
	// linkedlist.ReverseKGroup(head, 2)
	// fmt.Printf("greedyalgorithm.ScheduleCourse(mat): %v\n", greedyalgorithm.ScheduleCourse(mat))
	// dp.MinimumMoves(mat)
	// dist := []int{0, 0, 1, -1}
	// speed := []int{1, 1}
	// composition := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	// f := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	// cost := []int{5, 5}
	// fmt.Printf("graph.NetworkDelayTime(mat, 4, 2): %v\n", dp.MinSwaps("1001"))
	// bfs.QueensAttacktheKing(mat, speed)
	// grap := [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}
	// mysorts.FindKthLargest(grap)
	// graph.IsBipartite(grap)
	// fmt.Printf("binarysearch.FindUnsortedSubarray(speed): %v\n", binarysearch.FindUnsortedSubarray(speed))
	// fmt.Printf("pointer.SmallestSubsequence(\"bcabc\"): %v\n", pointer.MinSumOfLengths(speed, 20))
	// fmt.Printf("contest.MinLengthAfterRemovals(speed): %v\n", contest.CountPairs1(speed, 5))
	// fmt.Printf("contest.MaxNumberOfAlloys(3, 2, 15, composition, speed, cost): %v\n", contest.MaxNumberOfAlloys(2, 3, 10, composition, speed, cost))
	// bfs.WatchedVideosByFriends(composition, f, 0, 1)
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Printf("unionfindset.CountSubIslands(grid1, grid2): %v\n", unionfindset.CountSubIslands(grid1, grid2))
}

func makeLink(arr []int) *linkedlist.ListNode {
	pre := new(linkedlist.ListNode)
	head := pre
	for _, v := range arr {
		tmp := &linkedlist.ListNode{
			Val: v,
		}
		pre.Next = tmp
		pre = tmp
	}
	return head.Next
}
