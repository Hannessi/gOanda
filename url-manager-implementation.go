package gOanda

import (
	"net/url"
	"strconv"
	"strings"
)

type UrlManager struct {
	BaseUrl   string
	AccountId string
}

// accounts
func (m *UrlManager) GetAccounts() string {
	return m.BaseUrl + "/accounts"
}
func (m *UrlManager) GetAccount() string {
	return m.BaseUrl + "/accounts/" + m.AccountId
}
func (m *UrlManager) GetAccountSummary() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/summary"
}
func (m *UrlManager) GetAccountInstruments() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/instruments"
}
func (m *UrlManager) PatchAccountConfiguration() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/configuration"
}
func (m *UrlManager) GetAccountChanges(LastTransactionID string) string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/changes?sinceTransactionID=" + LastTransactionID
}

// instruments
type GetInstrumentCandlesRequestParameters struct {
	InstrumentName InstrumentName
	Count          int
	Granularity    CandlestickGranularity
}

func (m *UrlManager) GetInstrumentCandles(parameters GetInstrumentCandlesRequestParameters) string {
	url := m.BaseUrl + "/instruments/" + parameters.InstrumentName.String() + "/candles"

	additionalParameters := make([]string, 0)
	if parameters.Count > 0 {
		additionalParameters = append(additionalParameters, "count="+strconv.Itoa(parameters.Count))
	}
	if parameters.Granularity != "" {
		additionalParameters = append(additionalParameters, "granularity="+parameters.Granularity.String())
	}

	if len(additionalParameters) > 0 {
		url = url + "?" + strings.Join(additionalParameters, "&")
	}

	return url
}
func (m *UrlManager) GetInstrumentOrderBook(instrumentName InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/orderBook"
}
func (m *UrlManager) GetInstrumentPositionBook(instrumentName InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/positionBook"
}

// orders
func (m *UrlManager) PostOrder() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/orders"
}
func (m *UrlManager) GetOrders() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/orders"
}
func (m *UrlManager) GetPendingOrders() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/pendingOrders"
}
func (m *UrlManager) GetOrder(orderSpecifier OrderSpecifier) string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/order/" + orderSpecifier.String()
}
func (m *UrlManager) PutReplaceOrder() string {
	return "" // todo
}
func (m *UrlManager) PutCancelOrder() string {
	return "" // todo
}
func (m *UrlManager) PutUpdateOrderClientExtensions() string {
	return "" // todo
}

// trades
type GetTradesRequestParameters struct {
	Ids            []TradeID
	State          TradeStateFilter
	InstrumentName InstrumentName
	BeforeID       TradeID
	Count          int64
}

func (m *UrlManager) GetTrades(request GetTradesRequestParameters) string {
	url := m.BaseUrl + "/accounts/" + m.AccountId + "/trades"

	additionalParameters := make([]string, 0)
	if len(request.Ids) > 0 {
		additionalParameters = append(additionalParameters, "ids="+strings.Join(additionalParameters, ","))
	}
	if request.State != "" {
		additionalParameters = append(additionalParameters, "state="+request.State.String())
	}
	if request.InstrumentName != "" {
		additionalParameters = append(additionalParameters, "instrument="+request.InstrumentName.String())
	}
	if request.Count > 0 {
		additionalParameters = append(additionalParameters, "count="+string(request.Count))
	}
	if request.BeforeID != "" {
		additionalParameters = append(additionalParameters, "beforeID"+request.BeforeID.String())
	}

	if len(additionalParameters) > 0 {
		url = url + "?" + strings.Join(additionalParameters, "&")
	}

	return url
}
func (m *UrlManager) GetOpenTrades() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/openTrades"
}
func (m *UrlManager) GetTrade() string {
	return "" // todo
}
func (m *UrlManager) PutCloseTrade() string {
	return "" // todo
}
func (m *UrlManager) PutUpdateTradeClientExtensions() string {
	return "" // todo
}
func (m *UrlManager) PutUpdateTradeDependentOrders() string {
	return "" // todo
}

// positions
func (m *UrlManager) GetPositions() string {
	return "" // todo
}
func (m *UrlManager) GetOpenPositions() string {
	return "" // todo
}
func (m *UrlManager) GetInstrumentsPosition() string {
	return "" // todo
}
func (m *UrlManager) PutClosePosition() string {
	return "" // todo
}

// transactions
func (m *UrlManager) GetTransactions() string {
	return "" // todo
}
func (m *UrlManager) GetTransaction() string {
	return "" // todo
}
func (m *UrlManager) GetRangeOfTransactions() string {
	return "" // todo
}
func (m *UrlManager) GetTransactionsAfterTransaction() string {
	return "" // todo
}
func (m *UrlManager) GetTransactionsStream() string {
	return "" // todo
}

// pricing
type GetInstrumentPricingParameters struct {
	Instruments            []InstrumentName
	Since                  DateTime
	IncludeUnitsAvailable  bool
	IncludeHomeConversions bool
}

func (m *UrlManager) GetInstrumentPricing(params GetInstrumentPricingParameters) string {
	urlRaw := m.BaseUrl + "/accounts/" + m.AccountId + "/pricing"

	additionalParameters := make([]string, 0)
	if len(params.Instruments) > 0 {
		instrumentStrings := func(is []InstrumentName) []string {
			s := make([]string,0)
			for _, i := range is {
				s = append(s, i.String())
			}
			return s
		}(params.Instruments)
		additionalParameters = append(additionalParameters, "instruments="+url.QueryEscape(strings.Join(instrumentStrings, ",")))
	}
	if params.Since != "" {
		additionalParameters = append(additionalParameters, "since="+params.Since.String())
	}
	if params.IncludeHomeConversions {
		additionalParameters = append(additionalParameters, "includeHomeConversions=true")
	}
	if params.IncludeUnitsAvailable {
		additionalParameters = append(additionalParameters, "includeUnitsAvailable=true")
	}

	if len(additionalParameters) > 0 {
		urlRaw = urlRaw + "?" + strings.Join(additionalParameters, "&")
	}

	return urlRaw
}
func (m *UrlManager) GetPricingStream() string {
	return "" // todo
}
