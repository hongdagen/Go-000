package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql?charset=utf8")
}

type Result struct {
	Host     string
	Username string
}

func DoDao() (int, error) {
	var info Result
	errScan := db.QueryRow("select host,username from mysql.servers limit 1").Scan(&info.Host, &info.Username)
	if errScan != nil {
		return -1, errors.Wrap(errScan, "Scan error")
	}

	defer db.Close()
	return 1, nil
}
