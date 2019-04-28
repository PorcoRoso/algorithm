package dynamicPrograming

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res := minCostClimbingStairs(cost)
	if res == 6 {
		t.Log("pass")
	} else {
		t.Errorf("error res: %v", res)
	}
}

func TestWordBreak(t *testing.T) {
	s := "catsand"
	wordDict := []string{"cats", "and"}
	res := wordBreak(s, wordDict)
	if res {
		t.Log("pass")
	} else {
		t.Errorf("error")
	}
}