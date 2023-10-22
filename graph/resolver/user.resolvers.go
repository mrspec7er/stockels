package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/graph/service"
)

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string) (string, error) {
	token, err := service.Login(email, password)
	return token, err
}
