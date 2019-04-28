package greedy

import "testing"

func TestFindContentChildren(t *testing.T) {
	g := []int{1,2}
	s := []int{1,2,3}
	sum := findContentChildren(g, s)
	if sum != 2 {
		t.Errorf("false, sum is: %v", sum)
	}

}

func TestLemonadeChange(t *testing.T) {
	bills := []int{5,20,5,10,20}
	change := lemonadeChange(bills)
	if !change {
		t.Error("lemonnade change failed.")
	} else {
		t.Log("lemonade change success.")
	}
}

func TestRobotSim(t *testing.T) {
	commands := []int{4,-1,4,-2,4}
	obstacles := [][]int{{2,4}}
	square := robotSim(commands, obstacles)
	if square != 65 {
		t.Errorf("error square = %v", square)
	} else {
		i := -102 %4
		t.Logf("pass, %v", i)
	}
}

func TestJumpGame2(t *testing.T) {
	nums := []int{10,9,8,7,6,5,4,3,2,1,1,0}
	cnt := jump(nums)
	if cnt == 2 {
		t.Log("pass")
	} else {
		t.Errorf("failed %v", cnt)
	}
}