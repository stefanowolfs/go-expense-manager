package connector

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
	"os"
	"time"
)

var MySQLClient *sqlx.DB

func StartMySQLConn() *errs.AppError {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName))
	if err != nil {
		return errs.NewUnexpectedError("Failed to connect to database " + err.Error())
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	MySQLClient = client
	return nil
}
