package postgres

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"database/sql"
	"strconv"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) GetComments() ([]*model.Comment, error) {
	rows, err := r.db.Query("SELECT id, post_id, content, author_id, created_at, updated_at FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (r *CommentRepository) GetCommentByID(id string) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.QueryRow("SELECT id, post_id, content, author_id, created_at, updated_at FROM comments WHERE id = $1", id).
		Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetCommentsByPostID(postID string, pageSize int, offset int, maxDepth int, maxReplies int) ([]*model.Comment, error) {
	query := `
	WITH RECURSIVE
    root_comments AS (SELECT id,
                             post_id,
                             content,
                             author_id,
                             created_at,
                             updated_at,
                             parent_id
                      FROM comments
                      WHERE post_id = $1
                        AND parent_id IS NULL
                      ORDER BY created_at ASC
                      LIMIT $2 OFFSET $3),
    comment_tree AS (SELECT rc.id,
                            rc.post_id,
                            rc.content,
                            rc.author_id,
                            rc.created_at,
                            rc.updated_at,
                            rc.parent_id,
                            1 AS depth
                     FROM root_comments rc
                     UNION ALL
                     SELECT c.id,
                            c.post_id,
                            c.content,
                            c.author_id,
                            c.created_at,
                            c.updated_at,
                            c.parent_id,
                            ct.depth + 1
                     FROM comments c
                              INNER JOIN comment_tree ct ON c.parent_id = ct.id
                     WHERE ct.depth < $4),
    replies_count AS (SELECT parent_id,
                             COUNT(*) AS replies_count
                      FROM comments
                      WHERE parent_id IN (SELECT id FROM comment_tree)
                      GROUP BY parent_id)
	SELECT ct.id,
		   ct.post_id,
		   ct.content,
		   ct.author_id,
		   ct.created_at,
		   ct.updated_at,
		   ct.parent_id,
		   rc.replies_count
	FROM comment_tree ct
			 LEFT JOIN replies_count rc ON ct.id = rc.parent_id
	ORDER BY ct.created_at ASC`
	rows, err := r.db.Query(query, postID, pageSize, offset, maxDepth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentMap := make(map[string]*model.Comment)

	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt, &comment.ParentID, &comment.RepliesCount); err != nil {
			return nil, err
		}
		comment.Replies = []*model.Comment{}
		commentMap[comment.ID] = &comment
	}

	var nestedComments []*model.Comment
	for _, comment := range commentMap {
		if comment.ParentID == nil {
			nestedComments = append(nestedComments, comment)
		} else {
			parentComment, exists := commentMap[*comment.ParentID]
			if exists {
				parentComment.Replies = append(parentComment.Replies, comment)
			}
		}
	}

	return nestedComments, nil
}

func (r *CommentRepository) CreateComment(comment *model.NewComment) (*model.Comment, error) {
	var id int64
	var createdAt string

	query := "INSERT INTO comments (post_id, content, author_id, parent_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at"
	err := r.db.QueryRow(query, comment.PostID, comment.Content, comment.AuthorID, comment.ParentID).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}

	createdComment := &model.Comment{
		ID:        strconv.FormatInt(id, 10),
		PostID:    comment.PostID,
		Content:   comment.Content,
		AuthorID:  comment.AuthorID,
		ParentID:  comment.ParentID,
		CreatedAt: createdAt,
	}

	return createdComment, nil
}
