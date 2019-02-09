package V20

import (
	"github.com/hannessi/gOanda"
	"github.com/hannessi/gOanda/V20/definitions"
	"github.com/hannessi/gOanda/V20/requester/V20"
	"github.com/hannessi/gOanda/V20/requester/V20/api"
	"github.com/hannessi/gOanda/V20/requester/V20/url"
)

func New(accountNumber string, apiKey string) *Client {

	//instantiate URL Manager
	urlManager := url.Manager{
		BaseUrl:   gOanda.V20_BASE_PRACTICE_URL,
		AccountId: accountNumber,
	}

	//instantiate requester
	requester := api.New(accountNumber, apiKey, urlManager)

	return &Client{
		AccountNumber: accountNumber,
		ApiKey:        apiKey,
		Requester:     requester,
	}
}

type Client struct {
	AccountNumber string
	ApiKey        string
	Requester     V20.Requester
}

// account endpoints
type GetAccountsResponse struct {
	Account []definitions.Account
}

func (c *Client) GetAccounts() (*GetAccountsResponse, error) {
	getAccountResponse, err := c.Requester.GetAccounts(V20.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Account: getAccountResponse.Accounts,
	}, nil
}

type GetAccountResponse struct {
	Account definitions.Account
}

func (c *Client) GetAccount() (*GetAccountResponse, error) {
	getAccountResponse, err := c.Requester.GetAccount(V20.GetAccountRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: getAccountResponse.Account,
	}, nil
}

type GetAccountSummaryResponse struct {
	// todo
}

