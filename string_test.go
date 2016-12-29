package yphp

import "testing"

func Test_Lcfirst(t *testing.T) {
	str := "Test"
	exp := Lcfirst(str)
	if exp != "test" {
		t.Error("lcfirst error:", exp, " != test")
	}
}
