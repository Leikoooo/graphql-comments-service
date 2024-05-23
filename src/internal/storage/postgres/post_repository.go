package postgres

import (
	"OzonTest/src/internal/api/controllers/graph_controller/model"
	"database/sql"
	"strconv"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetPosts() ([]*model.Post, error) {
	rows, err := r.db.Query("SELECT id, title, content, author_id, created_at, updated_at, allow_comments FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &post.AllowComments); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostRepository) GetPostByID(id string) (*model.Post, error) {
	var post model.Post
	err := r.db.QueryRow("SELECT id, title, content, author_id, created_at, updated_at, allow_comments FROM posts WHERE id = $1", id).
		Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &post.AllowComments)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) CreatePost(post *model.NewPost) (*model.Post, error) {
	var id int64
	query := "INSERT INTO posts (title, content, author_id, allow_comments) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, post.Title, post.Content, post.AuthorID, post.AllowComments).Scan(&id)
	if err != nil {
		return nil, err
	}

	createdPost := &model.Post{
		ID:            strconv.FormatInt(id, 10),
		Title:         post.Title,
		Content:       post.Content,
		AuthorID:      post.AuthorID,
		AllowComments: post.AllowComments,
	}

	return createdPost, nil
}
