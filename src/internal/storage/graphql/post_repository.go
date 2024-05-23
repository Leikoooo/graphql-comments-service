package graphql

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"github.com/google/uuid"
)

type PostRepository struct {
	posts map[string]*model.Post
}

func NewPostRepository() *PostRepository {
	return &PostRepository{posts: make(map[string]*model.Post)}
}

func (r *PostRepository) GetPostByID(id string) (*model.Post, error) {
	return r.posts[id], nil
}

func (r *PostRepository) GetPosts() ([]*model.Post, error) {
	var result []*model.Post
	for _, post := range r.posts {
		result = append(result, post)
	}
	return result, nil
}

func (r *PostRepository) CreatePost(post *model.NewPost) (*model.Post, error) {
	id := uuid.New().String()

	var newPost = &model.Post{
		ID:            id,
		Title:         post.Title,
		Content:       post.Content,
		AuthorID:      post.AuthorID,
		AllowComments: post.AllowComments,
	}

	r.posts[id] = newPost

	return newPost, nil
}
