package postgres

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/domain/auth/entity"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/repositories"
)

var (
	queryCreate = `
		INSERT INTO users (username, email, img_url, password) VALUES($1, $2, $3, $4)
	`

	queryFindByUsername = `
		SELECT id, password from users where username = $1
	`

	queryUpdatePassword = `
		UPDATE users SET password = $1 WHERE id = $2
	`

	queryGetPassword = `
		SELECT password from users WHERE id = $1
	`
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) repositories.AuthRepo {
	return &authRepo{
		db: db,
	}
}

func (a *authRepo) Create(ctx context.Context, user *entity.User) error {
	stmt, err := a.db.Prepare(queryCreate)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.ImgUrl, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (a *authRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	stmt, err := a.db.Prepare(queryFindByUsername)

	if err != nil {
		return nil, err
	}

	rows := stmt.QueryRow(username)

	user := entity.User{}

	err = rows.Scan(
		&user.Id,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *authRepo) UpdatePassword(ctx context.Context, password string, userid int) error {
	stmt, err := a.db.Prepare(queryUpdatePassword)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(password, userid)

	if err != nil {
		return err
	}

	return nil
}

func (a *authRepo) GetPassword(ctx context.Context, userid int) (string, error) {
	stmt, err := a.db.Prepare(queryGetPassword)

	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(userid)

	var password string

	err = row.Scan(
		&password,
	)

	return password, nil
}