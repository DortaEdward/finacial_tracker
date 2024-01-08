package handlers

import (
	"github.com/dortaedward/financeTracker/db"
)

type UserHandlers struct{
	storage *db.Db
}

func CreateUserHandler(db *db.Db) *UserHandlers{
	return &UserHandlers{
		storage: db,	
	}
}

// create user
func (h *UserHandlers) CreateUser(){
	query := ""
	h.storage.Db.Query(query)
}


// get user
func (h *UserHandlers) FindUser(email string){

}


// update user


// delete user
