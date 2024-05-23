package postgres_controller

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"OzonTest/src/internal/service"
	"OzonTest/src/internal/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var CreatePost = func(w http.ResponseWriter, r *http.Request, postService *service.PostService) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newPost model.NewPost
	err = json.Unmarshal(bodyBytes, &newPost)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	var post *model.Post
	post, err = postService.CreatePost(newPost)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.Response(w, post, http.StatusOK)
	}
}

var GetPostById = func(w http.ResponseWriter, r *http.Request, postService *service.PostService) {
	postId := r.URL.Query().Get("id")
	result, err := postService.GetPostByID(postId)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.Response(w, result, http.StatusOK)
	}
}

var GetAllPosts = func(w http.ResponseWriter, r *http.Request, postService *service.PostService) {
	result, err := postService.GetPosts()
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.Response(w, result, http.StatusOK)
	}
}
