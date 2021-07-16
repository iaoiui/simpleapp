package main

import "testing"

func TestRun(t *testing.T) {
	ret := Run()
	if ret != 0 {
		t.Error("Failed: return code is", ret)
	}
}
