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
	fmt.Printf("%v\n", info.Precision.Amount)
}

func test(ex cxtgo.StreamingAPI) {
	params := cxtgo.Params{
		"limit": 50,
	}
	opts := []cxtgo.StreamOpt{
		cxtgo.WithStreamSymbol(cxtgo.NewSymbol("ETH", "BTC")),
		cxtgo.WithStreamParams(params),
	}
	ex.StreamOrderbook(
		func(s cxtgo.Summary) {
			// do something with summary
		},
		func(err error) {
			// do something with error
		},
		opts...,
	)
}
