package repositories

import (
	"awesomeProject3/internal/models"
	"database/sql"
	"fmt"
)

type PostRepository interface {
	GetPosts() (*[]models.Post, error)
	GetPost(id int) (*models.Post, error)
	StorePost(post models.Post) (*models.Post, error)
	UpdatePost(post models.Post, id int) (*models.Post, error)
	DeletePost(id int) error
}

type PostRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &PostRepositoryImpl{
		DB: db,
	}
}

func (r *PostRepositoryImpl) GetPosts() (*[]models.Post, error) {
	rows, err := r.DB.Query("SELECT id, title, body FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Body); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func (r *PostRepositoryImpl) GetPost(id int) (*models.Post, error) {
	query := `SELECT id, title, body FROM posts WHERE id = ?`

	row := r.DB.QueryRow(query, id)

	var post models.Post
	err := row.Scan(&post.Id, &post.Title, &post.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("пост с ID %d не найден", id)
		}
		return nil, err
	}

	return &post, nil
}

func (r *PostRepositoryImpl) StorePost(post models.Post) (*models.Post, error) {
	query := `INSERT INTO posts (title, body) VALUES (?, ?)`

	result, err := r.DB.Exec(query, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	post.Id = int(postId)

	return &post, nil
}

func (r *PostRepositoryImpl) UpdatePost(post models.Post, id int) (*models.Post, error) {
	query := `UPDATE posts SET title = ?, body = ? WHERE id = ?`

	_, err := r.DB.Exec(query, &post.Title, &post.Body, &id)
	if err != nil {
		return nil, err
	}

	post.Id = id

	return &post, nil
}

func (r *PostRepositoryImpl) DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id = ?`

	_, err := r.DB.Exec(query, &id)
	if err != nil {
		return err
	}

	return nil
}
