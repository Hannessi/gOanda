package gOanda

const INSTRUMENT_NAME_AUDCAD InstrumentName = "AUD_CAD"
const INSTRUMENT_NAME_AUDCHF InstrumentName = "AUD_CHF"
const INSTRUMENT_NAME_AUDJPY InstrumentName = "AUD_JPY"
const INSTRUMENT_NAME_AUDNZD InstrumentName = "AUD_NZD"
const INSTRUMENT_NAME_AUDUSD InstrumentName = "AUD_USD"
const INSTRUMENT_NAME_BCOUSD InstrumentName = "BCO_USD"
const INSTRUMENT_NAME_CADCHF InstrumentName = "CAD_CHF"
const INSTRUMENT_NAME_CADJPY InstrumentName = "CAD_JPY"
const INSTRUMENT_NAME_CHFJPY InstrumentName = "CHF_JPY"
const INSTRUMENT_NAME_CHFZAR InstrumentName = "CHF_ZAR"
const INSTRUMENT_NAME_EURAUD InstrumentName = "EUR_AUD"
const INSTRUMENT_NAME_EURCAD InstrumentName = "EUR_CAD"
const INSTRUMENT_NAME_EURCHF InstrumentName = "EUR_CHF"
const INSTRUMENT_NAME_EURGBP InstrumentName = "EUR_GBP"
const INSTRUMENT_NAME_EURJPY InstrumentName = "EUR_JPY"
const INSTRUMENT_NAME_EURNZD InstrumentName = "EUR_NZD"
const INSTRUMENT_NAME_EURUSD InstrumentName = "EUR_USD"
const INSTRUMENT_NAME_EURZAR InstrumentName = "EUR_ZAR"
const INSTRUMENT_NAME_GBPAUD InstrumentName = "GBP_AUD"
const INSTRUMENT_NAME_GBPCAD InstrumentName = "GBP_CAD"
const INSTRUMENT_NAME_GBPCHF InstrumentName = "GBP_CHF"
const INSTRUMENT_NAME_GBPJPY InstrumentName = "GBP_JPY"
const INSTRUMENT_NAME_GBPNZD InstrumentName = "GBP_NZD"
const INSTRUMENT_NAME_GBPUSD InstrumentName = "GBP_USD"
const INSTRUMENT_NAME_NZDCAD InstrumentName = "NZD_CAD"
const INSTRUMENT_NAME_NZDJPY InstrumentName = "NZD_JPY"
const INSTRUMENT_NAME_NZDUSD InstrumentName = "NZD_USD"
const INSTRUMENT_NAME_USDCAD InstrumentName = "USD_CAD"
const INSTRUMENT_NAME_USDCHF InstrumentName = "USD_CHF"
const INSTRUMENT_NAME_USDJPY InstrumentName = "USD_JPY"
const INSTRUMENT_NAME_USDTRY InstrumentName = "USD_TRY"
const INSTRUMENT_NAME_USDZAR InstrumentName = "USD_ZAR"
const INSTRUMENT_NAME_XAGUSD InstrumentName = "XAG_USD"
const INSTRUMENT_NAME_XAUUSD InstrumentName = "XAU_USD"

