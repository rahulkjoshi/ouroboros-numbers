package ouroboros

import (
	"bytes"
	"strconv"
	"strings"
)

func toString(digits []int, numDigits int) string {
	ret := make([]byte, 0, numDigits)
	for i := len(digits) - 1; i >= 0; i-- {
		ret = strconv.AppendInt(ret, int64(digits[i]), 10)
	}
	ret = append(ret, bytes.Repeat([]byte{'_'}, numDigits-len(digits))...)
	return string(ret)
}

func Solve(numDigits, mult int, allowLeadingZero bool) string {
	var min string
	for val := 1; val < 10; val++ {
		digits := make([]int, 1, numDigits)
		digits[0] = val
		r := 0
		for i := 1; i < numDigits; i++ {
			t := (digits[i-1] * mult) + r
			digits = append(digits, t%10)
			r = t / 10
		}
		if allowLeadingZero || digits[numDigits-1] != 0 {
			if (digits[numDigits-1]*mult)+r == digits[0] {
				solution := toString(digits, numDigits)
				if min == "" || strings.Compare(min, solution) > 0 {
					min = solution
				}
			}
		}
	}
	return min
}

func SolveToMax(max, mult int, allowLeadingZero bool) (string, int) {
	for i := 1; i <= max; i++ {
		val := Solve(i, mult, allowLeadingZero)
		if val != "" {
			return val, i
		}
	}
	return "", 0
}
