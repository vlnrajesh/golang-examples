package StringExamples

import "fmt"

// IntToString Converts  Integer Value to String
func IntToString(i int) string {
	return fmt.Sprintf("%v", i)
}

//ReverseString returns given string in reverse
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//SwapStrings to swap strings passed as parameters
func SwapStrings(a *string, b *string) {
	temp := *a
	*a = *b
	*b = temp
}
