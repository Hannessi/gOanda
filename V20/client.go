package gOanda

import (
	"github.com/hannessi/gOanda"
	"github.com/hannessi/gOanda/V20/requester"
	"github.com/hannessi/gOanda/V20/requester/api"
	"github.com/hannessi/gOanda/V20/requester/url"
)

func New(accountNumber string, apiKey string) *Client {

	//instantiate URL Manager
	urlManager := url.Manager{
		BaseUrl:   gOanda.V20_BASE_PRACTICE_URL,
		AccountId: accountNumber,
	}

	//instantiate requester
	requesterInstance := api.New(accountNumber, apiKey, urlManager)

	return &Client{
		AccountNumber: accountNumber,
		ApiKey:        apiKey,
		requester:     requesterInstance,
	}
}

type Client struct {
	AccountNumber string
	ApiKey        string
	requester     requester.Requester
}

// account endpoints
type GetAccountsResponse struct {
	Account []gOanda.AccountProperties
}

func (c *Client) GetAccounts() (*GetAccountsResponse, error) {
	getAccountResponse, err := c.requester.GetAccounts(requester.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Account: getAccountResponse.Accounts,
	}, nil
}

type GetAccountResponse struct {
	Account gOanda.Account
}

func (c *Client) GetAccount() (*GetAccountResponse, error) {
	getAccountResponse, err := c.requester.GetAccount(requester.GetAccountRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: getAccountResponse.Account,
	}, nil
}

type GetAccountSummaryResponse struct {
	Account           gOanda.AccountSummary
	LastTransactionID gOanda.TransactionID
}

