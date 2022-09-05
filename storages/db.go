package storages

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"gitlab.com/danmory/web-hashing-server/tools"
)

type databaseStorage struct {
	conn *pgx.Conn
}

func createDBStorage() *databaseStorage {
	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf(
			"user=%v password=%v host=%v port=%v dbname=%v",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE")))
	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to database")
	}
	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS Abbreviations (url varchar(2048) PRIMARY KEY, abbreviation varchar(64) NOT NULL)")
	if err != nil {
		fmt.Println(err)
		panic("Cannot create table")
	}
	return &databaseStorage{conn: conn}
}

func (dbstor *databaseStorage) Store(value string) (string, error) {
	if !tools.IsURL(value) {
		return "", &storageError{reason: "The value " + value + " is not URL"}
	}
	key := tools.StringConverter.Do(value)
	var exists bool
	err := dbstor.conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT url, abbreviation FROM Abbreviations where url=$1)", value).Scan(&exists)
	if err != nil {
		return "", err
	}
	if exists {
		return "", &storageError{reason: "The value " + value + " is already added"}
	}
	dbstor.conn.Exec(context.Background(), "INSERT INTO Abbreviations VALUES ($1, $2)", value, key)
	return key, nil
}

func (dbstor *databaseStorage) Find(key string) (string, error) {
	var value string
	err := dbstor.conn.QueryRow(context.Background(), "SELECT url FROM Abbreviations where abbreviation=$1", key).Scan(&value)
	if err == pgx.ErrNoRows {
		return "", &storageError{reason: "The key " + key + " does not exist"}
	}
	if err != nil {
		return "", err
	}
	return value, nil
}
