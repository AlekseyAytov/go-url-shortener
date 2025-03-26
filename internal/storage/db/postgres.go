package db

import (
	"context"
	"database/sql"

	"github.com/AlekseyAytov/go-url-shortener/internal/urlobject"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBStorage struct {
	db *sql.DB
}

func NewDBStorage(dbDSN string) (*DBStorage, error) {
	db, err := sql.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	result := &DBStorage{db: db}
	err = result.CheckDB()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DBStorage) CheckDB() error {
	err := d.db.PingContext(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (d *DBStorage) SaveObject(urlobject.URLObject) error {
	return nil
}

func (d *DBStorage) ReadObjects() ([]urlobject.URLObject, error) {
	return nil, nil
}
