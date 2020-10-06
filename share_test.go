package main

import (
	"testing"
)

type (
	testStruct struct {
		A string
		B int
		C testInner
		e int
		F map[string]int
	}
	testInner struct {
		D int
	}
)

func TestConvert(t *testing.T) {

	s := new(testStruct)

	s.A = "foo"
	s.B = 0101
	s.C = testInner{D: 2}
	s.e = 1337

	s.F = make(map[string]int)
	s.F["abc"] = 1
	s.F["def"] = -1

	share, err := ShareString(s)
	if err != nil {
		t.Error("Encode error", err)
	}
	u := new(testStruct)
	err = UnshareString(share, u)
	if err != nil {
		t.Error("Decode error", err)
	}
	if s.e == u.e {
		t.Error("Matching private information? Shouldn't be possible!")
	}
	if s.A != u.A || s.B != u.B || s.C != u.C {
		t.Error("Simple data doesn't match:", s, u)
	}
	uAbc, uAbcOk := u.F["abc"]
	uDef, uDefOk := u.F["def"]

	if !(uAbcOk && uDefOk) {
		t.Error("Unshare is missing map data:", s, u)
	}

	if uAbc != s.F["abc"] || uDef != s.F["def"] {
		t.Error("Map data doesn't match")
	}
}
