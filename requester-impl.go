package gOanda

import (
	"errors"
	"github.com/sirupsen/logrus"
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

type getAccountsResponse struct {
	Accounts     []AccountProperties `json:"accounts"`
	ErrorCode    string              `json:"errorCode"`
	ErrorMessage string              `json:"errorMessage"`
}

// accounts
func (r *HttpRequester) GetAccounts(request GetAccountsRequest) (*GetAccountsResponse, error) {
	response := &getAccountsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccounts(), nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}
	return &GetAccountsResponse{
		Accounts: response.Accounts,
	}, nil
}

type getAccountResponse struct {
	Account      Account `json:"account"`
	ErrorCode    string  `json:"errorCode"`
	ErrorMessage string  `json:"errorMessage"`
}

func (r *HttpRequester) GetAccount(request GetAccountRequest) (*GetAccountResponse, error) {
	response := &getAccountResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccount(), nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}
	return &GetAccountResponse{
		Account: response.Account,
	}, nil
}

type getAccountSummaryResponse struct {
	Account           AccountSummary `json:"account"`
	LastTransactionID TransactionID  `json:"lastTransactionID"`
	ErrorCode         string         `json:"errorCode"`
	ErrorMessage      string         `json:"errorMessage"`
}

func (r *HttpRequester) GetAccountSummary(request GetAccountSummaryRequest) (*GetAccountSummaryResponse, error) {
	response := &getAccountSummaryResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountSummary(), nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}
	return &GetAccountSummaryResponse{
		Account:           response.Account,
		LastTransactionID: response.LastTransactionID,
	}, nil
}

type getAccountInstrumentsResponse struct {
	Instruments       []Instrument  `json:"instruments"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
	ErrorCode         string        `json:"errorCode"`
	ErrorMessage      string        `json:"errorMessage"`
}

func (r *HttpRequester) GetAccountInstruments(request GetAccountInstrumentsRequest) (*GetAccountInstrumentsResponse, error) {
	response := &getAccountInstrumentsResponse{}
	if err := HttpRequestWrapper(GET, r.UrlManager.GetAccountInstruments(), nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}
	return &GetAccountInstrumentsResponse{
		Instruments:       response.Instruments,
		LastTransactionID: response.LastTransactionID,
	}, nil
}

func (r *HttpRequester) PatchAccountConfiguration(request PatchAccountConfigurationRequest) (*PatchAccountConfigurationResponse, error) {
	return nil, errors.New("not implemented yet")
}

// instruments

type getInstrumentCandlesResponse struct {
	Instrument   InstrumentName         `json:"instrument"`
	Granularity  CandlestickGranularity `json:"granularity"`
	Candles      []Candlestick          `json:"candles"`
	ErrorCode    string                 `json:"errorCode"`
	ErrorMessage string                 `json:"errorMessage"`
}

func (r *HttpRequester) GetInstrumentCandles(request GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error) {
	response := &getInstrumentCandlesResponse{}
	requestUrl := r.UrlManager.GetInstrumentCandles(GetInstrumentCandlesRequestParameters{
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		Granularity:    request.Granularity,
		From:           request.From,
		To:             request.To,
	})
	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}
	return &GetInstrumentCandlesResponse{
		Instrument:  response.Instrument,
		Granularity: response.Granularity,
		Candles:     response.Candles,
	}, nil
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

	if err := HttpRequestWrapper(POST, requestUrl, order, &response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return response, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	return response, nil
}

func (r *HttpRequester) GetOrders(request GetOrdersRequest) (*GetOrdersResponse, error) {
	response := &GetRawOrdersResponse{}
	requestUrl := r.UrlManager.GetOrders(GetOrdersRequestParameters{
		Ids:            request.Ids,
		State:          request.State,
		InstrumentName: request.InstrumentName,
		Count:          request.Count,
		BeforeID:       request.BeforeID,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}

	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	unmarshalledResponse, err := response.Unmarshal()
	if err != nil {
		logrus.Error("Could not unmarshal order:", err.Error())
		return nil, err
	}

	return unmarshalledResponse, nil
}

func (r *HttpRequester) GetPendingOrders(request GetPendingOrdersRequest) (*GetPendingOrdersResponse, error) {
	response := &GetRawPendingOrdersResponse{}
	requestUrl := r.UrlManager.GetPendingOrders()

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}

	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	unmarshalledResponse, err := response.Unmarshal()
	if err != nil {
		logrus.Error("Could not unmarshal order:", err.Error())
		return nil, err
	}

	return unmarshalledResponse, nil
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
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return response, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	return response, nil
}

func (r *HttpRequester) GetOpenTrades(request GetOpenTradesRequest) (*GetOpenTradesResponse, error) {
	response := &GetOpenTradesResponse{}
	requestUrl := r.UrlManager.GetOpenTrades()

	if err := HttpRequestWrapper(GET, requestUrl, nil, response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return response, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	return response, nil
}

func (r *HttpRequester) GetTrade(request GetTradeRequest) (*GetTradeResponse, error) {
	response := &GetTradeResponse{}
	requestUrl := r.UrlManager.GetTrade(GetTradeRequestParameters{
		TradeSpecifier: request.TradeSpecifier,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, &response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return response, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	return response, nil
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
	response := &GetRangeOfTransactionsResponse{}
	requestUrl := r.UrlManager.GetRangeOfTransactions(GetRangeOfTransactionsParameters{
		From: request.From,
		To:   request.To,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, &response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return response, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	return response, nil
}

type getTransactionsSinceIdResponse struct {
	Transactions      []RawTransaction `json:"transactions"`
	LastTransactionID TransactionID    `json:"lastTransactionID"`
	ErrorCode         string           `json:"errorCode"`
	ErrorMessage      string           `json:"errorMessage"`
}

func (g *getTransactionsSinceIdResponse) unmarshal() (*GetTransactionsSinceIdResponse, error) {
	transactions := make([]Transaction, 0)

	for _, raw := range g.Transactions {
		transaction, err := raw.ToTransaction()
		if err != nil {
			logrus.Error("Could not unmarshal transaction: ", raw)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return &GetTransactionsSinceIdResponse{
		Transactions:      transactions,
		LastTransactionID: g.LastTransactionID,
	}, nil
}

func (r *HttpRequester) GetTransactionsSinceId(request GetTransactionsSinceIdRequest) (*GetTransactionsSinceIdResponse, error) {
	response := &getTransactionsSinceIdResponse{}
	requestUrl := r.UrlManager.GetTransactionsSinceId(GetTransactionsSinceIdParameters{
		TransactionId: request.TransactionId,
	})

	if err := HttpRequestWrapper(GET, requestUrl, nil, &response, r.Token); err != nil {
		return nil, err
	}
	if response.ErrorCode != "" || response.ErrorMessage != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.ErrorMessage)
	}

	unmarshalledResponse, err := response.unmarshal()
	if err != nil {
		logrus.Error("Could not unmarshal response:", err.Error())
		return nil, err
	}
	return unmarshalledResponse, nil

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
