package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/app"
	"stockels/app/middleware"
	"stockels/app/module/subscribtion"
	"stockels/app/object"
)

// StockSubscribes is the resolver for the stockSubscribes field.
func (r *mutationResolver) StockSubscribes(ctx context.Context, stocks []*object.GetStockData) ([]*object.Subscribtion, error) {
	user, err := middleware.GetAuthContextMiddleware(ctx)
	if err != nil {
		return nil, err
	}
	return subscribtion.SubscribeMultipleStockService(stocks, user)
}

// GetStockSubscribe is the resolver for the getStockSubscribe field.
func (r *queryResolver) GetStockSubscribe(ctx context.Context) ([]*object.StockData, error) {
	user, err := middleware.GetAuthContextMiddleware(ctx)
	if err != nil {
		return nil, err
	}
	return subscribtion.GetSubscribtionStockService(*user)
}

// GenerateReportFile is the resolver for the generateReportFile field.
func (r *queryResolver) GenerateReportFile(ctx context.Context) (*object.GenerateReportResponse, error) {
	user, err := middleware.GetAuthContextMiddleware(ctx)
	if err != nil {
		return nil, err
	}
	return subscribtion.GenerateStockReportService(*user)
}

// Mutation returns app.MutationResolver implementation.
func (r *Resolver) Mutation() app.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
