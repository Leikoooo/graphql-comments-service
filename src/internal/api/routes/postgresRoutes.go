package routes

import (
	"OzonTest/config"
	"OzonTest/src/internal/api/controllers/postgres_controller"
	"OzonTest/src/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupPostgresRoutes(router *mux.Router, postService *service.PostService, commentService *service.CommentService, appConfig config.AppConfig) {
	router.HandleFunc("/savePost", func(w http.ResponseWriter, r *http.Request) {
		postgres_controller.CreatePost(w, r, postService)
	}).Methods("POST")

	router.HandleFunc("/getAllPosts", func(w http.ResponseWriter, r *http.Request) {
		postgres_controller.GetAllPosts(w, r, postService)
	}).Methods("GET")

	router.HandleFunc("/GetPostById", func(w http.ResponseWriter, r *http.Request) {
		postgres_controller.GetPostById(w, r, postService)
	}).Methods("GET")

	router.HandleFunc("/saveComment", func(w http.ResponseWriter, r *http.Request) {
		postgres_controller.CreateComment(w, r, commentService)
	}).Methods("POST")

	router.HandleFunc("/getCommentsByPostId", func(w http.ResponseWriter, r *http.Request) {
		postgres_controller.GetCommentsByPostId(w, r, commentService, appConfig)
	}).Methods("GET")

}
