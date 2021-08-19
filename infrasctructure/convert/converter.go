package convert

import (
	"context"
	"fmt"

	"github.com/MaJloe3Jlo/coinconverter/domain"
)

type Client interface {
	Convert(ctx context.Context, data domain.InputData) (string, error)
}

type Service struct {
	client Client
}

func New(client Client) *Service {
	return &Service{client: client}
}

func (s *Service) Convert(ctx context.Context, data domain.InputData) (string, error) {
	if data.Currency == data.CurrencyToConvert {
		return data.Amount, nil
	}

	result, err := s.client.Convert(ctx, data)
	if err != nil {
		return "", fmt.Errorf("Convert: %w", err)
	}

	return result, nil
}
