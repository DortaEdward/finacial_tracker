package types

type User struct{
  Id int `json:"id"`
  Username string `json:"username"`
  Email string  `json:"email"`
  Password string `json:"password"` 
}

func CreateUser(username, email, password string) *User{
  return &User{ 
    Username: username,
    Email: email,
    Password: password,
  }
}
