package api

import (
	"encoding/json"
	"errors"
	"github.com/hannessi/gOanda/V20/requester/V20"
	"github.com/hannessi/gOanda/V20/requester/V20/url"
	"io/ioutil"
	"net/http"
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

// accounts
func (r *Requester) GetAccounts(request V20.GetAccountsRequest) (*V20.GetAccountsResponse, error) {
	httpRequest, err := http.NewRequest("GET", r.UrlManager.GetAccounts(), nil)
	if err != nil {
		return nil, errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+r.Token)

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.New("could not read body of response: " + err.Error())
	}

	response := &V20.GetAccountsResponse{}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.New("could not unmarshal the body of the response: " + err.Error())
	}

	return response, nil
}

func (r *Requester) GetAccount(request V20.GetAccountRequest) (*V20.GetAccountResponse, error) {
	httpRequest, err := http.NewRequest("GET", r.UrlManager.GetAccount(), nil)
	if err != nil {
		return nil, errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+r.Token)

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.New("could not read body of response: " + err.Error())
	}

	response := &V20.GetAccountResponse{}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.New("could not unmarshal the body of the response: " + err.Error())
	}

	return response, nil
}

func (r *Requester) GetAccountSummary(request V20.GetAccountSummaryRequest) (*V20.GetAccountSummaryResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetAccountInstruments(request V20.GetAccountInstrumentsRequest) (*V20.GetAccountInstrumentsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PatchAccountConfiguration(request V20.PatchAccountConfigurationRequest) (*V20.PatchAccountConfigurationResponse, error) {
	return nil, errors.New("not implemented yet")
}

// instruments
func (r *Requester) GetInstrumentCandles(request V20.GetInstrumentCandlesRequest) (*V20.GetInstrumentCandlesResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetInstrumentOrderBook(request V20.GetInstrumentOrderBookRequest) (*V20.GetInstrumentOrderBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetInstrumentPositionBook(request V20.GetInstrumentPositionBookRequest) (*V20.GetInstrumentPositionBookResponse, error) {
	return nil, errors.New("not implemented yet")
}

// orders
func (r *Requester) PostOrder(request V20.PostOrderRequest) (*V20.PostOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOrders(request V20.GetOrdersRequest) (*V20.GetOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetPendingOrders(request V20.GetPendingOrdersRequest) (*V20.GetPendingOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOrder(request V20.GetOrderRequest) (*V20.GetOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutReplaceOrder(request V20.PutReplaceOrderRequest) (*V20.PutReplaceOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutCancelOrder(request V20.PutCancelOrderRequest) (*V20.PutCancelOrderResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateOrderClientExtensions(request V20.PutUpdateOrderClientExtensionsRequest) (*V20.PutUpdateOrderClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

// trades
func (r *Requester) GetTrades(request V20.GetTradesRequest) (*V20.GetTradesResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOpenTrades(request V20.GetOpenTradesRequest) (*V20.GetOpenTradesResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTrade(request V20.GetTradeRequest) (*V20.GetTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutCloseTrade(request V20.PutCloseTradeRequest) (*V20.PutCloseTradeResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateTradeClientExtensions(request V20.PutUpdateTradeClientExtensionsRequest) (*V20.PutUpdateTradeClientExtensionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutUpdateTradeDependentOrders(request V20.PutUpdateTradeDependentOrdersRequest) (*V20.PutUpdateTradeDependentOrdersResponse, error) {
	return nil, errors.New("not implemented yet")
}

// positions
func (r *Requester) GetPositions(request V20.GetPositionsRequest) (*V20.GetPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetOpenPositions(request V20.GetOpenPositionsRequest) (*V20.GetOpenPositionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetInstrumentsPosition(request V20.GetInstrumentsPositionRequest) (*V20.GetInstrumentsPositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) PutClosePosition(request V20.PutClosePositionRequest) (*V20.PutClosePositionResponse, error) {
	return nil, errors.New("not implemented yet")
}

// transactions
func (r *Requester) GetTransactions(request V20.GetTransactionsRequest) (*V20.GetTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransaction(request V20.GetTransactionRequest) (*V20.GetTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetRangeOfTransactions(request V20.GetRangeOfTransactionsRequest) (*V20.GetRangeOfTransactionsResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransactionsAfterTransaction(request V20.GetTransactionsAfterTransactionRequest) (*V20.GetTransactionsAfterTransactionResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetTransactionsStream(request V20.GetTransactionsStreamRequest) (*V20.GetTransactionsStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}

// pricing
func (r *Requester) GetInstrumentPricing(request V20.GetInstrumentPricingRequest) (*V20.GetInstrumentPricingResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (r *Requester) GetPricingStream(request V20.GetPricingStreamRequest) (*V20.GetPricingStreamResponse, error) {
	return nil, errors.New("not implemented yet")
}
