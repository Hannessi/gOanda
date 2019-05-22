package gOanda

import (
	"errors"
)

func NewHttpRequester(
	AccountId string,
	Token string,
	UrlManager UrlManager,
) *HttpRequester {
	return &HttpRequester{
		AccountId:  AccountId,
		Token:      Token,
		UrlManager: UrlManager,
	}
}

type HttpRequester struct {
	AccountId  string
	Token      string
	UrlManager UrlManager
}

// help

// accounts
func (r *HttpRequester) GetAccounts(request GetAccountsRequest) (*GetAccountsResponse, error) {
	response := &GetAccountsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccounts(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *HttpRequester) GetAccount(request GetAccountRequest) (*GetAccountResponse, error) {
	response := &GetAccountResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccount(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *HttpRequester) GetAccountSummary(request GetAccountSummaryRequest) (*GetAccountSummaryResponse, error) {
	response := &GetAccountSummaryResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountSummary(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *HttpRequester) GetAccountInstruments(request GetAccountInstrumentsRequest) (*GetAccountInstrumentsResponse, error) {
	response := &GetAccountInstrumentsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountInstruments(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *HttpRequester) PatchAccountConfiguration(request PatchAccountConfigurationRequest) (*PatchAccountConfigurationResponse, error) {
	return nil, errors.New("not implemented yet")
}

// instruments
func (r *HttpRequester) GetInstrumentCandles(request GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error) {
	response := &GetInstrumentCandlesResponse{}
	requestUrl := r.UrlManager.GetInstrumentCandles(GetInstrumentCandlesRequestParameters{
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		Granularity:    request.Granularity,
	})
	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *HttpRequester) GetInstrumentOrderBook(request GetInstrumentOrderBookRequest) (*GetInstrumentOrderBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetInstrumentPositionBook(request GetInstrumentPositionBookRequest) (*GetInstrumentPositionBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

// orders
func (r *HttpRequester) PostOrder(request PostOrderRequest) (*PostOrderResponse, error) {
	response := &PostOrderResponse{}
	requestUrl := r.UrlManager.PostOrder()

	order := struct {
		Order OrderRequest `json:"order"`
	}{
		Order: request.Order.ToOrderRequest(),
	}

	if err := HttpRequestWrapper(POST, requestUrl, order, response, r.Token); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *HttpRequester) GetOrders(request GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetPendingOrders(request GetPendingOrdersRequest) (*GetPendingOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetOrder(request GetOrderRequest) (*GetOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutReplaceOrder(request PutReplaceOrderRequest) (*PutReplaceOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutCancelOrder(request PutCancelOrderRequest) (*PutCancelOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutUpdateOrderClientExtensions(request PutUpdateOrderClientExtensionsRequest) (*PutUpdateOrderClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

// trades
func (r *HttpRequester) GetTrades(request GetTradesRequest) (*GetTradesResponse, error) {
	response := &GetTradesResponse{}
	requestUrl := r.UrlManager.GetTrades(GetTradesRequestParameters{
		Ids:            request.Ids,
		State:          request.State,
		InstrumentName: request.InstrumentName,
		BeforeID:       request.BeforeID,
		Count:          request.Count,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *HttpRequester) GetOpenTrades(request GetOpenTradesRequest) (*GetOpenTradesResponse, error) {
	response := &GetOpenTradesResponse{}
	requestUrl := r.UrlManager.GetOpenTrades()

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *HttpRequester) GetTrade(request GetTradeRequest) (*GetTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutCloseTrade(request PutCloseTradeRequest) (*PutCloseTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutUpdateTradeClientExtensions(request PutUpdateTradeClientExtensionsRequest) (*PutUpdateTradeClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutUpdateTradeDependentOrders(request PutUpdateTradeDependentOrdersRequest) (*PutUpdateTradeDependentOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

// positions
func (r *HttpRequester) GetPositions(request GetPositionsRequest) (*GetPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetOpenPositions(request GetOpenPositionsRequest) (*GetOpenPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetInstrumentsPosition(request GetInstrumentsPositionRequest) (*GetInstrumentsPositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) PutClosePosition(request PutClosePositionRequest) (*PutClosePositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

// transactions
func (r *HttpRequester) GetTransactions(request GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetTransaction(request GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetRangeOfTransactions(request GetRangeOfTransactionsRequest) (*GetRangeOfTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetTransactionsAfterTransaction(request GetTransactionsAfterTransactionRequest) (*GetTransactionsAfterTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *HttpRequester) GetTransactionsStream(request GetTransactionsStreamRequest) (*GetTransactionsStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}

// pricing
func (r *HttpRequester) GetInstrumentPricing(request GetInstrumentPricingRequest) (*GetInstrumentPricingResponse, error) {
	response := &GetInstrumentPricingResponse{}
	requestUrl := r.UrlManager.GetInstrumentPricing(GetInstrumentPricingParameters{
		Instruments:            request.Instruments,
		Since:                  request.Since,
		IncludeUnitsAvailable:  request.IncludeUnitsAvailable,
		IncludeHomeConversions: request.IncludeHomeConversions,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *HttpRequester) GetPricingStream(request GetPricingStreamRequest) (*GetPricingStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}
