package gOanda

import "errors"

type RawOrder struct {
	Id                         OrderID                      `json:"id"`
	CreateTime                 DateTime                     `json:"createTime"`
	State                      OrderState                   `json:"state"`
	Instrument                 InstrumentName               `json:"instrument"`
	ClientExtensions           ClientExtensions             `json:"clientExtensions"`
	Type                       OrderType                    `json:"type"`
	TradeClose                 MarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout       MarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout      MarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout             MarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose          MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Units                      DecimalNumber                `json:"units"`
	Price                      PriceValue                   `json:"price"`
	PriceBound                 PriceValue                   `json:"priceBound"`
	TimeInForce                TimeInForce                  `json:"timeInForce"`
	GtdTime                    DateTime                     `json:"gtdTime"`
	PositionFill               OrderPositionFill            `json:"positionFill"`
	TriggerCondition           OrderTriggerCondition        `json:"triggerCondition"`
	TakeProfitOnFill           TakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill             StopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill     TrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions      ClientExtensions             `json:"tradeClientExtensions"`
	FillingTransactionID       TransactionID                `json:"fillingTransactionID"`
	FilledTime                 DateTime                     `json:"filledTime"`
	TradeOpenedID              TradeID                      `json:"tradeOpenedID"`
	TradeReducedID             TradeID                      `json:"tradeReducedID"`
	TradeClosedIDs             []TradeID                    `json:"tradeClosedIDs"`
	CancellingTransactionID    TransactionID                `json:"cancellingTransactionID"`
	CancelledTime              DateTime                     `json:"cancelledTime"`
	ReplacesOrderID            OrderID                      `json:"replacesOrderID"`
	ReplacedByOrderID          OrderID                      `json:"replacedByOrderID"`
	TradeID                    TradeID                      `json:"tradeID"`
	ClientTradeID              ClientID                     `json:"clientTradeID"`
	GuaranteedExecutionPremium DecimalNumber                `json:"guaranteedExecutionPremium"`
	Distance                   DecimalNumber                `json:"distance"`
	Guaranteed                 bool                         `json:"guaranteed"`
}

