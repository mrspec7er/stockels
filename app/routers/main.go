package routers

import (
	"stockels/app"
	"stockels/app/handlers"
	"stockels/app/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(app.NewExecutableSchema(app.Config{Resolvers: &handlers.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Config()  {
	router := gin.Default()

	router.Use(middlewares.AuthContextMiddleware())
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	handlers.Routes(router)

	router.Run()
}