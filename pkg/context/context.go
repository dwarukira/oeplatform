package context

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ContextKeys holds the context keys throught the project
type ContextKeys struct {
	UserCtxKey    ContextKey // User db object in Auth
	GinContextKey ContextKey
}

type contextKey string

func (c contextKey) String() string {
	return "api context key " + string(c)
}

var (
	// ProjectContextKeys the project's context keys
	ProjectContextKeys = ContextKeys{
		UserCtxKey:    "auth_user",
		GinContextKey: "GinContextKey",
	}
)
