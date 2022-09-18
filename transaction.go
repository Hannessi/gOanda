package gOanda

type Transaction struct{} //todo

type CreateTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “CREATE” in a
	// CreateTransaction.
	Type TransactionType

	// The ID of the Division that the Account is in
	divisionID int64

	// The ID of the Site that the Account was created at
	siteID int64

	// The ID of the user that the Account was created for
	accountUserID int64

	// The number of the Account within the site/division/user
	accountNumber int64

	// The home currency of the Account
	homeCurrency Currency
}

type CloseTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “CLOSE” in a CloseTransaction.
	Type TransactionType
}

type ReopenTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “REOPEN” in a
	// ReopenTransaction.
	Type TransactionType
}

type ClientConfigureTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “CLIENT_CONFIGURE” in a
	// ClientConfigureTransaction.
	Type TransactionType

	// The client-provided alias for the Account.
	alias string

	// The margin rate override for the Account.
	marginRate DecimalNumber
}

type ClientConfigureRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “CLIENT_CONFIGURE_REJECT” in a
	// ClientConfigureRejectTransaction.
	Type TransactionType

	// The client-provided alias for the Account.
	alias string

	// The margin rate override for the Account.
	marginRate DecimalNumber

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type TransferFundsTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “TRANSFER_FUNDS” in a
	// TransferFundsTransaction.
	Type TransactionType

	// The amount to deposit/withdraw from the Account in the Account’s home
	// currency. A positive value indicates a deposit, a negative value
	// indicates a withdrawal.
	amount AccountUnits

	// The reason that an Account is being funded.
	fundingReason FundingReason

	// An optional comment that may be attached to a fund transfer for audit
	// purposes
	comment string

	// The Account’s balance after funds are transferred.
	accountBalance AccountUnits
}

type TransferFundsRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “TRANSFER_FUNDS_REJECT” in a
	// TransferFundsRejectTransaction.
	Type TransactionType

	// The amount to deposit/withdraw from the Account in the Account’s home
	// currency. A positive value indicates a deposit, a negative value
	// indicates a withdrawal.
	amount AccountUnits

	// The reason that an Account is being funded.
	fundingReason FundingReason

	// An optional comment that may be attached to a fund transfer for audit
	// purposes
	comment string

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

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

type FixedPriceOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “FIXED_PRICE_ORDER” in a
	// FixedPriceOrderTransaction.
	Type TransactionType

	// The Fixed Price Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the Fixed Price Order. A positive
	// number of units results in a long Order, and a negative number of units
	// results in a short Order.
	units DecimalNumber

	// The price specified for the Fixed Price Order. This price is the exact
	// price that the Fixed Price Order will be filled at.
	price PriceValue

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// The state that the trade resulting from the Fixed Price Order should be
	// set to.
	tradeState string

	// The reason that the Fixed Price Order was created
	reason FixedPriceOrderReason

	// The client extensions for the Fixed Price Order.
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions
}

type LimitOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “LIMIT_ORDER” in a
	// LimitOrderTransaction.
	Type TransactionType

	// The Limit Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the Limit Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	units DecimalNumber

	// The price threshold specified for the Limit Order. The Limit Order will
	// only be filled by a market price that is equal to or better than this
	// price.
	price PriceValue

	// The time-in-force requested for the Limit Order.
	timeInForce TimeInForce

	// The date/time when the Limit Order will be cancelled if its timeInForce
	// is “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Limit Order was initiated
	reason LimitOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type LimitOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “LIMIT_ORDER_REJECT” in a
	// LimitOrderRejectTransaction.
	Type TransactionType

	// The Limit Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the Limit Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	units DecimalNumber

	// The price threshold specified for the Limit Order. The Limit Order will
	// only be filled by a market price that is equal to or better than this
	// price.
	price PriceValue

	// The time-in-force requested for the Limit Order.
	timeInForce TimeInForce

	// The date/time when the Limit Order will be cancelled if its timeInForce
	// is “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Limit Order was initiated
	reason LimitOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type StopOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “STOP_ORDER” in a
	// StopOrderTransaction.
	Type TransactionType

	// The Stop Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the Stop Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	units DecimalNumber

	// The price threshold specified for the Stop Order. The Stop Order will
	// only be filled by a market price that is equal to or worse than this
	// price.
	price PriceValue

	// The worst market price that may be used to fill this Stop Order. If the
	// market gaps and crosses through both the price and the priceBound, the
	// Stop Order will be cancelled instead of being filled.
	priceBound PriceValue

	// The time-in-force requested for the Stop Order.
	timeInForce TimeInForce

	// The date/time when the Stop Order will be cancelled if its timeInForce is
	// “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Stop Order was initiated
	reason StopOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type StopOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “STOP_ORDER_REJECT” in a
	// StopOrderRejectTransaction.
	Type TransactionType

	// The Stop Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the Stop Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	units DecimalNumber

	// The price threshold specified for the Stop Order. The Stop Order will
	// only be filled by a market price that is equal to or worse than this
	// price.
	price PriceValue

	// The worst market price that may be used to fill this Stop Order. If the
	// market gaps and crosses through both the price and the priceBound, the
	// Stop Order will be cancelled instead of being filled.
	priceBound PriceValue

	// The time-in-force requested for the Stop Order.
	timeInForce TimeInForce

	// The date/time when the Stop Order will be cancelled if its timeInForce is
	// “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Stop Order was initiated
	reason StopOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type MarketIfTouchedOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “MARKET_IF_TOUCHED_ORDER” in a
	// MarketIfTouchedOrderTransaction.
	Type TransactionType

	// The MarketIfTouched Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the MarketIfTouched Order. A
	// positive number of units results in a long Order, and a negative number
	// of units results in a short Order.
	units DecimalNumber

	// The price threshold specified for the MarketIfTouched Order. The
	// MarketIfTouched Order will only be filled by a market price that crosses
	// this price from the direction of the market price at the time when the
	// Order was created (the initialMarketPrice). Depending on the value of the
	// Order’s price and initialMarketPrice, the MarketIfTouchedOrder will
	// behave like a Limit or a Stop Order.
	price PriceValue

	// The worst market price that may be used to fill this MarketIfTouched
	// Order.
	priceBound PriceValue

	// The time-in-force requested for the MarketIfTouched Order. Restricted to
	// “GTC”, “GFD” and “GTD” for MarketIfTouched Orders.
	timeInForce TimeInForce

	// The date/time when the MarketIfTouched Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Market-if-touched Order was initiated
	reason MarketIfTouchedOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type MarketIfTouchedOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “MARKET_IF_TOUCHED_ORDER_REJECT” in a
	// MarketIfTouchedOrderRejectTransaction.
	Type TransactionType

	// The MarketIfTouched Order’s Instrument.
	instrument InstrumentName

	// The quantity requested to be filled by the MarketIfTouched Order. A
	// positive number of units results in a long Order, and a negative number
	// of units results in a short Order.
	units DecimalNumber

	// The price threshold specified for the MarketIfTouched Order. The
	// MarketIfTouched Order will only be filled by a market price that crosses
	// this price from the direction of the market price at the time when the
	// Order was created (the initialMarketPrice). Depending on the value of the
	// Order’s price and initialMarketPrice, the MarketIfTouchedOrder will
	// behave like a Limit or a Stop Order.
	price PriceValue

	// The worst market price that may be used to fill this MarketIfTouched
	// Order.
	priceBound PriceValue

	// The time-in-force requested for the MarketIfTouched Order. Restricted to
	// “GTC”, “GFD” and “GTD” for MarketIfTouched Orders.
	timeInForce TimeInForce

	// The date/time when the MarketIfTouched Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	positionFill OrderPositionFill

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Market-if-touched Order was initiated
	reason MarketIfTouchedOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The specification of the Take Profit Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	takeProfitOnFill TakeProfitDetails

	// The specification of the Stop Loss Order that should be created for a
	// Trade opened when the Order is filled (if such a Trade is created).
	stopLossOnFill StopLossDetails

	// The specification of the Trailing Stop Loss Order that should be created
	// for a Trade that is opened when the Order is filled (if such a Trade is
	// created).
	trailingStopLossOnFill TrailingStopLossDetails

	// The specification of the Guaranteed Stop Loss Order that should be
	// created for a Trade that is opened when the Order is filled (if such a
	// Trade is created).
	guaranteedStopLossOnFill GuaranteedStopLossDetails

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created).  Do not set, modify, delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions ClientExtensions

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type TakeProfitOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “TAKE_PROFIT_ORDER” in a
	// TakeProfitOrderTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the TakeProfit Order. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	price PriceValue

	// The time-in-force requested for the TakeProfit Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TakeProfit Orders.
	timeInForce TimeInForce

	// The date/time when the TakeProfit Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Take Profit Order was initiated
	reason TakeProfitOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type TakeProfitOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “TAKE_PROFIT_ORDER_REJECT” in
	// a TakeProfitOrderRejectTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the TakeProfit Order. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	price PriceValue

	// The time-in-force requested for the TakeProfit Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TakeProfit Orders.
	timeInForce TimeInForce

	// The date/time when the TakeProfit Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Take Profit Order was initiated
	reason TakeProfitOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type StopLossOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “STOP_LOSS_ORDER” in a
	// StopLossOrderTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the Stop Loss Order. The associated
	// Trade will be closed by a market price that is equal to or worse than
	// this threshold.
	price PriceValue

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	distance DecimalNumber

	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for StopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// Flag indicating that the Stop Loss Order is guaranteed. The default value
	// depends on the GuaranteedStopLossOrderMode of the account, if it is
	// REQUIRED, the default will be true, for DISABLED or ENABLED the default
	// is false.
	// Deprecated: Will be removed in a future API update.
	guaranteed bool

	// The fee that will be charged if the Stop Loss Order is guaranteed and the
	// Order is filled at the guaranteed price. The value is determined at Order
	// creation time. It is in price units and is charged for each unit of the
	// Trade.
	// Deprecated: Will be removed in a future API update.
	guaranteedExecutionPremium DecimalNumber

	// The reason that the Stop Loss Order was initiated
	reason StopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type StopLossOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “STOP_LOSS_ORDER_REJECT” in a
	// StopLossOrderRejectTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the Stop Loss Order. The associated
	// Trade will be closed by a market price that is equal to or worse than
	// this threshold.
	price PriceValue

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	distance DecimalNumber

	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for StopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// Flag indicating that the Stop Loss Order is guaranteed. The default value
	// depends on the GuaranteedStopLossOrderMode of the account, if it is
	// REQUIRED, the default will be true, for DISABLED or ENABLED the default
	// is false.
	// Deprecated: Will be removed in a future API update.
	guaranteed bool

	// The reason that the Stop Loss Order was initiated
	reason StopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type GuaranteedStopLossOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “GUARANTEED_STOP_LOSS_ORDER”
	// in a GuaranteedStopLossOrderTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the Guaranteed Stop Loss Order. The
	// associated Trade will be closed at this price.
	price PriceValue

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Guaranteed Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	distance DecimalNumber

	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted
	// to “GTC”, “GFD” and “GTD” for GuaranteedStopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the GuaranteedStopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The fee that will be charged if the Guaranteed Stop Loss Order is filled
	// at the guaranteed price. The value is determined at Order creation time.
	// It is in price units and is charged for each unit of the Trade.
	guaranteedExecutionPremium DecimalNumber

	// The reason that the Guaranteed Stop Loss Order was initiated
	reason GuaranteedStopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type GuaranteedStopLossOrderReason string

type GuaranteedStopLossOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “GUARANTEED_STOP_LOSS_ORDER_REJECT” in a
	// GuaranteedStopLossOrderRejectTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price threshold specified for the Guaranteed Stop Loss Order. The
	// associated Trade will be closed at this price.
	price PriceValue

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Guaranteed Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	distance DecimalNumber

	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted
	// to “GTC”, “GFD” and “GTD” for GuaranteedStopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the GuaranteedStopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Guaranteed Stop Loss Order was initiated
	reason GuaranteedStopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type TrailingStopLossOrderTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “TRAILING_STOP_LOSS_ORDER” in
	// a TrailingStopLossOrderTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price distance (in price units) specified for the TrailingStopLoss
	// Order.
	distance DecimalNumber

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TrailingStopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Trailing Stop Loss Order was initiated
	reason TrailingStopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order replaces (only provided if this Order
	// replaces an existing Order).
	replacesOrderID OrderID

	// The ID of the Transaction that cancels the replaced Order (only provided
	// if this Order replaces an existing Order).
	cancellingTransactionID TransactionID
}

type TrailingStopLossOrderRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “TRAILING_STOP_LOSS_ORDER_REJECT” in a
	// TrailingStopLossOrderRejectTransaction.
	Type TransactionType

	// The ID of the Trade to close when the price threshold is breached.
	tradeID TradeID

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	clientTradeID ClientID

	// The price distance (in price units) specified for the TrailingStopLoss
	// Order.
	distance DecimalNumber

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TrailingStopLoss Orders.
	timeInForce TimeInForce

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	gtdTime DateTime

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a Guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	triggerCondition OrderTriggerCondition

	// The reason that the Trailing Stop Loss Order was initiated
	reason TrailingStopLossOrderReason

	// Client Extensions to add to the Order (only provided if the Order is
	// being created with client extensions).
	clientExtensions ClientExtensions

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	orderFillTransactionID TransactionID

	// The ID of the Order that this Order was intended to replace (only
	// provided if this Order was intended to replace an existing Order).
	intendedReplacesOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

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

type OrderCancelRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “ORDER_CANCEL_REJECT” for an
	// OrderCancelRejectTransaction.
	Type TransactionType

	// The ID of the Order intended to be cancelled
	orderID OrderID

	// The client ID of the Order intended to be cancelled (only provided if the
	// Order has a client Order ID).
	clientOrderID OrderID

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type OrderClientExtensionsModifyTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “ORDER_CLIENT_EXTENSIONS_MODIFY” for a
	// OrderClientExtensionsModifyTransaction.
	Type TransactionType

	// The ID of the Order who’s client extensions are to be modified.
	orderID OrderID

	// The original Client ID of the Order who’s client extensions are to be
	// modified.
	clientOrderID ClientID

	// The new Client Extensions for the Order.
	clientExtensionsModify ClientExtensions

	// The new Client Extensions for the Order’s Trade on fill.
	tradeClientExtensionsModify ClientExtensions
}

type OrderClientExtensionsModifyRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT” for a
	// OrderClientExtensionsModifyRejectTransaction.
	Type TransactionType

	// The ID of the Order who’s client extensions are to be modified.
	orderID OrderID

	// The original Client ID of the Order who’s client extensions are to be
	// modified.
	clientOrderID ClientID

	// The new Client Extensions for the Order.
	clientExtensionsModify ClientExtensions

	// The new Client Extensions for the Order’s Trade on fill.
	tradeClientExtensionsModify ClientExtensions

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type TradeClientExtensionsModifyTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “TRADE_CLIENT_EXTENSIONS_MODIFY” for a
	// TradeClientExtensionsModifyTransaction.
	Type TransactionType

	// The ID of the Trade who’s client extensions are to be modified.
	tradeID TradeID

	// The original Client ID of the Trade who’s client extensions are to be
	// modified.
	clientTradeID ClientID

	// The new Client Extensions for the Trade.
	tradeClientExtensionsModify ClientExtensions
}

type TradeClientExtensionsModifyRejectTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to
	// “TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT” for a
	// TradeClientExtensionsModifyRejectTransaction.
	Type TransactionType

	// The ID of the Trade who’s client extensions are to be modified.
	tradeID TradeID

	// The original Client ID of the Trade who’s client extensions are to be
	// modified.
	clientTradeID ClientID

	// The new Client Extensions for the Trade.
	tradeClientExtensionsModify ClientExtensions

	// The reason that the Reject Transaction was created
	rejectReason TransactionRejectReason
}

type MarginCallEnterTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “MARGIN_CALL_ENTER” for an
	// MarginCallEnterTransaction.
	Type TransactionType
}

type MarginCallExtendTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “MARGIN_CALL_EXTEND” for an
	// MarginCallExtendTransaction.
	Type TransactionType

	// The number of the extensions to the Account’s current margin call that
	// have been applied. This value will be set to 1 for the first
	// MarginCallExtend Transaction
	extensionNumber int64
}

type MarginCallExitTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “MARGIN_CALL_EXIT” for an
	// MarginCallExitTransaction.
	Type TransactionType
}

type DelayedTradeClosureTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “DELAYED_TRADE_CLOSURE” for an
	// DelayedTradeClosureTransaction.
	Type TransactionType

	// The reason for the delayed trade closure
	reason MarketOrderReason

	// List of Trade ID’s identifying the open trades that will be closed when
	// their respective instruments become tradeable
	tradeIDs TradeID
}

type DailyFinancingTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “DAILY_FINANCING” for a
	// DailyFinancingTransaction.
	Type TransactionType

	// The amount of financing paid/collected for the Account.
	financing AccountUnits

	// The Account’s balance after daily financing.
	accountBalance AccountUnits

	// The account financing mode at the time of the daily financing. This field
	// is no longer in use moving forward and was replaced by
	// accountFinancingMode in individual positionFinancings since the financing
	// mode could differ between instruments.
	// Deprecated: Will be removed in a future API update.
	accountFinancingMode AccountFinancingMode

	// The financing paid/collected for each Position in the Account.
	positionFinancings []PositionFinancing
}

type DividendAdjustmentTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID

	// The date/time when the Transaction was created.
	time DateTime

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “DIVIDEND_ADJUSTMENT” for a
	// DividendAdjustmentTransaction.
	Type TransactionType

	// The name of the instrument for the dividendAdjustment transaction
	instrument InstrumentName

	// The total dividend adjustment amount paid or collected in the Account’s
	// home currency for the Account as a result of applying the
	// DividendAdjustment Transaction. This is the sum of the dividend
	// adjustments paid/collected for each OpenTradeDividendAdjustment found
	// within the Transaction.
	dividendAdjustment AccountUnits

	// The total dividend adjustment amount paid or collected in the
	// Instrument’s quote currency for the Account as a result of applying the
	// DividendAdjustment Transaction. This is the sum of the quote dividend
	// adjustments paid/collected for each OpenTradeDividendAdjustment found
	// within the Transaction.
	quoteDividendAdjustment DecimalNumber

	// The HomeConversionFactors in effect at the time of the
	// DividendAdjustment.
	homeConversionFactors HomeConversionFactors

	// The Account balance after applying the DividendAdjustment Transaction
	accountBalance AccountUnits

	// The dividend adjustment payment/collection details for each open Trade,
	// within the Account, for which a dividend adjustment is to be paid or
	// collected.
	openTradeDividendAdjustments []OpenTradeDividendAdjustment
}

type ResetResettablePLTransaction struct {
	// The Transaction’s Identifier.
	id TransactionID `json:"id"`

	// The date/time when the Transaction was created.
	time DateTime `json:"time"`

	// The ID of the user that initiated the creation of the Transaction.
	userID int64

	// The ID of the Account the Transaction was created for.
	accountID AccountID

	// The ID of the “batch” that the Transaction belongs to. Transactions in
	// the same batch are applied to the Account simultaneously.
	batchID TransactionID

	// The Request ID of the request which generated the transaction.
	requestID RequestID

	// The Type of the Transaction. Always set to “RESET_RESETTABLE_PL” for a
	// ResetResettablePLTransaction.
	Type TransactionType `json:"type"`
}

type TransactionID string

func (t TransactionID) String() string { return string(t) }

type TransactionType string // todo

//Create Transaction
const TRANSACTION_TYPE_CREATE_ACCOUNT TransactionType = "CREATE_ACCOUNT"

//Close Transaction
const TRANSACTION_TYPE_CLOSE_ACCOUNT TransactionType = "CLOSE_ACCOUNT"

//Reopen Transaction
const TRANSACTION_TYPE_REOPEN_ACCOUNT TransactionType = "REOPEN_ACCOUNT"

