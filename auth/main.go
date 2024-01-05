package auth

import (
	"database/sql"

	"github.com/dortaedward/financeTracker/types"
)

func ComparePassword(password, hash_password string) bool {
  return false
}

func GenerateToken(){
 /*
 TODO
  - generate one of the following tokens
    - refresh token
    - access token
 */
}

func AuthenticateUser(user_creds types.SigninRequest, db *sql.DB){
  /*
  TODO
    - find user in db
    - compare passwords
    - find session  
      - if exist update to be active
      - if not create as active
    - create both tokens
    - return refresh and access tokens
  */
}

func ReissueRefreshToken(){
/*
  TODO
  - check if refresh token exists
    - if not return nil or err
    - if does create new token
*/ 
}

