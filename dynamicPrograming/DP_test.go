package dynamicPrograming

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res := minCostClimbingStairs(cost)
	if res == 15 {
		t.Log("pass")
	} else {
		t.Errorf("error res: %v", res)
	}
}
