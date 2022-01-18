package gOanda

// The string representation of an Account Identifier
type AccountID string

// The full details of a client’s Account. This includes full open Trade, open Position and pending Order representation.
type Account struct {
	Id                          AccountID                   `json:"id"`
	Alias                       string                      `json:"alias"`
	Currency                    Currency                    `json:"currency"`
	Balance                     AccountUnits                `json:"balance"`
	CreatedByUserID             int                         `json:"createdByUserID"`
	CreatedTime                 DateTime                    `json:"createdTime"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	Pl                          AccountUnits                `json:"pl"`
	ResettablePL                AccountUnits                `json:"resettablePL"`
	ResettablePLTime            DateTime                    `json:"resettablePLTime"`
	Financing                   AccountUnits                `json:"financing"`
	Commission                  AccountUnits                `json:"commission"`
	GuaranteedExecutionFees     AccountUnits                `json:"guaranteedExecutionFees"`
	MarginRate                  DecimalNumber               `json:"marginRate"`
	MarginCallEnterTime         DateTime                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                         `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DateTime                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                         `json:"openTradeCount"`
	OpenPositionCount           int                         `json:"openPositionCount"`
	PendingOrderCount           int                         `json:"pendingOrderCount"`
	HedgingEnabled              bool                        `json:"hedgingEnabled"`
	UnrealizedPL                AccountUnits                `json:"unrealizedPL"`
	NAV                         AccountUnits                `json:"NAV"`
	MarginUsed                  AccountUnits                `json:"marginUsed"`
	MarginAvailable             AccountUnits                `json:"marginAvailable"`
	PositionValue               AccountUnits                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnits                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnits                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnits                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumber               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumber               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnits                `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnits                `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumber               `json:"marginCallPercent"`
	LastTransactionID           TransactionID               `json:"lastTransactionID"`
	Trades                      []TradeSummary              `json:"trades"`
	Positions                   []Position                  `json:"positions"`
	Orders                      []RawOrder                  `json:"orders"`
}

// An AccountState Object is used to represent an Account’s current price-dependent state. Price-dependent Account state is dependent on OANDA’s current Prices, and includes things like unrealized PL, NAV and Trailing Stop Loss Order state. Fields will be omitted if their value has not changed since the specified transaction ID.
type AccountChangesState struct {
	UnrealizedPL                AccountUnits              `json:"unrealizedPL"`
	Nav                         AccountUnits              `json:"NAV"`
	MarginUsed                  AccountUnits              `json:"marginUsed"`
	MarginAvailable             AccountUnits              `json:"marginAvailable"`
	PositionValue               AccountUnits              `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnits              `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnits              `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnits              `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumber             `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumber             `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnits              `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnits              `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumber             `json:"marginCallPercent"`
	Orders                      []DynamicOrderState       `json:"orders"`
	Trades                      []CalculatedTradeState    `json:"trades"`
	Positions                   []CalculatedPositionState `json:"positions"`
}

// Properties related to an Account.
type AccountProperties struct {
	Id           AccountID `json:"id"`
	Mt4AccountID int       `json:"mt4AccountID"`
	Tags         []string  `json:"tags"`
}

// The current mutability and hedging settings related to guaranteed Stop Loss orders.
type GuaranteedStopLossOrderParameters struct {
	//# The current guaranteed Stop Loss Order mutability setting of the Account when market is open.
	MutabilityMarketOpen GuaranteedStopLossOrderMutability

	//# The current guaranteed Stop Loss Order mutability setting of the Account when market is halted.
	MutabilityMarketHalted GuaranteedStopLossOrderMutability
}

// For Accounts that support guaranteed Stop Loss Orders, describes the actions that can be be performed on guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMutability string

// The overall behaviour of the Account regarding guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMode string

// A summary representation of a client’s Account. The AccountSummary does not provide to full specification of pending Orders, open Trades and Positions.
type AccountSummary struct {
	Id                          AccountID                   `json:"id"`
	Alias                       string                      `json:"alias"`
	Currency                    Currency                    `json:"currency"`
	Balance                     AccountUnits                `json:"balance"`
	CreatedByUserID             int                         `json:"createdByUserID"`
	CreatedTime                 DateTime                    `json:"createdTime"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	Pl                          AccountUnits                `json:"pl"`
	ResettablePL                AccountUnits                `json:"resettablePL"`
	ResettablePLTime            DateTime                    `json:"resettablePLTime"`
	Financing                   AccountUnits                `json:"financing"`
	Commission                  AccountUnits                `json:"commission"`
	GuaranteedExecutionFees     AccountUnits                `json:"guaranteedExecutionFees"`
	MarginRate                  DecimalNumber               `json:"marginRate"`
	MarginCallEnterTime         DateTime                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                         `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DateTime                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                         `json:"openTradeCount"`
	OpenPositionCount           int                         `json:"openPositionCount"`
	PendingOrderCount           int                         `json:"pendingOrderCount"`
	HedgingEnabled              bool                        `json:"hedgingEnabled"`
	UnrealizedPL                AccountUnits                `json:"unrealizedPL"`
	NAV                         AccountUnits                `json:"NAV"`
	MarginUsed                  AccountUnits                `json:"marginUsed"`
	MarginAvailable             AccountUnits                `json:"marginAvailable"`
	PositionValue               AccountUnits                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnits                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnits                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnits                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumber               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumber               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnits                `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnits                `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumber               `json:"marginCallPercent"`
	LastTransactionID           TransactionID               `json:"lastTransactionID"`
}

