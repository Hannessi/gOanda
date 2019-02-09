package url

import (
	"github.com/hannessi/gOanda/V20/definitions"
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
func (m *Manager) GetInstrumentCandles(instrumentName definitions.InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/candles"
}
func (m *Manager) GetInstrumentOrderBook(instrumentName definitions.InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/orderBook"
}
func (m *Manager) GetInstrumentPositionBook(instrumentName definitions.InstrumentName) string {
	return m.BaseUrl + "/instruments/" + instrumentName.String() + "/positionBook"
}

// orders
func (m *Manager) PostOrder() string {
	return m.BaseUrl+"/accounts/"+m.AccountId+"/orders"
}
func (m *Manager) GetOrders() string {
	return m.BaseUrl+"/accounts/"+m.AccountId+"/orders"
}
func (m *Manager) GetPendingOrders() string {
	return m.BaseUrl+"/accounts/"+m.AccountId+"/pendingOrders"
}
func (m *Manager) GetOrder(orderSpecifier definitions.OrderSpecifier) string {
	return m.BaseUrl+"/accounts/"+m.AccountId+"/order/"+orderSpecifier.String()
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
func (m *Manager) GetTrades() string {
	return "" // todo
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

