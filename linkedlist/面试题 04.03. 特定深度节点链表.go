package linkedlist

/*
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。



示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
通过次数
35.9K
提交次数
44.5K
通过率
80.8%
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func listOfDepth(tree *TreeNode) (res []*ListNode) {

	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	res = append(res, &ListNode{Val: tree.Val})
	for len(queue) > 0 {
		size := len(queue)
		head := &ListNode{
			Val: queue[0].Val,
		}
		pre := head
		for i := 0; i < size; i++ {
			tnode := queue[i]
			if tnode.Left != nil {
				queue = append(queue, tnode.Left)
			}
			if tnode.Right != nil {
				queue = append(queue, tnode.Right)
			}
			if i > 0 {
				cur := &ListNode{
					Val: tnode.Val,
				}
				pre.Next = cur
				pre = cur
			}
		}
		queue = queue[size:]
		res = append(res, head)
	}
	return
}
