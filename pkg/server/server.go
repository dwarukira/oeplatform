package server

import (
	"oe/conf"
	"oe/internal/handlers"
	"oe/internal/handlers/middleware"
	"oe/internal/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(config *conf.GlobalConfiguration, orm *models.ORM) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"PUT", "PATCH"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},

		MaxAge: 12 * time.Hour,
	}))

	// r.Use(cors.Default())
	r.Use(middleware.Middleware(config, orm))
	r.Use(middleware.GinContextToContextMiddleware())

	g := r.Group(config.Graphql.Path)
	g.POST("", handlers.GraphqlHandler(orm))
	// logger.Info("GraphQL @ ", gqlPath)
	// Playground handler
	// fmt.Println(config.Graphql.Path.Path, config.Grapgql.EnablePlayground)

	if config.Graphql.EnablePlayground {
		// logger.Info("GraphQL Playground @ ", g.BasePath()+pgqlPath)
		g.GET(config.Graphql.PlaygroundPath, handlers.PlaygroundHandler(g.BasePath()))
	}

	// configCors := cors.DefaultConfig()
	// configCors.AllowAllOrigins = true
	// configCors.AllowCredentials = true
	// configCors.AddAllowHeaders("authorization")
	// configCors.AddAllowHeaders("Authorization")

	// r.Use(cors.New(configCors))

	r.Run()

}
