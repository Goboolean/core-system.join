package adapter

import (
	"github.com/Goboolean/query-server/internal/domain/value"
	"github.com/Goboolean/shared-packages/pkg/mongo"
)



type StockPersistenceAdapter struct {
	db mongo.Queries
}

func (a *StockPersistenceAdapter) FetchStock(stockChan chan<- value.StockAggs, errChan chan error, stock string) {
	mongoStockChan := make(chan mongo.StockAggregate)

	defer close(mongoStockChan)

	go func() {
		select {
		case mongoStock := <-mongoStockChan:
			stock := value.StockAggs{
				EventType: mongoStock.EventType,
				Average: mongoStock.Avg,
				Min: mongoStock.Min,
				Max: mongoStock.Max,
				Start: mongoStock.Start,
				End: mongoStock.End,
				StartTime: mongoStock.StartTime,
				EndTime: mongoStock.EndTime,
			}
			stockChan <- stock

		case _ = <-errChan:
			return
		}
	}()

	if err := a.db.FetchAllStockBatch(mongo.Transaction{}, stock, mongoStockChan); err != nil {
		errChan <- err
		return
	}
}