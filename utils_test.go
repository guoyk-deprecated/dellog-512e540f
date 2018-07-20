package main

import (
	"testing"
	"time"
)

func TestStringStrSliceContains(t *testing.T) {
	if !StrSliceContains([]string{"a", "b", "c"}, "c") {
		t.Fatal()
	}
	if StrSliceContains([]string{"a", "b", "c"}, "d") {
		t.Fatal()
	}
}

func TestFindDateMark(t *testing.T) {
	n, ok := FindDateMark("deep/dark/file/path.2018-08-19/hello.world.2018-07-10.log")
	if !ok {
		t.Fatal("failed")
	}
	if n.Day() != 10 || n.Month() != time.July {
		t.Fatal("failed")
	}
}
