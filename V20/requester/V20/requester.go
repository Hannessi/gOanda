package V20

import (
	"github.com/hannessi/gOanda/V20/definitions"
)

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
	Accounts []definitions.Account `json:"accounts"`
}

type GetAccountRequest struct{}

type GetAccountResponse struct {
	Account definitions.Account `json:"account"`
}

type GetAccountSummaryRequest struct{}

type GetAccountSummaryResponse struct{}

type GetAccountInstrumentsRequest struct{}

type GetAccountInstrumentsResponse struct{}

type PatchAccountConfigurationRequest struct{}

type PatchAccountConfigurationResponse struct{}

type GetInstrumentCandlesRequest struct{}

type GetInstrumentCandlesResponse struct{}

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

type GetTradeRequest struct{}

type GetTradeResponse struct{}

type PutCloseTradeRequest struct{}

type PutCloseTradeResponse struct{}

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
