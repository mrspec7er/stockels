package handlers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/app/object"
	"stockels/app/services/stock"
)

// GetStocks is the resolver for the getStocks field.
func (r *queryResolver) GetStocks(ctx context.Context, stocks []*object.GetStockData) ([]*object.StockData, error) {
	return stock.GetMultipleStockService(stocks)
}

// GetStockBySymbol is the resolver for the getStockBySymbol field.
func (r *queryResolver) GetStockBySymbol(ctx context.Context, symbol string, supportPrice int, resistancePrice int) (*object.StockData, error) {
	return stock.GetStockBySymbolService(symbol, supportPrice, resistancePrice)
}

// GetStockDetail is the resolver for the getStockDetail field.
func (r *queryResolver) GetStockDetail(ctx context.Context, symbol string, fromDate string, toDate string, supportPrice int, resistancePrice int) (*object.StockDetail, error) {
	return stock.GetStockDetailService(symbol, fromDate, toDate, supportPrice, resistancePrice)
}
