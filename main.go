package main

import (
	"flag"
	"fmt"

	"ouroboros"
)

var (
	max              = flag.Int("max", 2, "Max length of number.")
	multiplier       = flag.Int("multiplier", 6, "Number to multiply by.")
	allowLeadingZero = flag.Bool("allow_leading_zero", false,
		"Whether to accept solutions with a leading zero (e.g. 025641 * 4 = 102564)")
)

func main() {
	flag.Parse()

	val, digits := ouroboros.SolveToMax(*max, *multiplier, *allowLeadingZero)
	if val == "" {
		fmt.Printf("Cannot be solved in %d digits.\n", *max)
		return
	}
	fmt.Printf("Found value: %v (%d digits)\n", val, digits)
}
