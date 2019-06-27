package db

import (
	"fmt"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db DB

type DB struct {
	*sqlx.DB
}

func NewDB(dsn string) error {
	connectedDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db = DB{connectedDB}
	return nil
}

func (db *DB) RunInTx(txFunc func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

func CreateDataSourceName(port int, host, db, user, password string) (string, error) {
	if db == "" {
		return "", errors.New("DataBase Name Must be specified.")
	}
	if port == 0 {
		port = 3306
	}
	if host == "" {
		host = "localhost"
	}
	if user == "" {
		user = "root"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, db), nil
}
