package V20

type Transaction struct{} //todo

type CreateTransaction struct{} //todo

type CloseTransaction struct{} //todo

type ReopenTransaction struct{} //todo

type ClientConfigureTransaction struct{} //todo

type ClientConfigureRejectTransaction struct{} //todo

type TransferFundsTransaction struct{} //todo

type TransferFundsRejectTransaction struct{} //todo

type MarketOrderTransaction struct{} //todo

type MarketOrderRejectTransaction struct{} //todo

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

type OrderFillTransaction struct{} //todo

type OrderCancelTransaction struct{} //todo

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

type TakeProfitDetails struct{} // TODO

type StopLossDetails struct{} // TODO

type TrailingStopLossDetails struct{} // TODO

type TradeOpen struct{} // TODO

type TradeReduce struct{} // TODO

type MarketOrderTradeClose struct{} // TODO

type MarketOrderMarginCloseout struct{} // TODO

type MarketOrderMarginCloseoutReason string // todo const

type MarketOrderDelayedTradeClose struct{} // TODO

type MarketOrderPositionCloseout struct{} // TODO

type LiquidityRegenerationSchedule struct{} // TODO

type LiquidityRegenerationScheduleStep struct{} // TODO

type OpenTradeFinancing struct{} // TODO

type PositionFinancing struct{} // TODO

type RequestID string

type TransactionRequestReason string // todo const

type TransactionFilter string // TODO

type TransactionHeartbeat struct{} // TODO
