package repository

import (
	"aquafarm-management/app/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// ItemRepository mengandung informasi database dan mengandung metode-metode yang dibutuhkan untuk melakukan CRUD pada database
type ItemRepository struct {
	DB *sql.DB
}
type Repository interface {
	Fetch() ([]*model.Item, error)
	Get(id int) (*model.Item, error)
	Create(item *model.Item) error
	Update(item *model.Item) error
	Delete(id int) error
}

// NewDB membuat koneksi baru pada database
//func NewDB(driverName, cfg *config.Config) (DB, error) {
//	dataSourceName := ""
//	db, err := sql.Open(driverName, dataSourceName)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Ping(); err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}

// NewItemRepository membuat item repository baru
func NewItemRepository(db *DbRepository) *ItemRepository {
	return &ItemRepository{DB: db.DB}
}

// Fetch mengambil semua data pada tabel item
func (r *ItemRepository) Fetch() ([]model.Item, error) {
	rows, err := r.DB.Query("SELECT id, name, price, created_at, updated_at FROM items")
	if err != nil {
		return nil, fmt.Errorf("error querying items: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error fetching rows: %w", err)
	}

	return items, nil
}

// Get mengambil data item berdasarkan id
func (r *ItemRepository) Get(id string) (*model.Item, error) {
	row := r.DB.QueryRow("SELECT id, name, price, created_at, updated_at FROM items WHERE id = ?", id)

	var item model.Item
	if err := row.Scan(&item.ID, &item.Name, &item.Price, &item.CreatedAt, &item.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("model Item with id %d not found", id)
		}
		return nil, err
	}

	return &item, nil
}

// Store menyimpan data item baru pada tabel item
func (r *ItemRepository) Store(item *model.Item) (*model.Item, error) {
	// Buat perintah SQL untuk menyimpan data item baru
	query := `
		INSERT INTO items (name, price, quantity)
		VALUES (?, ?, ?)
	`

	// Jalankan perintah SQL
	result, err := r.DB.Exec(query, item.Name, item.Price, item.Quantity)
	if err != nil {
		return nil, err
	}

	// Ambil ID dari item baru yang disimpan
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Update ID pada struct item
	item.ID = strconv.FormatInt(id, 10)

	return item, nil
}

// Update updates an existing item in the database
func (ir *ItemRepository) Update(item *model.Item) error {
	query := `UPDATE items SET name=?, created_at=? WHERE id=?`
	stmt, err := ir.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(item.Name, item.CreatedAt, item.ID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("item not found")
	}
	return nil
}

// Delete deletes an existing item from the database
func (ir *ItemRepository) Delete(id string) error {
	query := `DELETE FROM items WHERE id=?`
	stmt, err := ir.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("item not found")
	}
	return nil
}
