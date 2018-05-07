package binance

import (
	"context"
	"strconv"
	"time"

	"github.com/barthr/cxtgo"
)

func (b *Binance) LimitOrder(ctx context.Context, symbol cxtgo.Symbol, side cxtgo.Side, offer cxtgo.Offer, params ...cxtgo.Params) (cxtgo.Order, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Order{}, err
	}
	req := b.http.R().SetResult(&createOrderResponse{})
	timeInForce := "GTC"
	if len(params) > 0 {
		val, ok := params[0].GetString("timeInForce")
		if ok {
			timeInForce = val
		}
		recvWindow, ok := params[0].GetInt("recvWindow")
		if ok {
			req.SetQueryParam("recvWindow", strconv.Itoa(recvWindow))
		}
	}

	resp, err := req.SetContext(ctx).SetQueryParams(map[string]string{
		"symbol":           symbol.String(),
		"side":             side.String(),
		"type":             cxtgo.LimitOrder.String(),
		"timeInForce":      timeInForce,
		"price":            strconv.FormatFloat(offer.Price, 'f', -1, 64),
		"quantity":         strconv.FormatFloat(offer.Amount, 'f', -1, 64),
		"timestamp":        strconv.FormatInt(time.Now().UnixNano(), 10),
		"newOrderRespType": "FULL",
	}).Post("/api/v3/order")

	if err != nil {
		binanceErr, ok := resp.Error().(*apiError)
		if ok {
			switch binanceErr.Code {
			case disconnected:
				return cxtgo.Order{}, cxtgo.ExchangeNotAvailableError{
					BaseError: cxtgo.NewError("binance", err),
				}
			}
		}
		return cxtgo.Order{}, cxtgo.ExchangeNotAvailableError{
			BaseError: cxtgo.NewError("binance", err),
		}
	}
	order := resp.Result().(*createOrderResponse)

	var raw []byte
	if b.base.Raw {
		raw = resp.Body()
	}

	return cxtgo.Order{
		Symbol:    symbol,
		ID:        strconv.Itoa(order.OrderID),
		Timestamp: time.Unix(order.TransactTime, 0),
		Status:    orderStatus[order.Status],
		Type:      orderType[order.Type],
		Price:     order.Price,
		Amount:    order.OrigQty,
		Filled:    order.ExecutedQty,
		Cost:      order.ExecutedQty * order.Price,
		Remaining: order.OrigQty - order.ExecutedQty,
		Raw:       raw,
	}, nil
}

func (b *Binance) MarketOrder(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) CancelAllOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) Order(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) Orders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) OpenOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) ClosedOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

// Helper function from converting to cxtgo specific types
var orderStatus = map[string]cxtgo.OrderStatus{
	"NEW":              cxtgo.OrderOpen,
	"PARTIALLY_FILLED": cxtgo.OrderOpen,
	"FILLED":           cxtgo.OrderClosed,
	"CANCELED":         cxtgo.OrderCanceled,
	"REJECTED":         cxtgo.OrderCanceled,
	"EXPIRED":          cxtgo.OrderCanceled,
}

var orderType = map[string]cxtgo.OrderType{
	"LIMIT":             cxtgo.LimitOrder,
	"MARKET":            cxtgo.MarketOrder,
	"STOP_LOSS":         cxtgo.LimitOrder,
	"STOP_LOSS_LIMIT":   cxtgo.LimitOrder,
	"TAKE_PROFIT":       cxtgo.LimitOrder,
	"TAKE_PROFIT_LIMIT": cxtgo.LimitOrder,
	"LIMIT_MAKER":       cxtgo.LimitOrder,
}

// Types for order calls to binance

type createOrderResponse struct {
	Symbol        string  `json:"symbol"`
	OrderID       int     `json:"orderId"`
	ClientOrderID string  `json:"clientOrderId"`
	TransactTime  int64   `json:"transactTime"`
	Price         float64 `json:"price,string"`
	OrigQty       float64 `json:"origQty,string"`
	ExecutedQty   float64 `json:"executedQty,string"`
	Status        string  `json:"status"`
	TimeInForce   string  `json:"timeInForce"`
	Type          string  `json:"type"`
	Side          string  `json:"side"`
}

type orderStatusResponse struct {
}

type cancelOrderResponse struct {
}

type closeOrderResponse struct {
}