func GetAllInstrumentNames() []InstrumentName {
	return []InstrumentName{
		INSTRUMENT_NAME_AUDCAD,
		INSTRUMENT_NAME_AUDCHF,
		INSTRUMENT_NAME_AUDJPY,
		INSTRUMENT_NAME_AUDNZD,
		INSTRUMENT_NAME_AUDUSD,
		INSTRUMENT_NAME_BCOUSD,
		INSTRUMENT_NAME_CADCHF,
		INSTRUMENT_NAME_CADJPY,
		INSTRUMENT_NAME_CHFJPY,
		INSTRUMENT_NAME_CHFZAR,
		INSTRUMENT_NAME_EURAUD,
		INSTRUMENT_NAME_EURCAD,
		INSTRUMENT_NAME_EURCHF,
		INSTRUMENT_NAME_EURGBP,
		INSTRUMENT_NAME_EURJPY,
		INSTRUMENT_NAME_EURNZD,
		INSTRUMENT_NAME_EURUSD,
		INSTRUMENT_NAME_EURZAR,
		INSTRUMENT_NAME_GBPAUD,
		INSTRUMENT_NAME_GBPCAD,
		INSTRUMENT_NAME_GBPCHF,
		INSTRUMENT_NAME_GBPJPY,
		INSTRUMENT_NAME_GBPNZD,
		INSTRUMENT_NAME_GBPUSD,
		INSTRUMENT_NAME_NZDCAD,
		INSTRUMENT_NAME_NZDJPY,
		INSTRUMENT_NAME_NZDUSD,
		INSTRUMENT_NAME_USDCAD,
		INSTRUMENT_NAME_USDCHF,
		INSTRUMENT_NAME_USDJPY,
		INSTRUMENT_NAME_USDTRY,
		INSTRUMENT_NAME_USDZAR,
		INSTRUMENT_NAME_XAGUSD,
		INSTRUMENT_NAME_XAUUSD,
	}
}

const CANDLESTICK_GRANULARITY_S5 CandlestickGranularity = "S5"
const CANDLESTICK_GRANULARITY_S10 CandlestickGranularity = "S10"
const CANDLESTICK_GRANULARITY_S15 CandlestickGranularity = "S15"
const CANDLESTICK_GRANULARITY_S30 CandlestickGranularity = "S30"
const CANDLESTICK_GRANULARITY_M1 CandlestickGranularity = "M1"
const CANDLESTICK_GRANULARITY_M2 CandlestickGranularity = "M2"
const CANDLESTICK_GRANULARITY_M4 CandlestickGranularity = "M4"
const CANDLESTICK_GRANULARITY_M5 CandlestickGranularity = "M5"
const CANDLESTICK_GRANULARITY_M10 CandlestickGranularity = "M10"
const CANDLESTICK_GRANULARITY_M15 CandlestickGranularity = "M15"
const CANDLESTICK_GRANULARITY_M30 CandlestickGranularity = "M30"
const CANDLESTICK_GRANULARITY_H1 CandlestickGranularity = "H1"
const CANDLESTICK_GRANULARITY_H2 CandlestickGranularity = "H2"
const CANDLESTICK_GRANULARITY_H3 CandlestickGranularity = "H3"
const CANDLESTICK_GRANULARITY_H4 CandlestickGranularity = "H4"
const CANDLESTICK_GRANULARITY_H6 CandlestickGranularity = "H6"
const CANDLESTICK_GRANULARITY_H8 CandlestickGranularity = "H8"
const CANDLESTICK_GRANULARITY_H12 CandlestickGranularity = "H12"
const CANDLESTICK_GRANULARITY_D CandlestickGranularity = "D"
const CANDLESTICK_GRANULARITY_W CandlestickGranularity = "W"
const CANDLESTICK_GRANULARITY_M CandlestickGranularity = "M"

func GetAllCandlestickGranularity() []CandlestickGranularity {
	return []CandlestickGranularity {
		CANDLESTICK_GRANULARITY_S5,
		CANDLESTICK_GRANULARITY_S10,
		CANDLESTICK_GRANULARITY_S15,
		CANDLESTICK_GRANULARITY_S30,
		CANDLESTICK_GRANULARITY_M1,
		CANDLESTICK_GRANULARITY_M2,
		CANDLESTICK_GRANULARITY_M4,
		CANDLESTICK_GRANULARITY_M5,
		CANDLESTICK_GRANULARITY_M10,
		CANDLESTICK_GRANULARITY_M15,
		CANDLESTICK_GRANULARITY_M30,
		CANDLESTICK_GRANULARITY_H1,
		CANDLESTICK_GRANULARITY_H2,
		CANDLESTICK_GRANULARITY_H3,
		CANDLESTICK_GRANULARITY_H4,
		CANDLESTICK_GRANULARITY_H6,
		CANDLESTICK_GRANULARITY_H8,
		CANDLESTICK_GRANULARITY_H12,
		CANDLESTICK_GRANULARITY_D,
		CANDLESTICK_GRANULARITY_W,
		CANDLESTICK_GRANULARITY_M,
	}
}

