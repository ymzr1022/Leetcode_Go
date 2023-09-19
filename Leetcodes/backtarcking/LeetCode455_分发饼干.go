package leetcodes

import (
	"sort"
)

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	m := len(g)
	if len(s) < len(g) {
		m = len(s)
	}
	for ; j < len(s); j++ {
		if i < m && g[i] <= s[j] {
			i++
		}
	}
	return i
}
