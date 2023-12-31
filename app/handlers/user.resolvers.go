package handlers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"stockels/app/object"
	"stockels/app/services/user"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, payload *object.Register) (*object.User, error) {
	return user.CreateUserService(payload)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*object.LoginResponse, error) {
	return user.LoginService(email, password)
}