type AccumulatedAccountState struct {
	//# The current balance of the account.
	Balance AccountUnits `json:"balance"`

	//# The total profit/loss realized over the lifetime of the Account.
	Pl AccountUnits `json:"pl"`

	//# The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`

	//# The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnits `json:"financing"`

	//# The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnits `json:"commission"`

	//# The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`

	//# The total amount of fees charged over the lifetime of the Account for the execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`

	//# The date/time when the Account entered a margin call state. Only provided if the Account is in a margin call.
	MarginCallEnterTime DateTime `json:"marginCallEnterTime"`

	//# The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount"`

	//# The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTime `json:"lastMarginCallExtensionTime"`
}

// The dynamically calculated state of a client’s Account.
type CalculatedAccountState struct {
	UnrealizedPL                AccountUnits  `json:"unrealizedPL"`
	NAV                         AccountUnits  `json:"NAV"`
	MarginUsed                  AccountUnits  `json:"marginUsed"`
	MarginAvailable             AccountUnits  `json:"marginAvailable"`
	PositionValue               AccountUnits  `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnits  `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnits  `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnits  `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumber `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumber `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnits  `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnits  `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumber `json:"marginCallPercent"`
}

// An AccountChanges Object is used to represent the changes to an Account’s Orders, Trades and Positions since a specified Account TransactionID in the past.
type AccountChanges struct {
	//# The Orders created. These Orders may have been filled, cancelled or triggered in the same period.
	OrdersCreated []Order `json:"ordersCreated"`

	//# The Orders cancelled.
	OrdersCancelled []Order `json:"ordersCancelled"`

	//# The Orders filled.
	OrdersFilled []Order `json:"ordersFilled"`

	//# The Orders triggered.
	OrdersTriggered []Order `json:"ordersTriggered"`

	//# The Trades opened.
	TradesOpened []TradeSummary `json:"tradesOpened"`

	//# The Trades reduced.
	TradesReduced []TradeSummary `json:"tradesReduced"`

	//# The Trades closed.
	TradesClosed []TradeSummary `json:"tradesClosed"`

	//# The Positions changed.
	Positions []Position `json:"positions"`

	//# The Transactions that have been generated.
	Transactions []Transaction `json:"transactions"`
}

// The financing mode of an Account
type AccountFinancingMode string

// No financing is paid/charged for open Trades in the Account
const ACCOUNT_FINANCING_MODE_NO_FINANCING AccountFinancingMode = "NO_FINANCING"

// Second-by-second financing is paid/charged for open Trades in the Account, both daily and when the the Trade is closed
const ACCOUNT_FINANCING_MODE_SECOND_BY_SECOND AccountFinancingMode = "SECOND_BY_SECOND"

// A full day’s worth of financing is paid/charged for open Trades in the Account daily at 5pm New York time
const ACCOUNT_FINANCING_MODE_DAILY AccountFinancingMode = "DAILY"

type UserAttributes struct {
	//# The user’s OANDA-assigned user ID.
	UserID int64 `json:"userID"`

	//# The user-provided username.
	Username string `json:"username"`

	//# The user’s title.
	Title string `json:"title"`

	//# The user’s name.
	Name string `json:"name"`

	//# The user’s email address.
	Email string `json:"email"`

	//# The OANDA division the user belongs to.
	DivisionAbbreviation string `json:"divisionAbbreviation"`

	//# The user’s preferred language.
	LanguageAbbreviation string `json:"languageAbbreviation"`

	//# The home currency of the Account.
	HomeCurrency Currency `json:"homeCurrency"`
}

// The way that position values for an Account are calculated and aggregated.
type PositionAggregationMode string

// The Position value or margin for each side (long and short) of the Position are computed independently and added together.
const POSITION_AGGREGATION_MODE_ABSOLUTE_SUM PositionAggregationMode = "ABSOLUTE_SUM"

// The Position value or margin for each side (long and short) of the Position are computed independently. The Position value or margin chosen is the maximal absolute value of the two.
const POSITION_AGGREGATION_MODE_MAXIMAL_SIDE PositionAggregationMode = "MAXIMAL_SIDE"

// The units for each side (long and short) of the Position are netted together and the resulting value (long or short) is used to compute the Position value or margin.
const POSITION_AGGREGATION_MODE_NET_SUM PositionAggregationMode = "NET_SUM"
