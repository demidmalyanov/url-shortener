package database

import (
	"database/sql"

	"github.com/demidmalyanov/url-shortener/models"
	_ "github.com/mattn/go-sqlite3"
)

type TokenDB struct {
	db *sql.DB
}

const dbFile string = "tokens.db"

const create string = `
  CREATE TABLE IF NOT EXISTS tokens (
  id INTEGER NOT NULL PRIMARY KEY,
  url TEXT,
  token TEXT
);`

func NewDB() (*TokenDB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &TokenDB{
		db: db,
	}, nil
}

func (c *TokenDB) Insert(token models.Token) (int, error) {
	res, err := c.db.Exec("INSERT INTO tokens VALUES(NULL,?,?);", token.Token, token.Url)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (c *TokenDB) GetUrlByToken(token string) (string, error) {
	rows, _ := c.db.Query("SELECT url FROM people WHERE token =?",token)

	founToken := models.Token{}

	for rows.Next() {
		rows.Scan(&founToken.Url)
	}

	return founToken.Url, nil
}
