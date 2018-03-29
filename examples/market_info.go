package main

import (
	"context"
	"fmt"

	"github.com/barthr/cxtgo"
)

func main() {
	binance := cxtgo.NewBinance(
		cxtgo.WithAPIKey("test"),
		cxtgo.WithAPISecret("test"),
	)

	market, err := binance.LoadMarkets(context.Background())

	switch err.(type) {
	case cxtgo.NetworkError:
		fmt.Println("network error")
	}
	info := market[cxtgo.NewSymbol("ETH", "BTC")]
	fmt.Printf("%#v\n", info)
}
