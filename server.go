package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SwArch-2025-1-2A/backend/graph"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

// graphqlHandler initializes and returns a GraphQL handler middleware for Gin.
// It sets up a new GraphQL handler with the following configurations:
//   - Configures basic schema with default resolvers
//   - Enables OPTIONS, GET and POST transport methods
//   - Implements query caching with LRU cache (1000 entries)
//   - Enables GraphQL introspection
//   - Enables Automatic Persisted Queries with LRU cache (100 entries)
//
// Returns a gin.HandlerFunc that serves GraphQL HTTP requests.
func graphqlHandler() gin.HandlerFunc {
	handler := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	handler.AddTransport(transport.Options{})
	handler.AddTransport(transport.GET{})
	handler.AddTransport(transport.POST{})

	handler.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	handler.Use(extension.Introspection{})
	handler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}

}

// playgroundHandler returns a Gin middleware function that serves the GraphQL Playground interface.
// The playground provides an interactive UI for testing GraphQL queries at the endpoint specified in Gin.
// It sets up the playground to send GraphQL operations to the "/query" endpoint.
//
// Returns a gin.HandlerFunc that serves GraphQL HTTP requests.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setup gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