func (c *Client) GetAccountSummary() (*GetAccountSummaryResponse, error) {
	getAccountSummaryResponse, err := c.requester.GetAccountSummary(requester.GetAccountSummaryRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountSummaryResponse{
		Account:           getAccountSummaryResponse.Account,
		LastTransactionID: getAccountSummaryResponse.LastTransactionID,
	}, nil
}

type GetAccountInstrumentsResponse struct {
	// todo
}

func (c *Client) GetAccountInstruments() (*GetAccountInstrumentsResponse, error) {
	_, err := c.requester.GetAccountInstruments(requester.GetAccountInstrumentsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountInstrumentsResponse{}, nil
}

type PatchAccountConfigurationResponse struct {
	// todo
}

func (c *Client) PatchAccountConfiguration() (*PatchAccountConfigurationResponse, error) {
	_, err := c.requester.PatchAccountConfiguration(requester.PatchAccountConfigurationRequest{})
	if err != nil {
		return nil, err
	}
	return &PatchAccountConfigurationResponse{}, nil
}

type GetInstrumentCandlesRequest struct {
	InstrumentName gOanda.InstrumentName
	Count          int
	Granularity    gOanda.CandlestickGranularity
}

type GetInstrumentCandlesResponse struct {
	Instrument  gOanda.InstrumentName         `json:"instrument"`
	Granularity gOanda.CandlestickGranularity `json:"granularity"`
	Candles     []gOanda.Candlestick          `json:"candles"`
}

func (c *Client) GetInstrumentCandles(request GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error) {
	getInstrumentCandlesResponse, err := c.requester.GetInstrumentCandles(requester.GetInstrumentCandlesRequest{
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		Granularity:    request.Granularity,
	})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentCandlesResponse{
		Instrument:  getInstrumentCandlesResponse.Instrument,
		Granularity: getInstrumentCandlesResponse.Granularity,
		Candles:     getInstrumentCandlesResponse.Candles,
	}, nil
}

type GetInstrumentOrderBookResponse struct {
	// todo
}

func (c *Client) GetInstrumentOrderBook() (*GetInstrumentOrderBookResponse, error) {
	_, err := c.requester.GetInstrumentOrderBook(requester.GetInstrumentOrderBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentOrderBookResponse{}, nil
}

type GetInstrumentPositionBookResponse struct {
	// todo
}

func (c *Client) GetInstrumentPositionBook() (*GetInstrumentPositionBookResponse, error) {
	_, err := c.requester.GetInstrumentPositionBook(requester.GetInstrumentPositionBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentPositionBookResponse{}, nil
}

type PostOrderRequest struct {
	Order gOanda.AnOrderRequest
}

type PostOrderResponse struct {
	OrderCreateTransaction        gOanda.Transaction
	OrderFillTransaction          gOanda.OrderFillTransaction
	OrderCancelTransaction        gOanda.OrderCancelTransaction
	OrderReissueTransaction       gOanda.Transaction
	OrderReissueRejectTransaction gOanda.Transaction
	RelatedTransactionIDs         []gOanda.TransactionID
	LastTransactionID             gOanda.TransactionID
	ErrorCode                     string
	ErrorMessage                  string
}

func (c *Client) PostOrder(request PostOrderRequest) (*PostOrderResponse, error) {
	response, err := c.requester.PostOrder(requester.PostOrderRequest{
		Order: request.Order,
	})
	if err != nil {
		return nil, err
	}
	return &PostOrderResponse{
		OrderCreateTransaction:        response.OrderCreateTransaction,
		OrderFillTransaction:          response.OrderFillTransaction,
		OrderCancelTransaction:        response.OrderCancelTransaction,
		OrderReissueTransaction:       response.OrderReissueTransaction,
		OrderReissueRejectTransaction: response.OrderReissueRejectTransaction,
		RelatedTransactionIDs:         response.RelatedTransactionIDs,
		LastTransactionID:             response.LastTransactionID,
		ErrorCode:                     response.ErrorCode,
		ErrorMessage:                  response.ErrorMessage,
	}, nil
}

type GetOrdersResponse struct {
	// todo
}

func (c *Client) GetOrders() (*GetOrdersResponse, error) {
	_, err := c.requester.GetOrders(requester.GetOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrdersResponse{}, nil
}

type GetPendingOrdersResponse struct {
	// todo
}

func (c *Client) GetPendingOrders() (*GetPendingOrdersResponse, error) {
	_, err := c.requester.GetPendingOrders(requester.GetPendingOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPendingOrdersResponse{}, nil
}

type GetOrderResponse struct {
	// todo
}

func (c *Client) GetOrder() (*GetOrderResponse, error) {
	_, err := c.requester.GetOrder(requester.GetOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrderResponse{}, nil
}

type PutReplaceOrderResponse struct {
	// todo
}

func (c *Client) PutReplaceOrder() (*PutReplaceOrderResponse, error) {
	_, err := c.requester.PutReplaceOrder(requester.PutReplaceOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutReplaceOrderResponse{}, nil
}

type PutCancelOrderResponse struct {
	// todo
}

func (c *Client) PutCancelOrder() (*PutCancelOrderResponse, error) {
	_, err := c.requester.PutCancelOrder(requester.PutCancelOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutCancelOrderResponse{}, nil
}

type PutUpdateOrderClientExtensionsResponse struct {
	// todo
}

func (c *Client) PutUpdateOrderClientExtensions() (*PutUpdateOrderClientExtensionsResponse, error) {
	_, err := c.requester.PutUpdateOrderClientExtensions(requester.PutUpdateOrderClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateOrderClientExtensionsResponse{}, nil
}

type GetTradesResponse struct {
	Trades            []gOanda.Trade
	LastTransactionID gOanda.TransactionID
}

type GetTradesRequest struct {
	Ids            []gOanda.TradeID
	State          gOanda.TradeStateFilter
	InstrumentName gOanda.InstrumentName
	Count          int64
	BeforeID       gOanda.TradeID
}

func (c *Client) GetTrades(request GetTradesRequest) (*GetTradesResponse, error) {
	response, err := c.requester.GetTrades(requester.GetTradesRequest{
		Ids:            request.Ids,
		State:          request.State,
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		BeforeID:       request.BeforeID,
	})
	if err != nil {
		return nil, err
	}
	return &GetTradesResponse{
		Trades:            response.Trades,
		LastTransactionID: response.LastTransactionID,
	}, nil
}

type GetOpenTradesResponse struct {
}

func (c *Client) GetOpenTrades() (*GetOpenTradesResponse, error) {
	_, err := c.requester.GetOpenTrades(requester.GetOpenTradesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOpenTradesResponse{}, nil
}

type GetTradeRequest struct {
	TradeSpecifier gOanda.TradeSpecifier
}

type GetTradeResponse struct {
	Trade             gOanda.Trade
	LastTransactionID gOanda.TransactionID
}

func (c *Client) GetTrade(request GetTradeRequest) (*GetTradeResponse, error) {
	getTradeResponse, err := c.requester.GetTrade(requester.GetTradeRequest{
		TradeSpecifier: request.TradeSpecifier,
	})
	if err != nil {
		return nil, err
	}
	return &GetTradeResponse{
		Trade:             getTradeResponse.Trade,
		LastTransactionID: getTradeResponse.LastTransactionID,
	}, nil
}

type PutCloseTradeRequest struct {
	TradeSpecifier gOanda.TradeSpecifier
	Units          int // todo in body of request
}

type PutCloseTradeResponse struct {
	OrderCreateTransaction gOanda.MarketOrderTransaction
	OrderFillTransaction   gOanda.OrderFillTransaction
	OrderCancelTransaction gOanda.OrderCancelTransaction
	OrderRejectTransaction gOanda.MarketOrderRejectTransaction
	RelatedTransactionsIDs []gOanda.TransactionID
	LastTransactionID      gOanda.TransactionID
}

func (c *Client) PutCloseTrade(request PutCloseTradeRequest) (*PutCloseTradeResponse, error) {
	putCloseTradeResponse, err := c.requester.PutCloseTrade(requester.PutCloseTradeRequest{
		TradeSpecifier: request.TradeSpecifier,
	})
	if err != nil {
		return nil, err
	}
	return &PutCloseTradeResponse{
		OrderCreateTransaction: putCloseTradeResponse.OrderCreateTransaction,
		OrderFillTransaction:   putCloseTradeResponse.OrderFillTransaction,
		OrderCancelTransaction: putCloseTradeResponse.OrderCancelTransaction,
		OrderRejectTransaction: putCloseTradeResponse.OrderRejectTransaction,
		RelatedTransactionsIDs: putCloseTradeResponse.RelatedTransactionsIDs,
		LastTransactionID:      putCloseTradeResponse.LastTransactionID,
	}, nil
}

type PutUpdateTradeClientExtensionsResponse struct {
	// todo
}

func (c *Client) PutUpdateTradeClientExtensions() (*PutUpdateTradeClientExtensionsResponse, error) {
	_, err := c.requester.PutUpdateTradeClientExtensions(requester.PutUpdateTradeClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeClientExtensionsResponse{}, nil
}

type PutUpdateTradeDependentOrdersResponse struct {
	// todo
}

func (c *Client) PutUpdateTradeDependentOrders() (*PutUpdateTradeDependentOrdersResponse, error) {
	_, err := c.requester.PutUpdateTradeDependentOrders(requester.PutUpdateTradeDependentOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeDependentOrdersResponse{}, nil
}

type GetPositionsResponse struct {
	// todo
}

func (c *Client) GetPositions() (*GetPositionsResponse, error) {
	_, err := c.requester.GetPositions(requester.GetPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPositionsResponse{}, nil
}

type GetOpenPositionsResponse struct {
	// todo
}

func (c *Client) GetOpenPositions() (*GetOpenPositionsResponse, error) {
	_, err := c.requester.GetOpenPositions(requester.GetOpenPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOpenPositionsResponse{}, nil
}

type GetInstrumentsPositionResponse struct {
	// todo
}

func (c *Client) GetInstrumentsPosition() (*GetInstrumentsPositionResponse, error) {
	_, err := c.requester.GetInstrumentsPosition(requester.GetInstrumentsPositionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentsPositionResponse{}, nil
}

type PutClosePositionResponse struct {
	// todo
}

func (c *Client) PutClosePosition() (*PutClosePositionResponse, error) {
	_, err := c.requester.PutClosePosition(requester.PutClosePositionRequest{})
	if err != nil {
		return nil, err
	}
	return &PutClosePositionResponse{}, nil
}

type GetTransactionsResponse struct {
	// todo
}

func (c *Client) GetTransactions() (*GetTransactionsResponse, error) {
	_, err := c.requester.GetTransactions(requester.GetTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsResponse{}, nil
}

type GetTransactionResponse struct {
	// todo
}

func (c *Client) GetTransaction() (*GetTransactionResponse, error) {
	_, err := c.requester.GetTransaction(requester.GetTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionResponse{}, nil
}

type GetRangeOfTransactionsResponse struct {
	// todo
}

func (c *Client) GetRangeOfTransactions() (*GetRangeOfTransactionsResponse, error) {
	_, err := c.requester.GetRangeOfTransactions(requester.GetRangeOfTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetRangeOfTransactionsResponse{}, nil
}

type GetTransactionsAfterTransactionResponse struct {
	// todo
}

func (c *Client) GetTransactionsAfterTransaction() (*GetTransactionsAfterTransactionResponse, error) {
	_, err := c.requester.GetTransactionsAfterTransaction(requester.GetTransactionsAfterTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsAfterTransactionResponse{}, nil
}

type GetTransactionsStreamResponse struct {
	// todo
}

func (c *Client) GetTransactionsStream() (*GetTransactionsStreamResponse, error) {
	_, err := c.requester.GetTransactionsStream(requester.GetTransactionsStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsStreamResponse{}, nil
}

type GetInstrumentPricingResponse struct {
	// todo
}

func (c *Client) GetInstrumentPricing() (*GetInstrumentPricingResponse, error) {
	_, err := c.requester.GetInstrumentPricing(requester.GetInstrumentPricingRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentPricingResponse{}, nil
}

type GetPricingStreamResponse struct {
	// todo
}

func (c *Client) GetPricingStream() (*GetPricingStreamResponse, error) {
	_, err := c.requester.GetPricingStream(requester.GetPricingStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPricingStreamResponse{}, nil
}
