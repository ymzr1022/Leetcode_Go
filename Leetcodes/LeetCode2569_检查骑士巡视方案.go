package leetcodes

func checkValidGrid(grid [][]int) bool {
	type pair struct{ i, j int }
	pos := make([]pair, len(grid)*len(grid))
	for i, row := range grid {
		for j, x := range row {
			pos[x] = pair{i, j} // 记录坐标
		}
	}
	if pos[0].i != 0 || pos[0].j != 0 {
		return false
	}
	for i := 1; i < len(pos); i++ {
		flg1 := abs(pos[i].i-pos[i-1].i) == 1 && abs(pos[i].j-pos[i-1].j) == 2
		flg2 := abs(pos[i].i-pos[i-1].i) == 2 && abs(pos[i].j-pos[i-1].j) == 1
		if flg1 || flg2 {
			continue
		} else {
			return false
		}
	}
	// go func() {
	// 	fmt.Println("1")
	// }()
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
