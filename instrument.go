package gOanda

type CandlestickGranularity string

func (c CandlestickGranularity) String() string { return string(c) }

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

func GetAllCandlestickGranularities() []CandlestickGranularity {
	return []CandlestickGranularity{
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

type WeeklyAlignment string

const WEEKLY_ALIGNMENT_MONDAY WeeklyAlignment = "Monday"
const WEEKLY_ALIGNMENT_TUESDAY WeeklyAlignment = "Tuesday"
const WEEKLY_ALIGNMENT_WEDNESDAY WeeklyAlignment = "Wednesday"
const WEEKLY_ALIGNMENT_THURSDAY WeeklyAlignment = "Thursday"
const WEEKLY_ALIGNMENT_FRIDAY WeeklyAlignment = "Friday"
const WEEKLY_ALIGNMENT_SATURDAY WeeklyAlignment = "Saturday"
const WEEKLY_ALIGNMENT_SUNDAY WeeklyAlignment = "Sunday"

type Candlestick struct {
	// The start time of the candlestick
	Time     DateTime        `json:"time"`
	Bid      CandlestickData `json:"bid"`
	Ask      CandlestickData `json:"ask"`
	Mid      CandlestickData `json:"mid"`
	Volume   int             `json:"volume"`
	Complete bool            `json:"complete"`
}

type CandlestickData struct {
	O PriceValue `json:"o"`
	H PriceValue `json:"h"`
	L PriceValue `json:"l"`
	C PriceValue `json:"c"`
}

// CandlestickResponse	Response containing instrument, granularity, and list of candles.
type CandlestickResponse struct {
	//# The instrument whose Prices are represented by the candlesticks.
	Instrument InstrumentName `json:"instrument"`

	//# The granularity of the candlesticks provided.
	Granularity CandlestickGranularity `json:"granularity"`

	//# The list of candlesticks that satisfy the request.
	Candles []Candlestick `json:"candles"`
}

type OrderBook struct {
	Instrument  InstrumentName    `json:"instrument"`
	Time        DateTime          `json:"time"`
	Price       PriceValue        `json:"price"`
	BucketWidth PriceValue        `json:"bucketWidth"`
	Buckets     []OrderBookBucket `json:"buckets"`
}

type OrderBookBucket struct {
	Price             PriceValue    `json:"price"`
	LongCountPercent  DecimalNumber `json:"LongCountPercent"`
	ShortCountPercent DecimalNumber `json:"shortCountPercent"`
}

type PositionBook struct {
	Instrument  InstrumentName       `json:"instrument"`
	Time        DateTime             `json:"time"`
	Price       PriceValue           `json:"price"`
	BucketWidth PriceValue           `json:"bucketWidth"`
	Buckets     []PositionBookBucket `json:"buckets"`
}

type PositionBookBucket struct {
	Price             PriceValue    `json:"price"`
	LongCountPercent  DecimalNumber `json:"longCountPercent"`
	ShortCountPercent DecimalNumber `json:"shortCountPercent"`
}
