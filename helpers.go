package cxtgo

import (
	"strings"
)

// Zeros calculates the amount of zero's after the splitter
// including the first item which is not '0'.
// For example 0,0001 returns the value 4.
func Zeros(input string, splitter string) int {
	parts := strings.Split(input, splitter)
	pricePrecision := 0
	if len(parts) != 2 {
		return 0
	}
	for _, item := range parts[1] {
		pricePrecision++
		if item != '0' {
			break
		}
	}
	return pricePrecision
}
