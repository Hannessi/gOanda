package api

import (
	"errors"
	"github.com/hannessi/gOanda"
	"github.com/hannessi/gOanda/V20/requester"
	"github.com/hannessi/gOanda/V20/requester/url"
)

func New(
	AccountId string,
	Token string,
	UrlManager url.Manager,
) *Requester {
	return &Requester{
		AccountId:  AccountId,
		Token:      Token,
		UrlManager: UrlManager,
	}
}

type Requester struct {
	AccountId  string
	Token      string
	UrlManager url.Manager
}

// help

// accounts
func (r *Requester) GetAccounts(request requester.GetAccountsRequest) (*requester.GetAccountsResponse, error) {
	response := &requester.GetAccountsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccounts(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) GetAccount(request requester.GetAccountRequest) (*requester.GetAccountResponse, error) {
	response := &requester.GetAccountResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccount(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) GetAccountSummary(request requester.GetAccountSummaryRequest) (*requester.GetAccountSummaryResponse, error) {
	response := &requester.GetAccountSummaryResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountSummary(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) GetAccountInstruments(request requester.GetAccountInstrumentsRequest) (*requester.GetAccountInstrumentsResponse, error) {
	response := &requester.GetAccountInstrumentsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountInstruments(), nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) PatchAccountConfiguration(request requester.PatchAccountConfigurationRequest) (*requester.PatchAccountConfigurationResponse, error) {
	return nil, errors.New("not implemented yet")
}

// instruments
func (r *Requester) GetInstrumentCandles(request requester.GetInstrumentCandlesRequest) (*requester.GetInstrumentCandlesResponse, error) {
	response := &requester.GetInstrumentCandlesResponse{}
	requestUrl := r.UrlManager.GetInstrumentCandles(url.GetInstrumentCandlesRequestParameters{
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		Granularity:    request.Granularity,
	})
	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) GetInstrumentOrderBook(request requester.GetInstrumentOrderBookRequest) (*requester.GetInstrumentOrderBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetInstrumentPositionBook(request requester.GetInstrumentPositionBookRequest) (*requester.GetInstrumentPositionBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

// orders
func (r *Requester) PostOrder(request requester.PostOrderRequest) (*requester.PostOrderResponse, error) {
	response := &requester.PostOrderResponse{}
	requestUrl := r.UrlManager.PostOrder()

	order := struct {
		Order gOanda.OrderRequest `json:"order"`
	}{
		Order: request.Order.ToOrderRequest(),
	}

	if err := HttpRequestWrapper(POST, requestUrl, order, response, r.Token); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Requester) GetOrders(request requester.GetOrdersRequest) (*requester.GetOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetPendingOrders(request requester.GetPendingOrdersRequest) (*requester.GetPendingOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOrder(request requester.GetOrderRequest) (*requester.GetOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutReplaceOrder(request requester.PutReplaceOrderRequest) (*requester.PutReplaceOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutCancelOrder(request requester.PutCancelOrderRequest) (*requester.PutCancelOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateOrderClientExtensions(request requester.PutUpdateOrderClientExtensionsRequest) (*requester.PutUpdateOrderClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

// trades
func (r *Requester) GetTrades(request requester.GetTradesRequest) (*requester.GetTradesResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOpenTrades(request requester.GetOpenTradesRequest) (*requester.GetOpenTradesResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTrade(request requester.GetTradeRequest) (*requester.GetTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutCloseTrade(request requester.PutCloseTradeRequest) (*requester.PutCloseTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateTradeClientExtensions(request requester.PutUpdateTradeClientExtensionsRequest) (*requester.PutUpdateTradeClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateTradeDependentOrders(request requester.PutUpdateTradeDependentOrdersRequest) (*requester.PutUpdateTradeDependentOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

// positions
func (r *Requester) GetPositions(request requester.GetPositionsRequest) (*requester.GetPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOpenPositions(request requester.GetOpenPositionsRequest) (*requester.GetOpenPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetInstrumentsPosition(request requester.GetInstrumentsPositionRequest) (*requester.GetInstrumentsPositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutClosePosition(request requester.PutClosePositionRequest) (*requester.PutClosePositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

// transactions
func (r *Requester) GetTransactions(request requester.GetTransactionsRequest) (*requester.GetTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransaction(request requester.GetTransactionRequest) (*requester.GetTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetRangeOfTransactions(request requester.GetRangeOfTransactionsRequest) (*requester.GetRangeOfTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransactionsAfterTransaction(request requester.GetTransactionsAfterTransactionRequest) (*requester.GetTransactionsAfterTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransactionsStream(request requester.GetTransactionsStreamRequest) (*requester.GetTransactionsStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}

// pricing
func (r *Requester) GetInstrumentPricing(request requester.GetInstrumentPricingRequest) (*requester.GetInstrumentPricingResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetPricingStream(request requester.GetPricingStreamRequest) (*requester.GetPricingStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}
