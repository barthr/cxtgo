package binance

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/barthr/cxtgo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestBinance_LimitOrder(t *testing.T) {
	r := require.New(t)
	binance := New(
		cxtgo.WithDebuglogger(os.Stdout),
		cxtgo.WithDebug(true),
	)

	httpmock.ActivateNonDefault(binance.http.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, baseURL+"/api/v3/order",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(http.StatusCreated, map[string]interface{}{
				"symbol":        "BTCUSDT",
				"orderId":       28,
				"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
				"transactTime":  1507725176595,
				"price":         "0.00000000",
				"origQty":       "10.00000000",
				"executedQty":   "10.00000000",
				"status":        "FILLED",
				"timeInForce":   "GTC",
				"type":          "MARKET",
				"side":          "SELL",
			})
		},
	)
	r.NoError(binance.LimitOrder(context.Background(), cxtgo.NewSymbol("BTC", "ETH"), cxtgo.Buy, 10, 10))

	httpmock.RegisterResponder(http.MethodPost, baseURL+"/api/v3/order",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("failing request")
		},
	)
	r.Error(binance.LimitOrder(context.Background(), cxtgo.NewSymbol("BTC", "ETH"), cxtgo.Buy, 10, 10))

}
