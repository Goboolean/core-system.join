package value

import "reflect"



var batchSize int = 1000

type StockAggs struct {
	EventType string	
	Average float64
	Min    float64
	Max    float64
	Start  float64
	End    float64

	StartTime int64
	EndTime   int64
}



type StockAggsMassive struct {
	stockCh <-chan StockAggs
	errCh <-chan error
	isFinished bool
}

func (s *StockAggsMassive) Next() interface{} {
	if s.isFinished {
		return struct{}{}
	}

	stockBatch := make([]StockAggs, batchSize)

	for {
		select {
		case err := <- s.errCh:
			if err != nil {
				return err
			} else {
				s.isFinished = true
				return stockBatch
			}

		case stock := <- s.stockCh:
			if !reflect.DeepEqual(stock, StockAggs{}) {
				stockBatch = append(stockBatch, stock)
				if len(stockBatch) == batchSize {
					stockBatch = stockBatch[:0]
					return stockBatch
				}
			}
		}
	}
}

func NewStockAggsMassave(ch chan StockAggs) *StockAggsMassive {
	return &StockAggsMassive{
		stockCh: make(chan StockAggs, batchSize),
		errCh: make(chan error),
	}
}