package postgres_controller

import (
	"OzonTest/config"
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"OzonTest/src/internal/service"
	"OzonTest/src/internal/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var CreateComment = func(w http.ResponseWriter, r *http.Request, commentService *service.CommentService) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newComment model.NewComment
	err = json.Unmarshal(bodyBytes, &newComment)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ограничение на 2000 символов в комментарии
	if len(newComment.Content) > 2000 {
		utils.Response(w, errors.New("the length of the comment must be less than 2000").Error(), 400)
	}

	var comment *model.Comment
	comment, err = commentService.CreateComment(newComment)

	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.Response(w, comment, http.StatusOK)
	}

}

var GetCommentsByPostId = func(w http.ResponseWriter, r *http.Request, commentService *service.CommentService, appConfig config.AppConfig) {
	page := utils.GetPage(r)
	pageSize := utils.GetPageSize(r, appConfig)
	offset := utils.GetPageOffset(page, pageSize)

	postID := r.URL.Query().Get("id")

	comments, err := commentService.GetCommentsByPostID(postID, pageSize, offset, appConfig.MaxReplyDepth, appConfig.MaxReplies)
	if err != nil {
		utils.Response(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
		"offset":   offset,
		"comments": comments,
	}

	utils.Response(w, response, 200)
}
