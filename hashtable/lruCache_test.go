package hashtable

import "testing"

func TestGet(t *testing.T) {
	lc := Constructor(2)
	lc.Put(12, 1)
	if 1 == lc.Get(12) {
		t.Log("get method success.")
	} else {
		t.Error("get method error.")
	}
}
