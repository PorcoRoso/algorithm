package greedy

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// 有一个例子没通过的原因是，obstacles存入map的值冲突，把系数从10000改为100000即可
func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][2]int{{0,1},{1,0},{0,-1},{-1,0}} // North, East, South, West

	obstacle := map[int]int{}
	for i, v := range obstacles {
		if len(v) != 0 {
			obstacle[obstacles[i][0]*100000 + obstacles[i][1]] = 1
		}
	}

	x, y, dir := 0, 0, 0
	square := 0
	for _, v := range commands {
		if v == -1 {
			dir++
		} else if v == -2 {
			dir--
		} else {
			for ; v > 0; v-- {
				d := dir % 4
				if d < 0 {
					d += 4
				}
				if _, ok := obstacle[(x+dirs[d][0])*100000 + y+dirs[d][1]]; ok {
					break
				}
				x += dirs[d][0]
				y += dirs[d][1]

				square = max(square, x*x + y*y)
			}
		}

	}

	return int(square)
}
