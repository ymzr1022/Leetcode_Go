package segmenttree

var Sum []int

func Init_(n, m int) {
	Sum = make([]int, 4*n)
}

//在idx位置的值加上val  add(1,1,n,idx,val)
func Add(o, l, r, idx, val int) {
	if l == r {
		Sum[l] += val
		return
	}
	mid := (l + r) / 2
	if idx <= mid {
		Add(o*2, l, mid, idx, val)
	} else {
		Add(0*2+1, mid+1, r, idx, val)
	}
	Sum[o] = Sum[o*2] + Sum[o*2+1]
}

//查询L-R的Sum  query(1,1,n,L,R)
func Query(o, l, r, L, R int) int {
	if L <= l && r <= R {
		return Sum[o]
	}
	Sum := 0
	mid := (l + r) / 2
	if L <= mid {
		Sum += Query(o*2, l, mid, L, R)
	}
	if R > mid {
		Sum += Query(0*2+1, mid+1, r, L, R)
	}
	return Sum
}

// func main() {
// 	Init_(10, 0)
// 	Sum = append(Sum, []int{0, 2, 3, 4, 5, 2, 6, 1, 9, 10}...)
// 	i := Query(1, 1, 10, 2, 3)
// 	fmt.Printf("i: %v\n", i)
// }
