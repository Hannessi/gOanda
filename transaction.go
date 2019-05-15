package gOanda

type Transaction struct{} //todo

type CreateTransaction struct{} //todo

type CloseTransaction struct{} //todo

type ReopenTransaction struct{} //todo

type ClientConfigureTransaction struct{} //todo

type ClientConfigureRejectTransaction struct{} //todo

type TransferFundsTransaction struct{} //todo

type TransferFundsRejectTransaction struct{} //todo

type MarketOrderTransaction struct {
	Id                     TransactionID                `json:"id"`
	Time                   DateTime                     `json:"time"`
	UserID                 int                          `json:"userID"`
	AccountID              AccountID                    `json:"accountID"`
	BatchId                TransactionID                `json:"batchID"`
	RequestID              RequestID                    `json:"requestID"`
	Type                   TransactionType              `json:"type"`
	Instrument             InstrumentName               `json:"instrument"`
	Units                  DecimalNumber                `json:"units"`
	TimeInForce            TimeInForce                  `json:"timeInForce"`
	PriceBound             PriceValue                   `json:"priceBound"`
	PositionFill           OrderPositionFill            `json:"positionFill"`
	TradeClose             MarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout   MarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout  MarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout         MarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose      MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Reason                 MarketOrderReason            `json:"reason"`
	ClientExtensions       ClientExtensions             `json:"clientExtension"`
	TakeProfitOnFill       TakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensions             `json:"tradeClientExtensions"`
}

