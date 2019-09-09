package ArrayExamples

import (
	"reflect"
	"testing"
)

func TestReverseIntArray(t *testing.T) {
	actualResult := ReverseIntArray([]int{4, 1, 2})
	expectedResult := []int{2, 1, 4}
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Reversal of Integer Array is failed")
	}
}
