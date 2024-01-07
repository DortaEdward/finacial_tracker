package types

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func ConnectDb() *sql.DB{ 
  db, err := sql.Open("sqlite3","database.db"); if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  user_table_exists := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
      email TEXT
			username TEXT,
			password TEXT
		)
	` 
  _, err = db.Query(user_table_exists); if err != nil {
    log.Fatal(err)
  }
  return db
}

