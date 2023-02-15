package postgres

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/domain/follow/entity"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/repositories"
)

var (
	queryCreate = `
		INSERT INTO follow (user_id, following_id, created_at) VALUES ($1, $2, $3)
	`

	queryDelete = `
		DELETE FROM follow WHERE user_id = $1 AND following_id = $2
	`
)


type followRepo struct {
	db *sql.DB
}

func NewFollowRepo(db *sql.DB) repositories.FollowRepo {
	return &followRepo{
		db: db,
	}
}

func (f *followRepo) Create(ctx context.Context, req *entity.Follow) error {
	stmt, err := f.db.Prepare(queryCreate)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.UserId, req.FollowingId, req.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (f *followRepo) Delete(ctx context.Context, userId int, followingId int) error {
	stmt, err := f.db.Prepare(queryDelete)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, followingId)

	if err != nil {
		return err
	}

	return nil
}