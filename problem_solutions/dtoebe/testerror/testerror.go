package testerror

import "testing"

func ErrorOut(fn, exp, res string, t *testing.T) {
	if exp != res {
		t.Errorf("Func: %s; Expected: %s; Got: %s\n", fn, exp, res)
	}
}
