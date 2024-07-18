package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// SideType define side type of order
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderExecutionType define order execution type
type OrderExecutionType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// SideEffectType define side effect type for orders
type SideEffectType string

// WorkingType define working type
type WorkingType string

// MarginType define margin type
type MarginType string

// ContractType define contract type
type ContractType string

// UserDataEventReasonType define reason type for user data event
type UserDataEventReasonType string

// ForceOrderCloseType define reason type for force order
type ForceOrderCloseType string

// Endpoints
const (
	baseApiMainUrl    = "https://fapi.binance.com"
	baseApiTestnetUrl = "https://testnet.binancefuture.com"
)

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"

	OrderExecutionTypeNew         OrderExecutionType = "NEW"
	OrderExecutionTypePartialFill OrderExecutionType = "PARTIAL_FILL"
	OrderExecutionTypeFill        OrderExecutionType = "FILL"
	OrderExecutionTypeCanceled    OrderExecutionType = "CANCELED"
	OrderExecutionTypeCalculated  OrderExecutionType = "CALCULATED"
	OrderExecutionTypeExpired     OrderExecutionType = "EXPIRED"
	OrderExecutionTypeTrade       OrderExecutionType = "TRADE"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusTypeNewInsurance    OrderStatusType = "NEW_INSURANCE"
	OrderStatusTypeNewADL          OrderStatusType = "NEW_ADL"

	SymbolTypeFuture SymbolType = "FUTURE"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize          SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePrice            SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice     SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMarketLotSize    SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders     SymbolFilterType = "MAX_NUM_ORDERS"
	SymbolFilterTypeMaxNumAlgoOrders SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	SymbolFilterTypeMinNotional      SymbolFilterType = "MIN_NOTIONAL"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCrossed  MarginType = "CROSSED"

	ContractTypePerpetual ContractType = "PERPETUAL"

	UserDataEventTypeListenKeyExpired    UserDataEventType = "listenKeyExpired"
	UserDataEventTypeMarginCall          UserDataEventType = "MARGIN_CALL"
	UserDataEventTypeAccountUpdate       UserDataEventType = "ACCOUNT_UPDATE"
	UserDataEventTypeOrderTradeUpdate    UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeAccountConfigUpdate UserDataEventType = "ACCOUNT_CONFIG_UPDATE"

	UserDataEventReasonTypeDeposit             UserDataEventReasonType = "DEPOSIT"
	UserDataEventReasonTypeWithdraw            UserDataEventReasonType = "WITHDRAW"
	UserDataEventReasonTypeOrder               UserDataEventReasonType = "ORDER"
	UserDataEventReasonTypeFundingFee          UserDataEventReasonType = "FUNDING_FEE"
	UserDataEventReasonTypeWithdrawReject      UserDataEventReasonType = "WITHDRAW_REJECT"
	UserDataEventReasonTypeAdjustment          UserDataEventReasonType = "ADJUSTMENT"
	UserDataEventReasonTypeInsuranceClear      UserDataEventReasonType = "INSURANCE_CLEAR"
	UserDataEventReasonTypeAdminDeposit        UserDataEventReasonType = "ADMIN_DEPOSIT"
	UserDataEventReasonTypeAdminWithdraw       UserDataEventReasonType = "ADMIN_WITHDRAW"
	UserDataEventReasonTypeMarginTransfer      UserDataEventReasonType = "MARGIN_TRANSFER"
	UserDataEventReasonTypeMarginTypeChange    UserDataEventReasonType = "MARGIN_TYPE_CHANGE"
	UserDataEventReasonTypeAssetTransfer       UserDataEventReasonType = "ASSET_TRANSFER"
	UserDataEventReasonTypeOptionsPremiumFee   UserDataEventReasonType = "OPTIONS_PREMIUM_FEE"
	UserDataEventReasonTypeOptionsSettleProfit UserDataEventReasonType = "OPTIONS_SETTLE_PROFIT"

	ForceOrderCloseTypeLiquidation ForceOrderCloseType = "LIQUIDATION"
	ForceOrderCloseTypeADL         ForceOrderCloseType = "ADL"
)

// NewCreateOrderService init creating order service
func (c *Client) NewFuturesCreateOrderService() *FuturesCreateOrderService {
	return &FuturesCreateOrderService{c: c}
}

func (c *Client) NewFuturesGetOrderService() *FuturesGetOrderService {
	return &FuturesGetOrderService{c: c}
}

func (c *Client) NewFuturesGetBalanceService() *FuturesGetBalanceService {
	return &FuturesGetBalanceService{c: c}
}

func (c *Client) NewFuturesGetPositionRiskService() *FuturesGetPositionRiskService {
	return &FuturesGetPositionRiskService{c: c}
}

func (c *Client) NewFuturesGetIncomeService() *FuturesGetIncomeService {
	return &FuturesGetIncomeService{c: c}
}

type FuturesCreateOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *string
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *string
	newOrderRespType NewOrderRespType
	closePosition    *string
}

// Symbol set symbol
func (s *FuturesCreateOrderService) Symbol(symbol string) *FuturesCreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *FuturesCreateOrderService) Side(side SideType) *FuturesCreateOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *FuturesCreateOrderService) PositionSide(positionSide PositionSideType) *FuturesCreateOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *FuturesCreateOrderService) Type(orderType OrderType) *FuturesCreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *FuturesCreateOrderService) TimeInForce(timeInForce TimeInForceType) *FuturesCreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *FuturesCreateOrderService) Quantity(quantity string) *FuturesCreateOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *FuturesCreateOrderService) ReduceOnly(reduceOnly bool) *FuturesCreateOrderService {
	reduceOnlyStr := strconv.FormatBool(reduceOnly)
	s.reduceOnly = &reduceOnlyStr
	return s
}

