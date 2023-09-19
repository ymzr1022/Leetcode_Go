package segmenttree

/*
给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）


示例 1：

输入：
["NumArray", "sumRange", "update", "sumRange"]
[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
输出：
[null, 9, null, 8]

解释：
NumArray numArray = new NumArray([1, 3, 5]);
numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
numArray.update(1, 2);   // nums = [1,2,5]
numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8


提示：

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
0 <= index < nums.length
-100 <= val <= 100
0 <= left <= right < nums.length
调用 update 和 sumRange 方法次数不大于 3 * 104
*/

type NumArray struct {
	Lens    int
	Array   []int
	Sums    []int
	Lazy    []int
	Change  []int
	Updatee []bool
}

func Constructor1(nums []int) NumArray {
	arr := NumArray{
		len(nums) + 1,
		make([]int, len(nums)+1),
		make([]int, (len(nums)+1)*4),
		make([]int, (len(nums)+1)*4),
		make([]int, (len(nums)+1)*4),
		make([]bool, (len(nums)+1)*4),
	}
	for i := 0; i < len(nums); i++ {
		arr.Array[i+1] = nums[i]
	}
	arr.Build(1, len(nums), 1)
	return arr
}

func (n *NumArray) Build(L, R, rt int) {
	if L == R {
		n.Sums[rt] = n.Array[L]
		return
	}
	mid := L + (R-L)>>1
	n.Build(L, mid, rt<<1)
	n.Build(mid+1, R, rt<<1|1)
	n.Sums[rt] = n.Sums[rt<<1] + n.Sums[rt<<1|1]
}

func (n *NumArray) Updates(l, r, L, R, C, rt int) {
	if L <= l && r <= R {
		n.Updatee[rt] = true
		n.Change[rt] = C
		n.Sums[rt] = C * (r - l + 1)
		n.Lazy[rt] = 0
		return
	}
	mid := (l + r) >> 1
	n.PushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		n.Updates(l, mid, L, R, C, rt<<1)
	}
	if R > mid {
		n.Updates(mid+1, r, L, R, C, rt<<1|1)
	}
	n.Sums[rt] = n.Sums[rt<<1] + n.Sums[rt<<1|1]
}

func (n *NumArray) PushDown(rt, ln, rn int) {
	if n.Updatee[rt] {
		n.Updatee[rt<<1] = true
		n.Updatee[rt<<1|1] = true
		n.Change[rt<<1] = n.Change[rt]
		n.Change[rt<<1|1] = n.Change[rt]
		n.Sums[rt<<1] = n.Change[rt] * ln
		n.Sums[rt<<1|1] = n.Change[rt] * rn
		n.Lazy[rt<<1] = 0
		n.Lazy[rt<<1|1] = 0
		n.Updatee[rt] = false
	}
	if n.Lazy[rt] != 0 {
		n.Lazy[rt<<1] = n.Lazy[rt]
		n.Lazy[rt<<1|1] = n.Lazy[rt]
		n.Sums[rt<<1] += n.Lazy[rt] * ln
		n.Sums[rt<<1|1] += n.Lazy[rt] * rn
		n.Lazy[rt] = 0
	}
}

func (n *NumArray) Query(l, r, L, R, rt int) int {
	if L <= l && r <= R {
		return n.Sums[rt]
	}
	mid := (l + r) >> 1
	sum := 0
	n.PushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		sum += n.Query(l, mid, L, R, rt<<1)
	}
	if R > mid {
		sum += n.Query(mid+1, r, L, R, rt<<1|1)
	}
	return sum
}

func (n *NumArray) Update(index int, val int) {
	n.Updates(1, n.Lens-1, index+1, index+1, val, 1)
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.Query(1, n.Lens-1, left+1, right+1, 1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
