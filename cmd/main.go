package main

import (
	"fmt"

	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/base"
)

func main() {
	binance := cxtgo.NewBinance(
		base.WithAPIKey("test"),
		base.WithAPISecret("test"),
		base.WithCountries("bart"),
	)

	fmt.Println(binance.Info().Countries)
}
