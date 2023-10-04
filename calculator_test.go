package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(2, 3)

	if total != 5 {
		t.Error("The result must be 5.")
	}
}

func TestSub(t *testing.T) {
	total := sub(3, 3)

	if total != 0 {
		t.Error("The result must be 0.")
	}
}

func TestTimes(t *testing.T) {
	total := times(3, 3)

	if total != 9 {
		t.Error("The result must be 9.")
	}
}

func TestSumX(t *testing.T) {
	total := sumX(3, 3)

	if total != 9 {
		t.Error("The result must be 9.")
	}
}
