package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/post/entity"
)

var (
	queryCreatePost = `
		INSERT INTO posts (user_id, body, created_at) VALUES($1, $2, $3)
	`

	queryDeletePost = `
		DELETE FROM posts WHERE id = $1 and user_id = $2 and created_at = $3
	`

	queryGetPost = `
		SELECT p.id, u.username, p.body, p.created_at from users as u left join posts as p 
		on u.id = p.user_id where p.id = $1 and u.username = $2 and p.created_at = $3
	`

	queryGetUserPost = `
		SELECT p.id, p.body, p.created_at from users as u left join posts as p
		on u.id = p.user_id where u.username = $1
	`

	queryGetAllPost = `
		SELECT u.username, p.body, p.created_at from users as u left join posts as p
		on u.id = p.user_id
	`
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (p *postRepo) Create(ctx context.Context, post *entity.Post) error {

	stmt, err := p.db.Prepare(queryCreatePost)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(post.UserId, post.Body, post.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (p *postRepo) Delete(ctx context.Context, post *entity.Post) error {
	stmt, err := p.db.Prepare(queryDeletePost)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(post.Id, post.UserId, post.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (p *postRepo) GetPost(ctx context.Context, username string, postId int, createdAt time.Time) (*entity.UserPost, error) {
	stmt, err := p.db.Prepare(queryGetPost)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(postId, username, createdAt)

	post := entity.UserPost{}

	err = row.Scan(
		&post.Id,
		&post.Username,
		&post.Body,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *postRepo) GetUserPosts(ctx context.Context, username string) ([]*entity.UserPost, error) {
	stmt, err := p.db.Prepare(queryGetUserPost)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)

		posts = append(posts, tempPost)
	}

	return posts, nil
}

func (p *postRepo) GetAllPost(ctx context.Context) ([]*entity.UserPost, error) {
	stmt, err := p.db.Prepare(queryGetAllPost)
	
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)

		posts = append(posts, tempPost)
	}

	return posts, nil
}