//Client Configuration Transaction
const TRANSACTION_TYPE_CLIENT_CONFIGURE TransactionType = "CLIENT_CONFIGURE"

//Client Configuration Reject Transaction
const TRANSACTION_TYPE_CLIENT_CONFIGURE_REJECT TransactionType = "CLIENT_CONFIGURE_REJECT"

//Funds Transaction
const TRANSACTION_TYPE_TRANSFER_FUNDSTransfer TransactionType = "TRANSFER_FUNDSTransfer"

//Transfer Funds Reject Transaction
const TRANSACTION_TYPE_TRANSFER_FUNDS_REJECT TransactionType = "TRANSFER_FUNDS_REJECT"

//Order Transaction
const TRANSACTION_TYPE_MARKET_ORDERMarket TransactionType = "MARKET_ORDERMarket"

//Market Order Reject Transaction
const TRANSACTION_TYPE_MARKET_ORDER_REJECT TransactionType = "MARKET_ORDER_REJECT"

//Fixed Price Order Transaction
const TRANSACTION_TYPE_FIXED_PRICE_ORDER TransactionType = "FIXED_PRICE_ORDER"

//Order Transaction
const TRANSACTION_TYPE_LIMIT_ORDERLimit TransactionType = "LIMIT_ORDERLimit"

//Limit Order Reject Transaction
const TRANSACTION_TYPE_LIMIT_ORDER_REJECT TransactionType = "LIMIT_ORDER_REJECT"

//Stop Order Transaction
const TRANSACTION_TYPE_STOP_ORDER TransactionType = "STOP_ORDER"

//Stop Order Reject Transaction
const TRANSACTION_TYPE_STOP_ORDER_REJECT TransactionType = "STOP_ORDER_REJECT"

//Market if Touched Order Transaction
const TRANSACTION_TYPE_MARKET_IF_TOUCHED_ORDER TransactionType = "MARKET_IF_TOUCHED_ORDER"

//Market if Touched Order Reject Transaction
const TRANSACTION_TYPE_MARKET_IF_TOUCHED_ORDER_REJECT TransactionType = "MARKET_IF_TOUCHED_ORDER_REJECT"

//Take Profit Order Transaction
const TRANSACTION_TYPE_TAKE_PROFIT_ORDER TransactionType = "TAKE_PROFIT_ORDER"

//Take Profit Order Reject Transaction
const TRANSACTION_TYPE_TAKE_PROFIT_ORDER_REJECT TransactionType = "TAKE_PROFIT_ORDER_REJECT"

//Stop Loss Order Transaction
const TRANSACTION_TYPE_STOP_LOSS_ORDER TransactionType = "STOP_LOSS_ORDER"

//Stop Loss Order Reject Transaction
const TRANSACTION_TYPE_STOP_LOSS_ORDER_REJECT TransactionType = "STOP_LOSS_ORDER_REJECT"

//Guaranteed Stop Loss Order Transaction
const TRANSACTION_TYPE_GUARANTEED_STOP_LOSS_ORDER TransactionType = "GUARANTEED_STOP_LOSS_ORDER"

//Guaranteed Stop Loss Order Reject Transaction
const TRANSACTION_TYPE_GUARANTEED_STOP_LOSS_ORDER_REJECT TransactionType = "GUARANTEED_STOP_LOSS_ORDER_REJECT"

//Trailing Stop Loss Order Transaction
const TRANSACTION_TYPE_TRAILING_STOP_LOSS_ORDER TransactionType = "TRAILING_STOP_LOSS_ORDER"

//Trailing Stop Loss Order Reject Transaction
const TRANSACTION_TYPE_TRAILING_STOP_LOSS_ORDER_REJECT TransactionType = "TRAILING_STOP_LOSS_ORDER_REJECT"

//Order Fill Transaction
const TRANSACTION_TYPE_ORDER_FILL TransactionType = "ORDER_FILL"

//Order Cancel Transaction
const TRANSACTION_TYPE_ORDER_CANCEL TransactionType = "ORDER_CANCEL"

