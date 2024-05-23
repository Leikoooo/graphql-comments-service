package resolvers

import (
	"OzonTest/src/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostService    *service.PostService
	CommentService *service.CommentService
}

type QueryResolver struct {
	*Resolver
}

type MutationResolver struct {
	*Resolver
}

func NewResolver(postService *service.PostService, commentService *service.CommentService) *Resolver {
	return &Resolver{
		PostService:    postService,
		CommentService: commentService,
	}
}

//func (r *QueryResolver) GetPosts(ctx context.Context) ([]*model.Post, error) {
//	return r.PostService.GetPosts()
//}
//
//func (r *QueryResolver) GetCommentsByPostID(ctx context.Context, postID string, page int, pageSize int, maxDepth int, maxReplies int) ([]*model.Comment, error) {
//	return r.CommentService.GetCommentsByPostID(postID, page, pageSize, maxDepth, maxReplies)
//}
//
//func (r *MutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
//	return r.PostService.CreatePost(input)
//}
//
//func (r *MutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
//	return r.CommentService.CreateComment(input)
//}
