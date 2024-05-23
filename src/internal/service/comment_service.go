package service

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"OzonTest/src/internal/repository"
	"errors"
)

type CommentService struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
}

func NewCommentService(commentRepo repository.CommentRepository, postRepo repository.PostRepository) *CommentService {
	return &CommentService{commentRepo, postRepo}
}

func (s *CommentService) GetCommentsByPostID(postID string, page int, pageSize int, maxDepth int, maxReplies int) ([]*model.Comment, error) {
	return s.commentRepo.GetCommentsByPostID(postID, page, pageSize, maxDepth, maxReplies)
}

func (s *CommentService) GetCommentByID(id string) (*model.Comment, error) {
	return s.commentRepo.GetCommentByID(id)
}

func (s *CommentService) CreateComment(comment model.NewComment) (*model.Comment, error) {
	post, err := s.postRepo.GetPostByID(comment.PostID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("no post found")
	}
	if !post.AllowComments {
		return nil, errors.New("comments is not allowed")
	}
	return s.commentRepo.CreateComment(&comment)
}

func (s *CommentService) GetComments() ([]*model.Comment, error) {
	return s.commentRepo.GetComments()
}
