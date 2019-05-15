package gOanda

func New(accountNumber string, apiKey string) *Client {

	//instantiate URL Manager
	urlManager := UrlManager{
		BaseUrl:   V20_BASE_PRACTICE_URL,
		AccountId: accountNumber,
	}

	//instantiate requester
	requesterInstance := NewHttpRequester(accountNumber, apiKey, urlManager)

	return &Client{
		AccountNumber: accountNumber,
		ApiKey:        apiKey,
		requester:     requesterInstance,
	}
}

type Client struct {
	AccountNumber string
	ApiKey        string
	requester     Requester
}

func (c *Client) GetAccounts() (*GetAccountsResponse, error) {
	getAccountResponse, err := c.requester.GetAccounts(GetAccountsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Accounts: getAccountResponse.Accounts,
	}, nil
}

func (c *Client) GetAccount() (*GetAccountResponse, error) {
	getAccountResponse, err := c.requester.GetAccount(GetAccountRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: getAccountResponse.Account,
	}, nil
}

func (c *Client) GetAccountSummary() (*GetAccountSummaryResponse, error) {
	getAccountSummaryResponse, err := c.requester.GetAccountSummary(GetAccountSummaryRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountSummaryResponse{
		Account:           getAccountSummaryResponse.Account,
		LastTransactionID: getAccountSummaryResponse.LastTransactionID,
	}, nil
}

func (c *Client) GetAccountInstruments() (*GetAccountInstrumentsResponse, error) {
	_, err := c.requester.GetAccountInstruments(GetAccountInstrumentsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountInstrumentsResponse{}, nil
}

func (c *Client) PatchAccountConfiguration() (*PatchAccountConfigurationResponse, error) {
	_, err := c.requester.PatchAccountConfiguration(PatchAccountConfigurationRequest{})
	if err != nil {
		return nil, err
	}
	return &PatchAccountConfigurationResponse{}, nil
}

func (c *Client) GetInstrumentCandles(request GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error) {
	response, err := c.requester.GetInstrumentCandles(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetInstrumentOrderBook() (*GetInstrumentOrderBookResponse, error) {
	_, err := c.requester.GetInstrumentOrderBook(GetInstrumentOrderBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentOrderBookResponse{}, nil
}

func (c *Client) GetInstrumentPositionBook() (*GetInstrumentPositionBookResponse, error) {
	_, err := c.requester.GetInstrumentPositionBook(GetInstrumentPositionBookRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentPositionBookResponse{}, nil
}

func (c *Client) PostOrder(request PostOrderRequest) (*PostOrderResponse, error) {
	response, err := c.requester.PostOrder(PostOrderRequest{
		Order: request.Order,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetOrders() (*GetOrdersResponse, error) {
	_, err := c.requester.GetOrders(GetOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrdersResponse{}, nil
}

func (c *Client) GetPendingOrders() (*GetPendingOrdersResponse, error) {
	_, err := c.requester.GetPendingOrders(GetPendingOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPendingOrdersResponse{}, nil
}

func (c *Client) GetOrder() (*GetOrderResponse, error) {
	_, err := c.requester.GetOrder(GetOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOrderResponse{}, nil
}

func (c *Client) PutReplaceOrder() (*PutReplaceOrderResponse, error) {
	_, err := c.requester.PutReplaceOrder(PutReplaceOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutReplaceOrderResponse{}, nil
}

func (c *Client) PutCancelOrder() (*PutCancelOrderResponse, error) {
	_, err := c.requester.PutCancelOrder(PutCancelOrderRequest{})
	if err != nil {
		return nil, err
	}
	return &PutCancelOrderResponse{}, nil
}

func (c *Client) PutUpdateOrderClientExtensions() (*PutUpdateOrderClientExtensionsResponse, error) {
	_, err := c.requester.PutUpdateOrderClientExtensions(PutUpdateOrderClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateOrderClientExtensionsResponse{}, nil
}

func (c *Client) GetTrades(request GetTradesRequest) (*GetTradesResponse, error) {
	response, err := c.requester.GetTrades(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetOpenTrades() (*GetOpenTradesResponse, error) {
	response, err := c.requester.GetOpenTrades(GetOpenTradesRequest{})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetTrade(request GetTradeRequest) (*GetTradeResponse, error) {
	getTradeResponse, err := c.requester.GetTrade(GetTradeRequest{
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

func (c *Client) PutCloseTrade(request PutCloseTradeRequest) (*PutCloseTradeResponse, error) {
	putCloseTradeResponse, err := c.requester.PutCloseTrade(PutCloseTradeRequest{
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

func (c *Client) PutUpdateTradeClientExtensions() (*PutUpdateTradeClientExtensionsResponse, error) {
	_, err := c.requester.PutUpdateTradeClientExtensions(PutUpdateTradeClientExtensionsRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeClientExtensionsResponse{}, nil
}

func (c *Client) PutUpdateTradeDependentOrders() (*PutUpdateTradeDependentOrdersResponse, error) {
	_, err := c.requester.PutUpdateTradeDependentOrders(PutUpdateTradeDependentOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return &PutUpdateTradeDependentOrdersResponse{}, nil
}

func (c *Client) GetPositions() (*GetPositionsResponse, error) {
	_, err := c.requester.GetPositions(GetPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPositionsResponse{}, nil
}

func (c *Client) GetOpenPositions() (*GetOpenPositionsResponse, error) {
	_, err := c.requester.GetOpenPositions(GetOpenPositionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetOpenPositionsResponse{}, nil
}

func (c *Client) GetInstrumentsPosition() (*GetInstrumentsPositionResponse, error) {
	_, err := c.requester.GetInstrumentsPosition(GetInstrumentsPositionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetInstrumentsPositionResponse{}, nil
}

func (c *Client) PutClosePosition() (*PutClosePositionResponse, error) {
	_, err := c.requester.PutClosePosition(PutClosePositionRequest{})
	if err != nil {
		return nil, err
	}
	return &PutClosePositionResponse{}, nil
}

func (c *Client) GetTransactions() (*GetTransactionsResponse, error) {
	_, err := c.requester.GetTransactions(GetTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsResponse{}, nil
}

func (c *Client) GetTransaction() (*GetTransactionResponse, error) {
	_, err := c.requester.GetTransaction(GetTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionResponse{}, nil
}

func (c *Client) GetRangeOfTransactions() (*GetRangeOfTransactionsResponse, error) {
	_, err := c.requester.GetRangeOfTransactions(GetRangeOfTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetRangeOfTransactionsResponse{}, nil
}

func (c *Client) GetTransactionsAfterTransaction() (*GetTransactionsAfterTransactionResponse, error) {
	_, err := c.requester.GetTransactionsAfterTransaction(GetTransactionsAfterTransactionRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsAfterTransactionResponse{}, nil
}

func (c *Client) GetTransactionsStream() (*GetTransactionsStreamResponse, error) {
	_, err := c.requester.GetTransactionsStream(GetTransactionsStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetTransactionsStreamResponse{}, nil
}

func (c *Client) GetInstrumentPricing(request GetInstrumentPricingRequest) (*GetInstrumentPricingResponse, error) {
	response, err := c.requester.GetInstrumentPricing(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetPricingStream() (*GetPricingStreamResponse, error) {
	_, err := c.requester.GetPricingStream(GetPricingStreamRequest{})
	if err != nil {
		return nil, err
	}
	return &GetPricingStreamResponse{}, nil
}
