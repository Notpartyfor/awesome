package common

import "math"

func UnHex(c string) int32 {
	var sum int32
	for i, v := range c {
		sum += unHexSingle(v) * int32(math.Pow(float64(16), float64(len(c)-i-1)))
	}
	return sum
}

func unHexSingle(c int32) int32 {

	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
