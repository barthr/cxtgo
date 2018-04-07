package precision

import "strings"

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
