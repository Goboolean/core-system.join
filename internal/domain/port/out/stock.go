package out

import "github.com/Goboolean/query-server/internal/domain/value"



type StockPort interface {
	GetStockBatch(string) (chan value.StockAggs, chan error)
}