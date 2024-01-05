package types

type Session struct{
  Id int `json:"id"`
  User_id int `json:"user_id"`
  Active bool `json:"active"`
  User_Agent string `json:"user_agent"`
}

