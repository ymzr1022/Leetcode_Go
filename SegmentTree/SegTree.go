package segmenttree

type SegTree struct {
	Maxn   int
	Arr    []int
	Sum    []int
	Lazy   []int
	Change []int
	Update []bool
}

func NewSegTree(orign []int) *SegTree {
	seg := new(SegTree)
	seg.Maxn = len(orign) + 1
	seg.Arr = make([]int, seg.Maxn)
	for i := 0; i < seg.Maxn-1; i++ {
		seg.Arr[i+1] = orign[i]
	}
	seg.Sum = make([]int, seg.Maxn<<2)
	seg.Lazy = make([]int, seg.Maxn<<2)
	seg.Change = make([]int, seg.Maxn<<2)
	seg.Update = make([]bool, seg.Maxn<<2)
	return seg
}

func (seg *SegTree) PushedUp(rt int) {
	seg.Sum[rt] = seg.Sum[rt<<1] + seg.Sum[rt<<1|1]
}

func (seg *SegTree) Build(l, r, rt int) {
	if l == r {
		seg.Sum[rt] = seg.Arr[l]
		return
	}
	mid := (l + r) >> 1
	seg.Build(l, mid, rt<<1)
	seg.Build(mid+1, r, rt<<1|1)
	seg.PushedUp(rt)
}

func (seg *SegTree) Add(l, r, C, L, R, rt int) {
	if L <= l && r <= R {
		seg.Sum[rt] += C * (r - l + 1)
		seg.Lazy[rt] += C
		return
	}
	mid := (l + r) >> 1
	seg.PushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		seg.Add(l, mid, C, L, R, rt<<1)
	}
	if R > mid {
		seg.Add(mid+1, r, C, L, R, rt<<1|1)
	}
	seg.PushedUp(rt)
}

func (seg *SegTree) Updates(l, r, C, L, R, rt int) {
	if L <= l && r <= R {
		seg.Update[rt] = true
		seg.Change[rt] = C
		seg.Sum[rt] = C * (r - l + 1)
		seg.Lazy[rt] = 0
		return
	}
	mid := (l + r) >> 1

	seg.PushDown(rt, mid-l+1, r-mid)

	if L <= mid {
		seg.Updates(l, mid, C, L, R, rt<<1)
	}
	if R > mid {
		seg.Updates(mid+1, r, C, L, R, rt<<1|1)
	}
	seg.PushedUp(rt)
}

func (seg *SegTree) Query(l, r, L, R, rt int) int {
	if L <= l && r <= R {
		return seg.Sum[rt]
	}
	sum := 0
	mid := (l + r) >> 1
	seg.PushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		sum += seg.Query(l, mid, L, R, rt<<1)
	}
	if R > mid {
		sum += seg.Query(mid+1, r, L, R, rt<<1|1)
	}
	return sum
}

func (seg *SegTree) PushDown(rt, ln, rn int) {
	if seg.Update[rt] {
		seg.Update[rt<<1] = true
		seg.Update[rt<<1|1] = true
		seg.Change[rt<<1] = seg.Change[rt]
		seg.Change[rt<<1|1] = seg.Change[rt]
		seg.Lazy[rt<<1] = 0
		seg.Lazy[rt<<1|1] = 0
		seg.Sum[rt<<1] = seg.Change[rt] * ln
		seg.Sum[rt<<1|1] = seg.Change[rt] * rn
		seg.Update[rt] = false
	}
	if seg.Lazy[rt] != 0 {
		seg.Lazy[rt<<1] += seg.Lazy[rt]
		seg.Sum[rt<<1] += seg.Lazy[rt] * ln
		seg.Lazy[rt<<1|1] += seg.Lazy[rt]
		seg.Sum[rt<<1|1] += seg.Lazy[rt] * rn
		seg.Lazy[rt] = 0
	}

}
