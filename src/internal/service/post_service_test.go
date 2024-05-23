// post_service_test.go
package service

import (
	"testing"

	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for testing
type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) GetPosts() ([]*model.Post, error) {
	args := m.Called()
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (m *MockPostRepository) GetPostByID(id string) (*model.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockPostRepository) CreatePost(newPost *model.NewPost) (*model.Post, error) {
	args := m.Called(newPost)
	return args.Get(0).(*model.Post), args.Error(1)
}

func TestGetPosts(t *testing.T) {
	mockRepo := new(MockPostRepository)
	mockRepo.On("GetPosts").Return([]*model.Post{}, nil)

	service := NewPostService(mockRepo)

	posts, err := service.GetPosts()

	assert.NoError(t, err)
	assert.NotNil(t, posts)
	mockRepo.AssertExpectations(t)
}

func TestGetPostByID(t *testing.T) {
	mockRepo := new(MockPostRepository)
	mockRepo.On("GetPostByID", "1").Return(&model.Post{ID: "1"}, nil)

	service := NewPostService(mockRepo)

	post, err := service.GetPostByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "1", post.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(MockPostRepository)
	newPost := model.NewPost{Title: "New Post"}
	mockRepo.On("CreatePost", &newPost).Return(&model.Post{ID: "1", Title: "New Post"}, nil)

	service := NewPostService(mockRepo)

	createdPost, err := service.CreatePost(newPost)

	assert.NoError(t, err)
	assert.NotNil(t, createdPost)
	assert.Equal(t, "1", createdPost.ID)
	mockRepo.AssertExpectations(t)
}
