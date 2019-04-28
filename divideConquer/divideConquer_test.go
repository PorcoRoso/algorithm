package divideConquer

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{6,-1,2,0,3,-6}
	s := maxSubArray(nums)
	if s != 10 {
		t.Errorf("max substring is : %v", s)
	}
}