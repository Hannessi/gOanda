package gOanda

import "github.com/sirupsen/logrus"

type Requester interface {
	// account
	GetAccounts(GetAccountsRequest) (*GetAccountsResponse, error)
	GetAccount(GetAccountRequest) (*GetAccountResponse, error)
	GetAccountSummary(GetAccountSummaryRequest) (*GetAccountSummaryResponse, error)
	GetAccountInstruments(GetAccountInstrumentsRequest) (*GetAccountInstrumentsResponse, error)
	PatchAccountConfiguration(PatchAccountConfigurationRequest) (*PatchAccountConfigurationResponse, error)

	// instruments
	GetInstrumentCandles(GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error)
	GetInstrumentOrderBook(GetInstrumentOrderBookRequest) (*GetInstrumentOrderBookResponse, error)
	GetInstrumentPositionBook(GetInstrumentPositionBookRequest) (*GetInstrumentPositionBookResponse, error)

	// orders
	PostOrder(PostOrderRequest) (*PostOrderResponse, error)
	GetOrders(GetOrdersRequest) (*GetOrdersResponse, error)
	GetPendingOrders(GetPendingOrdersRequest) (*GetPendingOrdersResponse, error)
	GetOrder(GetOrderRequest) (*GetOrderResponse, error)
	PutReplaceOrder(PutReplaceOrderRequest) (*PutReplaceOrderResponse, error)
	PutCancelOrder(PutCancelOrderRequest) (*PutCancelOrderResponse, error)
	PutUpdateOrderClientExtensions(PutUpdateOrderClientExtensionsRequest) (*PutUpdateOrderClientExtensionsResponse, error)

	// trades
	GetTrades(GetTradesRequest) (*GetTradesResponse, error)
	GetOpenTrades(GetOpenTradesRequest) (*GetOpenTradesResponse, error)
	GetTrade(GetTradeRequest) (*GetTradeResponse, error)
	PutCloseTrade(PutCloseTradeRequest) (*PutCloseTradeResponse, error)
	PutUpdateTradeClientExtensions(PutUpdateTradeClientExtensionsRequest) (*PutUpdateTradeClientExtensionsResponse, error)
	PutUpdateTradeDependentOrders(PutUpdateTradeDependentOrdersRequest) (*PutUpdateTradeDependentOrdersResponse, error)

	// positions
	GetPositions(GetPositionsRequest) (*GetPositionsResponse, error)
	GetOpenPositions(GetOpenPositionsRequest) (*GetOpenPositionsResponse, error)
	GetInstrumentsPosition(GetInstrumentsPositionRequest) (*GetInstrumentsPositionResponse, error)
	PutClosePosition(PutClosePositionRequest) (*PutClosePositionResponse, error)

	// transactions
	GetTransactions(GetTransactionsRequest) (*GetTransactionsResponse, error)
	GetTransaction(GetTransactionRequest) (*GetTransactionResponse, error)
	GetRangeOfTransactions(GetRangeOfTransactionsRequest) (*GetRangeOfTransactionsResponse, error)
	GetTransactionsAfterTransaction(GetTransactionsAfterTransactionRequest) (*GetTransactionsAfterTransactionResponse, error)
	GetTransactionsStream(GetTransactionsStreamRequest) (*GetTransactionsStreamResponse, error)

	// pricing
	GetInstrumentPricing(GetInstrumentPricingRequest) (*GetInstrumentPricingResponse, error)
	GetPricingStream(GetPricingStreamRequest) (*GetPricingStreamResponse, error)
}

type GetAccountsRequest struct{}

type GetAccountsResponse struct {
	Accounts []AccountProperties
}

type GetAccountRequest struct{}

type GetAccountResponse struct {
	Account Account
}

type GetAccountSummaryRequest struct{}

type GetAccountSummaryResponse struct {
	Account           AccountSummary
	LastTransactionID TransactionID
}

type GetAccountInstrumentsRequest struct{}

type GetAccountInstrumentsResponse struct {
	Instruments       []Instrument
	LastTransactionID TransactionID
}

type PatchAccountConfigurationRequest struct{}

type PatchAccountConfigurationResponse struct{}

type GetInstrumentCandlesRequest struct {
	InstrumentName InstrumentName
	Count          int
	Granularity    CandlestickGranularity
	From           DateTime
	To             DateTime
}

type GetInstrumentCandlesResponse struct {
	Instrument  InstrumentName
	Granularity CandlestickGranularity
	Candles     []Candlestick
}

type GetInstrumentOrderBookRequest struct{}

type GetInstrumentOrderBookResponse struct{}

type GetInstrumentPositionBookRequest struct{}

type GetInstrumentPositionBookResponse struct{}

type PostOrderRequest struct {
	Order AnOrderRequest
}

type PostOrderResponse struct {
	OrderCreateTransaction        Transaction            `json:"orderCreateTransaction"`
	OrderFillTransaction          OrderFillTransaction   `json:"orderFillTransaction"`
	OrderCancelTransaction        OrderCancelTransaction `json:"orderCancelTransaction"`
	OrderReissueTransaction       Transaction            `json:"orderReissueTransaction"`
	OrderReissueRejectTransaction Transaction            `json:"orderReissueRejectTransaction"`
	RelatedTransactionIDs         []TransactionID        `json:"relatedTransactionIDs"`
	LastTransactionID             TransactionID          `json:"lastTransactionID"`
	ErrorCode                     string                 `json:"errorCode"`
	ErrorMessage                  string                 `json:"errorMessage"`
}

