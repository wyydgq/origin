package main

import (
	"strconv"
	"math"
)

func fractionToDecimal(numerator int, denominator int) string {
	_helper := func(numerator int, denominator int) string {
		indexes := make(map[int]int)
		indexes[numerator] = 0
		ret, w := "", 1
		for numerator != 0 {
			numerator *= 10
			r := numerator / denominator
			numerator = numerator % denominator
			ret += strconv.Itoa(r)

			if i, ok := indexes[numerator]; ok {
				// Repeated
				repeated := ret[i:len(ret)]
				ret = ret[0:i]
				ret = ret + "(" + repeated + ")"
				break
			}
			indexes[numerator] = w
			w++
		}
		return ret
	}
	isNegative := false
	if numerator*denominator < 0 {
		isNegative = true
	}

	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))

	first := "0"
	if numerator >= denominator {
		first = strconv.Itoa(numerator / denominator)
		numerator %= denominator
	}

	second := _helper(numerator, denominator)
	if second != "" {
		first = first + "." + second
	}

	if isNegative && first != "0" {
		first = "-" + first
	}
	return first
}
