package gOanda

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Price struct{
	//
	// The string “PRICE”. Used to identify the a Price object when found in a
	// stream.
	//
	Type string `json:"type"`

	//
	// The Price’s Instrument.
	//
	Instrument InstrumentName `json:"instrument"`

	//
	// The date/time when the Price was created
	//
	Time DateTime `json:"time"`

	//
	// The status of the Price.
	//
	//
	// Deprecated: Will be removed in a future API update.
	//
	Status PriceStatus `json:"status"`

	//
	// Flag indicating if the Price is tradeable or not
	//
	Tradeable bool `json:"tradeable"`

	//
	// The list of prices and liquidity available on the Instrument’s bid side.
	// It is possible for this list to be empty if there is no bid liquidity
	// currently available for the Instrument in the Account.
	//
	Bids []PriceBucket `json:"bids"`

	//
	// The list of prices and liquidity available on the Instrument’s ask side.
	// It is possible for this list to be empty if there is no ask liquidity
	// currently available for the Instrument in the Account.
	//
	Asks []PriceBucket `json:"asks"`

	//
	// The closeout bid Price. This Price is used when a bid is required to
	// closeout a Position (margin closeout or manual) yet there is no bid
	// liquidity. The closeout bid is never used to open a new position.
	//
	CloseoutBid PriceValue `json:"closeoutBid"`

	//
	// The closeout ask Price. This Price is used when a ask is required to
	// closeout a Position (margin closeout or manual) yet there is no ask
	// liquidity. The closeout ask is never used to open a new position.
	//
	CloseoutAsk PriceValue `json:"closeoutAsk"`

	//
	// The factors used to convert quantities of this price’s Instrument’s quote
	// currency into a quantity of the Account’s home currency. When the
	// includeHomeConversions is present in the pricing request (regardless of
	// its value), this field will not be present.
	//
	//
	// Deprecated: Will be removed in a future API update.
	//
	QuoteHomeConversionFactors QuoteHomeConversionFactors `json:"quoteHomeConversionFactors"`

	//
	// Representation of how many units of an Instrument are available to be
	// traded by an Order depending on its positionFill option.
	//
	//
	// Deprecated: Will be removed in a future API update.
	//
	UnitsAvailable UnitsAvailable `json:"unitsAvailable"`
}

type PriceValue string

func (p PriceValue) ToFloat() float64 {
	float, err := strconv.ParseFloat(string(p), 10)
	if err != nil {
		log.Error("tried to convert '", p, "' to a float:", err)
		return 0
	}
	return float
}

func (p PriceValue) String() string {return string(p)}

type PriceBucket struct{
	//
	// The Price offered by the PriceBucket
	//
	Price PriceValue `json:"price"`

	//
	// The amount of liquidity offered by the PriceBucket
	//
	Liquidity int `json:"liquidity"`
}

type PriceStatus string

type QuoteHomeConversionFactors struct{
	//
	// The factor used to convert a positive amount of the Price’s Instrument’s
	// quote currency into a positive amount of the Account’s home currency.
	// Conversion is performed by multiplying the quote units by the conversion
	// factor.
	//
	PositiveUnits DecimalNumber `json:"positiveUnits"`

	//
	// The factor used to convert a negative amount of the Price’s Instrument’s
	// quote currency into a negative amount of the Account’s home currency.
	// Conversion is performed by multiplying the quote units by the conversion
	// factor.
	//
	NegativeUnits DecimalNumber `json:"negativeUnits"`
}

type HomeConversions struct{
	//
	// The currency to be converted into the home currency.
	//
	Currency Currency `json:"currency"`

	//
	// The factor used to convert any gains for an Account in the specified
	// currency into the Account’s home currency. This would include positive
	// realized P/L and positive financing amounts. Conversion is performed by
	// multiplying the positive P/L by the conversion factor.
	//
	AccountGain DecimalNumber `json:"accountGain"`

	//
	// The string representation of a decimal number.
	//
	AccountLoss DecimalNumber `json:"accountLoss"`

	//
	// The factor used to convert a Position or Trade Value in the specified
	// currency into the Account’s home currency. Conversion is performed by
	// multiplying the Position or Trade Value by the conversion factor.
	//
	PositionValue DecimalNumber `json:"positionValue"`
}

type ClientPrice struct{} // todo

type PricingHeartbeat struct{} // todo
