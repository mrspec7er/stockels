package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/app"
	"stockels/app/module/analytic"
	"stockels/app/object"
)

// GetDetailAnalytic is the resolver for the getDetailAnalytic field.
func (r *queryResolver) GetDetailAnalytic(ctx context.Context, symbol *string) (*object.Analytic, error) {
	return analytic.GetAnalyticFromAPI(*symbol)
}

// Query returns app.QueryResolver implementation.
func (r *Resolver) Query() app.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
