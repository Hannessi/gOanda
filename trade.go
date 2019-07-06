package gOanda

type TradeID string

func (t TradeID) String() string { return string(t) }

type TradeState string // TODO consts

//The Trades that are currently open
const TRADE_STATE_OPEN TradeState = "OPEN"

//The Trades that have been fully closed
const TRADE_STATE_CLOSED TradeState = "CLOSED"

type TradeStateFilter string

//The Trades that are currently open
const TRADE_STATE_FILTER_OPEN TradeStateFilter = "OPEN"

//The Trades that have been fully closed
const TRADE_STATE_FILTER_CLOSED TradeStateFilter = "CLOSED"

//The Trades that will be closed as soon as the trades’ instrument becomes tradeable
const TRADE_STATE_FILTER_CLOSE_WHEN_TRADEABLE TradeStateFilter = "CLOSE_WHEN_TRADEABLE"

//The Trades that are in any of the possible states listed above.
const TRADE_STATE_FILTER_ALL TradeStateFilter = "ALL"

func (t TradeStateFilter) String() string { return string(t) }

type TradeSpecifier string

func (t TradeSpecifier) String() string { return string(t) }

type Trade struct {
	Id                    TradeID               `json:"id"`
	Instrument            InstrumentName        `json:"instrument"`
	Price                 PriceValue            `json:"price"`
	OpenTime              DateTime              `json:"openTime"`
	State                 TradeState            `json:"state"`
	InitialUnits          DecimalNumber         `json:"initialUnits"`
	InitialMarginRequired AccountUnits          `json:"initialMarginRequired"`
	CurrentUnits          DecimalNumber         `json:"currentUnits"`
	RealizedPL            AccountUnits          `json:"realizedPL"`
	UnrealizedPL          AccountUnits          `json:"unrealizedPL"`
	MarginUsed            AccountUnits          `json:"marginUsed"`
	AverageClosePrice     PriceValue            `json:"averageClosePrice"`
	ClosingTransactionIDs []TransactionID       `json:"closingTransactionIDs"`
	Financing             AccountUnits          `json:"financing"`
	CloseTime             DateTime              `json:"closeTime"`
	ClientExtensions      ClientExtensions      `json:"clientExtensions"`
	TakeProfitOrder       TakeProfitOrder       `json:"takeProfitOrder"`
	StopLossOrder         StopLossOrder         `json:"stopLossOrder"`
	TrailingStopLossOrder TrailingStopLossOrder `json:"trailingStopLossOrder"`
}

type TradeSummary struct{} // todo

type CalculatedTradeState struct{} // todo

type TradePL string // todo consts
