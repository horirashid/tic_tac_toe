<<<<<<< HEAD
git
=======
package main

import "testing"

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0,0,0,0,0,0,0,0,0},
	}
	b.put(1,1, "o")
	if b.get(1,1) != "o" {
		t.Errorf("The test case failed")
	}
}
>>>>>>> 8005c32023d899ca73c5c12799b087561aeedf95