type GetOrdersRequest struct {
	Ids            []OrderID
	State          OrderStateFilter
	InstrumentName InstrumentName
	Count          int
	BeforeID       OrderID
}

type GetRawOrdersResponse struct {
	Orders            []RawOrder    `json:"orders"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetOrdersResponse struct {
	Orders            []Order       `json:"orders"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

func (g *GetRawOrdersResponse) Unmarshal() (*GetOrdersResponse, error) {
	unmarshalledOrders := make([]Order, 0)

	for _, marshalledOrder := range g.Orders {
		unmarshalledOrder, err := marshalledOrder.ToOrder()
		if err != nil {
			logrus.Error("Could not unmarshal order: ", marshalledOrder)
			return nil, err
		}
		unmarshalledOrders = append(unmarshalledOrders, unmarshalledOrder)
	}

	return &GetOrdersResponse{
		Orders:            unmarshalledOrders,
		LastTransactionID: g.LastTransactionID,
		ErrorCode:         g.ErrorCode,
		ErrorMessage:      g.ErrorMessage,
	}, nil
}

type GetPendingOrdersRequest struct{}

type GetRawPendingOrdersResponse struct {
	Orders            []RawOrder    `json:"orders"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

func (g *GetRawPendingOrdersResponse) Unmarshal() (*GetPendingOrdersResponse, error) {
	unmarshalledOrders := make([]Order, 0)

	for _, marshalledOrder := range g.Orders {
		unmarshalledOrder, err := marshalledOrder.ToOrder()
		if err != nil {
			logrus.Error("Could not unmarshal order: ", marshalledOrder)
			return nil, err
		}
		unmarshalledOrders = append(unmarshalledOrders, unmarshalledOrder)
	}

	return &GetPendingOrdersResponse{
		Orders:            unmarshalledOrders,
		LastTransactionID: g.LastTransactionID,
		ErrorCode:         g.ErrorCode,
		ErrorMessage:      g.ErrorMessage,
	}, nil
}

type GetPendingOrdersResponse struct {
	Orders            []Order       `json:"orders"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetOrderRequest struct{}

type GetOrderResponse struct{}

type PutReplaceOrderRequest struct{}

type PutReplaceOrderResponse struct{}

type PutCancelOrderRequest struct{}

type PutCancelOrderResponse struct{}

type PutUpdateOrderClientExtensionsRequest struct{}

type PutUpdateOrderClientExtensionsResponse struct{}

type GetTradesRequest struct {
	Ids            []TradeID        `json:"ids"`
	State          TradeStateFilter `json:"state"`
	InstrumentName InstrumentName   `json:"instrument"`
	Count          int64            `json:"count"`
	BeforeID       TradeID          `json:"beforeID"`
}

type GetTradesResponse struct {
	Trades            []Trade       `json:"trades"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetOpenTradesRequest struct{}

type GetOpenTradesResponse struct {
	Trades            []Trade       `json:"trades"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetTradeRequest struct {
	TradeSpecifier TradeSpecifier
}

type GetTradeResponse struct {
	Trade             Trade         `json:"trade"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type PutCloseTradeRequest struct {
	TradeSpecifier TradeSpecifier
}

type PutCloseTradeResponse struct {
	OrderCreateTransaction MarketOrderTransaction       `json:"orderCreateTransaction"`
	OrderFillTransaction   OrderFillTransaction         `json:"orderFillTransaction"`
	OrderCancelTransaction OrderCancelTransaction       `json:"orderCancelTransaction"`
	OrderRejectTransaction MarketOrderRejectTransaction `json:"orderRejectTransaction"`
	RelatedTransactionsIDs []TransactionID              `json:"relatedTransactionsIDs"`
	LastTransactionID      TransactionID                `json:"lastTransactionID"`
	ErrorCode              string                       `json:"errorCode"`
	ErrorMessage           string                       `json:"errorMessage"`
}

type PutUpdateTradeClientExtensionsRequest struct{}

type PutUpdateTradeClientExtensionsResponse struct{}

type PutUpdateTradeDependentOrdersRequest struct{}

type PutUpdateTradeDependentOrdersResponse struct{}

type GetPositionsRequest struct{}

type GetPositionsResponse struct{}

type GetOpenPositionsRequest struct{}

type GetOpenPositionsResponse struct{}

type GetInstrumentsPositionRequest struct{}

type GetInstrumentsPositionResponse struct{}

type PutClosePositionRequest struct{}

type PutClosePositionResponse struct{}

type GetTransactionsRequest struct{}

type GetTransactionsResponse struct{}

type GetTransactionRequest struct{}

type GetTransactionResponse struct{}

type GetRangeOfTransactionsRequest struct {
	To   TransactionID
	From TransactionID
}

type GetRangeOfTransactionsResponse struct {
	Transactions      []Transaction `json:"transactions"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetTransactionsAfterTransactionRequest struct {
	Id TransactionID
}

type GetTransactionsAfterTransactionResponse struct {
	Transactions      []Transaction `json:"transactions"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

type GetTransactionsStreamRequest struct{}

type GetTransactionsStreamResponse struct{}

type GetInstrumentPricingRequest struct {
	Instruments            []InstrumentName
	Since                  DateTime
	IncludeUnitsAvailable  bool
	IncludeHomeConversions bool
}

type GetInstrumentPricingResponse struct {
	Prices []Price
}

type GetPricingStreamRequest struct{}

type GetPricingStreamResponse struct{}
