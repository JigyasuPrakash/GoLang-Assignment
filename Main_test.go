package main

import "testing"

func TestMain(t *testing.T)  {
	if subString("this is test",2,9) != "his is t" {
		t.Error("Substring test failed!")
	}
}