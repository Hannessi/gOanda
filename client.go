package gOanda

func New(accountNumber string, apiKey string, live bool) *Client {

	url := V20_BASE_TRADE_URL
	if !live {
		url = V20_BASE_PRACTICE_URL
	}

	//instantiate URL Manager
	urlManager := UrlManager{
		BaseUrl:   url,
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

func (c *Client) GetOrders(request GetOrdersRequest) (*GetOrdersResponse, error) {
	response, err := c.requester.GetOrders(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetPendingOrders() (*GetPendingOrdersResponse, error) {
	response, err := c.requester.GetPendingOrders(GetPendingOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetOrder(request GetOrderRequest) (*GetOrderResponse, error) {
	response, err := c.requester.GetOrder(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PutReplaceOrder(request PutReplaceOrderRequest) (*PutReplaceOrderResponse, error) {
	response, err := c.requester.PutReplaceOrder(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PutCancelOrder(request PutCancelOrderRequest) (*PutCancelOrderResponse, error) {
	response, err := c.requester.PutCancelOrder(request)
	if err != nil {
		return nil, err
	}
	return response, nil
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
	response, err := c.requester.GetTrade(GetTradeRequest{
		TradeSpecifier: request.TradeSpecifier,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PutCloseTrade(request PutCloseTradeRequest) (*PutCloseTradeResponse, error) {
	putCloseTradeResponse, err := c.requester.PutCloseTrade(request)
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

func (c *Client) GetTransaction(request GetTransactionRequest) (*GetTransactionResponse, error) {
	response, err := c.requester.GetTransaction(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetRangeOfTransactions() (*GetRangeOfTransactionsResponse, error) {
	_, err := c.requester.GetRangeOfTransactions(GetRangeOfTransactionsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetRangeOfTransactionsResponse{}, nil
}

func (c *Client) GetTransactionsSinceId(request GetTransactionsSinceIdRequest) (*GetTransactionsSinceIdResponse, error) {
	response, err := c.requester.GetTransactionsSinceId(request)
	if err != nil {
		return nil, err
	}
	return &GetTransactionsSinceIdResponse{
		Transactions:      response.Transactions,
		LastTransactionID: response.LastTransactionID,
	}, nil
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
