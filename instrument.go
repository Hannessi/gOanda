package gOanda

// todo json tags

type CandlestickGranularity string

func (c CandlestickGranularity) String () string {return string(c)}

type WeeklyAlignment string

type Candlestick struct {
	Time     DateTime
	Bid      CandlestickData
	Ask      CandlestickData
	Mid      CandlestickData
	Volume   int
	Complete bool
}

type CandlestickData struct {
	O PriceValue
	H PriceValue
	L PriceValue
	C PriceValue
}

type OrderBook struct {
	Instrument  InstrumentName
	Time        DateTime
	Price       PriceValue
	BucketWidth PriceValue
	Buckets     []OrderBookBucket
}

type OrderBookBucket struct {
	Price             PriceValue
	LongCountPercent  DecimalNumber
	ShortCountPercent DecimalNumber
}

type PositionBook struct {
	Instrument  InstrumentName
	Time        DateTime
	Price       PriceValue
	BucketWidth PriceValue
	Buckets     []PositionBookBucket
}

type PositionBookBucket struct {
	Price             PriceValue
	LongCountPercent  DecimalNumber
	ShortCountPercent DecimalNumber
}
