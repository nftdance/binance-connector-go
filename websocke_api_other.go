package binance_connector

import (
	"context"
	"encoding/json"
	"strconv"
)

func (w *WebsocketAPIClient) NewPositionInfoService() *PositionInfoService {
	return &PositionInfoService{websocketAPI: w}
}

type PositionInfoService struct {
	websocketAPI *WebsocketAPIClient
	recvWindow   *int64
}

func (s *PositionInfoService) RecvWindow(recvWindow int64) *PositionInfoService {
	s.recvWindow = &recvWindow
	return s
}

func (s *PositionInfoService) Do(ctx context.Context) (*PositionResponse, error) {
	parameters := map[string]string{}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "account.position",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var accInfoResponse PositionResponse
		err = json.Unmarshal(response, &accInfoResponse)
		if err != nil {
			return nil, err
		}
		return &accInfoResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type PositionResponse struct {
	ID         string                `json:"id"`
	Status     int                   `json:"status"`
	Error      *WsAPIErrorResponse   `json:"error,omitempty"`
	Result     []PositionInformation `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit     `json:"rateLimits"`
}

type PositionInformation struct {
	EntryPrice       string `json:"entryPrice"`
	BreakEvenPrice   string `json:"breakEvenPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	UpdateTime       int    `json:"updateTime"`
}
