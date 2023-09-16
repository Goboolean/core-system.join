package inport

import (
	"github.com/Goboolean/core-system.join/internal/domain/value"
)

type StockPort interface {
	GetStockMassive(string) (value.StockAggsMassive, error)
}
