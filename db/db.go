package db

import (
	"database/sql"

	"github.com/dortaedward/financeTracker/types"
)

type Db struct{
	Db *sql.DB
}

func CreateDb(db *sql.DB) *Db{
	return &Db{
		Db: types.ConnectDb(),
	}
}

