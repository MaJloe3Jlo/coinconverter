package coinmarket

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MaJloe3Jlo/coinconverter/domain"
)

type Client struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) Convert(ctx context.Context, data domain.InputData) (string, error) {
	req, err := request(ctx, data)
	if err != nil {
		return "", fmt.Errorf("Making request failed. Reason: %s", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request to server failed. Reason: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	var price CoinmarketData
	err = json.Unmarshal(respBody, &price)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}

	if err := price.validate(); err != nil {
		return "", fmt.Errorf("Response failed. Reason: %w", err)
	}

	qu := price.Data[data.Currency].Quote[data.CurrencyToConvert]

	return fmt.Sprintf("%f", qu.Price), nil
}
