package precision

import "strings"

// FromString calculates the amount of zero's after the splitter.
// For example 0,0001 returns the value 4.
func FromString(input string, splitter string) int {
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
