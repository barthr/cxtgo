package main

import (
	"context"
	"fmt"

	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/exchanges/binance"
)

func main() {
	binance := binance.New(
		cxtgo.WithAPIKey("test"),
		cxtgo.WithAPISecret("test"),
	)

	market, err := binance.Markets(context.Background())

	switch err.(type) {
	case cxtgo.NetworkError:
		fmt.Println("network error")
	}
	info := market[cxtgo.NewSymbol("ETH", "BTC")]
	fmt.Printf("%#v\n", info)
}
