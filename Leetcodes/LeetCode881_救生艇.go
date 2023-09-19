package leetcodes

import "sort"

func NumRescueBoats(people []int, limit int) int {
	n := len(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})
	// sort.Ints(people)
	i := 0
	j := n - 1
	sum := 0
	for i <= j {
		if people[i]+people[j] <= limit {
			j--
		}
		i++
		sum++
	}

	return sum
}
