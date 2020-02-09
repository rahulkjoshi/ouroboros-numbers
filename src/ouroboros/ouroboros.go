package ouroboros

import (
	"bytes"
	"strconv"
	"strings"
)

func toBytes(digits []int, numDigits int) string {
	ret := make([]byte, 0, numDigits)
	for i := len(digits) - 1; i >= 0; i-- {
		ret = strconv.AppendInt(ret, int64(digits[i]), 10)
	}
	ret = append(ret, bytes.Repeat([]byte{'_'}, numDigits-len(digits))...)
	return string(ret)
}

func Solve(numDigits, mult int) string {
	var min string
	for val := 0; val < 10; val++ {
		digits := []int{val}
		i := 1
		r := 0
		for ; i < numDigits; i++ {
			t := (digits[i-1] * mult) + r
			digits = append(digits, t%10)
			r = t / 10
		}
		if digits[numDigits-1] != 0 && (digits[numDigits-1]*mult)+r == digits[0] {
			solution := toBytes(digits, numDigits)
			if min == "" || strings.Compare(min, solution) > 0 {
				min = solution
			}
		}
	}
	return min
}

func SolveToMax(max, mult int) (string, int) {
	for i := 1; i <= max; i++ {
		val := Solve(i, mult)
		if val != "" {
			return val, i
		}
	}
	return "", 0
}
