package adapter

import (
	"fmt"

	inport "github.com/Goboolean/query-server/internal/domain/port/in"
	"github.com/Goboolean/query-server/internal/domain/value"

	pb "github.com/Goboolean/query-server/api/grpc"
)

type StockFetcherAdapter struct {
	port inport.StockPort
}


func (a *StockFetcherAdapter) FetchStockAggs(in *pb.StockFetchRequest, stream pb.StockFetcher_FetchStockAggsServer) error {
	stockMass, err := a.port.GetStockMassive(in.StockName)
	if err != nil {
		return err
	}


	for {
		data := stockMass.Next()

		if err, ok := data.(error); ok {
			return err
		}

		if stockBatch, ok := data.([]value.StockAggs); ok {
			newStockBatch := make([]*pb.StockAggregate, len(stockBatch))

			for idx := range stockBatch {
				newStock := &pb.StockAggregate{
					EventType: stockBatch[idx].EventType,
					Average: float32(stockBatch[idx].Average),
					Min: float32(stockBatch[idx].Min),
					Max: float32(stockBatch[idx].Max),
					Start: float32(stockBatch[idx].Start),
					End: float32(stockBatch[idx].End),
					StartTime: stockBatch[idx].StartTime,
					EndTime: stockBatch[idx].EndTime,
				}
	
				newStockBatch[idx] = newStock
			}
	
			if err := stream.Send(&pb.StockFetchResponse{Stock: newStockBatch}); err != nil {
				return err
			}
		}

		if _, ok := data.(struct{}); ok {
			return nil
		}

		return fmt.Errorf("error: unexpected data type: %v", data)
	}
}
