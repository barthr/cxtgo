package main

import (
	"fmt"

	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/exchange"
)

func main() {
	binance := cxtgo.NewBinance(
		exchange.WithAPIKey("test"),
		exchange.WithAPISecret("test"),
		exchange.WithCountries("bart"),
	)

	fmt.Println(binance.Info().Countries)
}
