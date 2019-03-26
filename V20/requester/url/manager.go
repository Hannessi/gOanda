package url

import (
	"github.com/hannessi/gOanda"
	"strconv"
	"strings"
)

type Manager struct {
	BaseUrl   string
	AccountId string
}

// accounts
func (m *Manager) GetAccounts() string {
	return m.BaseUrl + "/accounts"
}
func (m *Manager) GetAccount() string {
	return m.BaseUrl + "/accounts/" + m.AccountId
}
func (m *Manager) GetAccountSummary() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/summary"
}
func (m *Manager) GetAccountInstruments() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/instruments"
}
func (m *Manager) PatchAccountConfiguration() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/configuration"
}
func (m *Manager) GetAccountChanges(LastTransactionID string) string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/changes?sinceTransactionID=" + LastTransactionID
}

// instruments
type GetInstrumentCandlesRequestParameters struct {
	InstrumentName gOanda.InstrumentName
	Count          int
	Granularity    gOanda.CandlestickGranularity
}

func (m *Manager) GetInstrumentCandles(parameters GetInstrumentCandlesRequestParameters) string {
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
func (m *Manager) GetInstrumentOrderBook(instrumentName gOanda.InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/orderBook"
}
func (m *Manager) GetInstrumentPositionBook(instrumentName gOanda.InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/positionBook"
}

// orders
func (m *Manager) PostOrder() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/orders"
}
func (m *Manager) GetOrders() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/orders"
}
func (m *Manager) GetPendingOrders() string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/pendingOrders"
}
func (m *Manager) GetOrder(orderSpecifier gOanda.OrderSpecifier) string {
	return m.BaseUrl + "/accounts/" + m.AccountId + "/order/" + orderSpecifier.String()
}
func (m *Manager) PutReplaceOrder() string {
	return "" // todo
}
func (m *Manager) PutCancelOrder() string {
	return "" // todo
}
func (m *Manager) PutUpdateOrderClientExtensions() string {
	return "" // todo
}

// trades
type GetTradesRequestParameters struct {
	Ids []gOanda.TradeID
	State gOanda.TradeStateFilter
	InstrumentName gOanda.InstrumentName
	BeforeID gOanda.TradeID
	Count int64
}

func (m *Manager) GetTrades(request GetTradesRequestParameters) string {
	url := m.BaseUrl + "/accounts/" + m.AccountId + "/trades"

	additionalParameters := make([]string,0)
	if len(request.Ids) > 0 {
		additionalParameters = append(additionalParameters, "ids="+strings.Join(additionalParameters,","))
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
func (m *Manager) GetOpenTrades() string {
	return "" // todo
}
func (m *Manager) GetTrade() string {
	return "" // todo
}
func (m *Manager) PutCloseTrade() string {
	return "" // todo
}
func (m *Manager) PutUpdateTradeClientExtensions() string {
	return "" // todo
}
func (m *Manager) PutUpdateTradeDependentOrders() string {
	return "" // todo
}

// positions
func (m *Manager) GetPositions() string {
	return "" // todo
}
func (m *Manager) GetOpenPositions() string {
	return "" // todo
}
func (m *Manager) GetInstrumentsPosition() string {
	return "" // todo
}
func (m *Manager) PutClosePosition() string {
	return "" // todo
}

// transactions
func (m *Manager) GetTransactions() string {
	return "" // todo
}
func (m *Manager) GetTransaction() string {
	return "" // todo
}
func (m *Manager) GetRangeOfTransactions() string {
	return "" // todo
}
func (m *Manager) GetTransactionsAfterTransaction() string {
	return "" // todo
}
func (m *Manager) GetTransactionsStream() string {
	return "" // todo
}

// pricing
func (m *Manager) GetInstrumentPricing() string {
	return "" // todo
}
func (m *Manager) GetPricingStream() string {
	return "" // todo
}
