package coinmarket

import (
	"errors"
)

type CoinmarketData struct {
	Status map[string]interface{} `json:"status"`
	Data   map[string]data        `json:"data"`
}

type status struct {
	ErrorMessage string `json:"error_message"`
}
type data struct {
	Symbol string           `json:""`
	Quote  map[string]quote `json:""`
}
type quote struct {
	Price float64 `json:"price"`
}

func (c *CoinmarketData) validate() error {
	if c.Status == nil {
		return errors.New("Cannot find status in response")
	}
	if c.Data == nil {
		return errors.New("Cannot find data in response")
	}
	return nil
}
