package postgres

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/domain/profile/entity"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/repositories"
)

var (
	queryFindById = `
		SELECT username, email, img_url from users WHERE id = $1
	`

	queryFindUser = `
		SELECT username, email, img_url from users WHERE username LIKE $1
	`

	queryGetAllUser = `
		SELECT username, email, img_url from users
	`

	queryFindByUsername = `
		SELECT username, email, img_url from users WHERE username = $1
	`

	queryGetFollowerByUserId = `
		SELECT u.username, u.email, u.img_url from users as u LEFT JOIN follow as f ON u.id = f.user_id WHERE f.following_id = $1
	`

	queryGetIdByUsername = `
	SELECT id from users where username = $1
	`
)

type profileRepo struct {
	db *sql.DB
}

func NewProfileRepo(db *sql.DB) repositories.ProfileRepo {
	return &profileRepo{
		db: db,
	}
}

func (r *profileRepo) FindById(ctx context.Context, userId int) (*entity.Profile, error) {
	stmt, err := r.db.Prepare(queryFindById)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userId)

	user := entity.Profile{}

	err = row.Scan(
		&user.Username,
		&user.Email,
		&user.ImgUrl,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *profileRepo) FindUser(ctx context.Context, username string) ([]*entity.Profile, error) {
	stmt, err := r.db.Prepare(queryFindUser)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query("%"+username+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.Profile

	for rows.Next() {
		tempUsers := new(entity.Profile)
		err := rows.Scan(
			&tempUsers.Username,
			&tempUsers.Email,
			&tempUsers.ImgUrl,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, tempUsers)
	}

	return users, nil
}

func (r *profileRepo) GetAllUser(ctx context.Context) ([]*entity.Profile, error) {
	stmt, err := r.db.Prepare(queryGetAllUser)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.Profile

	for rows.Next() {
		tempUser := new(entity.Profile)
		err := rows.Scan(
			&tempUser.Username,
			&tempUser.Email,
			&tempUser.ImgUrl,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, tempUser)
	}

	return users, nil
}

func (r *profileRepo) FindByUsername(ctx context.Context, username string) (*entity.Profile, error) {
	stmt, err := r.db.Prepare(queryFindByUsername)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	user := entity.Profile{}

	err = row.Scan(
		&user.Username,
		&user.Email,
		&user.ImgUrl,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *profileRepo) GetFollower(ctx context.Context, userId int) ([]*entity.Profile, error) {
	stmt, err := r.db.Prepare(queryGetFollowerByUserId)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)

	defer rows.Close()

	var users []*entity.Profile

	for rows.Next() {
		tempUser := new(entity.Profile)
		err := rows.Scan(
			&tempUser.Username,
			&tempUser.Email,
			&tempUser.ImgUrl,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, tempUser)
	}

	return users, nil
}

func(r *profileRepo) GetUserIdByUsername(ctx context.Context, username string) (int, error) {
	stmt, err := r.db.Prepare(queryGetIdByUsername)

	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(username)

	var id int

	err = row.Scan(
		&id,
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}