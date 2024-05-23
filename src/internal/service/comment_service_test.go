// comment_service_test.go
package service

import (
	"testing"

	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for testing
type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) GetCommentsByPostID(postID string, page, pageSize, maxDepth, maxReplies int) ([]*model.Comment, error) {
	args := m.Called(postID, page, pageSize, maxDepth, maxReplies)
	return args.Get(0).([]*model.Comment), args.Error(1)
}

func (m *MockCommentRepository) GetCommentByID(id string) (*model.Comment, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Comment), args.Error(1)
}

func (m *MockCommentRepository) CreateComment(comment *model.NewComment) (*model.Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(*model.Comment), args.Error(1)
}

func (m *MockCommentRepository) GetComments() ([]*model.Comment, error) {
	args := m.Called()
	return args.Get(0).([]*model.Comment), args.Error(1)
}

func TestGetCommentsByPostID(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	mockRepo2 := new(MockPostRepository)
	mockRepo.On("GetCommentsByPostID", "1", 1, 10, 2, 5).Return([]*model.Comment{}, nil)

	service := NewCommentService(mockRepo, mockRepo2)

	comments, err := service.GetCommentsByPostID("1", 1, 10, 2, 5)

	assert.NoError(t, err)
	assert.NotNil(t, comments)
	mockRepo.AssertExpectations(t)
}

func TestGetCommentByID(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	mockRepo2 := new(MockPostRepository)
	mockRepo.On("GetCommentByID", "1").Return(&model.Comment{ID: "1"}, nil)

	service := NewCommentService(mockRepo, mockRepo2)

	comment, err := service.GetCommentByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, "1", comment.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateComment(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	mockRepo2 := new(MockPostRepository)
	newComment := model.NewComment{Content: "New Comment"}
	mockRepo.On("CreateComment", &newComment).Return(&model.Comment{ID: "1", Content: "New Comment"}, nil)

	service := NewCommentService(mockRepo, mockRepo2)

	createdComment, err := service.CreateComment(newComment)

	assert.NoError(t, err)
	assert.NotNil(t, createdComment)
	assert.Equal(t, "1", createdComment.ID)
	mockRepo.AssertExpectations(t)
}
