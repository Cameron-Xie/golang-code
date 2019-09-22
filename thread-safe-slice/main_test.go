package main

import (
	"testing"
)

func TestLockAppend(t *testing.T) {
	num := 10
	s := LockAppend(make([]int, 0), num)

	if l := len(s); l != num {
		t.Errorf("expected %d, got %d", num, l)
	}
}

func TestChannelAppend(t *testing.T) {
	num := 10
	s := ChannelAppend(make([]int, 0), num)

	if l := len(s); l != num {
		t.Errorf("expected %d, got %d", num, l)
	}
}
