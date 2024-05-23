package repository

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
)

type CommentRepository interface {
	GetCommentsByPostID(postID string, page int, pageSize int, maxDepth int, maxReplies int) ([]*model.Comment, error)
	GetCommentByID(id string) (*model.Comment, error)
	CreateComment(comment *model.NewComment) (*model.Comment, error)
	GetComments() ([]*model.Comment, error)
}
