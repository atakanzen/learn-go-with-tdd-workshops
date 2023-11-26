package iteration

import "strings"

// Used in Repeat() function for repeat count
const repeatCount = 5

// Retrieves a string input and returns it by repeating it `repeatCount` times
// Uses string concatenation with += operator
func Repeat(input string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}
	return
}

// Retrieves a string input and an integer repeatCount and returns the repeated input repeatCount times
// Uses strings.Builder for more efficient processing
func RepeatWithSb(input string, repeatCount int) string {
	var sb strings.Builder
	for i := 0; i < repeatCount; i++ {
		sb.WriteString(input)
	}

	return sb.String()
}
