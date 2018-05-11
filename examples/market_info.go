package main

import (
	"context"
	"fmt"
	"os"

	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/exchanges/binance"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	binance := binance.New(
		cxtgo.WithAPIKey("test"),
		cxtgo.WithAPISecret("test"),
		cxtgo.WithDebug(true),
		cxtgo.WithDebuglogger(os.Stdout),
	)

	market, err := binance.Markets(context.Background())

	switch err.(type) {
	case cxtgo.NetworkError:
		fmt.Println("network error")
	}
	info := market[cxtgo.NewSymbol("ETH", "BTC")]
	fmt.Printf("%v\n", info.Precision.Amount)

	test(binance)
}

func test(ex cxtgo.StreamingAPI) {
	params := cxtgo.Params{
		"limit": 50,
	}
	opts := []cxtgo.StreamOpt{
		cxtgo.WithStreamSymbol(cxtgo.NewSymbol("ETH", "BTC")),
		cxtgo.WithStreamParams(params),
		cxtgo.WithStreamContext(context.Background()),
	}
	err := ex.StreamTicker(
		func(ticker cxtgo.Ticker) {
			spew.Dump(ticker)
		},
		func(err error) {
		},
		opts...,
	)

	if err != nil {
		panic(err)
	}
}
