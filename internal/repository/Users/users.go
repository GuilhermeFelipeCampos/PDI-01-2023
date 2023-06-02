package users

import (
	"context"

	"pdi-go-kafka-bd/internal/database"
	"pdi-go-kafka-bd/internal/entity"
)

type RepositoryUsers struct {
	conn database.DbConnection
}
type RepositoryUsersInterface interface {
	InsertUser(ctx context.Context, u entity.User) (msg []entity.User, err error)
	GetUsers(ctx context.Context) ([]entity.User, error)
}

func CreateRepository(db database.DbConnection) *RepositoryUsers {
	return &RepositoryUsers{
		conn: db,
	}
}

func (repo *RepositoryUsers) InsertUser(ctx context.Context, u entity.User) (*entity.User, error) {

	sql := `INSERT INTO users(name) VALUES ($1) RETURNING id, name;`
	rows, errQuery := repo.conn.Query(ctx, sql, u.Name)

	if errQuery != nil {
		return nil, errQuery
	}
	userReturn := entity.User{}

	for rows.Next() {
		err := rows.Scan(&userReturn.Id, &userReturn.Name)
		if err != nil {
			return nil, err
		}
	}
	return &userReturn, nil

}

func (repo *RepositoryUsers) GetUsers(ctx context.Context) ([]entity.User, error) {
	var us []entity.User
	sql := `SELECT * FROM users;`

	rows, errQuery := repo.conn.Query(ctx, sql)

	if errQuery != nil {
		return nil, errQuery
	}
	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			return nil, err
		}

		us = append(us, u)
	}
	return us, nil
}