// Price set price
func (s *FuturesCreateOrderService) Price(price string) *FuturesCreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *FuturesCreateOrderService) NewClientOrderID(newClientOrderID string) *FuturesCreateOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *FuturesCreateOrderService) StopPrice(stopPrice string) *FuturesCreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *FuturesCreateOrderService) WorkingType(workingType WorkingType) *FuturesCreateOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *FuturesCreateOrderService) ActivationPrice(activationPrice string) *FuturesCreateOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *FuturesCreateOrderService) CallbackRate(callbackRate string) *FuturesCreateOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *FuturesCreateOrderService) PriceProtect(priceProtect bool) *FuturesCreateOrderService {
	priceProtectStr := strconv.FormatBool(priceProtect)
	s.priceProtect = &priceProtectStr
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *FuturesCreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *FuturesCreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *FuturesCreateOrderService) ClosePosition(closePosition bool) *FuturesCreateOrderService {
	closePositionStr := strconv.FormatBool(closePosition)
	s.closePosition = &closePositionStr
	return s
}

func (s *FuturesCreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	for key, val := range m {
		r.setParam(key, val)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *FuturesCreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/fapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	Symbol                  string           `json:"symbol"`                      //
	OrderID                 int64            `json:"orderId"`                     //
	ClientOrderID           string           `json:"clientOrderId"`               //
	Price                   string           `json:"price"`                       //
	OrigQuantity            string           `json:"origQty"`                     //
	ExecutedQuantity        string           `json:"executedQty"`                 //
	CumQuote                string           `json:"cumQuote"`                    //
	ReduceOnly              bool             `json:"reduceOnly"`                  //
	Status                  OrderStatusType  `json:"status"`                      //
	StopPrice               string           `json:"stopPrice"`                   // please ignore when order type is TRAILING_STOP_MARKET
	TimeInForce             TimeInForceType  `json:"timeInForce"`                 //
	Type                    OrderType        `json:"type"`                        //
	Side                    SideType         `json:"side"`                        //
	UpdateTime              int64            `json:"updateTime"`                  // update time
	WorkingType             WorkingType      `json:"workingType"`                 //
	ActivatePrice           string           `json:"activatePrice"`               // activation price, only return with TRAILING_STOP_MARKET order
	PriceRate               string           `json:"priceRate"`                   // callback rate, only return with TRAILING_STOP_MARKET order
	AvgPrice                string           `json:"avgPrice"`                    //
	PositionSide            PositionSideType `json:"positionSide"`                //
	ClosePosition           bool             `json:"closePosition"`               // if Close-All
	PriceProtect            bool             `json:"priceProtect"`                // if conditional order trigger is protected
	PriceMatch              string           `json:"priceMatch"`                  // price match mode
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`     // self trading prevention mode
	GoodTillDate            int64            `json:"goodTillDate"`                // order pre-set auto cancel time for TIF GTD order
	CumQty                  string           `json:"cumQty"`                      //
	OrigType                OrderType        `json:"origType"`                    //
	RateLimitOrder10s       string           `json:"rateLimitOrder10s,omitempty"` //
	RateLimitOrder1m        string           `json:"rateLimitOrder1m,omitempty"`  //
}

// GetOrderService get an order
type FuturesGetOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *FuturesGetOrderService) Symbol(symbol string) *FuturesGetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *FuturesGetOrderService) OrderID(orderID int64) *FuturesGetOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *FuturesGetOrderService) OrigClientOrderID(origClientOrderID string) *FuturesGetOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *FuturesGetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type Order struct {
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

// GetBalanceService get account balance
type FuturesGetBalanceService struct {
	c *Client
}

// Balance define user balance of your account
type FuturesBalance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

// Do send request
func (s *FuturesGetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*FuturesBalance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*FuturesBalance{}, err
	}
	res = make([]*FuturesBalance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FuturesBalance{}, err
	}
	return res, nil
}

// GetPositionRiskService get account balance
type FuturesGetPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *FuturesGetPositionRiskService) Symbol(symbol string) *FuturesGetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *FuturesGetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
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
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
}

type FuturesGetIncomeService struct {
	c          *Client
	symbol     string
	incomeType string
	startTime  int64
	endTime    int64
	page       int64
	limit      int64
	recvWindow int64
}

// Symbol set symbol
func (s *FuturesGetIncomeService) Symbol(symbol string) *FuturesGetIncomeService {
	s.symbol = symbol
	return s
}

func (s *FuturesGetIncomeService) IncomeType(incomeType string) *FuturesGetIncomeService {
	s.incomeType = incomeType
	return s
}

func (s *FuturesGetIncomeService) StartTime(startTime int64) *FuturesGetIncomeService {
	s.startTime = startTime
	return s
}

func (s *FuturesGetIncomeService) EndTime(endTime int64) *FuturesGetIncomeService {
	s.endTime = endTime
	return s
}

func (s *FuturesGetIncomeService) Page(page int64) *FuturesGetIncomeService {
	s.page = page
	return s
}

func (s *FuturesGetIncomeService) Limit(limit int64) *FuturesGetIncomeService {
	s.limit = limit
	return s
}

type Income struct {
	Symbol     string `json:"symbol"`
	IncomeType string `json:"incomeType"`
	Income     string `json:"income"`
	Asset      string `json:"asset"`
	Info       string `json:"info"`
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`
	TradeID    int64  `json:"tradeId"`
}

// Do send request
func (s *FuturesGetIncomeService) Do(ctx context.Context, opts ...RequestOption) (res []*Income, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/income",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Income{}, err
	}
	res = make([]*Income, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Income{}, err
	}
	return res, nil
}
