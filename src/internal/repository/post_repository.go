package repository

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
)

type PostRepository interface {
	GetPosts() ([]*model.Post, error)
	GetPostByID(id string) (*model.Post, error)
	CreatePost(post *model.NewPost) (*model.Post, error)
}
