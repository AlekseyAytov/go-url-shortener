package db

import (
	"context"
	"database/sql"

	"github.com/AlekseyAytov/go-url-shortener/internal/urlobject"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbStorage struct {
	db *sql.DB
}

func NewDbStorage(dbDSN string) (*DbStorage, error) {
	db, err := sql.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	result := &DbStorage{db: db}
	err = result.CheckDB()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DbStorage) CheckDB() error {
	err := d.db.PingContext(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (d *DbStorage) SaveObject(urlobject.URLObject) error {
	return nil
}

func (d *DbStorage) ReadObjects() ([]urlobject.URLObject, error) {
	return nil, nil
}
