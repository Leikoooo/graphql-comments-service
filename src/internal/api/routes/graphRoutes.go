package routes

import (
	"OzonTest/src/internal/api/controllers/graph_controller/generated"
	"OzonTest/src/internal/api/controllers/graph_controller/resolvers"
	"OzonTest/src/internal/service"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

func SetupGraphQLRoutes(router *mux.Router, postService *service.PostService, commentService *service.CommentService) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		PostService:    postService,
		CommentService: commentService,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query")).Methods("GET")
	router.Handle("/query", srv).Methods("POST")
}
