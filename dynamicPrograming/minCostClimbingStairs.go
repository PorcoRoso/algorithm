package dynamicPrograming

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
// dp
// step[i] = min(step[i-1]+cost[i-1], step[i-2]+cost[i-2])
func minCostClimbingStairs(cost []int) int {
	step := []int{}
	step = append(step, 0)
	step = append(step, 0)
	for i := 2; i <= len(cost); i++ {
		tmp := min(step[i-1]+cost[i-1], step[i-2]+cost[i-2])
		step = append(step, tmp)
	}

	return step[len(cost)]
}