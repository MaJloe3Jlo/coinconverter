package coinmarket

import (
	"context"
	"net/http"
	"net/url"

	"github.com/MaJloe3Jlo/coinconverter/domain"
)

const (
	host   = "sandbox-api.coinmarketcap.com"
	apiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
	path   = "/v1/tools/price-conversion"
)

func request(ctx context.Context, data domain.InputData) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://"+host+path, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("amount", data.Amount)
	q.Add("symbol", data.Currency)
	q.Add("convert", data.CurrencyToConvert)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	return req, nil
}
