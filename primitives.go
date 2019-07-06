package gOanda

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO add comments

type DecimalNumber string

func (d *DecimalNumber) ToFloat() (float64, error) {
	float, err := strconv.ParseFloat(string(*d), 10)
	if err != nil {
		return 0, errors.New("tried to convert '"+ fmt.Sprintf("%.4f",*d)+ "' to a float:"+err.Error())
	}
	return float, nil
}

func (d *DecimalNumber) ToInt() (int64, error) {
	integer, err := strconv.ParseInt(string(*d), 10, 64)
	if err != nil {
		return 0, errors.New("tried to convert '"+fmt.Sprintf("%.4f", *d)+ "' to an int:"+ err.Error())
	}
	return integer, nil
}

type AccountUnits string

func (d *AccountUnits) ToFloat() (float64, error) {
	float, err := strconv.ParseFloat(string(*d), 10)
	if err != nil {
		return 0, errors.New("tried to convert '"+ fmt.Sprintf("%.4f",*d)+ "' to a float:"+ err.Error())
	}
	return float, nil
}

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

func (d DateTime) String() string {return string(d)}

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
