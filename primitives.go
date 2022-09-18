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
		return 0, errors.New("tried to convert '" + fmt.Sprintf("%.4f", *d) + "' to a float:" + err.Error())
	}
	return float, nil
}

func (d *DecimalNumber) ToInt() (int64, error) {
	integer, err := strconv.ParseInt(string(*d), 10, 64)
	if err != nil {
		return 0, errors.New("tried to convert '" + fmt.Sprintf("%.4f", *d) + "' to an int:" + err.Error())
	}
	return integer, nil
}

type AccountUnits string

func (d *AccountUnits) ToFloat() (float64, error) {
	float, err := strconv.ParseFloat(string(*d), 10)
	if err != nil {
		return 0, errors.New("tried to convert '" + fmt.Sprintf("%.4f", *d) + "' to a float:" + err.Error())
	}
	return float, nil
}

type Currency string

type InstrumentName string

func (i InstrumentName) String() string { return string(i) }

func (i InstrumentName) IsValid() bool {
	return contains(i, GetAllInstrumentNames())
}

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

func (d DateTime) String() string { return string(d) }

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

type ConversionFactor struct {
	// The factor by which to multiply the amount in the given currency to
	// obtain the amount in the home currency of the Account.
	Factor DecimalNumber `json:"factor"`
}

type HomeConversionFactors struct {
	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument quote units into units of the Account’s home
	// currency.
	GainQuoteHome ConversionFactor `json:"gainQuoteHome"`

	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument quote units into units of the Account’s home
	// currency.
	LossQuoteHome ConversionFactor `json:"lossQuoteHome"`

	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument base units into units of the Account’s home
	// currency.
	GainBaseHome ConversionFactor `json:"gainBaseHome"`

	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument base units into units of the Account’s home
	// currency.
	LossBaseHome ConversionFactor `json:"lossBaseHome"`
}

func contains(value InstrumentName, list []InstrumentName) bool {
	for _, listValue := range list {
		if value == listValue {
			return true
		}
	}
	return false
}
