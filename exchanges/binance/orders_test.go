package binance

import (
	"testing"
)

func TestBinance_LimitOrder(t *testing.T) {
	// r := require.New(t)
	// binance := New(
	// 	cxtgo.WithIncludeRaw(false),
	// )

	// mock := httpmock.NewMockTransport()
	// binance.http.SetTransport(mock)

	// mock.RegisterResponder(http.MethodPost, baseURL+"/api/v3/order",
	// 	func(req *http.Request) (*http.Response, error) {
	// 		return httpmock.NewJsonResponse(http.StatusCreated, map[string]interface{}{
	// 			"symbol":        "BTCUSDT",
	// 			"orderId":       28,
	// 			"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
	// 			"transactTime":  1507725176595,
	// 			"price":         "0.42300001",
	// 			"origQty":       "10.00000000",
	// 			"executedQty":   "10.00000000",
	// 			"status":        "FILLED",
	// 			"timeInForce":   "GTC",
	// 			"type":          "MARKET",
	// 			"side":          "SELL",
	// 		})
	// 	},
	// )

	// _, err := binance.LimitOrder(
	// 	context.Background(),
	// 	cxtgo.NewSymbol("BTC", "ETH"),
	// 	cxtgo.Buy,
	// 	cxtgo.Offer{
	// 		Price:  10,
	// 		Amount: 10,
	// 	},
	// 	cxtgo.Params{
	// 		"timeInForce": "FOK",
	// 	},
	// )
	// r.NoError(err)

	// mock.RegisterResponder(http.MethodPost, baseURL+"/api/v3/order",
	// 	func(req *http.Request) (*http.Response, error) {
	// 		return nil, errors.New("failing request")
	// 	},
	// )

	// _, err = binance.LimitOrder(
	// 	context.Background(),
	// 	cxtgo.NewSymbol("BTC", "ETH"),
	// 	cxtgo.Buy,
	// 	cxtgo.Offer{
	// 		Price:  10,
	// 		Amount: 10,
	// 	},
	// 	cxtgo.Params{
	// 		"timeInForce": "FOK",
	// 	},
	// )
	// r.Error(err)

	// mock.RegisterResponder(http.MethodPost, baseURL+"/api/v3/order",
	// 	func(req *http.Request) (*http.Response, error) {
	// 		return httpmock.NewJsonResponse(250, &apiError{
	// 			Code:    disconnected,
	// 			Message: "exchange wasn't able to respond correctly",
	// 		})
	// 	},
	// )

	// _, err = binance.LimitOrder(
	// 	context.Background(),
	// 	cxtgo.NewSymbol("BTC", "ETH"),
	// 	cxtgo.Buy,
	// 	cxtgo.Offer{
	// 		Price:  10,
	// 		Amount: 10,
	// 	},
	// 	cxtgo.Params{
	// 		"timeInForce": "FOK",
	// 	},
	// )
	// r.Error(err)
}
