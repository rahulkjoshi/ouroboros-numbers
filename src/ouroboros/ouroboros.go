package ouroboros

import (
	"fmt"
	"strings"
)

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func rotStrSlice(s []string, k int) []string {
	if k < 0 || len(s) == 0 {
		return s
	}
	r := len(s) - k%len(s)
	return append(s[r:], s[:r]...)
}

func PrintHeader(numDigits, mult int) {
	var numAlpha []string
	for i := 1; i <= numDigits; i++ {
		numAlpha = append(numAlpha, toCharStr(i))
	}
	resultAlpha := rotStrSlice(numAlpha, 1)

	fmt.Printf("Solving for:\n")
	fmt.Printf("  %v\n", strings.Join(numAlpha, ""))
	fmt.Printf("x %[1]*d\n", numDigits, mult)
	fmt.Printf("%v\n", strings.Repeat("-", numDigits + 2))
	fmt.Printf("  %v\n", strings.Join(resultAlpha, ""))
}
