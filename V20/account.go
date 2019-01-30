package V20

// TODO COMMENTS

// The string representation of an Account Identifier
type AccountID string

// The
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
	NAV                         AccountUnits                `json:"nAV"`
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
	Orders                      []Order                     `json:"orders"`
}

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

type AccountProperties struct {
	Id           []AccountID `json:"id"`
	Mt4AccountID int         `json:"mt4AccountID"`
	Tags         []string    `json:"tags"`
}

type GuaranteedStopLossOrderMode string

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

type AccountChanges struct {
	// todo
}

type AccountFinancingMode string

type PositionAggregationMode string