func (c *Client) GetAccountSummary() (*GetAccountSummaryResponse, error) {
	_, err := c.Requester.GetAccountSummary(V20.GetAccountSummaryRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountSummaryResponse{}, nil
}

type GetAccountInstrumentsResponse struct {
	// todo
}

func (c *Client) GetAccountInstruments() (*GetAccountInstrumentsResponse, error) {
	_, err := c.Requester.GetAccountInstruments(V20.GetAccountInstrumentsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountInstrumentsResponse{}, nil
}

type PatchAccountConfigurationResponse struct {
	// todo
}

func (c *Client) PatchAccountConfiguration() (*PatchAccountConfigurationResponse, error) {
	_, err := c.Requester.PatchAccountConfiguration(V20.PatchAccountConfigurationRequest{})
	if err != nil {
		return nil, err
	}
	return &PatchAccountConfigurationResponse{}, nil
}

type GetInstrumentCandlesResponse struct {
	// todo
}

func (c *Client) GetInstrumentCandles() (*GetInstrumentCandlesResponse, error) {
	_, err := c.Requester.GetInstrumentCandles(V20.GetInstrumentCandlesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentCandlesResponse{}, nil
}

type GetInstrumentOrderBookResponse struct {
	// todo
}

func (c *Client) GetInstrumentOrderBook() (*GetInstrumentOrderBookResponse, error) {
	_, err := c.Requester.GetInstrumentOrderBook(V20.GetInstrumentOrderBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentOrderBookResponse{}, nil
}

type GetInstrumentPositionBookResponse struct {
	// todo
}

func (c *Client) GetInstrumentPositionBook() (*GetInstrumentPositionBookResponse, error) {
	_, err := c.Requester.GetInstrumentPositionBook(V20.GetInstrumentPositionBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentPositionBookResponse{}, nil
}

type PostOrderResponse struct {
	// todo
}

func (c *Client) PostOrder() (*PostOrderResponse, error) {
	_, err := c.Requester.PostOrder(V20.PostOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PostOrderResponse{}, nil
}

type GetOrdersResponse struct {
	// todo
}

func (c *Client) GetOrders() (*GetOrdersResponse, error) {
	_, err := c.Requester.GetOrders(V20.GetOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrdersResponse{}, nil
}

type GetPendingOrdersResponse struct {
	// todo
}

func (c *Client) GetPendingOrders() (*GetPendingOrdersResponse, error) {
	_, err := c.Requester.GetPendingOrders(V20.GetPendingOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPendingOrdersResponse{}, nil
}

type GetOrderResponse struct {
	// todo
}

func (c *Client) GetOrder() (*GetOrderResponse, error) {
	_, err := c.Requester.GetOrder(V20.GetOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrderResponse{}, nil
}

type PutReplaceOrderResponse struct {
	// todo
}

func (c *Client) PutReplaceOrder() (*PutReplaceOrderResponse, error) {
	_, err := c.Requester.PutReplaceOrder(V20.PutReplaceOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutReplaceOrderResponse{}, nil
}

type PutCancelOrderResponse struct {
	// todo
}

func (c *Client) PutCancelOrder() (*PutCancelOrderResponse, error) {
	_, err := c.Requester.PutCancelOrder(V20.PutCancelOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutCancelOrderResponse{}, nil
}

type PutUpdateOrderClientExtensionsResponse struct {
	// todo
}

func (c *Client) PutUpdateOrderClientExtensions() (*PutUpdateOrderClientExtensionsResponse, error) {
	_, err := c.Requester.PutUpdateOrderClientExtensions(V20.PutUpdateOrderClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateOrderClientExtensionsResponse{}, nil
}

type GetTradesResponse struct {
	// todo
}

func (c *Client) GetTrades() (*GetTradesResponse, error) {
	_, err := c.Requester.GetTrades(V20.GetTradesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTradesResponse{}, nil
}

type GetOpenTradesResponse struct {
	// todo
}

func (c *Client) GetOpenTrades() (*GetOpenTradesResponse, error) {
	_, err := c.Requester.GetOpenTrades(V20.GetOpenTradesRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOpenTradesResponse{}, nil
}

type GetTradeResponse struct {
	// todo
}

func (c *Client) GetTrade() (*GetTradeResponse, error) {
	_, err := c.Requester.GetTrade(V20.GetTradeRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTradeResponse{}, nil
}

type PutCloseTradeResponse struct {
	// todo
}

func (c *Client) PutCloseTrade() (*PutCloseTradeResponse, error) {
	_, err := c.Requester.PutCloseTrade(V20.PutCloseTradeRequest{})
	if err != nil {
		return nil, err
	}
	return &PutCloseTradeResponse{}, nil
}

type PutUpdateTradeClientExtensionsResponse struct {
	// todo
}

func (c *Client) PutUpdateTradeClientExtensions() (*PutUpdateTradeClientExtensionsResponse, error) {
	_, err := c.Requester.PutUpdateTradeClientExtensions(V20.PutUpdateTradeClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeClientExtensionsResponse{}, nil
}

type PutUpdateTradeDependentOrdersResponse struct {
	// todo
}

func (c *Client) PutUpdateTradeDependentOrders() (*PutUpdateTradeDependentOrdersResponse, error) {
	_, err := c.Requester.PutUpdateTradeDependentOrders(V20.PutUpdateTradeDependentOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeDependentOrdersResponse{}, nil
}

type GetPositionsResponse struct {
	// todo
}

func (c *Client) GetPositions() (*GetPositionsResponse, error) {
	_, err := c.Requester.GetPositions(V20.GetPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPositionsResponse{}, nil
}

type GetOpenPositionsResponse struct {
	// todo
}

func (c *Client) GetOpenPositions() (*GetOpenPositionsResponse, error) {
	_, err := c.Requester.GetOpenPositions(V20.GetOpenPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOpenPositionsResponse{}, nil
}

type GetInstrumentsPositionResponse struct {
	// todo
}

func (c *Client) GetInstrumentsPosition() (*GetInstrumentsPositionResponse, error) {
	_, err := c.Requester.GetInstrumentsPosition(V20.GetInstrumentsPositionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentsPositionResponse{}, nil
}

type PutClosePositionResponse struct {
	// todo
}

func (c *Client) PutClosePosition() (*PutClosePositionResponse, error) {
	_, err := c.Requester.PutClosePosition(V20.PutClosePositionRequest{})
	if err != nil {
		return nil, err
	}
	return &PutClosePositionResponse{}, nil
}

type GetTransactionsResponse struct {
	// todo
}

func (c *Client) GetTransactions() (*GetTransactionsResponse, error) {
	_, err := c.Requester.GetTransactions(V20.GetTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsResponse{}, nil
}

type GetTransactionResponse struct {
	// todo
}

func (c *Client) GetTransaction() (*GetTransactionResponse, error) {
	_, err := c.Requester.GetTransaction(V20.GetTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionResponse{}, nil
}

type GetRangeOfTransactionsResponse struct {
	// todo
}

func (c *Client) GetRangeOfTransactions() (*GetRangeOfTransactionsResponse, error) {
	_, err := c.Requester.GetRangeOfTransactions(V20.GetRangeOfTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetRangeOfTransactionsResponse{}, nil
}

type GetTransactionsAfterTransactionResponse struct {
	// todo
}

func (c *Client) GetTransactionsAfterTransaction() (*GetTransactionsAfterTransactionResponse, error) {
	_, err := c.Requester.GetTransactionsAfterTransaction(V20.GetTransactionsAfterTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsAfterTransactionResponse{}, nil
}

type GetTransactionsStreamResponse struct {
	// todo
}

func (c *Client) GetTransactionsStream() (*GetTransactionsStreamResponse, error) {
	_, err := c.Requester.GetTransactionsStream(V20.GetTransactionsStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsStreamResponse{}, nil
}

type GetInstrumentPricingResponse struct {
	// todo
}

func (c *Client) GetInstrumentPricing() (*GetInstrumentPricingResponse, error) {
	_, err := c.Requester.GetInstrumentPricing(V20.GetInstrumentPricingRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentPricingResponse{}, nil
}

type GetPricingStreamResponse struct {
	// todo
}

func (c *Client) GetPricingStream() (*GetPricingStreamResponse, error) {
	_, err := c.Requester.GetPricingStream(V20.GetPricingStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPricingStreamResponse{}, nil
}
