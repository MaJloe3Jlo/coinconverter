package usecase

import (
	"context"

	converter "github.com/MaJloe3Jlo/coinconverter"
	"github.com/MaJloe3Jlo/coinconverter/domain"
	"github.com/shopspring/decimal"
)

type Commission struct {
	service converter.Converter
	percent decimal.Decimal
}

func NewComission(percent float64, service converter.Converter) *Commission {
	return &Commission{
		percent: decimal.NewFromFloat(percent),
		service: service,
	}
}

func (c *Commission) Convert(ctx context.Context, data domain.InputData) (string, error) {
	result, err := c.service.Convert(ctx, data)
	if err != nil {
		return "", err
	}

	resultDecimal, err := decimal.NewFromString(result)
	if err != nil {
		return "", nil
	}

	return resultDecimal.Mul(c.percent).String(), nil
}
