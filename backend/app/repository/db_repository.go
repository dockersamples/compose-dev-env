package repository

import (
	"aquafarm-management/app/config"
	"database/sql"
	"fmt"
	"log"
)

// DbRepository menangani akses data item
type DbRepository struct {
	DB *sql.DB
}

func NewDB(cfg *config.Config) *DbRepository {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	fmt.Println("Test Koneksi : " + cfg.DB.Host + " | " + cfg.DB.Name + " | " + cfg.DB.User + " | " + cfg.DB.Password + " | ")

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)

	return &DbRepository{
		DB: db,
	}
}
