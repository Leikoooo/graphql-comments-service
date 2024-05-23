package Init

import (
	"OzonTest/config"
	"OzonTest/src/internal/api/routes"
	"OzonTest/src/internal/repository"
	"OzonTest/src/internal/service"
	"OzonTest/src/internal/storage/graphql"
	"OzonTest/src/internal/storage/postgres"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

func initPostgres(DB_HOST string, DB_PORT string, DB_NAME string, DB_PASSWORD string, DB_USER string) (repository.PostRepository, repository.CommentRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)

	db, err := sql.Open("postgres", connStr)
	if errPing := db.Ping(); errPing != nil || err != nil {
		return nil, nil, fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	log.Printf("Connected to postgres database")
	return postgres.NewPostRepository(db), postgres.NewCommentRepository(db), nil
}

func initGraph() (repository.PostRepository, repository.CommentRepository, error) {
	return graphql.NewPostRepository(), graphql.NewCommentRepository(), nil
}

func InitServices(storageType string, appConfig config.AppConfig) (*service.PostService, *service.CommentService, error) {
	var postRepository repository.PostRepository
	var commentRepository repository.CommentRepository
	var err error

	switch storageType {
	case "graph":
		postRepository, commentRepository, err = initGraph()
	case "postgres":
		postRepository, commentRepository, err = initPostgres(appConfig.DB_HOST, appConfig.DB_PORT, appConfig.DB_NAME, appConfig.DB_PASSWORD, appConfig.DB_USER)
	default:
		return nil, nil, fmt.Errorf("unknown storage type: %s", storageType)
	}

	if err != nil {
		return nil, nil, err
	}

	postService := service.NewPostService(postRepository)
	commentService := service.NewCommentService(commentRepository, postRepository)

	return postService, commentService, nil
}

func SetupRoutes(router *mux.Router, postService *service.PostService, commentService *service.CommentService, appConfig config.AppConfig) error {
	switch appConfig.StorageType {
	case "graph":
		routes.SetupGraphQLRoutes(router, postService, commentService)
	case "postgres":
		routes.SetupPostgresRoutes(router, postService, commentService, appConfig)
	default:
		return fmt.Errorf("unknown storage type: %s", appConfig.StorageType)
	}
	return nil
}
