package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stefanowolf/go-expense-manager/internal/core/domain"
	"github.com/stefanowolf/go-expense-manager/internal/infra/connector"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
	"github.com/stefanowolf/go-expense-manager/internal/infra/logger"
	"strconv"
)

var dbClient = connector.MySQLClient

type UserRepositoryDB struct {
	client *sqlx.DB
}

func (d UserRepositoryDB) FindAll() ([]domain.User, *errs.AppError) {
	findAllSQL := "select id, name, created_at from users"
	users := make([]domain.User, 0)
	err := d.client.Select(&users, findAllSQL)
	if err != nil {
		logger.Error("error while querying user table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return users, nil
}

func (d UserRepositoryDB) Save(user domain.User) (*domain.User, *errs.AppError) {
	sqlInsert := "insert into users (name) values (?)"

	result, err := d.client.Exec(sqlInsert, user.Name)
	if err != nil {
		logger.Error("error while inserting user " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last inserted id " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	user.Id = strconv.FormatInt(id, 10)
	return &user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{dbClient}
}
