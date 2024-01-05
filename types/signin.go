package types

type SigninRequest struct{
  email string `json:"email" bind:"required"`
  password string `json:"password" bind:"required"`
}
