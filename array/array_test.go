package array

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{2,3,-2,4}
	res := maxProductDP(nums)
	if res == 6 {
		t.Log("pass")
	} else {
		t.Errorf("error %v", res)
	}
}
