package test

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := GetHello("kngfjsr")
	expect := "hello,kngfjsr"
	if result != expect {
		t.Error("\n result:", result, "expect:", expect)
	}
	t.Log("finish TestHello")
}
