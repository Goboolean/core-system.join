package inport

import (
	"github.com/Goboolean/query-server/internal/domain/value"
)



type StockPort interface {
	GetStockMassive(string) (value.StockAggsMassive, error)
}