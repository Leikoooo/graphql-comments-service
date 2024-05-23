package graphql

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"errors"
	"github.com/google/uuid"
	"sort"
	"time"
)

type CommentRepository struct {
	comments map[string]*model.Comment
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{comments: make(map[string]*model.Comment)}
}

func (r *CommentRepository) GetCommentsByPostID(postID string, page int, pageSize int, maxDepth int, maxReplies int) ([]*model.Comment, error) {
	if page < 1 || pageSize < 1 {
		return nil, errors.New("invalid page or pageSize")
	}

	var rootComments []*model.Comment
	commentMap := make(map[string][]*model.Comment)

	for _, comment := range r.comments {
		if comment.PostID == postID {
			if *comment.ParentID == "" {
				rootComments = append(rootComments, comment)
			} else {
				commentMap[*comment.ParentID] = append(commentMap[*comment.ParentID], comment)
			}
		}
	}

	sort.Slice(rootComments, func(i, j int) bool {
		t1, _ := time.Parse(time.RFC3339, rootComments[i].CreatedAt)
		t2, _ := time.Parse(time.RFC3339, rootComments[j].CreatedAt)
		return t1.Before(t2)
	})

	startIndex := (page - 1) * pageSize
	if startIndex >= len(rootComments) {
		return []*model.Comment{}, nil
	}
	endIndex := startIndex + pageSize
	if endIndex > len(rootComments) {
		endIndex = len(rootComments)
	}
	paginatedRootComments := rootComments[startIndex:endIndex]

	for _, rootComment := range paginatedRootComments {
		buildCommentTree(rootComment, commentMap, maxDepth, maxReplies, 1)
	}

	return paginatedRootComments, nil
}

func buildCommentTree(comment *model.Comment, commentMap map[string][]*model.Comment, maxDepth int, maxReplies int, currentDepth int) {

	replies, exists := commentMap[comment.ID]
	if !exists {
		return
	}

	sort.Slice(replies, func(i, j int) bool {
		t1, _ := time.Parse(time.RFC3339, replies[i].CreatedAt)
		t2, _ := time.Parse(time.RFC3339, replies[j].CreatedAt)
		return t1.Before(t2)
	})

	if len(replies) > maxReplies {
		replies = replies[:maxReplies]
	}

	for _, reply := range replies {
		buildCommentTree(reply, commentMap, maxDepth, maxReplies, currentDepth+1)
	}

	repliesCount := len(comment.Replies)
	comment.RepliesCount = &repliesCount

	if currentDepth >= maxDepth {
		return
	}

	comment.Replies = []*model.Comment{}
	comment.Replies = replies

}

func (r *CommentRepository) buildTree(parent *model.Comment, commentMap map[string]*model.Comment, maxDepth int, currentDepth int, maxReplies int) {
	if currentDepth >= maxDepth {
		return
	}

	repliesCount := 0
	for _, comment := range commentMap {
		if *comment.ParentID != "" && *comment.ParentID == parent.ID {
			if repliesCount >= maxReplies {
				break
			}
			repliesCount++
			parent.Replies = append(parent.Replies, comment)
			r.buildTree(comment, commentMap, maxDepth, currentDepth+1, maxReplies)
		}
	}
}

func (r *CommentRepository) GetCommentByID(id string) (*model.Comment, error) {
	comment, exists := r.comments[id]
	if !exists {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}

func (r *CommentRepository) CreateComment(comment *model.NewComment) (*model.Comment, error) {
	id := uuid.New().String()
	now := time.Now()

	if *comment.ParentID != "" {
		_, exists := r.comments[*comment.ParentID]
		if !exists {
			return nil, errors.New("no parent Comment")
		}
	}

	var newComment = &model.Comment{
		ID:        id,
		PostID:    comment.PostID,
		Content:   comment.Content,
		AuthorID:  comment.AuthorID,
		ParentID:  comment.ParentID,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	r.comments[id] = newComment
	return newComment, nil
}

func (r *CommentRepository) GetComments() ([]*model.Comment, error) {
	var result []*model.Comment
	for _, post := range r.comments {
		result = append(result, post)
	}
	return result, nil
}
