package main

import (
	"context"
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

	market, err := binance.LoadMarkets(context.Background())

	switch err.(type) {
	case cxtgo.NetworkError:
		fmt.Println("network error")
	}
	info := market[exchange.NewSymbol("ETH", "BTC")]
	fmt.Printf("%#v\n", info)
}