//Order Cancel Reject Transaction
const TRANSACTION_TYPE_ORDER_CANCEL_REJECT TransactionType = "ORDER_CANCEL_REJECT"

//Order Client Extensions Modify Transaction
const TRANSACTION_TYPE_ORDER_CLIENT_EXTENSIONS_MODIFY TransactionType = "ORDER_CLIENT_EXTENSIONS_MODIFY"

//Order Client Extensions Modify Reject Transaction
const TRANSACTION_TYPE_ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionType = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

//Trade Client Extensions Modify Transaction
const TRANSACTION_TYPE_TRADE_CLIENT_EXTENSIONS_MODIFY TransactionType = "TRADE_CLIENT_EXTENSIONS_MODIFY"

//Trade Client Extensions Modify Reject Transaction
const TRANSACTION_TYPE_TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionType = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

//Margin Call Enter Transaction
const TRANSACTION_TYPE_MARGIN_CALL_ENTER TransactionType = "MARGIN_CALL_ENTER"

//Margin Call Extend Transaction
const TRANSACTION_TYPE_MARGIN_CALL_EXTEND TransactionType = "MARGIN_CALL_EXTEND"

//Margin Call Exit Transaction
const TRANSACTION_TYPE_MARGIN_CALL_EXIT TransactionType = "MARGIN_CALL_EXIT"

//Delayed Trade Closure Transaction
const TRANSACTION_TYPE_DELAYED_TRADE_CLOSURE TransactionType = "DELAYED_TRADE_CLOSURE"

//Daily Financing Transaction
const TRANSACTION_TYPE_DAILY_FINANCING TransactionType = "DAILY_FINANCING"

//Dividend Adjustment Transaction
const TRANSACTION_TYPE_DIVIDEND_ADJUSTMENT TransactionType = "DIVIDEND_ADJUSTMENT"

//Reset Resettable PL Transaction
const TRANSACTION_TYPE_RESET_RESETTABLE_PL TransactionType = "RESET_RESETTABLE_PL"

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

type OpenTradeDividendAdjustment struct {
	// The ID of the Trade for which the dividend adjustment is to be paid or
	// collected.
	TradeID TradeID `json:"tradeID"`

	// The dividend adjustment amount to pay or collect for the Trade.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`

	// The dividend adjustment amount to pay or collect for the Trade, in the
	// Instrument’s quote currency.
	QuoteDividendAdjustment DecimalNumber `json:"quoteDividendAdjustment"`
}

type ClientID string

type ClientTag string

type ClientComment string

type ClientExtensions struct {
	// The Client ID of the Order/Trade
	Id ClientID `json:"id,omitempty"`

	// A tag associated with the Order/Trade
	Tag ClientTag `json:"tag,omitempty"`

	// A comment associated with the Order/Trade
	Comment ClientComment `json:"comment,omitempty"`
}

type TakeProfitDetails struct {
	// The price that the Take Profit Order will be triggered at. Only one of
	// the price and distance fields may be specified.
	Price PriceValue `json:"price"`

	// The time in force for the created Take Profit Order. This may only be
	// GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`

	// The date when the Take Profit Order will be cancelled on if timeInForce
	// is GTD.
	GtdTime DateTime `json:"gtdTime,omitempty"`

	// The Client Extensions to add to the Take Profit Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`
}

type StopLossDetails struct {
	// The price that the Stop Loss Order will be triggered at. Only one of the
	// price and distance fields may be specified.
	Price PriceValue `json:"price,omitempty"`

	// Specifies the distance (in price units) from the Trade’s open price to
	// use as the Stop Loss Order price. Only one of the distance and price
	// fields may be specified.
	Distance DecimalNumber `json:"distance,omitempty"`

	// The time in force for the created Stop Loss Order. This may only be GTC,
	// GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`

	// The date when the Stop Loss Order will be cancelled on if timeInForce is
	// GTD.
	GtdTime DateTime `json:"gtdTime,omitempty"`

	// The Client Extensions to add to the Stop Loss Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// Flag indicating that the price for the Stop Loss Order is guaranteed. The
	// default value depends on the GuaranteedStopLossOrderMode of the account,
	// if it is REQUIRED, the default will be true, for DISABLED or ENABLED the
	// default is false.
	Guaranteed bool `json:"guaranteed,omitempty"`
}

