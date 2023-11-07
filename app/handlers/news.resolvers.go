package handlers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/app/object"
	"stockels/app/services/news"
)

// GetArticles is the resolver for the getArticles field.
func (r *queryResolver) GetArticles(ctx context.Context) ([]*object.Article, error) {
	return news.GetNewsFromAPI()
}