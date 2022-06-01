package db

import (
	"database/sql"
	"sail/internal/store/shared/migrate/mysql"
	"sail/internal/store/shared/migrate/sqlite"
	"time"

	"github.com/jmoiron/sqlx"
)

func Connect(driver, datasource string, maxOpenConnections int) (*DB, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, err
	}
	switch driver {
	case "mysql":
		db.SetMaxIdleConns(0)
	}
	var engine Driver
	var locker Locker

	return &DB{
		conn:   sqlx.NewDb(db, driver),
		lock:   locker,
		driver: engine,
	}, nil
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

func setupDatabase(db *sql.DB, driver string) error {
	switch driver {
	case "mysql":
		return mysql.Migrate(db)
	default:
		return sqlite.Migrate(db)
	}
}
