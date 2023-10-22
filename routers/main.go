package routers

import (
	"stockels/graph"
	"stockels/graph/resolver"
	"stockels/middleware"
	"stockels/module/stock"
	"stockels/module/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

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
	stock.Routes(router)
	user.Routes(router)
	router.Use(middleware.GinContextToContextMiddleware())
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run()
}