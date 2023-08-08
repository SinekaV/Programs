//WAP that identifies the no of occurances of the substring in a given string
package main

import (
	"fmt"
	"strings"
)

func SubstringOcc(input string, substringLength int) string {
	if len(input) < substringLength {
		return ""
	}

	substringCounts := make(map[string]int)
	maxSubstring := ""
	maxCount := 0

	for i := 0; i <= len(input)-substringLength; i++ {
		substr := input[i : i+substringLength]
		substringCounts[substr]++
		if substringCounts[substr] > maxCount {
			maxCount = substringCounts[substr]
			maxSubstring = substr
		}
	}

	return maxSubstring
}

func main() {
	input := "abcbcadefghabceijab"
	substringLength := 3

	maxSubstring := SubstringOcc(input, substringLength)
	if maxSubstring != "" {
		fmt.Printf("The most frequent %d-character substring is: %s\n", substringLength, maxSubstring)
	} else {
		fmt.Printf("No valid %d-character substring found.\n", substringLength)
	}
}
