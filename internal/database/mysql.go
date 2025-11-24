package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("MariaDB connected successfully!")
	return db, nil
}
