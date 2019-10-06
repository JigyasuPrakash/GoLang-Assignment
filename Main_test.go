package main

import "testing"

func TestMain(t *testing.T) {
	value := "this is test"
	start := 2
	end := 9
	mySubString := subString(value, start, end)
	myLength := length(value)
	myReverse := reverse(value)
	var expectedResult = "is is t"
	var expectedReverse = "tset si siht"
	if mySubString != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, mySubString)
	}
	if myLength != 12 {
		t.Fatalf("Expected %d but got %d", 12, myLength)
	}
	if myReverse != expectedReverse {
		t.Fatalf("Expected '%s' but got '%s'", expectedReverse, myReverse)
	}
}