type MarketOrderRejectTransaction struct {
	Id                     TransactionID                `json:"id"`
	Time                   DateTime                     `json:"time"`
	UserID                 int                          `json:"userID"`
	AccountID              AccountID                    `json:"accountID"`
	BatchId                TransactionID                `json:"batchID"`
	RequestID              RequestID                    `json:"requestID"`
	Type                   TransactionType              `json:"type"`
	Instrument             InstrumentName               `json:"instrument"`
	Units                  DecimalNumber                `json:"units"`
	TimeInForce            TimeInForce                  `json:"timeInForce"`
	PriceBound             PriceValue                   `json:"priceBound"`
	PositionFill           OrderPositionFill            `json:"positionFill"`
	TradeClose             MarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout   MarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout  MarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout         MarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose      MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Reason                 MarketOrderReason            `json:"reason"`
	ClientExtensions       ClientExtensions             `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensions             `json:"tradeClientExtensions"`
	RejectReason           TransactionRejectReason      `json:"rejectReason"`
}

type FixedPriceOrderTransaction struct{} //todo

type LimitOrderTransaction struct{} //todo

type LimitOrderRejectTransaction struct{} //todo

type StopOrderTransaction struct{} //todo

type StopOrderRejectTransaction struct{} //todo

type MarketIfTouchedOrderTransaction struct{} //todo

type MarketIfTouchedOrderRejectTransaction struct{} //todo

type TakeProfitOrderTransaction struct{} //todo

type TakeProfitOrderRejectTransaction struct{} //todo

type StopLossOrderTransaction struct{} //todo

type StopLossOrderRejectTransaction struct{} //todo

type TrailingStopLossOrderTransaction struct{} //todo

type TrailingStopLossOrderRejectTransaction struct{} //todo

type OrderFillTransaction struct {
	Id                            TransactionID   `json:"id"`
	Time                          DateTime        `json:"time"`
	UserID                        int             `json:"userID"`
	AccountID                     AccountID       `json:"accountID"`
	BatchID                       TransactionID   `json:"batchID"`
	RequestID                     RequestID       `json:"requestID"`
	Type                          TransactionType `json:"type"`
	OrderID                       OrderID         `json:"orderID"`
	ClientOrderID                 ClientID        `json:"clientOrderID"`
	Instrument                    InstrumentName  `json:"instrument"`
	Units                         DecimalNumber   `json:"units"`
	GainQuoteHomeConversionFactor DecimalNumber   `json:"gainQuoteHomeConversionFactor"`
	LossQuoteHomeConversionFactor DecimalNumber   `json:"lossQuoteHomeConversionFactor"`
	Price                         PriceValue      `json:"price"` // fixme deprecated
	FullPrice                     ClientPrice     `json:"fullPrice"`
	Reason                        OrderFillReason `json:"reason"`
	Pl                            AccountUnits    `json:"pl"`
	Financing                     AccountUnits    `json:"financing"`
	Commission                    AccountUnits    `json:"commission"`
	GuaranteedExecutionFee        AccountUnits    `json:"guaranteedExecutionFee"`
	AccountBalance                AccountUnits    `json:"accountBalance"`
	TradeOpened                   TradeOpen       `json:"tradeOpened"`
	TradesClosed                  []TradeReduce   `json:"tradesClosed"`
	TradeReduced                  TradeReduce     `json:"tradeReduced"`
	HalfSpreadCost                AccountUnits    `json:"halfSpreadCost"`
}

type OrderCancelTransaction struct {
	Id                TransactionID     `json:"id"`
	Time              DateTime          `json:"time"`
	UserID            int               `json:"userID"`
	AccountID         AccountID         `json:"accountID"`
	BatchId           TransactionID     `json:"batchID"`
	RequestID         RequestID         `json:"requestID"`
	Type              TransactionType   `json:"type"`
	OrderID           OrderID           `json:"orderID"`
	ClientOrderID     OrderID           `json:"clientOrderID"`
	Reason            OrderCancelReason `json:"reason"`
	ReplacedByOrderID OrderID           `json:"replacedByOrderID"`
}

type OrderCancelRejectTransaction struct{} //todo

type OrderClientExtensionsModifyTransaction struct{} //todo

type OrderClientExtensionsModifyRejectTransaction struct{} //todo

type TradeClientExtensionsModifyTransaction struct{} //todo

type TradeClientExtensionsModifyRejectTransaction struct{} //todo

type MarginCallEnterTransaction struct{} //todo

type MarginCallExtendTransaction struct{} //todo

type MarginCallExitTransaction struct{} //todo

type DelayedTradeClosureTransaction struct{} //todo

type DailyFinancingTransaction struct{} //todo

type ResetResettablePLTransaction struct{} //todo

type TransactionID string

func (t TransactionID) String() string { return string(t) }

type TransactionType string // todo

type FundingReason string // todo

type MarketOrderReason string // todo consts

type FixedPriceOrderReason string // todo const

type LimitOrderReason string // todo consts

type StopOrderReason string // todo consts

type MarketIfTouchedOrderReason string // todo consts

type TakeProfitOrderReason string // todo const

type StopLossOrderReason string // todo const

type TrailingStopLossOrderReason string // todo const

type OrderFillReason string // todo const

type OrderCancelReason string // todo const

type ClientID string

type ClientTag string

type ClientComment string

type ClientExtensions struct{} // TODO

type TakeProfitDetails struct {
	//
	// The price that the Take Profit Order will be triggered at. Only one of
	// the price and distance fields may be specified.
	//
	Price PriceValue `json:"price"`

	//
	// The time in force for the created Take Profit Order. This may only be
	// GTC, GTD or GFD.
	//
	TimeInForce TimeInForce `json:"timeInForce"`

	//
	// The date when the Take Profit Order will be cancelled on if timeInForce
	// is GTD.
	//
	GtdTime DateTime `json:"gtdTime"`

	//
	// The Client Extensions to add to the Take Profit Order when created.
	//
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type StopLossDetails struct {
	//
	// The price that the Stop Loss Order will be triggered at. Only one of the
	// price and distance fields may be specified.
	//
	Price PriceValue `json:"price"`

	//
	// Specifies the distance (in price units) from the Trade’s open price to
	// use as the Stop Loss Order price. Only one of the distance and price
	// fields may be specified.
	//
	Distance DecimalNumber `json:"distance"`

	//
	// The time in force for the created Stop Loss Order. This may only be GTC,
	// GTD or GFD.
	//
	TimeInForce TimeInForce `json:"timeInForce"`

	//
	// The date when the Stop Loss Order will be cancelled on if timeInForce is
	// GTD.
	//
	GtdTime DateTime `json:"gtdTime"`

	//
	// The Client Extensions to add to the Stop Loss Order when created.
	//
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	//
	// Flag indicating that the price for the Stop Loss Order is guaranteed. The
	// default value depends on the GuaranteedStopLossOrderMode of the account,
	// if it is REQUIRED, the default will be true, for DISABLED or ENABLED the
	// default is false.
	//
	Guaranteed bool `json:"guaranteed"`
}

type TrailingStopLossDetails struct {
	//
	// The distance (in price units) from the Trade’s fill price that the
	// Trailing Stop Loss Order will be triggered at.
	//
	Distance DecimalNumber `json:"distance"`

	//
	// The time in force for the created Trailing Stop Loss Order. This may only
	// be GTC, GTD or GFD.
	//
	TimeInForce TimeInForce `json:"timeInForce"`

	//
	// The date when the Trailing Stop Loss Order will be cancelled on if
	// timeInForce is GTD.
	//
	GtdTime DateTime `json:"gtdTime"`

	//
	// The Client Extensions to add to the Trailing Stop Loss Order when
	// created.
	//
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type TradeOpen struct {
	TradeID                TradeID          `json:"tradeID"`
	Units                  DecimalNumber    `json:"units"`
	Price                  PriceValue       `json:"price"`
	GuaranteedExecutionFee AccountUnits     `json:"guaranteedExecutionFee"`
	ClientExtensions       ClientExtensions `json:"clientExtensions"`
	HalfSpreadCost         AccountUnits     `json:"halfSpreadCost"`
	InitialMarginRequired  AccountUnits     `json:"initialMarginRequired"`
}

type TradeReduce struct {
	TradeID                TradeID       `json:"tradeID"`
	Units                  DecimalNumber `json:"units"`
	Price                  PriceValue    `json:"price"`
	RealizedPL             AccountUnits  `json:"realizedPL"`
	Financing              AccountUnits  `json:"financing"`
	GuaranteedExecutionFee AccountUnits  `json:"guaranteedExecutionFee"`
	HalfSpreadCost         AccountUnits  `json:"halfSpreadCost"`
}

type MarketOrderTradeClose struct{} // TODO

type MarketOrderMarginCloseout struct{} // TODO

type MarketOrderMarginCloseoutReason string // todo const

type MarketOrderDelayedTradeClose struct{} // TODO

type MarketOrderPositionCloseout struct {
	Instrument InstrumentName `json:"instrument"`
	Units      string         `json:"units"`
}

type LiquidityRegenerationSchedule struct{} // TODO

type LiquidityRegenerationScheduleStep struct{} // TODO

type OpenTradeFinancing struct{} // TODO

type PositionFinancing struct{} // TODO

type RequestID string

type TransactionRejectReason string // todo const

type TransactionFilter string // TODO

type TransactionHeartbeat struct{} // TODO
