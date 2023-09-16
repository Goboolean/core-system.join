package out

import "github.com/Goboolean/core-system.join/internal/domain/value"

type StockPort interface {
	GetStockBatch(string) (value.StockAggsMassive, error)
}
