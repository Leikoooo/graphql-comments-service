package service

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"OzonTest/src/internal/repository"
)

type PostService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) *PostService {
	return &PostService{postRepo}
}

func (s *PostService) GetPosts() ([]*model.Post, error) {
	return s.postRepo.GetPosts()
}

func (s *PostService) GetPostByID(id string) (*model.Post, error) {
	return s.postRepo.GetPostByID(id)
}

func (s *PostService) CreatePost(newPost model.NewPost) (*model.Post, error) {
	return s.postRepo.CreatePost(&newPost)
}
