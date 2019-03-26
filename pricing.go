package gOanda

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Price struct{} // todo

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

type PriceBucket struct{} // todo

type PriceStatus string // todo const

type QuoteHomeConversionFactors struct{} // todo

type HomeConversions struct{} // todo

type ClientPrice struct{} // todo

type PricingHeartbeat struct{} // todo
