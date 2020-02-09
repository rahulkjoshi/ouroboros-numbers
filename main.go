package main

import (
	"flag"
	"fmt"
	"os"

	"ouroboros"
)

var (
	numDigits  = flag.Int("num_digits", 2, "Length of the number.")
	multiplier = flag.Int("multiplier", 6, "Number to multiply by.")
)

func main() {
	flag.Parse()
	if *numDigits > 26 {
		fmt.Printf("--num_digits must be less than or equal to 26: %v\n", *numDigits)
		os.Exit(1)
	}

	ouroboros.PrintHeader(*numDigits, *multiplier)
}
