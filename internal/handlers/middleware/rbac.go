package middleware

import (
	"context"
	"fmt"
	"oe/internal/models"
	"oe/pkg/logger"

	"github.com/99designs/gqlgen/graphql"
)

func contains(slice []models.Permission, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		logger.Info(s)
		set[s.Tag] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ContainsPermissions(slice []models.Permission, scopes []*string) bool {
	for _, i := range scopes {
		logger.Info(*i, contains(slice, *i))
		if contains(slice, *i) {
			return true
		}
	}

	return false
}

func containsRole(slice []models.Role, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		logger.Info(s)
		set[s.Name] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ContainsPermissionsRole(slice []models.Role, roles []*string) bool {
	for _, i := range roles {
		logger.Info(*i, containsRole(slice, *i))
		if containsRole(slice, *i) {
			return true
		}
	}

	return false
}

func AuthorizeScope(ctx context.Context, obj interface{}, next graphql.Resolver, scopes []*string) (interface{}, error) {
	user := getCurrentUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("Access denied")
	}
	// logger.Info(use, "{{{{{{{{{{{{{{{{{{{{{{{{{{{{{")
	if !ContainsPermissions(user.Permissions, scopes) {
		// block calling the next resolver
		return nil, fmt.Errorf("Access denied")
	}

	// or let it pass through
	return next(ctx)
}

func AuthorizeRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*string) (interface{}, error) {
	user := getCurrentUser(ctx)
	logger.Info(user.Roles, "==================><==========")
	if user == nil {
		return nil, fmt.Errorf("Access denied")
	}
	// logger.Info(use, "{{{{{{{{{{{{{{{{{{{{{{{{{{{{{")
	if user == nil || !ContainsPermissionsRole(user.Roles, roles) {
		// block calling the next resolver
		return nil, fmt.Errorf("Access denied")
	}

	// or let it pass through
	return next(ctx)
}

func AuthorizeAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := getCurrentUser(ctx)

	logger.Info(user, "{{{{{{{{{{{{{{{{{{{{{{{{{{{{{")
	if user == nil {
		// block calling the next resolver
		return nil, fmt.Errorf("Access denied")
	}

	// or let it pass through
	return next(ctx)
}
