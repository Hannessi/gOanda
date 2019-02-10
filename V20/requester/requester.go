package requester

import "github.com/hannessi/gOanda"

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
	Accounts     []gOanda.AccountProperties `json:"accounts"`
	ErrorCode    string                     `json:"errorCode"`
	ErrorMessage string                     `json:"errorMessage"`
}

type GetAccountRequest struct{}

type GetAccountResponse struct {
	Account      gOanda.Account `json:"account"`
	ErrorCode    string         `json:"errorCode"`
	ErrorMessage string         `json:"errorMessage"`
}

type GetAccountSummaryRequest struct{}

type GetAccountSummaryResponse struct {
	Account           gOanda.AccountSummary `json:"account"`
	LastTransactionID gOanda.TransactionID  `json:"lastTransactionID"`
	ErrorCode         string                `json:"errorCode"`
	ErrorMessage      string                `json:"errorMessage"`
}

type GetAccountInstrumentsRequest struct{}

type GetAccountInstrumentsResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type PatchAccountConfigurationRequest struct{}

type PatchAccountConfigurationResponse struct{}

type GetInstrumentCandlesRequest struct {
	InstrumentName gOanda.InstrumentName
	Count          int
	Granularity    gOanda.CandlestickGranularity
}

type GetInstrumentCandlesResponse struct {
	Instrument   gOanda.InstrumentName         `json:"instrument"`
	Granularity  gOanda.CandlestickGranularity `json:"granularity"`
	Candles      []gOanda.Candlestick          `json:"candles"`
	ErrorCode    string                        `json:"errorCode"`
	ErrorMessage string                        `json:"errorMessage"`
}

type GetInstrumentOrderBookRequest struct{}

type GetInstrumentOrderBookResponse struct{}

type GetInstrumentPositionBookRequest struct{}

type GetInstrumentPositionBookResponse struct{}

type PostOrderRequest struct{}

type PostOrderResponse struct{}

type GetOrdersRequest struct{}

type GetOrdersResponse struct{}

type GetPendingOrdersRequest struct{}

type GetPendingOrdersResponse struct{}

type GetOrderRequest struct{}

type GetOrderResponse struct{}

type PutReplaceOrderRequest struct{}

type PutReplaceOrderResponse struct{}

type PutCancelOrderRequest struct{}

type PutCancelOrderResponse struct{}

type PutUpdateOrderClientExtensionsRequest struct{}

type PutUpdateOrderClientExtensionsResponse struct{}

type GetTradesRequest struct{}

type GetTradesResponse struct{}

type GetOpenTradesRequest struct{}

type GetOpenTradesResponse struct{}

type GetTradeRequest struct {
	TradeSpecifier gOanda.TradeSpecifier
}

type GetTradeResponse struct {
	Trade             gOanda.Trade         `json:"trade"`
	LastTransactionID gOanda.TransactionID `json:"lastTransactionID"`
	ErrorCode         string               `json:"errorCode"`
	ErrorMessage      string               `json:"errorMessage"`
}

type PutCloseTradeRequest struct {
	TradeSpecifier gOanda.TradeSpecifier
}

type PutCloseTradeResponse struct {
	OrderCreateTransaction gOanda.MarketOrderTransaction       `json:"orderCreateTransaction"`
	OrderFillTransaction   gOanda.OrderFillTransaction         `json:"orderFillTransaction"`
	OrderCancelTransaction gOanda.OrderCancelTransaction       `json:"orderCancelTransaction"`
	OrderRejectTransaction gOanda.MarketOrderRejectTransaction `json:"orderRejectTransaction"`
	RelatedTransactionsIDs []gOanda.TransactionID              `json:"relatedTransactionsIDs"`
	LastTransactionID      gOanda.TransactionID                `json:"lastTransactionID"`
	ErrorCode              string                              `json:"errorCode"`
	ErrorMessage           string                              `json:"errorMessage"`
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

type GetRangeOfTransactionsRequest struct{}

type GetRangeOfTransactionsResponse struct{}

type GetTransactionsAfterTransactionRequest struct{}

type GetTransactionsAfterTransactionResponse struct{}

type GetTransactionsStreamRequest struct{}

type GetTransactionsStreamResponse struct{}

type GetInstrumentPricingRequest struct{}

type GetInstrumentPricingResponse struct{}

type GetPricingStreamRequest struct{}

type GetPricingStreamResponse struct{}
