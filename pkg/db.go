package pkg

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/nqvinh00/CakeAssignment/model"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(config model.Database) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", config.Username, config.Password, config.Host, config.DB, config.Args)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Duration(config.MaxConnLifeTime) * time.Minute)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	return
}