func (o *RawOrder) ToOrder() (Order, error) {
	switch o.Type {
	case ORDER_TYPE_STOP:
		return &StopOrder{
			Id:                      o.Id,
			CreateTime:              o.CreateTime,
			State:                   o.State,
			ClientExtensions:        o.ClientExtensions,
			Type:                    o.Type,
			Instrument:              o.Instrument,
			Units:                   o.Units,
			Price:                   o.Price,
			PriceBound:              o.PriceBound,
			TimeInForce:             o.TimeInForce,
			GtdTime:                 o.GtdTime,
			PositionFill:            o.PositionFill,
			TriggerCondition:        o.TriggerCondition,
			TakeProfitOnFill:        o.TakeProfitOnFill,
			StopLossOnFill:          o.StopLossOnFill,
			TrailingStopLossOnFill:  o.TrailingStopLossOnFill,
			TradeClientExtensions:   o.TradeClientExtensions,
			FillingTransactionID:    o.FillingTransactionID,
			FilledTime:              o.FilledTime,
			TradeOpenedID:           o.TradeOpenedID,
			TradeReducedID:          o.TradeReducedID,
			TradeClosedIDs:          o.TradeClosedIDs,
			CancellingTransactionID: o.CancellingTransactionID,
			CancelledTime:           o.CancelledTime,
			ReplacesOrderID:         o.ReplacesOrderID,
			ReplacedByOrderID:       o.ReplacedByOrderID,
		}, nil

	case ORDER_TYPE_LIMIT:
		return &LimitOrder{
			Id:                      o.Id,
			CreateTime:              o.CreateTime,
			State:                   o.State,
			ClientExtensions:        o.ClientExtensions,
			Type:                    o.Type,
			Instrument:              o.Instrument,
			Units:                   o.Units,
			Price:                   o.Price,
			TimeInForce:             o.TimeInForce,
			GtdTime:                 o.GtdTime,
			PositionFill:            o.PositionFill,
			TriggerCondition:        o.TriggerCondition,
			TakeProfitOnFill:        o.TakeProfitOnFill,
			StopLossOnFill:          o.StopLossOnFill,
			TrailingStopLossOnFill:  o.TrailingStopLossOnFill,
			TradeClientExtensions:   o.TradeClientExtensions,
			FillingTransactionID:    o.FillingTransactionID,
			FilledTime:              o.FilledTime,
			TradeOpenedID:           o.TradeOpenedID,
			TradeReducedID:          o.TradeReducedID,
			TradeClosedIDs:          o.TradeClosedIDs,
			CancellingTransactionID: o.CancellingTransactionID,
			CancelledTime:           o.CancelledTime,
			ReplacesOrderID:         o.ReplacesOrderID,
			ReplacedByOrderID:       o.ReplacedByOrderID,
		}, nil

	// TODO fix consistency between exported & non exported fields of order
	case ORDER_TYPE_MARKET:
		return &MarketOrder{
			Id:                      o.Id,
			CreateTime:              o.CreateTime,
			State:                   o.State,
			ClientExtensions:        o.ClientExtensions,
			Type:                    o.Type,
			Instrument:              o.Instrument,
			Units:                   o.Units,
			TimeInForce:             o.TimeInForce,
			PriceBound:              o.PriceBound,
			PositionFill:            o.PositionFill,
			TradeClose:              o.TradeClose,
			LongPositionCloseout:    o.LongPositionCloseout,
			ShortPositionCloseout:   o.ShortPositionCloseout,
			MarginCloseout:          o.MarginCloseout,
			DelayedTradeClose:       o.DelayedTradeClose,
			TakeProfitOnFill:        o.TakeProfitOnFill,
			StopLossOnFill:          o.StopLossOnFill,
			TrailingStopLossOnFill:  o.TrailingStopLossOnFill,
			TradeClientExtensions:   o.TradeClientExtensions,
			FillingTransactionID:    o.FillingTransactionID,
			FilledTime:              o.FilledTime,
			TradeOpenedID:           o.TradeOpenedID,
			TradeReducedID:          o.TradeReducedID,
			TradeClosedIDs:          o.TradeClosedIDs,
			CancellingTransactionID: o.CancellingTransactionID,
			CancelledTime:           o.CancelledTime,
		}, nil

	case ORDER_TYPE_TAKE_PROFIT:
		return &TakeProfitOrder{
			Id:                      o.Id,
			CreateTime:              o.CreateTime,
			State:                   o.State,
			ClientExtensions:        o.ClientExtensions,
			Type:                    o.Type,
			TradeID:                 o.TradeID,
			ClientTradeID:           o.ClientTradeID,
			Price:                   o.Price,
			TimeInForce:             o.TimeInForce,
			GtdTime:                 o.GtdTime,
			TriggerCondition:        o.TriggerCondition,
			FillingTransactionID:    o.FillingTransactionID,
			FilledTime:              o.FilledTime,
			TradeOpenedID:           o.TradeOpenedID,
			TradeReducedID:          o.TradeReducedID,
			TradeClosedIDs:          o.TradeClosedIDs,
			CancellingTransactionID: o.CancellingTransactionID,
			CancelledTime:           o.CancelledTime,
			ReplacesOrderID:         o.ReplacesOrderID,
			ReplacedByOrderID:       o.ReplacedByOrderID,
		}, nil

	case ORDER_TYPE_STOP_LOSS:
		return &StopLossOrder{
			Id:                         o.Id,
			CreateTime:                 o.CreateTime,
			State:                      o.State,
			ClientExtensions:           o.ClientExtensions,
			Type:                       o.Type,
			GuaranteedExecutionPremium: o.GuaranteedExecutionPremium,
			TradeID:                    o.TradeID,
			ClientTradeID:              o.ClientTradeID,
			Price:                      o.Price,
			Distance:                   o.Distance,
			TimeInForce:                o.TimeInForce,
			GtdTime:                    o.GtdTime,
			TriggerCondition:           o.TriggerCondition,
			Guaranteed:                 o.Guaranteed,
			FillingTransactionID:       o.FillingTransactionID,
			FilledTime:                 o.FilledTime,
			TradeOpenedID:              o.TradeOpenedID,
			TradeReducedID:             o.TradeReducedID,
			TradeClosedIDs:             o.TradeClosedIDs,
			CancellingTransactionID:    o.CancellingTransactionID,
			CancelledTime:              o.CancelledTime,
			ReplacesOrderID:            o.ReplacesOrderID,
			ReplacedByOrderID:          o.ReplacedByOrderID,
		}, nil
	}
	return nil, errors.New("OrderType not implemented for: " + string(o.Type))
}

type Order interface {
	GetType() OrderType
}

type MarketOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “MARKET” for Market Orders.
	//default=MARKET
	Type OrderType `json:"type"`

	// The Market Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Market Order. A posititive
	// number of units results in a long Order, and a negative number of units
	// results in a short Order.
	Units DecimalNumber `json:"units"`

	// The time-in-force requested for the Market Order. Restricted to FOK or
	// IOC for a MarketOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The worst price that the client is willing to have the Market Order
	// filled at.
	PriceBound PriceValue `json:"priceBound"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

	// Details of the Trade requested to be closed, only provided when the
	// Market Order is being used to explicitly close a Trade.
	TradeClose MarketOrderTradeClose `json:"tradeClose"`

	// Details of the long Position requested to be closed out, only provided
	// when a Market Order is being used to explicitly closeout a long Position.
	LongPositionCloseout MarketOrderPositionCloseout `json:"longPositionCloseout"`

	// Details of the short Position requested to be closed out, only provided
	// when a Market Order is being used to explicitly closeout a short Position.
	ShortPositionCloseout MarketOrderPositionCloseout `json:"shortPositionCloseout"`

	// Details of the Margin Closeout that this Market Order was created for
	MarginCloseout MarketOrderMarginCloseout `json:"marginCloseout"`

	// Details of the delayed Trade close that this Market Order was created for
	DelayedTradeClose MarketOrderDelayedTradeClose `json:"delayedTradeClose"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
}

func (m *MarketOrder) GetType() OrderType {
	return m.Type
}

type FixedPriceOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “FIXED_PRICE” for Fixed Price
	// Orders.
	Type OrderType `json:"type"`

	// The Fixed Price Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Fixed Price Order. A
	// posititive number of units results in a long Order, and a negative number
	// of units results in a short Order.
	Units DecimalNumber `json:"units"`

	// The price specified for the Fixed Price Order. This price is the exact
	// price that the Fixed Price Order will be filled at.
	Price PriceValue `json:"price"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

	// The state that the trade resulting from the Fixed Price Order should be
	// set to.
	TradeState string `json:"tradeState"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
}

func (f *FixedPriceOrder) GetType() OrderType {
	return f.Type
}

type LimitOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “LIMIT” for Limit Orders.
	Type OrderType `json:"type"`

	// The Limit Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Limit Order. A posititive
	// number of units results in a long Order, and a negative number of units
	// results in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the Limit Order. The Limit Order will
	// only be filled by a market price that is equal to or better than this
	// price.
	Price PriceValue `json:"price"`

	// The time-in-force requested for the Limit Order.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Limit Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (l *LimitOrder) GetType() OrderType {
	return l.Type
}

type StopOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “STOP” for Stop Orders.
	Type OrderType `json:"type"`

	// The Stop Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Stop Order. A posititive
	// number of units results in a long Order, and a negative number of units
	// results in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the Stop Order. The Stop Order will
	// only be filled by a market price that is equal to or worse than this
	// price.
	Price PriceValue `json:"price"`

	// The worst market price that may be used to fill this Stop Order. If the
	// market gaps and crosses through both the price and the priceBound, the
	// Stop Order will be cancelled instead of being filled.
	PriceBound PriceValue `json:"priceBound"`

	// The time-in-force requested for the Stop Order.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Stop Order will be cancelled if its timeInForce is
	// “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (s *StopOrder) GetType() OrderType {
	return s.Type
}

type MarketIfTouchedOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “MARKET_IF_TOUCHED” for Market If
	// Touched Orders.
	Type OrderType `json:"type"`

	// The MarketIfTouched Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the MarketIfTouched Order. A
	// posititive number of units results in a long Order, and a negative number
	// of units results in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the MarketIfTouched Order. The
	// MarketIfTouched Order will only be filled by a market price that crosses
	// this price from the direction of the market price at the time when the
	// Order was created (the initialMarketPrice). Depending on the value of the
	// Order’s price and initialMarketPrice, the MarketIfTouchedOrder will
	// behave like a Limit or a Stop Order.
	Price PriceValue `json:"price"`

	// The worst market price that may be used to fill this MarketIfTouched
	// Order.
	PriceBound PriceValue `json:"priceBound"`

	// The time-in-force requested for the MarketIfTouched Order. Restricted to
	// “GTC”, “GFD” and “GTD” for MarketIfTouched Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the MarketIfTouched Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The Market price at the time when the MarketIfTouched Order was created.
	InitialMarketPrice PriceValue `json:"initialMarketPrice"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (m *MarketIfTouchedOrder) GetType() OrderType {
	return m.Type
}

type TakeProfitOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “TAKE_PROFIT” for Take Profit
	// Orders.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the TakeProfit Order. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	Price PriceValue `json:"price"`

	// The time-in-force requested for the TakeProfit Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TakeProfit Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the TakeProfit Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (t *TakeProfitOrder) GetType() OrderType {
	return t.Type
}

type StopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “STOP_LOSS” for Stop Loss Orders.
	Type OrderType `json:"type"`

	// The premium that will be charged if the Stop Loss Order is guaranteed and
	// the Order is filled at the guaranteed price. It is in price units and is
	// charged for each unit of the Trade.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the Stop Loss Order. If the guaranteed
	// flag is false, the associated Trade will be closed by a market price that
	// is equal to or worse than this threshold. If the flag is true the
	// associated Trade will be closed at this price.
	Price PriceValue `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for StopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// Flag indicating that the Stop Loss Order is guaranteed. The default value
	// depends on the GuaranteedStopLossOrderMode of the account, if it is
	// REQUIRED, the default will be true, for DISABLED or ENABLED the default
	// is false.
	Guaranteed bool `json:"guaranteed"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (s *StopLossOrder) GetType() OrderType {
	return s.Type
}

type GuaranteedStopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “GUARANTEED_STOP_LOSS” for
	// Guaranteed Stop Loss Orders.
	Type OrderType `json:"type"`

	// The premium that will be charged if the Guaranteed Stop Loss Order is
	// filled at the guaranteed price. It is in price units and is charged for
	// each unit of the Trade.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the Guaranteed Stop Loss Order. The
	// associated Trade will be closed at this price.
	Price PriceValue `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Guaranteed Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted
	// to “GTC”, “GFD” and “GTD” for GuaranteedStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the GuaranteedStopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

type TrailingStopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “TRAILING_STOP_LOSS” for Trailing
	// Stop Loss Orders.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price distance (in price units) specified for the TrailingStopLoss
	// Order.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TrailingStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The trigger price for the Trailing Stop Loss Order. The trailing stop
	// value will trail (follow) the market price by the TSL order’s configured
	// “distance” as the market price moves in the winning direction. If the
	// market price moves to a level that is equal to or worse than the trailing
	// stop value, the order will be filled and the Trade will be closed.
	TrailingStopValue PriceValue `json:"trailingStopValue"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (t *TrailingStopLossOrder) GetType() OrderType {
	return t.Type
}

type AnOrderRequest interface {
	ToOrderRequest() OrderRequest
}

type OrderRequest struct {
	Type                   OrderType                `json:"type"`
	Instrument             InstrumentName           `json:"instrument"`
	Units                  DecimalNumber            `json:"units"`
	Price                  *PriceValue              `json:"price,omitempty"`
	TimeInForce            TimeInForce              `json:"timeInForce,omitempty"`
	GtdTime                DateTime                 `json:"gtdTime,omitempty"`
	PriceBound             PriceValue               `json:"priceBound,omitempty"`
	PositionFill           OrderPositionFill        `json:"positionFill,omitempty"`
	ClientExtensions       *ClientExtensions        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetails       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetails         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensions        `json:"tradeClientExtensions,omitempty"`
}

type MarketOrderRequest struct {
	Type                   OrderType                `json:"type"`
	Instrument             InstrumentName           `json:"instrument"`
	Units                  DecimalNumber            `json:"units"`
	TimeInForce            TimeInForce              `json:"timeInForce"`
	PriceBound             PriceValue               `json:"priceBound"`
	PositionFill           OrderPositionFill        `json:"positionFill"`
	ClientExtensions       *ClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       *TakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         *StopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  *ClientExtensions        `json:"tradeClientExtensions"`
}

func (m *MarketOrderRequest) ToOrderRequest() OrderRequest {

	return OrderRequest{
		Type:                   ORDER_TYPE_MARKET,
		Instrument:             m.Instrument,
		Units:                  m.Units,
		TimeInForce:            m.TimeInForce,
		PriceBound:             m.PriceBound,
		PositionFill:           m.PositionFill,
		ClientExtensions:       m.ClientExtensions,
		TakeProfitOnFill:       m.TakeProfitOnFill,
		StopLossOnFill:         m.StopLossOnFill,
		TrailingStopLossOnFill: m.TrailingStopLossOnFill,
		TradeClientExtensions:  m.TradeClientExtensions,
	}
}

type LimitOrderRequest struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`

	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “LIMIT” for Limit Orders.
	Type OrderType `json:"type"`

	// The Limit Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Limit Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the Limit Order. The Limit Order will
	// only be filled by a market price that is equal to or better than this
	// price.
	Price PriceValue `json:"price"`

	// The time-in-force requested for the Limit Order.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Limit Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the
	// Order’s state is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state
	// is FILLED)
	FilledTime DateTime `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the
	// fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the
	// fill)
	TradeReducedID TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a
	// result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the
	// Order’s state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of
	// the Order is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	// this Order was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

func (l *LimitOrderRequest) ToOrderRequest() OrderRequest {

	return OrderRequest{
		Type:                   ORDER_TYPE_LIMIT,
		Instrument:             l.Instrument,
		Units:                  l.Units,
		Price:                  &l.Price,
		TimeInForce:            l.TimeInForce,
		GtdTime:                l.GtdTime,
		PositionFill:           l.PositionFill,
		ClientExtensions:       l.ClientExtensions,
		TakeProfitOnFill:       l.TakeProfitOnFill,
		StopLossOnFill:         l.StopLossOnFill,
		TrailingStopLossOnFill: l.TrailingStopLossOnFill,
		TradeClientExtensions:  l.TradeClientExtensions,
	}
}

type StopOrderRequest struct {
	// The type of the Order to Create. Must be set to “STOP” when creating a
	// Stop Order.
	Type OrderType `json:"type"`

	// The Stop Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the Stop Order. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the Stop Order. The Stop Order will
	// only be filled by a market price that is equal to or worse than this
	// price.
	Price PriceValue `json:"price"`

	// The worst market price that may be used to fill this Stop Order. If the
	// market gaps and crosses through both the price and the priceBound, the
	// Stop Order will be cancelled instead of being filled.
	PriceBound PriceValue `json:"priceBound"`

	// The time-in-force requested for the Stop Order.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Stop Order will be cancelled if its timeInForce is
	// “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	// restriction applies when creating a guaranteed Stop Loss Order. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a Stop Loss Order for
	// a long trade valid values are “DEFAULT” and “BID”, and for short trades
	// “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
}

func (s *StopOrderRequest) ToOrderRequest() OrderRequest {

	return OrderRequest{
		Type:                   ORDER_TYPE_STOP,
		Instrument:             s.Instrument,
		Units:                  s.Units,
		TimeInForce:            s.TimeInForce,
		GtdTime:                s.GtdTime,
		Price:                  &s.Price,
		PositionFill:           s.PositionFill,
		ClientExtensions:       s.ClientExtensions,
		TakeProfitOnFill:       s.TakeProfitOnFill,
		StopLossOnFill:         s.StopLossOnFill,
		TrailingStopLossOnFill: s.TrailingStopLossOnFill,
		TradeClientExtensions:  s.TradeClientExtensions,
	}
}

type MarketIfTouchedOrderRequest struct {
	// The type of the Order to Create. Must be set to “MARKET_IF_TOUCHED” when
	// creating a Market If Touched Order.
	Type OrderType `json:"type"`

	// The MarketIfTouched Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`

	// The quantity requested to be filled by the MarketIfTouched Order. A
	// positive number of units results in a long Order, and a negative number
	// of units results in a short Order.
	Units DecimalNumber `json:"units"`

	// The price threshold specified for the MarketIfTouched Order. The
	// MarketIfTouched Order will only be filled by a market price that crosses
	// this price from the direction of the market price at the time when the
	// Order was created (the initialMarketPrice). Depending on the value of the
	// Order’s price and initialMarketPrice, the MarketIfTouchedOrder will
	// behave like a Limit or a Stop Order.
	Price PriceValue `json:"price"`

	// The worst market price that may be used to fill this MarketIfTouched
	// Order.
	PriceBound PriceValue `json:"priceBound"`

	// The time-in-force requested for the MarketIfTouched Order. Restricted to
	// “GTC”, “GFD” and “GTD” for MarketIfTouched Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the MarketIfTouched Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`
}

