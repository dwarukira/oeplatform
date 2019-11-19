package middleware

import (
	"context"
	"fmt"
	"oe/internal/models"
	ocontext "oe/pkg/context"

	"github.com/gin-gonic/gin"
)

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(ocontext.ProjectContextKeys.UserCtxKey).(*models.User)
	// raw, _ := c.Request.Context().Value(ocontext.ProjectContextKeys.UserCtxKey).(*models.User)
	return raw
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ocontext.ProjectContextKeys.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ocontext.ProjectContextKeys.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
