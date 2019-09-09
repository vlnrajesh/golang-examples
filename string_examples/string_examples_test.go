package StringExamples

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	actualResult := ReverseString("HI")
	var expectedResult = "IH"
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestIntToString(t *testing.T) {
	actualResult := IntToString(4)
	expectedResult := "4"
	if actualResult != expectedResult {
		t.Fatalf("Expected %T Type but got %T", expectedResult, actualResult)
	}
}
