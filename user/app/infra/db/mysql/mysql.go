package mysql

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type MySQL struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func NewDB() (*MySQL, error) {

	master, err := sqlx.Open("mysql", masterDSN())
	if err != nil {
		return nil, xerrors.Errorf("sqlx open error: %w", err)
	}

	slave, err := sqlx.Open("mysql", masterDSN()) // TODO: Prepare a DSN for slave
	if err != nil {
		return nil, xerrors.Errorf("sqlx open error: %w", err)
	}

	return &MySQL{
		Master: master,
		Slave:  slave,
	}, nil
}

func masterDSN() string {
	cfg := mysql.NewConfig()

	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASS")
	cfg.DBName = os.Getenv("MYSQL_DB_NAME")
	cfg.ParseTime = true

	return cfg.FormatDSN()
}
