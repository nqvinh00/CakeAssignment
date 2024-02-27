package pkg

import (
	"fmt"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(config model.Database) (gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", config.Username, config.Password, config.Host, config.DB, config.Args)
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return
	}

	db, err := gormDB.DB()
	db.SetConnMaxLifetime(time.Duration(config.MaxConnLifeTime) * time.Minute)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.Close()
	return
}
