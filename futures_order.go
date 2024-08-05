package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Order define order info
type FuturesOrder struct {
	Symbol                  string           `json:"symbol"`
	OrderID                 int64            `json:"orderId"`
	ClientOrderID           string           `json:"clientOrderId"`
	Price                   string           `json:"price"`
	ReduceOnly              bool             `json:"reduceOnly"`
	OrigQuantity            string           `json:"origQty"`
	ExecutedQuantity        string           `json:"executedQty"`
	CumQuantity             string           `json:"cumQty"`
	CumQuote                string           `json:"cumQuote"`
	Status                  OrderStatusType  `json:"status"`
	TimeInForce             TimeInForceType  `json:"timeInForce"`
	Type                    OrderType        `json:"type"`
	Side                    SideType         `json:"side"`
	StopPrice               string           `json:"stopPrice"`
	Time                    int64            `json:"time"`
	UpdateTime              int64            `json:"updateTime"`
	WorkingType             WorkingType      `json:"workingType"`
	ActivatePrice           string           `json:"activatePrice"`
	PriceRate               string           `json:"priceRate"`
	AvgPrice                string           `json:"avgPrice"`
	OrigType                OrderType        `json:"origType"`
	PositionSide            PositionSideType `json:"positionSide"`
	PriceProtect            bool             `json:"priceProtect"`
	ClosePosition           bool             `json:"closePosition"`
	PriceMatch              string           `json:"priceMatch"`
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`
	GoodTillDate            int64            `json:"goodTillDate"`
}

func (c *Client) NewFuturesListOrdersService() *FuturesListOrdersService {
	return &FuturesListOrdersService{c: c}
}

// ListOrdersService all account orders; active, canceled, or filled
type FuturesListOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *FuturesListOrdersService) Symbol(symbol string) *FuturesListOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *FuturesListOrdersService) OrderID(orderID int64) *FuturesListOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *FuturesListOrdersService) StartTime(startTime int64) *FuturesListOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *FuturesListOrdersService) EndTime(endTime int64) *FuturesListOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *FuturesListOrdersService) Limit(limit int) *FuturesListOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *FuturesListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*FuturesOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*FuturesOrder{}, err
	}
	res = make([]*FuturesOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FuturesOrder{}, err
	}
	return res, nil
}
