package main

import "testing"

func Test01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("The test case failed")
	}
}

//isTripleEqual
func Test02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 1, 1, 1, 0, 0, 0},
	}

	if b.isTripleEqual(3, 4, 5) != true {
		t.Errorf("The test case failed")
	}
}

//wo jiaoyixia2 check
func Test03(t *testing.T) {
	b := &Board{
		tokens: []int{1, 1, 1, 0, 0, 0, 0, 0, 0},
	}
	if b.check() != true {
		t.Errorf("The test case failed")
	}
}
