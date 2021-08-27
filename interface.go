package converter

import (
	"context"

	"github.com/MaJloe3Jlo/coinconverter/domain"
)

type Converter interface {
	Convert(ctx context.Context, data domain.InputData) (string, error)
}
