package definitions

// TODO add comments

type DecimalNumber string

type AccountUnits string

type Currency string

type InstrumentName string

func (i InstrumentName) String() string {return string(i)}

type InstrumentType string

type Instrument struct {
	Name                        InstrumentName       `json:"name"`
	Type                        InstrumentType       `json:"type"`
	DisplayName                 string               `json:"displayName"`
	PipLocation                 int                  `json:"pipLocation"`
	DisplayPrecision            int                  `json:"displayPrecision"`
	TradeUnitsPrecision         int                  `json:"tradeUnitsPrecision"`
	MinimumTradeSize            DecimalNumber        `json:"minimumTradeSize"`
	MaximumTrailingStopDistance DecimalNumber        `json:"maximumTrailingStopDistance"`
	MinimumTrailingStopDistance DecimalNumber        `json:"minimumTrailingStopDistance"`
	MaximumPositionSize         DecimalNumber        `json:"maximumPositionSize"`
	MaximumOrderUnits           DecimalNumber        `json:"maximumOrderUnits"`
	MarginRate                  DecimalNumber        `json:"marginRate"`
	Commission                  InstrumentCommission `json:"commission"`
}

type DateTime string

type AcceptDatetimeFormat string

type InstrumentCommission struct {
	Commission        DecimalNumber `json:"commission"`
	UnitsTraded       DecimalNumber `json:"unitsTraded"`
	MinimumCommission DecimalNumber `json:"minimumCommission"`
}

type GuaranteedStopLossOrderLevelRestriction struct {
	Volume     DecimalNumber `json:"volume"`
	PriceRange DecimalNumber `json:"priceRange"`
}

type Direction string