type TakeProfitOrderRequest struct {
	// The type of the Order to Create. Must be set to “TAKE_PROFIT” when
	// creating a Take Profit Order.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the TakeProfit Order. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	Price PriceValue `json:"price"`

	// The time-in-force requested for the TakeProfit Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TakeProfit Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the TakeProfit Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type StopLossOrderRequest struct {
	// The type of the Order to Create. Must be set to “STOP_LOSS” when creating
	// a Stop Loss Order.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the Stop Loss Order. The associated
	// Trade will be closed by a market price that is equal to or worse than
	// this threshold.
	Price PriceValue `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for StopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// Flag indicating that the Stop Loss Order is guaranteed. The default value
	// depends on the GuaranteedStopLossOrderMode of the account, if it is
	// REQUIRED, the default will be true, for DISABLED or ENABLED the default
	// is false.
	// Deprecated: Will be removed in a future API update.
	Guaranteed bool `json:"guaranteed"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type GuaranteedStopLossOrderRequest struct {
	// The type of the Order to Create. Must be set to “GUARANTEED_STOP_LOSS”
	// when creating a Guaranteed Stop Loss Order.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the Guaranteed Stop Loss Order. The
	// associated Trade will be closed at this price.
	Price PriceValue `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the Guaranteed Stop Loss Order price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted
	// to “GTC”, “GFD” and “GTD” for GuaranteedStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the GuaranteedStopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type TrailingStopLossOrderRequest struct {
	// The type of the Order to Create. Must be set to “TRAILING_STOP_LOSS” when
	// creating a Trailing Stop Loss Order.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is
	// breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price distance (in price units) specified for the TrailingStopLoss
	// Order.
	Distance DecimalNumber `json:"distance"`

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to
	// “GTC”, “GFD” and “GTD” for TrailingStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the StopLoss Order will be cancelled if its
	// timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

type OrderID string

func (o OrderID) String() string { return string(o) }

type OrderType string // todo do consts

type CancellableOrderType string // todo do consts

type OrderState string // todo consts

type OrderStateFilter string // todo consts

//The Orders that are currently pending execution
const PENDING OrderStateFilter = "PENDING"

//The Orders that have been filled
const FILLED OrderStateFilter = "FILLED"

//The Orders that have been triggered
const TRIGGERED OrderStateFilter = "TRIGGERED"

//The Orders that have been cancelled
const CANCELLED OrderStateFilter = "CANCELLED"

//The Orders that are in any of the possible states listed above
const ALL OrderStateFilter = "ALL"

func (o OrderStateFilter) String() string { return string(o) }

type OrderIdentifier struct {
	// The OANDA-assigned Order ID
	OrderID OrderID `json:"orderID"`

	// The client-provided client Order ID
	ClientOrderID ClientID `json:"clientOrderID"`
}

type OrderSpecifier string

func (o OrderSpecifier) String() string { return string(o) }

type TimeInForce string

type OrderPositionFill string // todo consts

type OrderTriggerCondition string // todo consts

type DynamicOrderState struct {
	// The Order’s ID.
	Id OrderID `json:"id"`

	// The Order’s calculated trailing stop value.
	TrailingStopValue PriceValue `json:"trailingStopValue"`

	// The distance between the Trailing Stop Loss Order’s trailingStopValue and
	// the current Market Price. This represents the distance (in price units)
	// of the Order from a triggering price. If the distance could not be
	// determined, this value will not be set.
	TriggerDistance PriceValue `json:"triggerDistance"`

	// True if an exact trigger distance could be calculated. If false, it means
	// the provided trigger distance is a best estimate. If the distance could
	// not be determined, this value will not be set.
	IsTriggerDistanceExact bool `json:"isTriggerDistanceExact"`
}

type UnitsAvailableDetails struct {
	// The units available for long Orders.
	Long DecimalNumber `json:"long"`

	// The units available for short Orders.
	Short DecimalNumber `json:"short"`
}

type UnitsAvailable struct {
	// The number of units that are available to be traded using an Order with a
	// positionFill option of “DEFAULT”. For an Account with hedging enabled,
	// this value will be the same as the “OPEN_ONLY” value. For an Account
	// without hedging enabled, this value will be the same as the
	// “REDUCE_FIRST” value.
	Default UnitsAvailableDetails `json:"default"`

	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “REDUCE_FIRST”.
	reduceFirst UnitsAvailableDetails `json:"reduceFirst"`

	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “REDUCE_ONLY”.
	reduceOnly UnitsAvailableDetails `json:"reduceOnly"`

	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “OPEN_ONLY”.
	openOnly UnitsAvailableDetails `json:"openOnly"`
}

type GuaranteedStopLossOrderEntryData struct {
	// The minimum distance allowed between the Trade’s fill price and the
	// configured price for guaranteed Stop Loss Orders created for this
	// instrument. Specified in price units.
	minimumDistance DecimalNumber `json:"minimumDistance"`

	// The amount that is charged to the account if a guaranteed Stop Loss Order
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade.
	premium DecimalNumber `json:"premium"`

	// The guaranteed Stop Loss Order level restriction for this instrument.
	levelRestriction GuaranteedStopLossOrderLevelRestriction `json:"levelRestriction"`
}