type GuaranteedStopLossDetails struct{} //todo

type TrailingStopLossDetails struct {
	// The distance (in price units) from the Trade’s fill price that the
	// Trailing Stop Loss Order will be triggered at.
	Distance DecimalNumber `json:"distance,omitempty"`

	// The time in force for the created Trailing Stop Loss Order. This may only
	// be GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`

	// The date when the Trailing Stop Loss Order will be cancelled on if
	// timeInForce is GTD.
	GtdTime DateTime `json:"gtdTime,omitempty"`

	// The Client Extensions to add to the Trailing Stop Loss Order when
	// created.
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`
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

type MarketOrderTradeClose struct {
	// The reason the Market Order was created to perform a margin closeout
	reason MarketOrderMarginCloseoutReason
}

type MarketOrderMarginCloseout struct {
	//The reason the Market Order was created to perform a margin closeout
	Reason MarketOrderMarginCloseoutReason `json:"reason"`
}

type MarketOrderMarginCloseoutReason string // todo const

type MarketOrderDelayedTradeClose struct {
	// The ID of the Trade being closed
	tradeID TradeID

	// The Client ID of the Trade being closed
	clientTradeID TradeID

	// The Transaction ID of the DelayedTradeClosure transaction to which this
	// Delayed Trade Close belongs to
	sourceTransactionID TransactionID
}

type MarketOrderPositionCloseout struct {
	Instrument InstrumentName `json:"instrument"`
	Units      string         `json:"units"`
}

type LiquidityRegenerationSchedule struct {
	// The steps in the Liquidity Regeneration Schedule
	steps []LiquidityRegenerationScheduleStep
}

type LiquidityRegenerationScheduleStep struct {
	// The timestamp of the schedule step.
	timestamp DateTime

	// The amount of bid liquidity used at this step in the schedule.
	bidLiquidityUsed DecimalNumber

	// The amount of ask liquidity used at this step in the schedule.
	askLiquidityUsed DecimalNumber
}

type OpenTradeFinancing struct {
	// The ID of the Trade that financing is being paid/collected for.
	tradeID TradeID

	// The amount of financing paid/collected for the Trade.
	financing AccountUnits

	// The amount of financing paid/collected in the Instrument’s base currency
	// for the Trade.
	baseFinancing DecimalNumber

	// The amount of financing paid/collected in the Instrument’s quote currency
	// for the Trade.
	quoteFinancing DecimalNumber

	// The financing rate in effect for the instrument used to calculate the the
	// amount of financing paid/collected for the Trade. This field will only be
	// set if the AccountFinancingMode at the time of the daily financing is
	// DAILY_INSTRUMENT or SECOND_BY_SECOND_INSTRUMENT. The value is in decimal
	// rather than percentage points, e.g. 5% is represented as 0.05.
	financingRate DecimalNumber
}

type PositionFinancing struct {
	// The instrument of the Position that financing is being paid/collected
	// for.
	instrument InstrumentName

	// The amount of financing paid/collected for the Position.
	financing AccountUnits

	// The amount of base financing paid/collected for the Position.
	baseFinancing DecimalNumber

	// The amount of quote financing paid/collected for the Position.
	quoteFinancing DecimalNumber

	// The HomeConversionFactors in effect for the Position’s Instrument at the
	// time of the DailyFinancing.
	homeConversionFactors HomeConversionFactors

	// The financing paid/collected for each open Trade within the Position.
	openTradeFinancings []OpenTradeFinancing

	// The account financing mode at the time of the daily financing.
	accountFinancingMode AccountFinancingMode
}

type RequestID string

type TransactionRejectReason string // todo const

type TransactionFilter string // TODO

type TransactionHeartbeat struct {
	// The string “HEARTBEAT”
	Type string

	// The ID of the most recent Transaction created for the Account
	lastTransactionID TransactionID

	// The date/time when the TransactionHeartbeat was created.
	time DateTime
}
