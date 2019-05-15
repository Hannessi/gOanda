package gOanda

const INSTRUMENT_NAME_EURUSD = "EUR_USD"

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
