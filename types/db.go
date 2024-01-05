package types

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func ConnectDb() *sql.DB{ 
  db, err := sql.Open("sqlite3",""); if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  return db
}

