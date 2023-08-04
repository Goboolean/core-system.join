package adapter

import (
	"sync"

	"github.com/Goboolean/query-server/internal/domain/value"
	"github.com/Goboolean/shared-packages/pkg/mongo"
)



type StockPersistenceAdapter struct {
	db mongo.Queries
}

func (a *StockPersistenceAdapter) GetStockBatch(name string) (*value.StockAggsMassive, error) {

	stockChan := make(chan value.StockAggs)
	errChan := make(chan error)

	stock := value.NewStockAggsMassave(stockChan)

	mongoStockChan := make(chan mongo.StockAggregate)

	defer close(mongoStockChan)
	defer close(stockChan)
	defer close(errChan)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		if err := a.db.FetchAllStockBatch(mongo.Transaction{}, name, mongoStockChan); err != nil {
			errChan <- err
			return
		}
		wg.Done()

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

		case <-errChan:
			return
		}
	}()
	wg.Wait()

	return stock, nil
}