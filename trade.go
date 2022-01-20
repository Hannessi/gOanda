package gOanda

type TradeID string

func (t TradeID) String() string { return string(t) }

type TradeState string

//The Trades that are currently open
const TRADE_STATE_OPEN TradeState = "OPEN"

//The Trades that have been fully closed
const TRADE_STATE_CLOSED TradeState = "CLOSED"

//The Trade will be closed as soon as the trade’s instrument becomes tradeable
const TRADE_STATE_CLOSE_WHEN_TRADEABLE TradeState = "CLOSE_WHEN_TRADEABLE"

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

type TradeSummary struct {
	// The Trade’s identifier, unique within the Trade’s Account.
	Id TradeID `json:"id"`

	// The Trade’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The execution price of the Trade.
	Price PriceValue `json:"price"`

	// The date/time when the Trade was opened.
	OpenTime DateTime `json:"openTime"`

	// The current state of the Trade.
	State TradeState `json:"state"`

	// The initial size of the Trade. Negative values indicate a short Trade,
	// and positive values indicate a long Trade.
	InitialUnits DecimalNumber `json:"initialUnits"`

	// The margin required at the time the Trade was created. Note, this is the
	// ‘pure’ margin required, it is not the ‘effective’ margin used that
	// factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired AccountUnits `json:"initialMarginRequired"`

	// The number of units currently open for the Trade. This value is reduced
	// to 0.0 as the Trade is closed.
	CurrentUnits DecimalNumber `json:"currentUnits"`

	// The total profit/loss realized on the closed portion of the Trade.
	RealizedPL AccountUnits `json:"realizedPL"`

	// The unrealized profit/loss on the open portion of the Trade.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`

	// Margin currently used by the Trade.
	MarginUsed AccountUnits `json:"marginUsed"`

	// The average closing price of the Trade. Only present if the Trade has
	// been closed or reduced at least once.
	AverageClosePrice PriceValue `json:"averageClosePrice"`

	// The IDs of the Transactions that have closed portions of this Trade.
	ClosingTransactionIDs []TransactionID `json:"closingTransactionIDs"`

	// The financing paid/collected for this Trade.
	Financing AccountUnits `json:"financing"`

	// The dividend adjustment paid for this Trade.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`

	// The date/time when the Trade was fully closed. Only provided for Trades
	// whose state is CLOSED.
	CloseTime DateTime `json:"closeTime"`

	// The client extensions of the Trade.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// ID of the Trade’s Take Profit Order, only provided if such an Order
	// exists.
	TakeProfitOrderID OrderID `json:"takeProfitOrderID"`

	// ID of the Trade’s Stop Loss Order, only provided if such an Order exists.
	StopLossOrderID OrderID `json:"stopLossOrderID"`

	// ID of the Trade’s Guaranteed Stop Loss Order, only provided if such an
	// Order exists.
	GuaranteedStopLossOrderID OrderID `json:"guaranteedStopLossOrderID"`

	// ID of the Trade’s Trailing Stop Loss Order, only provided if such an
	// Order exists.
	TrailingStopLossOrderID OrderID `json:"trailingStopLossOrderID"`
}

type CalculatedTradeState struct {
	// The Trade’s ID.
	Id TradeID `json:"id"`

	// The Trade’s unrealized profit/loss.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`

	// Margin currently used by the Trade.
	MarginUsed AccountUnits `json:"marginUsed"`
}

type TradePL string

//An open Trade currently has a positive (profitable) unrealized P/L, or a closed Trade realized a positive amount of P/L.
const TRADE_PL_POSITIVE TradePL = "POSITIVE"

//An open Trade currently has a negative (losing) unrealized P/L, or a closed Trade realized a negative amount of P/L.
const TRADE_PL_NEGATIVE TradePL = "NEGATIVE"

//An open Trade currently has unrealized P/L of zero (neither profitable nor losing), or a closed Trade realized a P/L amount of zero.
const TRADE_PL_ZERO TradePL = "ZERO"