const INSTRUMENT_TYPE_CURRENCY = "CURRENCY"
const INSTRUMENT_TYPE_CFD = "CFD"
const INSTRUMENT_TYPE_METAL = "METAL"

const ACCEPT_DATETIME_FORMAT_UNIX = "UNIX"
const ACCEPT_DATETIME_FORMAT_RFC3339 = "RFC3999"

const DIRECTION_LONG = "LONG"
const DIRECTION_SHORT = "SHORT"

const GUARANTEED_STOP_LOSS_ORDER_MODE_DISABLED = "DISABLED"
const GUARANTEED_STOP_LOSS_ORDER_MODE_ALLOWED = "ALLOWED"
const GUARANTEED_STOP_LOSS_ORDER_MODE_REQUIRED = "REQUIRED"

const ACCOUNT_FINANCING_MODE_NO_FINANCING = "NO_FINANCING"
const ACCOUNT_FINANCING_MODE_SECOND_BY_SECOND = "SECOND_BY_SECOND"
const ACCOUNT_FINANCING_MODE_DAILY = "DAILY"

const POSITION_AGGREGATION_MODE_ABSOLUTE_SUM = "ABSOLUTE_SUM"
const POSITION_AGGREGATION_MODE_MAXIMAL_SIDE = "MAXIMAL_SIDE"
const POSITION_AGGREGATION_MODE_NET_SUM = "NET_SUM"

//The Order is “Good unTil Cancelled”
const TIME_IN_FORCE_GTC = "GTC"

//The Order is “Good unTil Date” and will be cancelled at the provided time
const TIME_IN_FORCE_GTD = "GTD"

//The Order is “Good For Day” and will be cancelled at 5pm New York time
const TIME_IN_FORCE_GFD = "GFD"

//The Order must be immediately “Filled Or Killed”
const TIME_IN_FORCE_FOK = "FOK"

//The Order must be “Immediately partially filled Or Cancelled”
const TIME_IN_FORCE_IOC = "IOC"

//When the Order is filled, only allow Positions to be opened or extended.
const ORDER_POSITION_FILL_OPEN_ONLY = "OPEN_ONLY"

//When the Order is filled, always fully reduce an existing Position before opening a new Position.
const ORDER_POSITION_FILL_REDUCE_FIRST = "REDUCE_FIRST"

//When the Order is filled, only reduce an existing Position.
const ORDER_POSITION_FILL_REDUCE_ONLY = "REDUCE_ONLY"

//When the Order is filled, use REDUCE_FIRST behaviour for non-client hedging Accounts, and OPEN_ONLY behaviour for client hedging Accounts.
const ORDER_POSITION_FILL_DEFAULT = "DEFAULT"

//The Instrument’s price is tradeable.
const PRICE_STATE_TRADEABLE = "tradeable"

//The Instrument’s price is not tradeable.
const PRICE_STATE_NON_TRADEABLE = "non-tradeable"

// The Instrument of the price is invalid or there is no valid Price for the Instrument.
const PRICE_STATE_INVALID = "invalid"

//	A Market Order
const ORDER_TYPE_MARKET = "MARKET"

//	A Limit Order
const ORDER_TYPE_LIMIT = "LIMIT"

//	A Stop Order
const ORDER_TYPE_STOP = "STOP"

//	A Market-if-touched Order
const ORDER_TYPE_MARKET_IF_TOUCHED = "MARKET_IF_TOUCHED"

//	A Take Profit Order
const ORDER_TYPE_TAKE_PROFIT = "TAKE_PROFIT"

//	A Stop Loss Order
const ORDER_TYPE_STOP_LOSS = "STOP_LOSS"

//	A Trailing Stop Loss Order
const ORDER_TYPE_TRAILING_STOP_LOSS = "TRAILING_STOP_LOSS"

//	A Fixed
const ORDER_TYPE_FIXED_PRICE = "FIXED_PRICE"
